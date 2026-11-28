package client_test

import (
	"context"
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
						WantRequest: &apiv1.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv1.VersionServiceGetResponse{
								Version: &apiv1.Version{Version: "1.0"},
							})
						},
					},
				}),
			},
			Log: log,
		})

		require.NoError(t, err)
		v, err := c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)

		time.Sleep(1 * time.Second)
		v, err = c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)

		time.Sleep(3 * time.Second)
		v, err = c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
	})
}

func Test_Client_ApiKeySecret(t *testing.T) {
	var (
		log       = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		secret    = client.ApiKeySecretPrefix + "v1_abcd1234"
		authCalls []string
	)

	c, err := client.New(&client.DialConfig{
		BaseURL: "http://localhost",
		Token:   secret,

		Interceptors: []connect.Interceptor{
			&authorizationRecorder{authorization: &authCalls},
			client.NewTestInterceptor(t, []client.ClientCall{
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
		Log: log,
	})
	require.NoError(t, err)

	v, err := c.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Version.Version)

	require.Equal(t, []string{"Bearer " + secret}, authCalls)
}

// authorizationRecorder records the authorization header of outgoing requests.
type authorizationRecorder struct {
	authorization *[]string
}

func (r *authorizationRecorder) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		*r.authorization = append(*r.authorization, request.Header().Get("Authorization"))
		return next(ctx, request)
	}
}

func (r *authorizationRecorder) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (r *authorizationRecorder) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
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
