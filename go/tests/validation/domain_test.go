package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
)

func TestValidateDomain(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.Domain{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - fqdn: must be within 2 and 128 characters
 - tenant: must be within 2 and 128 characters`,
		},
		{
			name: "valid",
			msg: &mvmv1.Domain{
				Uuid:   testUUID,
				Fqdn:   "domain.example.com",
				Tenant: "tenant",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateDomainServiceListRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.DomainServiceListRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - tenant: must be within 2 and 128 characters
 - location_uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:             "invalid location uuid",
			msg:              &mvmv1.DomainServiceListRequest{Tenant: "tenant", LocationUuid: "not-a-uuid"},
			wantErr:          true,
			wantErrorMessage: `validation error: location_uuid: must be a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &mvmv1.DomainServiceListRequest{Tenant: "tenant", LocationUuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
