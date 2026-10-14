package validation

import (
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

func TestValidateTenant(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.Tenant{},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name:    "valid",
			msg:     &v1.Tenant{Name: "tenant", Description: "a tenant"},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateTenantServiceListRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "invalid name",
			msg:              &v1.TenantServiceListRequest{Name: proto.String("a")},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name:    "valid",
			msg:     &v1.TenantServiceListRequest{Name: proto.String("tenant")},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateTenantServiceUpdateRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "missing update_meta",
			msg:              &v1.TenantServiceUpdateRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: update_meta: value is required`,
		},
		{
			name: "invalid email",
			msg: &v1.TenantServiceUpdateRequest{
				Email: proto.String("not-an-email"),
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - update_meta: value is required
 - email: must be a valid email address`,
		},
		{
			name: "valid",
			msg: &v1.TenantServiceUpdateRequest{
				UpdateMeta: &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
				Email:      proto.String("user@example.com"),
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
