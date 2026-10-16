package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
)

type (
	prototest struct {
		name             string
		msg              proto.Message
		wantErr          bool
		wantErrorMessage string
	}

	prototests []prototest
)

func validateProtos(t *testing.T, tests prototests) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := protovalidate.Validate(tt.msg)
			if !tt.wantErr {
				if err != nil {
					t.Errorf("validate error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err == nil {
				t.Fatalf("validate error = nil, wantErr %v", tt.wantErr)
			}
			if diff := cmp.Diff(err.Error(), tt.wantErrorMessage); diff != "" {
				t.Errorf("validate error = %v, diff %v", err, diff)
			}
		})
	}
}
