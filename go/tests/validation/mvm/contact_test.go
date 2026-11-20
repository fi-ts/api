package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/fi-ts/api/go/tests/validation"
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

	tests := validation.Prototests{
		{
			Name:    "valid",
			Msg:     request("+4915123456789"),
			WantErr: false,
		},
		{
			Name:             "missing phone",
			Msg:              request(""),
			WantErr:          true,
			WantErrorMessage: "validation error: phone: " + msgPhone,
		},
		{
			Name:             "phone without country code",
			Msg:              request("015123456789"),
			WantErr:          true,
			WantErrorMessage: "validation error: phone: " + msgPhone,
		},
		{
			Name:             "phone with whitespace",
			Msg:              request("+49 151 23456789"),
			WantErr:          true,
			WantErrorMessage: "validation error: phone: " + msgPhone,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateContactServiceUpdateRequest(t *testing.T) {
	request := func(phone string) *mvmv1.ContactServiceUpdateRequest {
		return &mvmv1.ContactServiceUpdateRequest{
			Tenant:     "tenant",
			Uuid:       validation.TestUUID,
			UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			FirstName:  "Jane",
			LastName:   "Doe",
			Department: "IT",
			Phone:      phone,
			Email:      "jane.doe@example.com",
		}
	}

	tests := validation.Prototests{
		{
			Name:    "valid",
			Msg:     request("+4915123456789"),
			WantErr: false,
		},
		{
			Name:             "invalid phone",
			Msg:              request("0151-23456789"),
			WantErr:          true,
			WantErrorMessage: "validation error: phone: " + msgPhone,
		},
	}

	validation.ValidateProtos(t, tests)
}
