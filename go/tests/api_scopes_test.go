package apitests

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/fi-ts/api/go/tests/protoparser"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/descriptorpb"
)

type (
	tenant     struct{}
	project    struct{}
	admin      struct{}
	visibility struct{}
)

func (tenant) Get(methodOpts []*descriptorpb.UninterpretedOption) (scopes []string) {
	scopes = getScopes(methodOpts, []string{
		v1.TenantRole_TENANT_ROLE_OWNER.String(),
		v1.TenantRole_TENANT_ROLE_EDITOR.String(),
		v1.TenantRole_TENANT_ROLE_VIEWER.String(),
	})
	return
}

func (project) Get(methodOpts []*descriptorpb.UninterpretedOption) (scopes []string) {
	scopes = getScopes(methodOpts, []string{
		v1.ProjectRole_PROJECT_ROLE_OWNER.String(),
		v1.ProjectRole_PROJECT_ROLE_EDITOR.String(),
		v1.ProjectRole_PROJECT_ROLE_VIEWER.String(),
	})
	return
}

func (admin) Get(methodOpts []*descriptorpb.UninterpretedOption) (scopes []string) {
	scopes = getScopes(methodOpts, []string{
		v1.AdminRole_ADMIN_ROLE_EDITOR.String(),
		v1.AdminRole_ADMIN_ROLE_VIEWER.String(),
	})
	return
}

func (visibility) Get(methodOpts []*descriptorpb.UninterpretedOption) (scopes []string) {
	scopes = getScopes(methodOpts, []string{
		v1.Visibility_VISIBILITY_PUBLIC.String(),
		v1.Visibility_VISIBILITY_SELF.String(),
	})
	return
}

func getScopes(methodOpts []*descriptorpb.UninterpretedOption, identifiers []string) (scopes []string) {
	for _, methodOpt := range methodOpts {
		for _, namePart := range methodOpt.GetName() {
			if !namePart.GetIsExtension() {
				continue
			}
			if slices.Contains(identifiers, methodOpt.GetIdentifierValue()) {
				scopes = append(scopes, methodOpt.GetIdentifierValue())
			}
		}
	}
	return
}

func getProtos(root string) ([]string, error) {
	var (
		walk = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func Test_APIScopes(t *testing.T) {
	err := validateProto("../../proto")
	require.NoError(t, err)
	// err = validateProto("./testproto")

	// errs := errors.Join(
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Add\" has apiv1.TenantRole but request payload \"WrongProjectServiceAddRequest\" does not have a login field"),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Get\" has apiv1.ProjectRole but request payload \"WrongProjectServiceGetRequest\" does not have a project field"),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/List\" has no scope defined. one scope needs to be defined though. use one of the following scopes: [apiv1.AdminRole apiv1.InfraRole apiv1.MachineRole apiv1.ProjectRole apiv1.TenantRole apiv1.Visibility]"),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Update\" does not have a update_meta field in WrongProjectServiceUpdateRequest"),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Update\" can not have apiv1.AdminRole ([ADMIN_ROLE_VIEWER]) and apiv1.ProjectRole ([PROJECT_ROLE_OWNER]) at the same time. only one scope is allowed."),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Delete\" can not have apiv1.AdminRole ([ADMIN_ROLE_VIEWER]) and apiv1.Visibility ([VISIBILITY_PUBLIC]) at the same time. only one scope is allowed."),
	// 	errors.New("api service method: \"/fits.api.v1.WrongProjectService/Charge\" has no scope defined. one scope needs to be defined though. use one of the following scopes: [apiv1.AdminRole apiv1.InfraRole apiv1.MachineRole apiv1.ProjectRole apiv1.TenantRole apiv1.Visibility]"),
	// )

	// require.Equal(t, err, errs)
}

func Test_FieldNumbering(t *testing.T) {
	files, err := getProtos("../../proto")
	require.NoError(t, err)
	var errs []error

	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		require.NoError(t, err)

		for _, mt := range fd.GetMessageType() {
			var (
				firstField = true
				lastNumber int32
			)

			for _, field := range mt.GetField() {
				if field.Number != nil {
					if firstField {
						// Enable if the first field must start with 1
						// if *field.Number != 1 {
						// 	errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *mt.Name, *field.Name, 1, *field.Number))
						// }
						firstField = false
					} else {
						if lastNumber+1 != *field.Number {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *mt.Name, *field.Name, lastNumber+1, *field.Number))
						}
					}
					lastNumber = *field.Number
				}
			}
		}

		for _, et := range fd.GetEnumType() {
			var (
				firstField = true
				lastNumber int32
			)

			for _, value := range et.GetValue() {
				if value.Number != nil {
					if firstField {
						if *value.Number != 0 {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *et.Name, *value.Name, 0, *value.Number))
						}
						firstField = false
					} else {
						if lastNumber+1 != *value.Number {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *et.Name, *value.Name, lastNumber+1, *value.Number))
						}
					}
					lastNumber = *value.Number
				}
			}
		}
	}
	require.NoError(t, errors.Join(errs...))
}

func validateProto(root string) error {
	var (
		tr v1.TenantRole
		pr v1.ProjectRole
		ar v1.AdminRole
		vr v1.Visibility

		trs = fmt.Sprintf("%T", tr)
		prs = fmt.Sprintf("%T", pr)
		ars = fmt.Sprintf("%T", ar)
		vrs = fmt.Sprintf("%T", vr)

		// add all *rs from above here
		scopeKeys = []string{
			trs, prs, ars, vrs,
		}
	)
	slices.Sort(scopeKeys)
	files, err := getProtos(root)
	if err != nil {
		return err
	}

	var errs []error
	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return err
		}
		for _, serviceDesc := range fd.GetService() {
			for _, method := range serviceDesc.GetMethod() {
				var (
					methodName  = fmt.Sprintf("/%s.%s/%s", fd.GetPackage(), serviceDesc.GetName(), method.GetName())
					methodOpts  = method.GetOptions().GetUninterpretedOption()
					methodScope = ""
					scopes      = map[string][]string{
						trs: tenant{}.Get(methodOpts),
						prs: project{}.Get(methodOpts),
						ars: admin{}.Get(methodOpts),
						vrs: visibility{}.Get(methodOpts),
					}
					allScopeNames = func() (names []string) {
						for name := range scopes {
							names = append(names, name)
						}
						return
					}()
				)

				var auditingSpecified bool
				for _, methodOpt := range methodOpts {
					if *methodOpt.IdentifierValue == v1.Auditing_AUDITING_EXCLUDED.String() || *methodOpt.IdentifierValue == v1.Auditing_AUDITING_INCLUDED.String() {
						auditingSpecified = true
					}
				}
				if !auditingSpecified {
					errs = append(errs, fmt.Errorf("api service method: %q does not have auditing specified", methodName))
				}

				// Sort all to have stable results
				slices.Sort(allScopeNames)

				for _, mt := range fd.GetMessageType() {
					if mt.GetName() != method.GetInputType() {
						continue
					}
					var (
						updateRequest string
					)
					if strings.Contains(mt.GetName(), "UpdateRequest") {
						var (
							hasUpdatedAtField bool
						)
						for _, field := range mt.GetField() {
							if field.GetName() == "update_meta" {
								hasUpdatedAtField = true
							}
						}
						updateRequest = mt.GetName()
						if !hasUpdatedAtField {
							errs = append(errs, fmt.Errorf("api service method: %q does not have a update_meta field in %s", methodName, updateRequest))
						}
					}
				}

				for _, name := range scopeKeys {
					s := scopes[name]
					if len(s) == 0 {
						continue
					}
					if methodScope != "" {
						errs = append(errs, fmt.Errorf("api service method: %q can not have %s and %s (%s) at the same time. only one scope is allowed.", methodName, methodScope, name, s))
					}
					methodScope = fmt.Sprintf("%s (%s)", name, s)

					if name == prs {
						projectFound := false
						projectRequest := ""
						for _, mt := range fd.GetMessageType() {
							if mt.GetName() != method.GetInputType() {
								continue
							}
							for _, field := range mt.GetField() {
								if field.GetName() == "project" {
									projectFound = true
								}
							}
							projectRequest = mt.GetName()
						}
						if !projectFound {
							errs = append(errs, fmt.Errorf("api service method: %q has %s but request payload %q does not have a project field", methodName, prs, projectRequest))
						}
					}
					if name == trs {
						tenantFound := false
						tenantRequest := ""
						for _, mt := range fd.GetMessageType() {
							if mt.GetName() != method.GetInputType() {
								continue
							}
							for _, field := range mt.GetField() {
								if field.GetName() == "login" {
									tenantFound = true
								}
							}
							tenantRequest = mt.GetName()
						}
						if !tenantFound {
							errs = append(errs, fmt.Errorf("api service method: %q has %s but request payload %q does not have a login field", methodName, trs, tenantRequest))
						}
					}
				}

				if methodScope == "" {
					errs = append(errs, fmt.Errorf("api service method: %q has no scope defined. one scope needs to be defined though. use one of the following scopes: %s", methodName, allScopeNames))
				}
			}
		}
	}
	return errors.Join(errs...)
}
