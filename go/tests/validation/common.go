// Package validation contains shared helpers for the protovalidate tests.
// The tests for fits.api.v1 live in this package, product-specific tests
// live in subpackages mirroring the proto packages (e.g. mvm, pcvm).
package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
)

// TestUUID is a valid UUID to be used in tests.
const TestUUID = "e266fcc6-f6de-4ee1-ba26-baa17bf47b13"

type (
	// Prototest is a single validation test case.
	Prototest struct {
		// Name of the test case.
		Name string
		// Msg is the message to validate.
		Msg proto.Message
		// WantErr is true if the validation is expected to fail.
		WantErr bool
		// WantErrorMessage is the expected error message if WantErr is true.
		WantErrorMessage string
	}

	// Prototests is a list of validation test cases.
	Prototests []Prototest
)

// ValidateProtos runs protovalidate for every test case and compares the result.
func ValidateProtos(t *testing.T, tests Prototests) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := protovalidate.Validate(tt.Msg)
			if !tt.WantErr {
				if err != nil {
					t.Errorf("validate error = %v, wantErr %v", err, tt.WantErr)
				}
				return
			}
			if err == nil {
				t.Fatalf("validate error = nil, wantErr %v", tt.WantErr)
			}
			if diff := cmp.Diff(err.Error(), tt.WantErrorMessage); diff != "" {
				t.Errorf("validate error = %v, diff %v", err, diff)
			}
		})
	}
}
