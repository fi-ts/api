package client_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log/slog"
	"os"
	"testing"
	"testing/synctest"

	"time"

	"connectrpc.com/connect"
	"github.com/fi-ts/api/go/client"
	apiv1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	var (
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	)

	synctest.Test(t, func(t *testing.T) {
		tokenString, err := generateToken(2 * time.Second)
		require.NoError(t, err)
		var renewedToken string

		c, err := client.New(&client.DialConfig{
			BaseURL: "http://localhost",
			Token:   tokenString,

			Interceptors: []connect.Interceptor{
				client.NewTestInterceptor(t, []client.ClientCall{
					{
						WantRequest: &apiv1.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv1.VersionServiceGetResponse{
								Version: &apiv1.Version{Version: "1.0"},
							})
						},
					},
					{
						WantRequest: &apiv1.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv1.VersionServiceGetResponse{
								Version: &apiv1.Version{Version: "1.0"},
							})
						},
					},
					{
						WantRequest: &apiv1.TokenServiceRefreshRequest{},
						WantResponse: func() connect.AnyResponse {
							tokenString, err := generateToken(2 * time.Second)
							require.NoError(t, err)

							return connect.NewResponse(&apiv1.TokenServiceRefreshResponse{
								Secret: tokenString,
							})
						},
					},
					{
						WantRequest: &apiv1.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv1.VersionServiceGetResponse{
								Version: &apiv1.Version{Version: "1.0"},
							})
						},
					},
				}),
			},
			TokenRenewal: &client.TokenRenewal{
				PersistTokenFn: func(token string) error {
					renewedToken = token
					return nil
				},
			},
			Log: log,
		})

		require.NoError(t, err)
		v, err := c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
		require.Empty(t, renewedToken)

		time.Sleep(1 * time.Second)
		v, err = c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
		require.Empty(t, renewedToken)

		time.Sleep(3 * time.Second)
		v, err = c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
		require.NotEmpty(t, renewedToken)
		require.NotEqual(t, renewedToken, tokenString, "haven't changed")
	})
}

func generateToken(duration time.Duration) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
