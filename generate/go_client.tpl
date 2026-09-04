// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"

{{ range $name, $api := . -}}
	"github.com/fi-ts/api/go{{ $api.Path }}/{{ $api.ConnectName }}connect"
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
	}

    {{ $name }} struct {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }} {{ $api.ConnectName }}connect.{{ $svc.Name }}Client
{{ end }}
    }

{{ end }}
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
