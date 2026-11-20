package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateDomain(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.Domain{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - fqdn: must be within 2 and 128 characters
 - tenant: must be within 2 and 128 characters`,
		},
		{
			Name: "valid",
			Msg: &mvmv1.Domain{
				Uuid:   validation.TestUUID,
				Fqdn:   "domain.example.com",
				Tenant: "tenant",
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateDomainServiceListRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.DomainServiceListRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - tenant: must be within 2 and 128 characters
 - location_uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:             "invalid location uuid",
			Msg:              &mvmv1.DomainServiceListRequest{Tenant: "tenant", LocationUuid: "not-a-uuid"},
			WantErr:          true,
			WantErrorMessage: `validation error: location_uuid: must be a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.DomainServiceListRequest{Tenant: "tenant", LocationUuid: validation.TestUUID},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
