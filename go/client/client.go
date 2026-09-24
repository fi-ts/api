// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"

	"github.com/fi-ts/api/go/fits/api/mvm/v1/mvmv1connect"
	"github.com/fi-ts/api/go/fits/api/v1/apiv1connect"
)

type (
	Client interface {
		Apimvmv1() Apimvmv1
		Apiv1() Apiv1
	}
	client struct {
		config *DialConfig

		interceptors []connect.Interceptor
	}
	Apimvmv1 interface {
		Location() mvmv1connect.LocationServiceClient
		MVM() mvmv1connect.MVMServiceClient
		OS() mvmv1connect.OSServiceClient
		StageType() mvmv1connect.StageTypeServiceClient
		VLAN() mvmv1connect.VLANServiceClient
	}

	apimvmv1 struct {
		locationservice  mvmv1connect.LocationServiceClient
		mvmservice       mvmv1connect.MVMServiceClient
		osservice        mvmv1connect.OSServiceClient
		stagetypeservice mvmv1connect.StageTypeServiceClient
		vlanservice      mvmv1connect.VLANServiceClient
	}

	Apiv1 interface {
		Health() apiv1connect.HealthServiceClient
		IP() apiv1connect.IPServiceClient
		Method() apiv1connect.MethodServiceClient
		Project() apiv1connect.ProjectServiceClient
		Tenant() apiv1connect.TenantServiceClient
		Token() apiv1connect.TokenServiceClient
		Version() apiv1connect.VersionServiceClient
	}

	apiv1 struct {
		healthservice  apiv1connect.HealthServiceClient
		ipservice      apiv1connect.IPServiceClient
		methodservice  apiv1connect.MethodServiceClient
		projectservice apiv1connect.ProjectServiceClient
		tenantservice  apiv1connect.TenantServiceClient
		tokenservice   apiv1connect.TokenServiceClient
		versionservice apiv1connect.VersionServiceClient
	}
)

func (c *client) Apimvmv1() Apimvmv1 {
	a := &apimvmv1{
		locationservice: mvmv1connect.NewLocationServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		mvmservice: mvmv1connect.NewMVMServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		osservice: mvmv1connect.NewOSServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		stagetypeservice: mvmv1connect.NewStageTypeServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		vlanservice: mvmv1connect.NewVLANServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *apimvmv1) Location() mvmv1connect.LocationServiceClient {
	return c.locationservice
}
func (c *apimvmv1) MVM() mvmv1connect.MVMServiceClient {
	return c.mvmservice
}
func (c *apimvmv1) OS() mvmv1connect.OSServiceClient {
	return c.osservice
}
func (c *apimvmv1) StageType() mvmv1connect.StageTypeServiceClient {
	return c.stagetypeservice
}
func (c *apimvmv1) VLAN() mvmv1connect.VLANServiceClient {
	return c.vlanservice
}

func (c *client) Apiv1() Apiv1 {
	a := &apiv1{
		healthservice: apiv1connect.NewHealthServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		ipservice: apiv1connect.NewIPServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		methodservice: apiv1connect.NewMethodServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		projectservice: apiv1connect.NewProjectServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: apiv1connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		tokenservice: apiv1connect.NewTokenServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		versionservice: apiv1connect.NewVersionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
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
