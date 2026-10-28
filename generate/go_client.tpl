// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"
	"github.com/redpanda-data/protoc-gen-go-mcp/pkg/runtime"

{{ range $name, $api := . -}}
	"github.com/fi-ts/api/go{{ $api.Path }}/{{ $api.ConnectName }}connect"
	"github.com/fi-ts/api/go{{ $api.Path }}/{{ $api.ConnectName }}mcp"
{{ end }}
)

type (
	Client interface {
{{ range $name, $api := . -}}
	{{ $name | title }}() {{ $name | title }}
{{ end }}
	}
	client struct {
		config *DialConfig

		interceptors []connect.Interceptor
	}
{{ range $name, $api := . -}}
	{{ $name | title }} interface {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | trimSuffix "Service" }}() {{ $api.ConnectName }}connect.{{ $svc.Name }}Client
{{ end }}
	EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option)
	}

    {{ $name }} struct {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }} {{ $api.ConnectName }}connect.{{ $svc.Name }}Client
{{ end }}
    }

{{ end }}
)

{{ range $name, $api := . -}}
func (c *client) {{ $name | title }}() {{ $name | title }} {
	a := &{{ $name }}{
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }}:  {{ $api.ConnectName }}connect.New{{ $svc.Name }}Client(
		c.config.HttpClient(),
		c.config.BaseURL,
		connect.WithInterceptors(c.interceptors...),
		compress.WithAll(compress.LevelBalanced),
	),
{{ end }}
	}
	return a
}

{{ range $svc := $api.Services -}}
func (c  *{{ $name }} ) {{ $svc.Name | trimSuffix "Service" }}() {{ $api.ConnectName }}connect.{{ $svc.Name }}Client {
	return c.{{ $svc.Name | lower }}
}
{{ end }}

{{ end }}


{{ range $name, $api := . -}}

func (c  *{{ $name }} ) EnableMCP(mcps runtime.MCPServer, opts ...runtime.Option) {
{{ range $svc := $api.Services -}}
	{{ $api.ConnectName }}mcp.ForwardTo{{ $svc.Name }}Client(mcps, c.{{ $svc.Name | trimSuffix "Service" }}(), opts...)
{{ end -}}
}

{{ end }}