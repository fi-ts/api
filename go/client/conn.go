package client

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenEnvName     = "FCO_APIV1_TOKEN"
	TokenFileEnvName = "FCO_APIV1_TOKEN_FILE"
	BaseURLEnvName   = "FCO_APIV1_URL"

	// ApiKeySecretPrefix is the common prefix of all api key secrets, e.g. fco_v1_abcd12....
	// Secrets of api keys are opaque strings and not jwts, the version within the
	// prefix selects the generation and hash scheme of the secret, detection of api
	// key secrets stays version independent via this common prefix.
	ApiKeySecretPrefix = "fco_"
)

type (
	// DialConfig is the configuration to create a api-server connection
	DialConfig struct {
		// BaseUrl points to the apiv1 url where the apiserver is reachable
		BaseURL string
		// Token to be used to talk to the apiserver, the string representation of the token.
		// This can be a jwt or the secret of an api key, e.g. fco_v1_abcd12....
		// If Token is specified, TokenFile cannot be specified.
		// It is possible to renew a jwt automatically, see TokenRenewal.
		// Api key secrets are managed by the apiserver, renewal does not apply to them.
		Token string
		// Tokenfile path to a file containing the string representation of the token.
		// If Tokenfile is specified, Token cannot be specified.
		TokenFile string

		// Optional client Interceptors
		Interceptors []connect.Interceptor

		UserAgent string

		// Transport optional, can be used to configure how the http transport works.
		Transport http.RoundTripper

		Log *slog.Logger

		expiresAt time.Time
		issuedAt  time.Time
	}
)

func New(config *DialConfig) (Client, error) {
	err := config.parse()
	if err != nil {
		return nil, err
	}

	c := &client{
		config:       config,
		interceptors: []connect.Interceptor{},
	}

	if config.Token != "" {
		authInterceptor := &authInterceptor{config: config}
		c.interceptors = append(c.interceptors, authInterceptor)
	}

	if config.TokenFile != "" {
		authInterceptor := &authInterceptor{config: config}
		c.interceptors = append(c.interceptors, authInterceptor)
	}

	if config.Log != nil {
		loggingInterceptor := &loggingInterceptor{config: config}
		c.interceptors = append(c.interceptors, loggingInterceptor)
	}
	c.interceptors = append(c.interceptors, config.Interceptors...)

	return c, nil
}

func (d *DialConfig) HttpClient() *http.Client {
	transport := http.DefaultTransport
	if d.Transport != nil {
		transport = d.Transport
	}

	return &http.Client{
		Transport: transport,
	}
}

func (dc *DialConfig) parse() error {
	if dc.BaseURL == "" {
		dc.BaseURL = os.Getenv(BaseURLEnvName)
		if dc.BaseURL == "" {
			return fmt.Errorf("neither BaseURL nor %s were given", BaseURLEnvName)
		}
		if _, err := url.Parse(dc.BaseURL); err != nil {
			return err
		}
	}

	if dc.Token == "" {
		dc.Token = os.Getenv(TokenEnvName)
	}

	if dc.TokenFile == "" {
		dc.TokenFile = os.Getenv(TokenFileEnvName)
	}

	if dc.Token != "" && dc.TokenFile != "" {
		return fmt.Errorf("either token or tokenfile must be specified, not both")
	}

	if dc.Token == "" && dc.TokenFile != "" {
		content, err := os.ReadFile(dc.TokenFile)
		if err != nil {
			return err
		}
		dc.Token = string(content)
	}

	if dc.Token == "" && dc.TokenFile == "" {
		return nil
	}

	return dc.parseTokenClaims()
}

// parseTokenClaims extracts expiresAt and issuedAt from the current Token.
// It is called both when initially parsing the config and when a tokenfile
// token is re-read, in which case only the claims need to be refreshed.
// Secrets of api keys, e.g. fco_v1_abcd12..., are opaque and not jwts, for
// them there are no claims to parse and token renewal does not apply.
func (dc *DialConfig) parseTokenClaims() error {
	if !strings.Contains(dc.Token, ".") || strings.HasPrefix(dc.Token, ApiKeySecretPrefix) {
		// not a jwt, most likely the secret of an api key, nothing to parse
		return nil
	}
	parsed, err := jwt.Parse(dc.Token, nil)
	if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
		return fmt.Errorf("unable to parse token:%w", err)
	}
	expiresAt, err := parsed.Claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("unable to extract expiresAt from token:%w", err)
	}
	if expiresAt != nil {
		dc.expiresAt = expiresAt.Time
	}

	issuedAt, err := parsed.Claims.GetIssuedAt()
	if err != nil {
		return fmt.Errorf("unable to extract issuedAt from token:%w", err)
	}
	if issuedAt != nil {
		dc.issuedAt = issuedAt.Time
	}
	if dc.issuedAt.IsZero() {
		dc.issuedAt = time.Now()
	}

	return nil
}
