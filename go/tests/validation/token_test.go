package validation

import (
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateTokenServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.TokenServiceGetRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.TokenServiceGetRequest{Uuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateTokenServiceRevokeRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.TokenServiceRevokeRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.TokenServiceRevokeRequest{Uuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateToken(t *testing.T) {
	tests := prototests{
		{
			name: "user token with permissions",
			msg: &v1.Token{
				TokenType:   v1.TokenType_TOKEN_TYPE_USER,
				Permissions: []*v1.PermissionsByVisibility{{}},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - token type user must not have permissions
 - uuid: value is empty, which is not a valid UUID
 - user: must be at least 2 characters`,
		},
		{
			name: "valid user token",
			msg: &v1.Token{
				Uuid:      testUUID,
				User:      "user",
				TokenType: v1.TokenType_TOKEN_TYPE_USER,
			},
			wantErr: false,
		},
		{
			name: "valid api token with permissions",
			msg: &v1.Token{
				Uuid:        testUUID,
				User:        "user",
				TokenType:   v1.TokenType_TOKEN_TYPE_API,
				Permissions: []*v1.PermissionsByVisibility{{}},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateProjectPermissions(t *testing.T) {
	tests := prototests{
		{
			name:             "invalid project",
			msg:              &v1.ProjectPermissions{Project: "not-uuid"},
			wantErr:          true,
			wantErrorMessage: `validation error: project: must be '*' or a uuid`,
		},
		{
			name:    "wildcard",
			msg:     &v1.ProjectPermissions{Project: "*"},
			wantErr: false,
		},
		{
			name:    "uuid",
			msg:     &v1.ProjectPermissions{Project: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateTenantPermissions(t *testing.T) {
	tests := prototests{
		{
			name:             "invalid login",
			msg:              &v1.TenantPermissions{Login: "a"},
			wantErr:          true,
			wantErrorMessage: `validation error: login: must be '*' or within 2 and 128 characters`,
		},
		{
			name:    "wildcard",
			msg:     &v1.TenantPermissions{Login: "*"},
			wantErr: false,
		},
		{
			name:    "login",
			msg:     &v1.TenantPermissions{Login: "tenant"},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMethodPermission(t *testing.T) {
	tests := prototests{
		{
			name:             "subject too long",
			msg:              &v1.MethodPermission{Subject: strings.Repeat("x", 257)},
			wantErr:          true,
			wantErrorMessage: `validation error: subject: must be at most 256 characters`,
		},
		{
			name:    "valid",
			msg:     &v1.MethodPermission{Subject: "project", Methods: []string{"List"}},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
