package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateLocation(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.Location{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - title: must be within 2 and 128 characters`,
		},
		{
			Name:             "invalid uuid",
			Msg:              &mvmv1.Location{Uuid: "not-a-uuid", Title: "Nürnberg (RZ1)"},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: must be a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.Location{Uuid: validation.TestUUID, Title: "Nürnberg (RZ1)"},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateLocationServiceListRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.LocationServiceListRequest{},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
