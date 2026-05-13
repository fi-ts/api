// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	"connectrpc.com/connect"
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
			Admin:   Admin{
				{{- range $role, $methods := .Roles.Admin }}
					"{{ $role }}": []string{
						{{- range $method := $methods }}
							"{{ $method }}",
						{{- end }}
					},
				{{- end }}
			},
			Tenant:  Tenant{
				{{- range $role, $methods := .Roles.Tenant }}
					"{{ $role }}": []string{
						{{- range $method := $methods }}
							"{{ $method }}",
						{{- end }}
					},
				{{- end }}
			},
			Project: Project{
				{{- range $role, $methods := .Roles.Project }}
					"{{ $role }}": []string{
						{{- range $method := $methods }}
							"{{ $method }}",
						{{- end }}
					},
				{{- end }}
			},
		},
		Methods:    map[string]bool{
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
			Admin:    map[string]bool{
{{- range $key, $value := .Visibility.Admin }}
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

func IsAdminScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Admin[req.Spec().Procedure]
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
	switch rq := req.Any().(type) {
	case interface{ GetLogin() string }:
		return rq.GetLogin(), true
	}
	return "", false
}

func GetProjectFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsProjectScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetProject() string }:
		return rq.GetProject(), true
	}
	return "", false
}
