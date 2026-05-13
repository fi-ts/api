// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	apiclient "github.com/fi-ts/api/go/client"
{{ range $name, $api := . -}}
	"github.com/fi-ts/api/go{{ $api.Path }}/{{ $api.Name }}connect"
	{{ $name }}mocks "github.com/fi-ts/api/go/tests/mocks{{ $api.Path }}/{{ $api.Name }}connect"
{{ end }}
	"github.com/stretchr/testify/mock"
)

type (
	client struct {
{{ range $name, $api := . -}}
	{{ $name }}service *{{ $name}}
{{ end }}
	}

	ClientMockFns struct {
{{ range $name, $api := . -}}
	{{ $name | title }}Mocks *{{ $name | title }}MockFns
{{ end }}
	}

	wrapper struct {
		t *testing.T
	}
{{ range $name, $api := . -}}
    {{ $name }} struct {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }} *{{ $name }}mocks.{{ $svc.Name }}Client
{{ end }}
    }

	{{ $name | title }}MockFns struct {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | trimSuffix "Service" }}  func(m *mock.Mock)
{{ end }}
	}
{{ end }}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
{{ range $name, $api := . -}}
	{{ $name }}service: w.{{ $name |title }}(fns.{{ $name | title }}Mocks),
{{ end }}
	}
}

{{ range $name, $api := . -}}
func (c *client) {{ $name | title }}() apiclient.{{ $name | title }} {
	return c.{{ $name }}service
}
{{ end }}

{{ range $name, $api := . -}}
func (w wrapper) {{ $name | title }}(fns *{{ $name | title }}MockFns) *{{ $name }} {
	return new{{ $name }}(w.t, fns)
}

func new{{ $name }}(t *testing.T, fns *{{ $name | title }}MockFns) *{{ $name }} {
	a := &{{ $name }}{
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }}:  {{ $name }}mocks.New{{ $svc.Name }}Client(t),
{{ end }}
	}

	if fns != nil {
{{ range $svc := $api.Services -}}
		if fns.{{ $svc.Name | trimSuffix "Service" }} != nil {
			fns.{{ $svc.Name | trimSuffix "Service" }}(&a.{{ $svc.Name | lower }}.Mock)
		}
{{ end }}
	}

	return a
}

{{ range $svc := $api.Services -}}
func (c  *{{ $name }} ) {{ $svc.Name | trimSuffix "Service" }}() {{ $name }}connect.{{ $svc.Name }}Client {
	return c.{{ $svc.Name | lower }}
}
{{ end }}

{{ end }}
