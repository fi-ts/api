// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	"connectrpc.com/connect"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func GetServices() []string {
	return []string{
		"fits.api.mvm.v1.ContactService",
		"fits.api.mvm.v1.DomainService",
		"fits.api.mvm.v1.LDAPService",
		"fits.api.mvm.v1.LocationService",
		"fits.api.mvm.v1.MVMService",
		"fits.api.mvm.v1.OSService",
		"fits.api.mvm.v1.StageService",
		"fits.api.mvm.v1.VLANService",
		"fits.api.pcvm.v1.PublicCloudVMService",
		"fits.api.v1.HealthService",
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
			Operator: Operator{},
			Tenant: Tenant{
				v1.TenantRole_TENANT_ROLE_ADMIN: map[string]struct{}{
					"/fits.api.mvm.v1.ContactService/Create": {},
					"/fits.api.mvm.v1.ContactService/List":   {},
					"/fits.api.mvm.v1.ContactService/Update": {},
					"/fits.api.mvm.v1.DomainService/List":    {},
					"/fits.api.mvm.v1.LDAPService/List":      {},
					"/fits.api.mvm.v1.OSService/List":        {},
					"/fits.api.mvm.v1.VLANService/List":      {},
					"/fits.api.v1.ProjectService/Create":     {},
					"/fits.api.v1.TenantService/Get":         {},
					"/fits.api.v1.TenantService/Update":      {},
				},
				v1.TenantRole_TENANT_ROLE_EDITOR: map[string]struct{}{
					"/fits.api.mvm.v1.ContactService/Create": {},
					"/fits.api.mvm.v1.ContactService/List":   {},
					"/fits.api.mvm.v1.ContactService/Update": {},
					"/fits.api.mvm.v1.DomainService/List":    {},
					"/fits.api.mvm.v1.LDAPService/List":      {},
					"/fits.api.mvm.v1.OSService/List":        {},
					"/fits.api.mvm.v1.VLANService/List":      {},
					"/fits.api.v1.ProjectService/Create":     {},
					"/fits.api.v1.TenantService/Get":         {},
					"/fits.api.v1.TenantService/Update":      {},
				},
				v1.TenantRole_TENANT_ROLE_VIEWER: map[string]struct{}{
					"/fits.api.mvm.v1.ContactService/List": {},
					"/fits.api.mvm.v1.DomainService/List":  {},
					"/fits.api.mvm.v1.LDAPService/List":    {},
					"/fits.api.mvm.v1.OSService/List":      {},
					"/fits.api.mvm.v1.VLANService/List":    {},
					"/fits.api.v1.TenantService/Get":       {},
				},
				v1.TenantRole_TENANT_ROLE_GUEST: map[string]struct{}{
					"/fits.api.v1.TenantService/Get": {},
				},
			},
			Project: Project{
				v1.ProjectRole_PROJECT_ROLE_ADMIN: map[string]struct{}{
					"/fits.api.mvm.v1.MVMService/AddDisks":                   {},
					"/fits.api.mvm.v1.MVMService/AddIPAddress":               {},
					"/fits.api.mvm.v1.MVMService/AddSnapshot":                {},
					"/fits.api.mvm.v1.MVMService/Create":                     {},
					"/fits.api.mvm.v1.MVMService/Delete":                     {},
					"/fits.api.mvm.v1.MVMService/DeleteDisk":                 {},
					"/fits.api.mvm.v1.MVMService/DeleteIPAddress":            {},
					"/fits.api.mvm.v1.MVMService/DeleteSnapshot":             {},
					"/fits.api.mvm.v1.MVMService/Get":                        {},
					"/fits.api.mvm.v1.MVMService/List":                       {},
					"/fits.api.mvm.v1.MVMService/Move":                       {},
					"/fits.api.mvm.v1.MVMService/MoveIPAddress":              {},
					"/fits.api.mvm.v1.MVMService/Reinstall":                  {},
					"/fits.api.mvm.v1.MVMService/Restart":                    {},
					"/fits.api.mvm.v1.MVMService/Update":                     {},
					"/fits.api.mvm.v1.MVMService/UpdateDisk":                 {},
					"/fits.api.mvm.v1.MVMService/UpdateServiceClass":         {},
					"/fits.api.mvm.v1.MVMService/ValidateAddDisks":           {},
					"/fits.api.mvm.v1.MVMService/ValidateCreate":             {},
					"/fits.api.mvm.v1.MVMService/ValidateReinstall":          {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdate":             {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdateDisk":         {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdateServiceClass": {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Create":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Delete":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Get":             {},
					"/fits.api.pcvm.v1.PublicCloudVMService/List":            {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Update":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/ValidateCreate":  {},
					"/fits.api.pcvm.v1.PublicCloudVMService/ValidateUpdate":  {},
					"/fits.api.v1.ProjectService/Delete":                     {},
					"/fits.api.v1.ProjectService/Get":                        {},
					"/fits.api.v1.ProjectService/Update":                     {},
				},
				v1.ProjectRole_PROJECT_ROLE_EDITOR: map[string]struct{}{
					"/fits.api.mvm.v1.MVMService/AddDisks":                   {},
					"/fits.api.mvm.v1.MVMService/AddIPAddress":               {},
					"/fits.api.mvm.v1.MVMService/AddSnapshot":                {},
					"/fits.api.mvm.v1.MVMService/Create":                     {},
					"/fits.api.mvm.v1.MVMService/Get":                        {},
					"/fits.api.mvm.v1.MVMService/List":                       {},
					"/fits.api.mvm.v1.MVMService/Move":                       {},
					"/fits.api.mvm.v1.MVMService/MoveIPAddress":              {},
					"/fits.api.mvm.v1.MVMService/Restart":                    {},
					"/fits.api.mvm.v1.MVMService/Update":                     {},
					"/fits.api.mvm.v1.MVMService/UpdateDisk":                 {},
					"/fits.api.mvm.v1.MVMService/UpdateServiceClass":         {},
					"/fits.api.mvm.v1.MVMService/ValidateAddDisks":           {},
					"/fits.api.mvm.v1.MVMService/ValidateCreate":             {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdate":             {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdateDisk":         {},
					"/fits.api.mvm.v1.MVMService/ValidateUpdateServiceClass": {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Create":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Delete":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Get":             {},
					"/fits.api.pcvm.v1.PublicCloudVMService/List":            {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Update":          {},
					"/fits.api.pcvm.v1.PublicCloudVMService/ValidateCreate":  {},
					"/fits.api.pcvm.v1.PublicCloudVMService/ValidateUpdate":  {},
					"/fits.api.v1.ProjectService/Get":                        {},
					"/fits.api.v1.ProjectService/Update":                     {},
				},
				v1.ProjectRole_PROJECT_ROLE_VIEWER: map[string]struct{}{
					"/fits.api.mvm.v1.MVMService/Get":             {},
					"/fits.api.mvm.v1.MVMService/List":            {},
					"/fits.api.pcvm.v1.PublicCloudVMService/Get":  {},
					"/fits.api.pcvm.v1.PublicCloudVMService/List": {},
					"/fits.api.v1.ProjectService/Get":             {},
				},
			},
			Resource: Resource{},
		},
		Methods: map[string]struct{}{
			"/fits.api.mvm.v1.ContactService/Create":                         {},
			"/fits.api.mvm.v1.ContactService/List":                           {},
			"/fits.api.mvm.v1.ContactService/Update":                         {},
			"/fits.api.mvm.v1.DomainService/List":                            {},
			"/fits.api.mvm.v1.LDAPService/List":                              {},
			"/fits.api.mvm.v1.LocationService/List":                          {},
			"/fits.api.mvm.v1.MVMService/AddDisks":                           {},
			"/fits.api.mvm.v1.MVMService/AddIPAddress":                       {},
			"/fits.api.mvm.v1.MVMService/AddSnapshot":                        {},
			"/fits.api.mvm.v1.MVMService/Create":                             {},
			"/fits.api.mvm.v1.MVMService/Delete":                             {},
			"/fits.api.mvm.v1.MVMService/DeleteDisk":                         {},
			"/fits.api.mvm.v1.MVMService/DeleteIPAddress":                    {},
			"/fits.api.mvm.v1.MVMService/DeleteSnapshot":                     {},
			"/fits.api.mvm.v1.MVMService/Get":                                {},
			"/fits.api.mvm.v1.MVMService/List":                               {},
			"/fits.api.mvm.v1.MVMService/Move":                               {},
			"/fits.api.mvm.v1.MVMService/MoveIPAddress":                      {},
			"/fits.api.mvm.v1.MVMService/Reinstall":                          {},
			"/fits.api.mvm.v1.MVMService/Restart":                            {},
			"/fits.api.mvm.v1.MVMService/Update":                             {},
			"/fits.api.mvm.v1.MVMService/UpdateDisk":                         {},
			"/fits.api.mvm.v1.MVMService/UpdateServiceClass":                 {},
			"/fits.api.mvm.v1.MVMService/ValidateAddDisks":                   {},
			"/fits.api.mvm.v1.MVMService/ValidateCreate":                     {},
			"/fits.api.mvm.v1.MVMService/ValidateReinstall":                  {},
			"/fits.api.mvm.v1.MVMService/ValidateUpdate":                     {},
			"/fits.api.mvm.v1.MVMService/ValidateUpdateDisk":                 {},
			"/fits.api.mvm.v1.MVMService/ValidateUpdateServiceClass":         {},
			"/fits.api.mvm.v1.OSService/List":                                {},
			"/fits.api.mvm.v1.StageService/List":                             {},
			"/fits.api.mvm.v1.VLANService/List":                              {},
			"/fits.api.pcvm.v1.PublicCloudVMService/Create":                  {},
			"/fits.api.pcvm.v1.PublicCloudVMService/Delete":                  {},
			"/fits.api.pcvm.v1.PublicCloudVMService/Get":                     {},
			"/fits.api.pcvm.v1.PublicCloudVMService/List":                    {},
			"/fits.api.pcvm.v1.PublicCloudVMService/Update":                  {},
			"/fits.api.pcvm.v1.PublicCloudVMService/ValidateCreate":          {},
			"/fits.api.pcvm.v1.PublicCloudVMService/ValidateUpdate":          {},
			"/fits.api.v1.HealthService/Get":                                 {},
			"/fits.api.v1.MethodService/List":                                {},
			"/fits.api.v1.MethodService/TokenScopedList":                     {},
			"/fits.api.v1.ProjectService/Create":                             {},
			"/fits.api.v1.ProjectService/Delete":                             {},
			"/fits.api.v1.ProjectService/Get":                                {},
			"/fits.api.v1.ProjectService/List":                               {},
			"/fits.api.v1.ProjectService/Update":                             {},
			"/fits.api.v1.TenantService/Get":                                 {},
			"/fits.api.v1.TenantService/List":                                {},
			"/fits.api.v1.TenantService/Update":                              {},
			"/fits.api.v1.TokenService/Get":                                  {},
			"/fits.api.v1.TokenService/List":                                 {},
			"/fits.api.v1.TokenService/Revoke":                               {},
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
				"/fits.api.mvm.v1.LocationService/List":      true,
				"/fits.api.mvm.v1.StageService/List":         true,
				"/fits.api.v1.MethodService/TokenScopedList": true,
				"/fits.api.v1.ProjectService/List":           true,
				"/fits.api.v1.TenantService/List":            true,
				"/fits.api.v1.TokenService/Get":              true,
				"/fits.api.v1.TokenService/List":             true,
				"/fits.api.v1.TokenService/Revoke":           true,
			},
			Operator: map[string]bool{},
			Tenant: map[string]bool{
				"/fits.api.mvm.v1.ContactService/Create": true,
				"/fits.api.mvm.v1.ContactService/List":   true,
				"/fits.api.mvm.v1.ContactService/Update": true,
				"/fits.api.mvm.v1.DomainService/List":    true,
				"/fits.api.mvm.v1.LDAPService/List":      true,
				"/fits.api.mvm.v1.OSService/List":        true,
				"/fits.api.mvm.v1.VLANService/List":      true,
				"/fits.api.v1.ProjectService/Create":     true,
				"/fits.api.v1.TenantService/Get":         true,
				"/fits.api.v1.TenantService/Update":      true,
			},
			Project: map[string]bool{
				"/fits.api.mvm.v1.MVMService/AddDisks":                   true,
				"/fits.api.mvm.v1.MVMService/AddIPAddress":               true,
				"/fits.api.mvm.v1.MVMService/AddSnapshot":                true,
				"/fits.api.mvm.v1.MVMService/Create":                     true,
				"/fits.api.mvm.v1.MVMService/Delete":                     true,
				"/fits.api.mvm.v1.MVMService/DeleteDisk":                 true,
				"/fits.api.mvm.v1.MVMService/DeleteIPAddress":            true,
				"/fits.api.mvm.v1.MVMService/DeleteSnapshot":             true,
				"/fits.api.mvm.v1.MVMService/Get":                        true,
				"/fits.api.mvm.v1.MVMService/List":                       true,
				"/fits.api.mvm.v1.MVMService/Move":                       true,
				"/fits.api.mvm.v1.MVMService/MoveIPAddress":              true,
				"/fits.api.mvm.v1.MVMService/Reinstall":                  true,
				"/fits.api.mvm.v1.MVMService/Restart":                    true,
				"/fits.api.mvm.v1.MVMService/Update":                     true,
				"/fits.api.mvm.v1.MVMService/UpdateDisk":                 true,
				"/fits.api.mvm.v1.MVMService/UpdateServiceClass":         true,
				"/fits.api.mvm.v1.MVMService/ValidateAddDisks":           true,
				"/fits.api.mvm.v1.MVMService/ValidateCreate":             true,
				"/fits.api.mvm.v1.MVMService/ValidateReinstall":          true,
				"/fits.api.mvm.v1.MVMService/ValidateUpdate":             true,
				"/fits.api.mvm.v1.MVMService/ValidateUpdateDisk":         true,
				"/fits.api.mvm.v1.MVMService/ValidateUpdateServiceClass": true,
				"/fits.api.pcvm.v1.PublicCloudVMService/Create":          true,
				"/fits.api.pcvm.v1.PublicCloudVMService/Delete":          true,
				"/fits.api.pcvm.v1.PublicCloudVMService/Get":             true,
				"/fits.api.pcvm.v1.PublicCloudVMService/List":            true,
				"/fits.api.pcvm.v1.PublicCloudVMService/Update":          true,
				"/fits.api.pcvm.v1.PublicCloudVMService/ValidateCreate":  true,
				"/fits.api.pcvm.v1.PublicCloudVMService/ValidateUpdate":  true,
				"/fits.api.v1.ProjectService/Delete":                     true,
				"/fits.api.v1.ProjectService/Get":                        true,
				"/fits.api.v1.ProjectService/Update":                     true,
			},
			Resource: map[string]bool{},
		},
		Auditable: map[string]bool{
			"/fits.api.mvm.v1.ContactService/Create":                 true,
			"/fits.api.mvm.v1.ContactService/List":                   false,
			"/fits.api.mvm.v1.ContactService/Update":                 true,
			"/fits.api.mvm.v1.DomainService/List":                    false,
			"/fits.api.mvm.v1.LDAPService/List":                      false,
			"/fits.api.mvm.v1.LocationService/List":                  false,
			"/fits.api.mvm.v1.MVMService/AddDisks":                   true,
			"/fits.api.mvm.v1.MVMService/AddIPAddress":               true,
			"/fits.api.mvm.v1.MVMService/AddSnapshot":                true,
			"/fits.api.mvm.v1.MVMService/Create":                     true,
			"/fits.api.mvm.v1.MVMService/Delete":                     true,
			"/fits.api.mvm.v1.MVMService/DeleteDisk":                 true,
			"/fits.api.mvm.v1.MVMService/DeleteIPAddress":            true,
			"/fits.api.mvm.v1.MVMService/DeleteSnapshot":             true,
			"/fits.api.mvm.v1.MVMService/Get":                        false,
			"/fits.api.mvm.v1.MVMService/List":                       false,
			"/fits.api.mvm.v1.MVMService/Move":                       true,
			"/fits.api.mvm.v1.MVMService/MoveIPAddress":              true,
			"/fits.api.mvm.v1.MVMService/Reinstall":                  true,
			"/fits.api.mvm.v1.MVMService/Restart":                    true,
			"/fits.api.mvm.v1.MVMService/Update":                     true,
			"/fits.api.mvm.v1.MVMService/UpdateDisk":                 true,
			"/fits.api.mvm.v1.MVMService/UpdateServiceClass":         true,
			"/fits.api.mvm.v1.MVMService/ValidateAddDisks":           false,
			"/fits.api.mvm.v1.MVMService/ValidateCreate":             false,
			"/fits.api.mvm.v1.MVMService/ValidateReinstall":          false,
			"/fits.api.mvm.v1.MVMService/ValidateUpdate":             false,
			"/fits.api.mvm.v1.MVMService/ValidateUpdateDisk":         false,
			"/fits.api.mvm.v1.MVMService/ValidateUpdateServiceClass": false,
			"/fits.api.mvm.v1.OSService/List":                        false,
			"/fits.api.mvm.v1.StageService/List":                     false,
			"/fits.api.mvm.v1.VLANService/List":                      false,
			"/fits.api.pcvm.v1.PublicCloudVMService/Create":          true,
			"/fits.api.pcvm.v1.PublicCloudVMService/Delete":          true,
			"/fits.api.pcvm.v1.PublicCloudVMService/Get":             false,
			"/fits.api.pcvm.v1.PublicCloudVMService/List":            false,
			"/fits.api.pcvm.v1.PublicCloudVMService/Update":          true,
			"/fits.api.pcvm.v1.PublicCloudVMService/ValidateCreate":  false,
			"/fits.api.pcvm.v1.PublicCloudVMService/ValidateUpdate":  false,
			"/fits.api.v1.HealthService/Get":                         false,
			"/fits.api.v1.MethodService/List":                        false,
			"/fits.api.v1.MethodService/TokenScopedList":             false,
			"/fits.api.v1.ProjectService/Create":                     true,
			"/fits.api.v1.ProjectService/Delete":                     true,
			"/fits.api.v1.ProjectService/Get":                        false,
			"/fits.api.v1.ProjectService/List":                       false,
			"/fits.api.v1.ProjectService/Update":                     true,
			"/fits.api.v1.TenantService/Get":                         false,
			"/fits.api.v1.TenantService/List":                        false,
			"/fits.api.v1.TenantService/Update":                      true,
			"/fits.api.v1.TokenService/Get":                          true,
			"/fits.api.v1.TokenService/List":                         true,
			"/fits.api.v1.TokenService/Revoke":                       true,
			"/fits.api.v1.VersionService/Get":                        false,
		},
		Resources: map[string]v1.ResourceType{},
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

// GetResourceType returns the resource a resource scoped method acts on.
func GetResourceType(req connect.AnyRequest) (v1.ResourceType, bool) {
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

// GetResourceFromRequest returns the resource instance id from a resource scoped request.
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
