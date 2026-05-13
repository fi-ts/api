package client

import (
	"context"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/testing/protocmp"
)

type TestClientInterceptor struct {
	t     *testing.T
	calls []ClientCall
	count int
}

type ClientCall struct {
	WantRequest  proto.Message
	WantResponse func() connect.AnyResponse
	WantError    *connect.Error
}

func NewTestInterceptor(t *testing.T, calls []ClientCall) *TestClientInterceptor {
	return &TestClientInterceptor{
		t:     t,
		calls: calls,
	}
}

func (t *TestClientInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		defer func() { t.count++ }()

		if t.count >= len(t.calls) {
			t.t.Errorf("received an unexpected client call of type %T: %v", ar.Any(), ar.Any())
			t.t.FailNow()
		}

		call := t.calls[t.count]

		if diff := cmp.Diff(call.WantRequest, ar.Any(), protocmp.Transform(), IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
			t.t.Errorf("request diff (+got -want):\n %s", diff)
			t.t.FailNow()
		}

		if call.WantError != nil {
			return nil, call.WantError
		}

		return call.WantResponse(), nil
	}
}

func (t *TestClientInterceptor) WrapStreamingClient(connect.StreamingClientFunc) connect.StreamingClientFunc {
	t.t.Errorf("streaming not supported")
	return nil
}

func (t *TestClientInterceptor) WrapStreamingHandler(connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	t.t.Errorf("streaming not supported")
	return nil
}

func IgnoreUnexported() cmp.Option {
	// the exporter opt allows all unexported fields: https://github.com/google/go-cmp/pull/176
	return cmp.Exporter(func(reflect.Type) bool { return true })
}
