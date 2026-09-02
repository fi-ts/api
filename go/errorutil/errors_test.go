package errorutil

import (
	"errors"
	"fmt"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNotFound(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		args    []any
		wantErr error
	}{
		{
			name:    "error gets created as expected",
			format:  "something went wrong: %s",
			args:    []any{"abc"},
			wantErr: errors.New("not_found: something went wrong: abc"),
		},
		{
			name:    "error gets wrapped",
			format:  "something went wrong: %w",
			args:    []any{fmt.Errorf("abc")},
			wantErr: errors.New("not_found: something went wrong: abc"),
		},
		{
			name:    "error gets wrapped",
			format:  "something went wrong: %w",
			args:    []any{InvalidArgument("abc")},
			wantErr: errors.New("not_found: something went wrong: invalid_argument: abc"),
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			err := NotFound(tt.format, tt.args...)
			require.Error(t, err)

			assert.EqualError(t, tt.wantErr, err.Error())
		})
	}
}

func TestIsErrFns(t *testing.T) {
	for _, fns := range []struct {
		name  string
		errFn func(err error) bool
		code  connect.Code
	}{
		{
			name:  "not found",
			errFn: IsNotFound,
			code:  connect.CodeNotFound,
		},
		{
			name:  "conflict",
			errFn: IsConflict,
			code:  connect.CodeAlreadyExists,
		},
		{
			name:  "invalid",
			errFn: IsInvalidArgument,
			code:  connect.CodeInvalidArgument,
		},
		{
			name:  "unauthenticated",
			errFn: IsUnauthenticated,
			code:  connect.CodeUnauthenticated,
		},
		{
			name:  "unimplemented",
			errFn: IsUnimplemented,
			code:  connect.CodeUnimplemented,
		},
	} {
		tests := []struct {
			name    string
			errorFn func(err error) bool
			err     error
			want    bool
		}{
			{
				name:    "plain error",
				errorFn: fns.errFn,
				err:     errors.New("Some other Error"),
				want:    false,
			},
			{
				name:    "connect error",
				errorFn: fns.errFn,
				err:     connect.NewError(fns.code, errors.New("")),
				want:    true,
			},
			{
				name:    "nil",
				errorFn: fns.errFn,
				err:     nil,
				want:    false,
			},
			{
				name:    "wrapped",
				errorFn: fns.errFn,
				err:     fmt.Errorf("wrapped: %w", connect.NewError(fns.code, errors.New(""))),
				want:    true,
			},
			{
				name:    "grpc",
				errorFn: fns.errFn,
				err:     fmt.Errorf("wrapped: %w", status.Error(codes.Code(fns.code), "")),
				want:    true,
			},
		}
		for i := range tests {
			tt := tests[i]
			t.Run(fns.name+" "+tt.name, func(t *testing.T) {
				if got := tt.errorFn(tt.err); got != tt.want {
					t.Errorf("%T = %v, want %v", tt.errorFn, got, tt.want)
				}
			})
		}
	}
}

func TestIsInternal(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "plain error",
			err:  errors.New("Some other Error"),
			want: true, // a plain error is interpreted as internal (see default of convert)
		},
		{
			name: "connect error",
			err:  connect.NewError(connect.CodeInternal, errors.New("")),
			want: true,
		},
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "wrapped",
			err:  fmt.Errorf("wrapped: %w", connect.NewError(connect.CodeInternal, errors.New(""))),
			want: true,
		},
		{
			name: "grpc",
			err:  fmt.Errorf("wrapped: %w", status.Error(codes.Internal, "")),
			want: true,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInternal(tt.err); got != tt.want {
				t.Errorf("%v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrappedInternal(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name:    "one error",
			err:     Internal("something went wrong"),
			wantErr: connect.NewError(connect.CodeInternal, errors.New("something went wrong")),
		},
		{
			name:    "two errors",
			err:     Internal("wrapping error: %w", fmt.Errorf("root error")),
			wantErr: connect.NewError(connect.CodeInternal, fmt.Errorf("wrapping error: %w", errors.New("root error"))),
		},
		{
			name:    "two errors upside down",
			err:     fmt.Errorf("wrapping error: %w", Internal("root error")),
			wantErr: fmt.Errorf("wrapping error: %w", connect.NewError(connect.CodeInternal, errors.New("root error"))),
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantErr, tt.err, ErrorStringComparer()); diff != "" {
				t.Errorf("wrapping() %v", diff)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name:    "non specific error",
			err:     fmt.Errorf("something went wrong"),
			wantErr: connect.NewError(connect.CodeInternal, errors.New("something went wrong")),
		},
		{
			name:    "passing connect error",
			err:     Internal("something went wrong"),
			wantErr: connect.NewError(connect.CodeInternal, errors.New("something went wrong")),
		},
		{
			name:    "wrapped connect error",
			err:     Internal("wrapping error: %w", fmt.Errorf("root error")),
			wantErr: connect.NewError(connect.CodeInternal, fmt.Errorf("wrapping error: %w", errors.New("root error"))),
		},
		{
			name:    "wrapped connect error gets unwrapped",
			err:     fmt.Errorf("wrapping error: %w", Internal("root error")),
			wantErr: connect.NewError(connect.CodeInternal, fmt.Errorf("wrapping error: root error")),
		},
		{
			name:    "doubly-wrapped connect error gets unwraps the first and maintains the last",
			err:     fmt.Errorf("wrapping error: %w", Internal("doubly-wrapped: %w", InvalidArgument("invalid"))),
			wantErr: connect.NewError(connect.CodeInternal, fmt.Errorf("wrapping error: doubly-wrapped: invalid_argument: invalid")),
		},
		{
			name:    "grpc error gets converted",
			err:     status.Errorf(codes.AlreadyExists, "project already exists"), // TODO: check if this is really returned by a grpc client like that
			wantErr: connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("project already exists")),
		},
		{
			name:    "wrapped grpc error gets unwrapped",
			err:     fmt.Errorf("wrapping error: %w", status.Errorf(codes.AlreadyExists, "project already exists")),
			wantErr: connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("wrapping error: project already exists")),
		},
		{
			name:    "doubly-wrapped grpc error unwraps the first and maintains the last",
			err:     fmt.Errorf("wrapping error: %w", status.Errorf(codes.AlreadyExists, "doubly-wrapped: %s", status.Errorf(codes.FailedPrecondition, "precondition"))),
			wantErr: connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("wrapping error: doubly-wrapped: rpc error: code = FailedPrecondition desc = precondition")),
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantErr, Convert(tt.err), ErrorStringComparer()); diff != "" {
				t.Errorf("wrapping() %v", diff)
			}
		})
	}
}
