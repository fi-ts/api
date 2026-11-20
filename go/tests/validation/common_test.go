package validation

import (
	"strings"
	"testing"

	v1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateLabels(t *testing.T) {
	tests := Prototests{
		{
			Name:             "empty key",
			Msg:              &v1.Labels{Labels: map[string]string{"": "value"}},
			WantErr:          true,
			WantErrorMessage: `validation error: labels[""]: must be at least 1 characters`,
		},
		{
			Name:             "value too long",
			Msg:              &v1.Labels{Labels: map[string]string{"key": strings.Repeat("x", 257)}},
			WantErr:          true,
			WantErrorMessage: `validation error: labels["key"]: must be at most 256 characters`,
		},
		{
			Name:             "untrimmed key and value",
			Msg:              &v1.Labels{Labels: map[string]string{" key ": " value "}},
			WantErr:          true,
			WantErrorMessage: `validation error: labels: keys and values must not start or end with whitespace`,
		},
		{
			Name:    "valid",
			Msg:     &v1.Labels{Labels: map[string]string{"key": "value"}},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}

func TestValidateUpdateMeta(t *testing.T) {
	tests := Prototests{
		{
			Name:             "undefined locking strategy",
			Msg:              &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy(99)},
			WantErr:          true,
			WantErrorMessage: `validation error: locking_strategy: value must be one of the defined enum values`,
		},
		{
			Name:    "client locking strategy",
			Msg:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_CLIENT},
			WantErr: false,
		},
		{
			Name:    "server locking strategy",
			Msg:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
			WantErr: false,
		},
	}

	ValidateProtos(t, tests)
}
