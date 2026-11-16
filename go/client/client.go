// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"
	"github.com/redpanda-data/protoc-gen-go-mcp/pkg/runtime"

	"github.com/fi-ts/api/go/fits/api/mvm/v1/mvmv1connect"
	"github.com/fi-ts/api/go/fits/api/mvm/v1/mvmv1mcp"
	"github.com/fi-ts/api/go/fits/api/pcvm/v1/pcvmv1connect"
	"github.com/fi-ts/api/go/fits/api/pcvm/v1/pcvmv1mcp"
	"github.com/fi-ts/api/go/fits/api/v1/apiv1connect"
	"github.com/fi-ts/api/go/fits/api/v1/apiv1mcp"
)

type (
	Client interface {
		Apimvmv1() Apimvmv1
		Apipcvmv1() Apipcvmv1
		Apiv1() Apiv1
	}
	client struct {
		config *DialConfig

		interceptors []connect.Interceptor
	}
	Apimvmv1 interface {
		Contact() mvmv1connect.ContactServiceClient
		Domain() mvmv1connect.DomainServiceClient
		LDAP() mvmv1connect.LDAPServiceClient
		Location() mvmv1connect.LocationServiceClient
		MVM() mvmv1connect.MVMServiceClient
		OS() mvmv1connect.OSServiceClient
		VLAN() mvmv1connect.VLANServiceClient

		EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option)
	}

	apimvmv1 struct {
		contactservice  mvmv1connect.ContactServiceClient
		domainservice   mvmv1connect.DomainServiceClient
		ldapservice     mvmv1connect.LDAPServiceClient
		locationservice mvmv1connect.LocationServiceClient
		mvmservice      mvmv1connect.MVMServiceClient
		osservice       mvmv1connect.OSServiceClient
		vlanservice     mvmv1connect.VLANServiceClient
	}

	Apipcvmv1 interface {
		PublicCloudVM() pcvmv1connect.PublicCloudVMServiceClient

		EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option)
	}

	apipcvmv1 struct {
		publiccloudvmservice pcvmv1connect.PublicCloudVMServiceClient
	}

	Apiv1 interface {
		Health() apiv1connect.HealthServiceClient
		Method() apiv1connect.MethodServiceClient
		Project() apiv1connect.ProjectServiceClient
		Tenant() apiv1connect.TenantServiceClient
		Token() apiv1connect.TokenServiceClient
		Version() apiv1connect.VersionServiceClient

		EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option)
	}

	apiv1 struct {
		healthservice  apiv1connect.HealthServiceClient
		methodservice  apiv1connect.MethodServiceClient
		projectservice apiv1connect.ProjectServiceClient
		tenantservice  apiv1connect.TenantServiceClient
		tokenservice   apiv1connect.TokenServiceClient
		versionservice apiv1connect.VersionServiceClient
	}
)

func (c *client) Apimvmv1() Apimvmv1 {
	a := &apimvmv1{
		contactservice: mvmv1connect.NewContactServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		domainservice: mvmv1connect.NewDomainServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		ldapservice: mvmv1connect.NewLDAPServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
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
		vlanservice: mvmv1connect.NewVLANServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *apimvmv1) Contact() mvmv1connect.ContactServiceClient {
	return c.contactservice
}
func (c *apimvmv1) Domain() mvmv1connect.DomainServiceClient {
	return c.domainservice
}
func (c *apimvmv1) LDAP() mvmv1connect.LDAPServiceClient {
	return c.ldapservice
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
func (c *apimvmv1) VLAN() mvmv1connect.VLANServiceClient {
	return c.vlanservice
}

func (c *client) Apipcvmv1() Apipcvmv1 {
	a := &apipcvmv1{
		publiccloudvmservice: pcvmv1connect.NewPublicCloudVMServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *apipcvmv1) PublicCloudVM() pcvmv1connect.PublicCloudVMServiceClient {
	return c.publiccloudvmservice
}

func (c *client) Apiv1() Apiv1 {
	a := &apiv1{
		healthservice: apiv1connect.NewHealthServiceClient(
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

func (c *apimvmv1) EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option) {
	mvmv1mcp.ForwardToContactServiceClient(mcps, c.Contact(), opts...)
	mvmv1mcp.ForwardToDomainServiceClient(mcps, c.Domain(), opts...)
	mvmv1mcp.ForwardToLDAPServiceClient(mcps, c.LDAP(), opts...)
	mvmv1mcp.ForwardToLocationServiceClient(mcps, c.Location(), opts...)
	mvmv1mcp.ForwardToMVMServiceClient(mcps, c.MVM(), opts...)
	mvmv1mcp.ForwardToOSServiceClient(mcps, c.OS(), opts...)
	mvmv1mcp.ForwardToVLANServiceClient(mcps, c.VLAN(), opts...)
}

func (c *apipcvmv1) EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option) {
	pcvmv1mcp.ForwardToPublicCloudVMServiceClient(mcps, c.PublicCloudVM(), opts...)
}

func (c *apiv1) EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option) {
	apiv1mcp.ForwardToHealthServiceClient(mcps, c.Health(), opts...)
	apiv1mcp.ForwardToMethodServiceClient(mcps, c.Method(), opts...)
	apiv1mcp.ForwardToProjectServiceClient(mcps, c.Project(), opts...)
	apiv1mcp.ForwardToTenantServiceClient(mcps, c.Tenant(), opts...)
	apiv1mcp.ForwardToTokenServiceClient(mcps, c.Token(), opts...)
	apiv1mcp.ForwardToVersionServiceClient(mcps, c.Version(), opts...)
}
