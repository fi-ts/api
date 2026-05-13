package client

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
)

const tokenRenewChecksDuringLifetime = 4

type (
	// DialConfig is the configuration to create a api-server connection
	DialConfig struct {
		BaseURL string
		Token   string

		// Optional client Interceptors
		Interceptors []connect.Interceptor

		UserAgent string
		// TokenRenewal defines if and how the token should be renewed
		TokenRenewal *TokenRenewal

		Transport http.RoundTripper

		Log *slog.Logger

		expiresAt time.Time
		issuedAt  time.Time
	}

	TokenRenewal struct {
		// PersistTokenFn is called to persist the newly fetched token
		// token will not be persisted if not specified
		PersistTokenFn PersistTokenFn
	}

	PersistTokenFn func(token string) error
)

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
	if dc.Token == "" {
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
