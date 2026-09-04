// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"

	"github.com/fi-ts/api/go/fits/api/v1/apiv1connect"
	"github.com/fi-ts/api/go/fits/api/vm/v1/vmv1connect"
)

type (
	Client interface {
		Apiv1() Apiv1
		Apivmv1() Apivmv1
	}
	client struct {
		config *DialConfig

		interceptors []connect.Interceptor
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

	Apivmv1 interface {
		VM() vmv1connect.VMServiceClient
	}

	apivmv1 struct {
		vmservice vmv1connect.VMServiceClient
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

		if config.TokenRenewal != nil {
			tokenRenewingInterceptor := &tokenRenewingInterceptor{config: config, client: c}
			c.interceptors = append(c.interceptors, tokenRenewingInterceptor)
		}
	}
	if config.Log != nil {
		loggingInterceptor := &loggingInterceptor{config: config}
		c.interceptors = append(c.interceptors, loggingInterceptor)
	}
	c.interceptors = append(c.interceptors, config.Interceptors...)

	return c, nil
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

func (c *client) Apivmv1() Apivmv1 {
	a := &apivmv1{
		vmservice: vmv1connect.NewVMServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *apivmv1) VM() vmv1connect.VMServiceClient {
	return c.vmservice
}
