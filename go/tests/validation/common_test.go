package validation

import (
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateLabels(t *testing.T) {
	tests := prototests{
		{
			name:             "empty key",
			msg:              &v1.Labels{Labels: map[string]string{"": "value"}},
			wantErr:          true,
			wantErrorMessage: `validation error: labels[""]: must be at least 1 characters`,
		},
		{
			name:             "value too long",
			msg:              &v1.Labels{Labels: map[string]string{"key": strings.Repeat("x", 257)}},
			wantErr:          true,
			wantErrorMessage: `validation error: labels["key"]: must be at most 256 characters`,
		},
		{
			name:             "untrimmed key and value",
			msg:              &v1.Labels{Labels: map[string]string{" key ": " value "}},
			wantErr:          true,
			wantErrorMessage: `validation error: labels: keys and values must not start or end with whitespace`,
		},
		{
			name:    "valid",
			msg:     &v1.Labels{Labels: map[string]string{"key": "value"}},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateUpdateMeta(t *testing.T) {
	tests := prototests{
		{
			name:             "undefined locking strategy",
			msg:              &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy(99)},
			wantErr:          true,
			wantErrorMessage: `validation error: locking_strategy: value must be one of the defined enum values`,
		},
		{
			name:    "client locking strategy",
			msg:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_CLIENT},
			wantErr: false,
		},
		{
			name:    "server locking strategy",
			msg:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}
