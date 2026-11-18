package validation

import (
	"strings"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

var testTime = time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)

func TestValidateApiKeyServiceCreateRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.ApiKeyServiceCreateRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name: "name too short",
			msg: &v1.ApiKeyServiceCreateRequest{
				Name: "k",
			},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
		{
			name: "valid without expiry",
			msg: &v1.ApiKeyServiceCreateRequest{
				Name: "my key",
			},
			wantErr: false,
		},
		{
			name: "valid with expiry",
			msg: &v1.ApiKeyServiceCreateRequest{
				Name:      "my key",
				ExpiresAt: timestamppb.New(testTime),
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateApiKey(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &v1.ApiKey{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - id: value is empty, which is not a valid UUID
 - user_id: value is empty, which is not a valid UUID
 - realm: must be within 2 and 128 characters
 - name: must be within 2 and 128 characters
 - hash: does not match regex pattern ` + "`" + `^[0-9a-f]{64}$` + "`" + `
 - prefix: must be 12 characters`,
		},
		{
			name:    "valid",
			msg:     validApiKey(),
			wantErr: false,
		},
		{
			name: "invalid hash",
			msg: func() *v1.ApiKey {
				key := validApiKey()
				key.Hash = strings.Repeat("g", 64)
				return key
			}(),
			wantErr:          true,
			wantErrorMessage: "validation error: hash: does not match regex pattern `^[0-9a-f]{64}$`",
		},
		{
			name: "short prefix",
			msg: func() *v1.ApiKey {
				key := validApiKey()
				key.Prefix = "fco_v1_ab"
				return key
			}(),
			wantErr:          true,
			wantErrorMessage: `validation error: prefix: must be 12 characters`,
		},
		{
			name: "revoked key",
			msg: func() *v1.ApiKey {
				key := validApiKey()
				key.RevokedAt = timestamppb.New(testTime)
				key.LastUsedAt = timestamppb.New(testTime)
				return key
			}(),
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateApiKeyServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.ApiKeyServiceGetRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: id: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.ApiKeyServiceGetRequest{Id: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateApiKeyServiceRevokeRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.ApiKeyServiceRevokeRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: id: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.ApiKeyServiceRevokeRequest{Id: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateApiKeyServiceDeleteRequest(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &v1.ApiKeyServiceDeleteRequest{},
			wantErr:          true,
			wantErrorMessage: `validation error: id: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &v1.ApiKeyServiceDeleteRequest{Id: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func validApiKey() *v1.ApiKey {
	return &v1.ApiKey{
		Id:        testUUID,
		UserId:    testUUID,
		Realm:     "fi-ts",
		Name:      "my key",
		Hash:      strings.Repeat("a", 64),
		Prefix:    "fco_v1_abc12",
		CreatedAt: timestamppb.New(testTime),
	}
}
