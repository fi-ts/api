package validation

import (
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateTokenServiceGetRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty",
			Msg:              &v1.TokenServiceGetRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &v1.TokenServiceGetRequest{Uuid: TestUUID},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateTokenServiceRevokeRequest(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty",
			Msg:              &v1.TokenServiceRevokeRequest{},
			WantErr:          true,
			WantErrorMessage: `validation error: uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &v1.TokenServiceRevokeRequest{Uuid: TestUUID},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateToken(t *testing.T) {
	tests := Prototests{
		{
			Name: "user token with permissions",
			Msg: &v1.Token{
				TokenType:   v1.TokenType_TOKEN_TYPE_USER,
				Permissions: []*v1.PermissionsByVisibility{{}},
			},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - token type user must not have permissions
 - uuid: value is empty, which is not a valid UUID
 - user: must be at least 2 characters`,
		},
		{
			Name: "valid user token",
			Msg: &v1.Token{
				Uuid:      TestUUID,
				User:      "user",
				TokenType: v1.TokenType_TOKEN_TYPE_USER,
			},
			WantErr: false,
		},
		{
			Name: "valid api token with permissions",
			Msg: &v1.Token{
				Uuid:        TestUUID,
				User:        "user",
				TokenType:   v1.TokenType_TOKEN_TYPE_API,
				Permissions: []*v1.PermissionsByVisibility{{}},
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateProjectPermissions(t *testing.T) {
	tests := Prototests{
		{
			Name:             "invalid project",
			Msg:              &v1.ProjectPermissions{Project: "not-uuid"},
			WantErr:          true,
			WantErrorMessage: `validation error: project: must be '*' or a uuid`,
		},
		{
			Name:    "wildcard",
			Msg:     &v1.ProjectPermissions{Project: "*"},
			WantErr: false,
		},
		{
			Name:    "uuid",
			Msg:     &v1.ProjectPermissions{Project: TestUUID},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateTenantPermissions(t *testing.T) {
	tests := Prototests{
		{
			Name:             "invalid login",
			Msg:              &v1.TenantPermissions{Login: "a"},
			WantErr:          true,
			WantErrorMessage: `validation error: login: must be '*' or within 2 and 128 characters`,
		},
		{
			Name:    "wildcard",
			Msg:     &v1.TenantPermissions{Login: "*"},
			WantErr: false,
		},
		{
			Name:    "login",
			Msg:     &v1.TenantPermissions{Login: "tenant"},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateMethodPermission(t *testing.T) {
	tests := Prototests{
		{
			Name:             "subject too long",
			Msg:              &v1.MethodPermission{Subject: strings.Repeat("x", 257)},
			WantErr:          true,
			WantErrorMessage: `validation error: subject: must be at most 256 characters`,
		},
		{
			Name:    "valid",
			Msg:     &v1.MethodPermission{Subject: "project", Methods: []string{"List"}},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}
