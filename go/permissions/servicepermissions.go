// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	"connectrpc.com/connect"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func GetServices() []string {
	return []string{
		"fits.api.v1.HealthService",
		"fits.api.v1.IPService",
		"fits.api.v1.MethodService",
		"fits.api.v1.ProjectService",
		"fits.api.v1.TenantService",
		"fits.api.v1.TokenService",
		"fits.api.v1.VersionService",
	}
}

func GetServicePermissions() *ServicePermissions {
	return &ServicePermissions{
		Roles: Roles{
			Admin: Admin{},
			Tenant: Tenant{
				v1.TenantRole_TENANT_ROLE_OWNER: map[string]struct{}{
					"/fits.api.v1.ProjectService/Create": {},
					"/fits.api.v1.TenantService/Delete":  {},
					"/fits.api.v1.TenantService/Get":     {},
					"/fits.api.v1.TenantService/Update":  {},
				},
				v1.TenantRole_TENANT_ROLE_EDITOR: map[string]struct{}{
					"/fits.api.v1.ProjectService/Create": {},
					"/fits.api.v1.TenantService/Delete":  {},
					"/fits.api.v1.TenantService/Get":     {},
					"/fits.api.v1.TenantService/Update":  {},
				},
				v1.TenantRole_TENANT_ROLE_VIEWER: map[string]struct{}{
					"/fits.api.v1.TenantService/Get": {},
				},
				v1.TenantRole_TENANT_ROLE_GUEST: map[string]struct{}{
					"/fits.api.v1.TenantService/Get": {},
				},
			},
			Project: Project{
				v1.ProjectRole_PROJECT_ROLE_OWNER: map[string]struct{}{
					"/fits.api.v1.IPService/Create":      {},
					"/fits.api.v1.IPService/Delete":      {},
					"/fits.api.v1.IPService/Get":         {},
					"/fits.api.v1.IPService/List":        {},
					"/fits.api.v1.IPService/Update":      {},
					"/fits.api.v1.ProjectService/Delete": {},
					"/fits.api.v1.ProjectService/Get":    {},
					"/fits.api.v1.ProjectService/Update": {},
				},
				v1.ProjectRole_PROJECT_ROLE_EDITOR: map[string]struct{}{
					"/fits.api.v1.IPService/Create":      {},
					"/fits.api.v1.IPService/Delete":      {},
					"/fits.api.v1.IPService/Get":         {},
					"/fits.api.v1.IPService/List":        {},
					"/fits.api.v1.IPService/Update":      {},
					"/fits.api.v1.ProjectService/Get":    {},
					"/fits.api.v1.ProjectService/Update": {},
				},
				v1.ProjectRole_PROJECT_ROLE_VIEWER: map[string]struct{}{
					"/fits.api.v1.IPService/Get":      {},
					"/fits.api.v1.IPService/List":     {},
					"/fits.api.v1.ProjectService/Get": {},
				},
			},
		},
		Methods: map[string]struct{}{
			"/fits.api.v1.HealthService/Get":                                 {},
			"/fits.api.v1.IPService/Create":                                  {},
			"/fits.api.v1.IPService/Delete":                                  {},
			"/fits.api.v1.IPService/Get":                                     {},
			"/fits.api.v1.IPService/List":                                    {},
			"/fits.api.v1.IPService/Update":                                  {},
			"/fits.api.v1.MethodService/List":                                {},
			"/fits.api.v1.MethodService/TokenScopedList":                     {},
			"/fits.api.v1.ProjectService/Create":                             {},
			"/fits.api.v1.ProjectService/Delete":                             {},
			"/fits.api.v1.ProjectService/Get":                                {},
			"/fits.api.v1.ProjectService/List":                               {},
			"/fits.api.v1.ProjectService/Update":                             {},
			"/fits.api.v1.TenantService/Create":                              {},
			"/fits.api.v1.TenantService/Delete":                              {},
			"/fits.api.v1.TenantService/Get":                                 {},
			"/fits.api.v1.TenantService/List":                                {},
			"/fits.api.v1.TenantService/Update":                              {},
			"/fits.api.v1.TokenService/Create":                               {},
			"/fits.api.v1.TokenService/Get":                                  {},
			"/fits.api.v1.TokenService/List":                                 {},
			"/fits.api.v1.TokenService/Refresh":                              {},
			"/fits.api.v1.TokenService/Revoke":                               {},
			"/fits.api.v1.TokenService/Update":                               {},
			"/fits.api.v1.VersionService/Get":                                {},
			"/grpc.reflection.v1.ServerReflection/ServerReflectionInfo":      {},
			"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo": {},
		},
		Visibility: Visibility{
			Public: map[string]bool{
				"/fits.api.v1.HealthService/Get":                                 true,
				"/fits.api.v1.MethodService/List":                                true,
				"/fits.api.v1.VersionService/Get":                                true,
				"/grpc.reflection.v1.ServerReflection/ServerReflectionInfo":      true,
				"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo": true,
			},
			Self: map[string]bool{
				"/fits.api.v1.MethodService/TokenScopedList": true,
				"/fits.api.v1.ProjectService/List":           true,
				"/fits.api.v1.TenantService/Create":          true,
				"/fits.api.v1.TenantService/List":            true,
				"/fits.api.v1.TokenService/Create":           true,
				"/fits.api.v1.TokenService/Get":              true,
				"/fits.api.v1.TokenService/List":             true,
				"/fits.api.v1.TokenService/Refresh":          true,
				"/fits.api.v1.TokenService/Revoke":           true,
				"/fits.api.v1.TokenService/Update":           true,
			},
			Admin: map[string]bool{},
			Tenant: map[string]bool{
				"/fits.api.v1.ProjectService/Create": true,
				"/fits.api.v1.TenantService/Delete":  true,
				"/fits.api.v1.TenantService/Get":     true,
				"/fits.api.v1.TenantService/Update":  true,
			},
			Project: map[string]bool{
				"/fits.api.v1.IPService/Create":      true,
				"/fits.api.v1.IPService/Delete":      true,
				"/fits.api.v1.IPService/Get":         true,
				"/fits.api.v1.IPService/List":        true,
				"/fits.api.v1.IPService/Update":      true,
				"/fits.api.v1.ProjectService/Delete": true,
				"/fits.api.v1.ProjectService/Get":    true,
				"/fits.api.v1.ProjectService/Update": true,
			},
		},
		Auditable: map[string]bool{
			"/fits.api.v1.HealthService/Get":             false,
			"/fits.api.v1.IPService/Create":              true,
			"/fits.api.v1.IPService/Delete":              true,
			"/fits.api.v1.IPService/Get":                 false,
			"/fits.api.v1.IPService/List":                false,
			"/fits.api.v1.IPService/Update":              true,
			"/fits.api.v1.MethodService/List":            false,
			"/fits.api.v1.MethodService/TokenScopedList": false,
			"/fits.api.v1.ProjectService/Create":         true,
			"/fits.api.v1.ProjectService/Delete":         true,
			"/fits.api.v1.ProjectService/Get":            false,
			"/fits.api.v1.ProjectService/List":           false,
			"/fits.api.v1.ProjectService/Update":         true,
			"/fits.api.v1.TenantService/Create":          true,
			"/fits.api.v1.TenantService/Delete":          true,
			"/fits.api.v1.TenantService/Get":             false,
			"/fits.api.v1.TenantService/List":            false,
			"/fits.api.v1.TenantService/Update":          true,
			"/fits.api.v1.TokenService/Create":           true,
			"/fits.api.v1.TokenService/Get":              true,
			"/fits.api.v1.TokenService/List":             true,
			"/fits.api.v1.TokenService/Refresh":          true,
			"/fits.api.v1.TokenService/Revoke":           true,
			"/fits.api.v1.TokenService/Update":           true,
			"/fits.api.v1.VersionService/Get":            false,
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
