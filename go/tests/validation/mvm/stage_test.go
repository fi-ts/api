package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateStage(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.Stage{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - title: must be within 2 and 128 characters`,
		},
		{
			Name:             "invalid uuid",
			Msg:              &mvmv1.Stage{Uuid: "not-a-uuid", Title: "production"},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: must be a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.Stage{Uuid: validation.TestUUID, Title: "production"},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateStageServiceListRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.StageServiceListRequest{},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
