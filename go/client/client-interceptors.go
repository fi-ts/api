package client

import (
	"context"

	"connectrpc.com/connect"
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
