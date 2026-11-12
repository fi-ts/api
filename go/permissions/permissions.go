package permissions

import (
	_ "embed"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

type (
	ServicePermissions struct {
		Roles      Roles      `json:"roles"`
		Methods    Methods    `json:"methods"`
		Visibility Visibility `json:"visibility"`
		Auditable  Auditable  `json:"auditable,omitempty"`
		Services   []string   `json:"services,omitempty"`
		Resources  Resources  `json:"products,omitempty"`
	}

	Methods map[string]struct{}

	Auditable map[string]bool

	Operator  map[v1.OperatorRole]Methods
	Tenant    map[v1.TenantRole]Methods
	Project   map[v1.ProjectRole]Methods
	Resource  map[v1.ResourceRole]Methods
	Resources map[string]v1.Resource

	// Roles
	Roles struct {
		Operator Operator `json:"operator,omitempty"`
		Tenant   Tenant   `json:"tenant,omitempty"`
		Project  Project  `json:"project,omitempty"`
		Resource Resource `json:"resource,omitempty"`
	}

	Visibility struct {
		Public   map[string]bool `json:"public,omitempty"`
		Self     map[string]bool `json:"self,omitempty"`
		Operator map[string]bool `json:"operator,omitempty"`
		Tenant   map[string]bool `json:"tenant,omitempty"`
		Project  map[string]bool `json:"project,omitempty"`
		Resource map[string]bool `json:"resource,omitempty"`
	}
)
