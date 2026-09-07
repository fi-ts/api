package client

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"connectrpc.com/connect"
	apiv1models "github.com/fi-ts/api/go/fits/api/v1"
)

// authinterceptor adds the required auth headers
type authInterceptor struct {
	config *DialConfig
}

func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		request.Header().Add("Authorization", "Bearer "+i.config.Token)
		return next(ctx, request)
	})
}

func (i *authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return &streamingInterceptorConn{
			StreamingClientConn: next(ctx, spec),
			token:               i.config.Token,
		}
	}
}

func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

type streamingInterceptorConn struct {
	connect.StreamingClientConn
	token string
}

func (conn *streamingInterceptorConn) Send(m any) error {
	conn.RequestHeader().Add("Authorization", "Bearer "+conn.token)
	return conn.StreamingClientConn.Send(m)
}

type loggingInterceptor struct {
	config *DialConfig
}

func (i *loggingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		i.config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "body", request.Any())
		response, err := next(ctx, request)
		if err != nil {
			return nil, err
		}
		i.config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "response", response.Any())
		return response, err
	})
}

func (i *loggingInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	// TODO also log here
	return next
}

func (i *loggingInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

type tokenRenewingInterceptor struct {
	config *DialConfig
	client *client

	renewing atomic.Bool

	sync.Mutex
}

func (i *tokenRenewingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		err := i.renewTokenIfNeeded()
		if err != nil {
			return nil, err
		}
		return next(ctx, request)
	})
}

func (i *tokenRenewingInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (i *tokenRenewingInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (i *tokenRenewingInterceptor) renewTokenIfNeeded() error {
	if i.config.expiresAt.IsZero() {
		return nil
	}
	if i.renewing.Load() {
		return nil
	}
	if i.config.Log == nil {
		i.config.Log = slog.Default()
	}

	replaceBefore := i.config.expiresAt.Sub(i.config.issuedAt) / tokenRenewChecksDuringLifetime

	if time.Until(i.config.expiresAt) > replaceBefore {
		return nil
	}

	i.renewing.Store(true)
	defer i.renewing.Store(false)

	i.config.Log.Info("call token refresh, current token expires soon", "expires", i.config.expiresAt.String())

	i.Lock()
	defer i.Unlock()

	resp, err := i.client.Apiv1().Token().Refresh(context.Background(), &apiv1models.TokenServiceRefreshRequest{})
	if err != nil {
		return fmt.Errorf("unable to refresh token %w", err)
	}

	i.config.Token = resp.Secret
	err = i.config.parse()
	if err != nil {
		return fmt.Errorf("unable to parse token %w", err)
	}

	if i.config.TokenRenewal.PersistTokenFn == nil {
		return nil
	}

	err = i.config.TokenRenewal.PersistTokenFn(i.config.Token)
	if err != nil {
		return fmt.Errorf("unable to persist token %w", err)
	}

	i.config.Log.Info("token refreshed, new token expires in", "expires", i.config.expiresAt.String())
	return nil
}
