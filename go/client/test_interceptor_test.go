package client_test

import (
	"log/slog"
	"testing"

	"connectrpc.com/connect"
	"github.com/fi-ts/api/go/client"
	apiv1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestInterceptor(t *testing.T) {
	cl, err := client.New(&client.DialConfig{
		BaseURL: "http://this-is-just-for-testing",
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
			}),
		},
		UserAgent: "cli-test",
		Log:       slog.Default(),
	})
	require.NoError(t, err)

	resp, err := cl.Apiv1().Version().Get(t.Context(), &apiv1.VersionServiceGetRequest{})
	require.NoError(t, err)

	if diff := cmp.Diff(&apiv1.VersionServiceGetResponse{
		Version: &apiv1.Version{
			Version: "1.0",
		},
	}, resp, protocmp.Transform(), client.IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
		t.Errorf("diff = %s", diff)
	}
}
