package validation

import (
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateVersion(t *testing.T) {
	tests := Prototests{
		{
			Name:             "version too long",
			Msg:              &v1.Version{Version: strings.Repeat("x", 257)},
			WantErr:          true,
			WantErrorMessage: `validation error: version: must be shorter than 256 characters`,
		},
		{
			Name: "valid",
			Msg: &v1.Version{
				Version:   "1.0.0",
				Revision:  "abc123",
				GitSha1:   "deadbeef",
				BuildDate: "2026-01-01T00:00:00Z",
			},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}
