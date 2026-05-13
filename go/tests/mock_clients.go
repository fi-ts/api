// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	apiclient "github.com/fi-ts/api/go/client"
	"github.com/fi-ts/api/go/fits/api/v1/apiv1connect"
	apiv1mocks "github.com/fi-ts/api/go/tests/mocks/fits/api/v1/apiv1connect"

	"github.com/stretchr/testify/mock"
)

type (
	client struct {
		apiv1service *apiv1
	}

	ClientMockFns struct {
		Apiv1Mocks *Apiv1MockFns
	}

	wrapper struct {
		t *testing.T
	}
	apiv1 struct {
		healthservice  *apiv1mocks.HealthServiceClient
		ipservice      *apiv1mocks.IPServiceClient
		methodservice  *apiv1mocks.MethodServiceClient
		projectservice *apiv1mocks.ProjectServiceClient
		tenantservice  *apiv1mocks.TenantServiceClient
		tokenservice   *apiv1mocks.TokenServiceClient
		versionservice *apiv1mocks.VersionServiceClient
	}

	Apiv1MockFns struct {
		Health  func(m *mock.Mock)
		IP      func(m *mock.Mock)
		Method  func(m *mock.Mock)
		Project func(m *mock.Mock)
		Tenant  func(m *mock.Mock)
		Token   func(m *mock.Mock)
		Version func(m *mock.Mock)
	}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
		apiv1service: w.Apiv1(fns.Apiv1Mocks),
	}
}

func (c *client) Apiv1() apiclient.Apiv1 {
	return c.apiv1service
}

func (w wrapper) Apiv1(fns *Apiv1MockFns) *apiv1 {
	return newapiv1(w.t, fns)
}

func newapiv1(t *testing.T, fns *Apiv1MockFns) *apiv1 {
	a := &apiv1{
		healthservice:  apiv1mocks.NewHealthServiceClient(t),
		ipservice:      apiv1mocks.NewIPServiceClient(t),
		methodservice:  apiv1mocks.NewMethodServiceClient(t),
		projectservice: apiv1mocks.NewProjectServiceClient(t),
		tenantservice:  apiv1mocks.NewTenantServiceClient(t),
		tokenservice:   apiv1mocks.NewTokenServiceClient(t),
		versionservice: apiv1mocks.NewVersionServiceClient(t),
	}

	if fns != nil {
		if fns.Health != nil {
			fns.Health(&a.healthservice.Mock)
		}
		if fns.IP != nil {
			fns.IP(&a.ipservice.Mock)
		}
		if fns.Method != nil {
			fns.Method(&a.methodservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}
		if fns.Version != nil {
			fns.Version(&a.versionservice.Mock)
		}

	}

	return a
}

func (c *apiv1) Health() apiv1connect.HealthServiceClient {
	return c.healthservice
}
func (c *apiv1) IP() apiv1connect.IPServiceClient {
	return c.ipservice
}
func (c *apiv1) Method() apiv1connect.MethodServiceClient {
	return c.methodservice
}
func (c *apiv1) Project() apiv1connect.ProjectServiceClient {
	return c.projectservice
}
func (c *apiv1) Tenant() apiv1connect.TenantServiceClient {
	return c.tenantservice
}
func (c *apiv1) Token() apiv1connect.TokenServiceClient {
	return c.tokenservice
}
func (c *apiv1) Version() apiv1connect.VersionServiceClient {
	return c.versionservice
}
