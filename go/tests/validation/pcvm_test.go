package validation

import (
	"testing"

	pcvmv1 "github.com/fi-ts/api/go/fits/api/pcvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidatePublicCloudVMServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &pcvmv1.PublicCloudVMServiceGetRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &pcvmv1.PublicCloudVMServiceGetRequest{Uuid: testUUID, ProjectUuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceCreateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &pcvmv1.PublicCloudVMServiceCreateRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - project_uuid: value is empty, which is not a valid UUID
 - name: must be within 2 and 128 characters`,
		},
		{
			name: "valid",
			msg: &pcvmv1.PublicCloudVMServiceCreateRequest{
				ProjectUuid: testUUID,
				Name:        "my-vm",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceUpdateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &pcvmv1.PublicCloudVMServiceUpdateRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID
 - update_meta: value is required`,
		},
		{
			name: "valid",
			msg: &pcvmv1.PublicCloudVMServiceUpdateRequest{
				Uuid:        testUUID,
				ProjectUuid: testUUID,
				UpdateMeta:  &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceValidateCreateRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &pcvmv1.PublicCloudVMServiceValidateCreateRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: create: value is required`,
		},
		{
			name: "valid",
			msg: &pcvmv1.PublicCloudVMServiceValidateCreateRequest{
				Create: &pcvmv1.PublicCloudVMServiceCreateRequest{
					ProjectUuid: testUUID,
					Name:        "my-vm",
				},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceValidateUpdateRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &pcvmv1.PublicCloudVMServiceValidateUpdateRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: update: value is required`,
		},
		{
			name: "valid",
			msg: &pcvmv1.PublicCloudVMServiceValidateUpdateRequest{
				Update: &pcvmv1.PublicCloudVMServiceUpdateRequest{
					Uuid:        testUUID,
					ProjectUuid: testUUID,
					UpdateMeta:  &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
				},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
