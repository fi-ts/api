package validation

import (
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

func TestValidateProject(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty",
			Msg:              &v1.Project{},
			WantErr:          true,
			WantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			Name: "invalid name",
			Msg: &v1.Project{
				Name:        "a",
				Description: "a project",
			},
			WantErr:          true,
			WantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			Name: "valid",
			Msg: &v1.Project{
				Name:        "project",
				Description: "a project",
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateProjectServiceGetRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty",
			Msg:              &v1.ProjectServiceGetRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: project: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &v1.ProjectServiceGetRequest{Project: TestUUID},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateProjectServiceListRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "invalid name",
			Msg:              &v1.ProjectServiceListRequest{Name: proto.String("a")},
			WantErr:          true,
			WantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			Name:    "valid",
			Msg:     &v1.ProjectServiceListRequest{Name: proto.String("project")},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateProjectServiceCreateRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:    "empty",
			Msg:     &v1.ProjectServiceCreateRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - tenant: must be within 2 and 128 characters
 - name: must be within 2 and 128 characters`,
		},
		{
			Name: "valid",
			Msg: &v1.ProjectServiceCreateRequest{
				Tenant:      "tenant",
				Name:        "project",
				Description: "a project",
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateProjectServiceUpdateRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:    "empty",
			Msg:     &v1.ProjectServiceUpdateRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - project: value is empty, which is not a valid UUID
 - update_meta: value is required`,
		},
		{
			Name: "valid",
			Msg: &v1.ProjectServiceUpdateRequest{
				Project:    TestUUID,
				UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}
