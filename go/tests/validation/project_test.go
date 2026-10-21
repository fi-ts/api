package validation

import (
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

const testUUID = "e266fcc6-f6de-4ee1-ba26-baa17bf47b13"

func TestValidateProject(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.Project{},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name: "invalid name",
			msg: &v1.Project{
				Name:        "a",
				Description: "a project",
			},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name: "valid",
			msg: &v1.Project{
				Name:        "project",
				Description: "a project",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateProjectServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.ProjectServiceGetRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: project: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.ProjectServiceGetRequest{Project: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateProjectServiceListRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "invalid name",
			msg:              &v1.ProjectServiceListRequest{Name: proto.String("a")},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name:    "valid",
			msg:     &v1.ProjectServiceListRequest{Name: proto.String("project")},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateProjectServiceCreateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &v1.ProjectServiceCreateRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - tenant: must be within 2 and 128 characters
 - name: must be within 2 and 128 characters`,
		},
		{
			name: "valid",
			msg: &v1.ProjectServiceCreateRequest{
				Tenant:      "tenant",
				Name:        "project",
				Description: "a project",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateProjectServiceUpdateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &v1.ProjectServiceUpdateRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - project: value is empty, which is not a valid UUID
 - update_meta: value is required`,
		},
		{
			name: "valid",
			msg: &v1.ProjectServiceUpdateRequest{
				Project:    testUUID,
				UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
