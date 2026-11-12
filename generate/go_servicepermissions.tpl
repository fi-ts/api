// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	"connectrpc.com/connect"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func GetServices() []string {
	return []string{
{{- range $s := .Services }}
	"{{ $s }}",
{{- end }}
	}
}

func GetServicePermissions() *ServicePermissions {
	return &ServicePermissions{
		Roles:      Roles{
			Operator: Operator{
				{{- range $role, $methods := .Roles.Operator }}
					v1.OperatorRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Tenant:  Tenant{
				{{- range $role, $methods := .Roles.Tenant }}
					v1.TenantRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Project: Project{
				{{- range $role, $methods := .Roles.Project }}
					v1.ProjectRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Resource: Resource{
				{{- range $role, $methods := .Roles.Resource }}
					v1.ResourceRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
		},
		Methods:    map[string]struct{}{
{{- range $key, $value := .Methods }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
		},
		Visibility: Visibility{
			Public:  map[string]bool{
{{- range $key, $value := .Visibility.Public }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Self:    map[string]bool{
{{- range $key, $value := .Visibility.Self }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Operator:    map[string]bool{
{{- range $key, $value := .Visibility.Operator }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Tenant:    map[string]bool{
{{- range $key, $value := .Visibility.Tenant }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Project:    map[string]bool{
{{- range $key, $value := .Visibility.Project }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Resource:    map[string]bool{
{{- range $key, $value := .Visibility.Resource }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
		},
		Auditable:  map[string]bool{
{{- range $key, $value := .Auditable }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
		},
		Resources:  map[string]v1.Resource{
{{- range $key, $value := .Resources }}
	"{{ $key }}": v1.Resource_{{ $value }} ,
{{- end }}
		},
	}
}

func IsPublicScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Public[req.Spec().Procedure]
	return ok
}

func IsSelfScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Self[req.Spec().Procedure]
	return ok
}

func IsOperatorScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Operator[req.Spec().Procedure]
	return ok
}

func IsTenantScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Tenant[req.Spec().Procedure]
	return ok
}

func IsProjectScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Project[req.Spec().Procedure]
	return ok
}

func IsResourceScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Resource[req.Spec().Procedure]
	return ok
}

// GetResource returns the resource a resource scoped method acts on.
func GetResource(req connect.AnyRequest) (v1.Resource, bool) {
	resource, ok := GetServicePermissions().Resources[req.Spec().Procedure]
	return resource, ok
}

func IsAuditable(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Auditable[req.Spec().Procedure]
	return ok
}

func GetTenantFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsTenantScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetTenant() string }:
		return rq.GetTenant(), true
	}
	return "", false
}

func GetProjectFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsProjectScope(req) && !IsResourceScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetProjectUuid() string }:
		return rq.GetProjectUuid(), true
		// TODO: all services should use project_uuid; drop the GetProject
		// fallback once every request is streamlined.
	case interface{ GetProject() string }:
		return rq.GetProject(), true
	}
	return "", false
}

// GetResourceFromRequest returns the product instance id from a product scoped request.
func GetResourceFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsResourceScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetUuid() string }:
		return rq.GetUuid(), true
	case interface{ GetId() string }:
		return rq.GetId(), true
	}
	return "", false
}
