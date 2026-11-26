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
		},
		Auditable:  map[string]bool{
{{- range $key, $value := .Auditable }}
	"{{ $key }}": {{ $value }} ,
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

func IsTenantScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Tenant[req.Spec().Procedure]
	return ok
}

func IsProjectScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Project[req.Spec().Procedure]
	return ok
}

func IsAuditable(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Auditable[req.Spec().Procedure]
	return ok
}

func GetTenantFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsTenantScope(req) {
		return "", false
	}
	return getTenantFromMessage(req.Any())
}

// getTenantFromMessage returns the tenant of a request message, or ok=false
// if the message has no tenant getter. It is the lookup GetTenantFromRequest
// uses, split out so it can be tested against every request type without a
// running connect handler.
func getTenantFromMessage(msg any) (string, bool) {
	switch rq := msg.(type) {
	case interface{ GetTenant() string }:
		return rq.GetTenant(), true
	}
	return "", false
}

func GetProjectFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsProjectScope(req) {
		return "", false
	}
	return getProjectFromMessage(req.Any())
}

// getProjectFromMessage returns the project of a request message, or ok=false
// if the message has no project getter. It is the lookup GetProjectFromRequest
// uses, split out so it can be tested against every request type without a
// running connect handler.
//
// Requests that only wrap another request (e.g. Validate*) have no top-level
// project field and need a hand-written GetProjectUuid in a *_project_scope.go
// file next to the generated code.
func getProjectFromMessage(msg any) (string, bool) {
	switch rq := msg.(type) {
	case interface{ GetProjectUuid() string }:
		return rq.GetProjectUuid(), true
		// TODO: all services should use project_uuid; drop the GetProject
		// fallback once every request is streamlined.
	case interface{ GetProject() string }:
		return rq.GetProject(), true
	}
	return "", false
}
