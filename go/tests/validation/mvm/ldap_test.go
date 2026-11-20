package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateLDAP(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.LDAP{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - fqdn: must be within 2 and 128 characters
 - tenant: must be within 2 and 128 characters`,
		},
		{
			Name:             "invalid uuid",
			Msg:              &mvmv1.LDAP{Uuid: "not-a-uuid", Fqdn: "ldap.example.com", Tenant: "tenant"},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: must be a valid UUID`,
		},
		{
			Name: "valid",
			Msg: &mvmv1.LDAP{
				Uuid:   validation.TestUUID,
				Fqdn:   "ldap.example.com",
				Tenant: "tenant",
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateLDAPServiceListRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "empty",
			Msg:              &mvmv1.LDAPServiceListRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: tenant: must be within 2 and 128 characters`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.LDAPServiceListRequest{Tenant: "tenant"},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}
