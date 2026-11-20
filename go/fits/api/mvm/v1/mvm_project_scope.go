package mvmv1

// This file is HAND-WRITTEN and must NOT be deleted by code generation.
// It lives next to the generated *.pb.go files on purpose: the methods below
// have to be defined in the same Go package as the types they extend.
//
// Background: the authorization layer resolves the project of a request via a
// type assertion against interface{ GetProjectUuid() string } on the top-level
// request message (see fits/api/permissions.GetProjectFromRequest). The
// Validate* requests only wrap another request and therefore do not carry a
// top-level project_uuid field, so the generated code does not provide a
// GetProjectUuid method on them.
//
// GetProjectUuid returns the project of the wrapped create request so the
// request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateCreateRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetCreate().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped add-disk request so the
// request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateAddDisksRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetAddDisks().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped update-disk request so the
// request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateUpdateDiskRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetUpdateDisk().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped reinstall request so the
// request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateReinstallRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetReinstall().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped update request so the
// request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateUpdateRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetUpdate().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped update-service-class
// request so the request satisfies the project-scope resolution used by authorization.
func (x *MVMServiceValidateUpdateServiceClassRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetUpdateServiceClass().GetProjectUuid()
}

// The OnBehalf* requests wrap a public request and therefore do not carry a
// top-level project_uuid field either; the same forwarding applies.

// GetProjectUuid returns the project of the wrapped get request so the
// request satisfies the project-scope resolution used by authorization.
func (x *OnBehalfMVMServiceGetRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetRequest().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped list request so the
// request satisfies the project-scope resolution used by authorization.
func (x *OnBehalfMVMServiceListRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetRequest().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped create request so the
// request satisfies the project-scope resolution used by authorization.
func (x *OnBehalfMVMServiceCreateRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetRequest().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped delete request so the
// request satisfies the project-scope resolution used by authorization.
func (x *OnBehalfMVMServiceDeleteRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetRequest().GetProjectUuid()
}
