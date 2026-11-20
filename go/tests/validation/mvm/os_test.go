package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateOS(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.OS{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - title: must be within 2 and 128 characters`,
		},
		{
			Name:             "invalid uuid",
			Msg:              &mvmv1.OS{Uuid: "not-a-uuid", Title: "Ubuntu 24.04", Type: mvmv1.OSType_OS_TYPE_LINUX},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: must be a valid UUID`,
		},
		{
			Name:    "valid linux",
			Msg:     &mvmv1.OS{Uuid: validation.TestUUID, Title: "Ubuntu 24.04", Type: mvmv1.OSType_OS_TYPE_LINUX},
			WantErr: false,
		},
		{
			Name:    "valid windows",
			Msg:     &mvmv1.OS{Uuid: validation.TestUUID, Title: "Windows Server 2022", Type: mvmv1.OSType_OS_TYPE_WINDOWS},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateOSServiceListRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "empty",
			Msg:              &mvmv1.OSServiceListRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: tenant: must be within 2 and 128 characters`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.OSServiceListRequest{Tenant: "tenant"},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
