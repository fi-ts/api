package validation

import (
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

func TestValidateTenant(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty",
			Msg:              &v1.Tenant{},
			WantErr:          true,
			WantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			Name:    "valid",
			Msg:     &v1.Tenant{Name: "tenant", Description: "a tenant"},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateTenantServiceListRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "invalid name",
			Msg:              &v1.TenantServiceListRequest{Name: proto.String("a")},
			WantErr:          true,
			WantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			Name:    "valid",
			Msg:     &v1.TenantServiceListRequest{Name: proto.String("tenant")},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateTenantServiceUpdateRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "missing update_meta",
			Msg:              &v1.TenantServiceUpdateRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: update_meta: value is required`,
		},
		{
			Name: "invalid email",
			Msg: &v1.TenantServiceUpdateRequest{
				Email: proto.String("not-an-email"),
			},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - update_meta: value is required
 - email: must be a valid email address`,
		},
		{
			Name: "valid",
			Msg: &v1.TenantServiceUpdateRequest{
				UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
				Email:      proto.String("user@example.com"),
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}
