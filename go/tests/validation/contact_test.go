package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

const msgPhone = "must be a valid phone number in E.164 format, e.g. +4915123456789"

func TestValidateContactServiceCreateRequest(t *testing.T) {
	request := func(phone string) *mvmv1.ContactServiceCreateRequest {
		return &mvmv1.ContactServiceCreateRequest{
			Tenant:     "tenant",
			FirstName:  "Jane",
			LastName:   "Doe",
			Department: "IT",
			Phone:      phone,
			Email:      "jane.doe@example.com",
		}
	}

	tests := prototests{
		{
			name:    "valid",
			msg:     request("+4915123456789"),
			wantErr: false,
		},
		{
			name:             "missing phone",
			msg:              request(""),
			wantErr:          true,
			wantErrorMessage: "validation error: phone: " + msgPhone,
		},
		{
			name:             "phone without country code",
			msg:              request("015123456789"),
			wantErr:          true,
			wantErrorMessage: "validation error: phone: " + msgPhone,
		},
		{
			name:             "phone with whitespace",
			msg:              request("+49 151 23456789"),
			wantErr:          true,
			wantErrorMessage: "validation error: phone: " + msgPhone,
		},
	}

	validateProtos(t, tests)
}

func TestValidateContactServiceUpdateRequest(t *testing.T) {
	request := func(phone string) *mvmv1.ContactServiceUpdateRequest {
		return &mvmv1.ContactServiceUpdateRequest{
			Tenant:     "tenant",
			Uuid:       testUUID,
			UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			FirstName:  "Jane",
			LastName:   "Doe",
			Department: "IT",
			Phone:      phone,
			Email:      "jane.doe@example.com",
		}
	}

	tests := prototests{
		{
			name:    "valid",
			msg:     request("+4915123456789"),
			wantErr: false,
		},
		{
			name:             "invalid phone",
			msg:              request("0151-23456789"),
			wantErr:          true,
			wantErrorMessage: "validation error: phone: " + msgPhone,
		},
	}

	validateProtos(t, tests)
}
