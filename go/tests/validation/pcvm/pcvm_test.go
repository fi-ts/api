package pcvm

import (
	"testing"

	pcvmv1 "github.com/fi-ts/api/go/fits/api/pcvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidatePublicCloudVMServiceGetRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &pcvmv1.PublicCloudVMServiceGetRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &pcvmv1.PublicCloudVMServiceGetRequest{Uuid: validation.TestUUID, ProjectUuid: validation.TestUUID},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceCreateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &pcvmv1.PublicCloudVMServiceCreateRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - project_uuid: value is empty, which is not a valid UUID
 - name: must be within 2 and 128 characters`,
		},
		{
			Name: "valid",
			Msg: &pcvmv1.PublicCloudVMServiceCreateRequest{
				ProjectUuid: validation.TestUUID,
				Name:        "my-vm",
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceUpdateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &pcvmv1.PublicCloudVMServiceUpdateRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID
 - update_meta: value is required`,
		},
		{
			Name: "valid",
			Msg: &pcvmv1.PublicCloudVMServiceUpdateRequest{
				Uuid:        validation.TestUUID,
				ProjectUuid: validation.TestUUID,
				UpdateMeta:  &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceValidateCreateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "empty",
			Msg:              &pcvmv1.PublicCloudVMServiceValidateCreateRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: create: value is required`,
		},
		{
			Name: "valid",
			Msg: &pcvmv1.PublicCloudVMServiceValidateCreateRequest{
				Create: &pcvmv1.PublicCloudVMServiceCreateRequest{
					ProjectUuid: validation.TestUUID,
					Name:        "my-vm",
				},
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidatePublicCloudVMServiceValidateUpdateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "empty",
			Msg:              &pcvmv1.PublicCloudVMServiceValidateUpdateRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: update: value is required`,
		},
		{
			Name: "valid",
			Msg: &pcvmv1.PublicCloudVMServiceValidateUpdateRequest{
				Update: &pcvmv1.PublicCloudVMServiceUpdateRequest{
					Uuid:        validation.TestUUID,
					ProjectUuid: validation.TestUUID,
					UpdateMeta:  &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
				},
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
