package pcvmv1

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

// GetProjectUuid returns the project of the wrapped create request so the
// request satisfies the project-scope resolution used by authorization.
func (x *PublicCloudVMServiceValidateCreateRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetCreate().GetProjectUuid()
}

// GetProjectUuid returns the project of the wrapped update request so the
// request satisfies the project-scope resolution used by authorization.
func (x *PublicCloudVMServiceValidateUpdateRequest) GetProjectUuid() string {
	if x == nil {
		return ""
	}
	return x.GetUpdate().GetProjectUuid()
}
