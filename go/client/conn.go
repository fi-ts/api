package client

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenEnvName     = "FCO_APIV1_TOKEN"
	TokenFileEnvName = "FCO_APIV1_TOKEN_FILE"
	BaseURLEnvName   = "FCO_APIV1_URL"
)

type (
	// DialConfig is the configuration to create a api-server connection
	DialConfig struct {
		// BaseUrl points to the apiv1 url where the apiserver is reachable
		BaseURL string
		// Token to be used to talk to the apiserver, the string representation of the token.
		// If Token is specified, TokenFile cannot be specified.
		// It is possible to renew this token automatically, see TokenRenewal.
		Token string
		// Tokenfile path to a file containing the string representation of the token.
		// If Tokenfile is specified, Token cannot be specified.
		// Token renewal must be done from outside
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
func (dc *DialConfig) parseTokenClaims() error {
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
