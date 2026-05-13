package validation

import (
	"testing"

	apiv1 "github.com/fi-ts/api/go/fits/api/v1"
)

func TestValidateIP(t *testing.T) {
	tests := prototests{
		{
			name: "Invalid IP",
			msg: &apiv1.IP{
				Uuid: "abc",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: must be a valid UUID
 - ip: value is empty, which is not a valid IP address
 - name: must be within 2 and 128 characters
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP, but valid uuid",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Description: "A IP",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - ip: value is empty, which is not a valid IP address
 - name: must be within 2 and 128 characters
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP, but valid uuid and ipv4",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Ip:          "1.2.3.4",
				Description: "A IP",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - name: must be within 2 and 128 characters
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP, but valid uuid and ipv6",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Ip:          "fe80:db8::1",
				Description: "A IP",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - name: must be within 2 and 128 characters
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP, but valid uuid, name and ipv6",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:        "Test IPv6",
				Ip:          "fe80:db8::1",
				Description: "A IP",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP, but valid uuid, name, network and ipv6",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:        "Test IPv6",
				Ip:          "fe80:db8::1",
				Network:     "Internet",
				Description: "A IP",
			},
			wantErr:          true,
			wantErrorMessage: `validation error: project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Invalid IP with invalid type",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:        "Test IPv6",
				Ip:          "fe80:db8::1",
				Network:     "Internet",
				Project:     "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Type:        apiv1.IPType(99),
				Description: "A IP",
			},
			wantErr:          true,
			wantErrorMessage: `validation error: type: value must be one of the defined enum values`,
		},
		{
			name: "Valid IP",
			msg: &apiv1.IP{
				Uuid:        "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:        "Test IPv6",
				Ip:          "fe80:db8::1",
				Network:     "Internet",
				Project:     "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Type:        apiv1.IPType_IP_TYPE_EPHEMERAL,
				Description: "A IP",
			},
			wantErr: false,
		},
		// IPServiceGetRequest
		{
			name:    "invalid IPServiceGetRequest",
			msg:     &apiv1.IPServiceGetRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - ip: value is empty, which is not a valid IP address
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Valid IPServiceGetRequest",
			msg: &apiv1.IPServiceGetRequest{
				Ip:      "fe80:db8::1",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
			},
			wantErr: false,
		},
		// IPServiceCreateRequest
		{
			name:    "invalid IPServiceCreateRequest",
			msg:     &apiv1.IPServiceCreateRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - network: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID`,
		},
		{
			name: "Valid IPServiceCreateRequest",
			msg: &apiv1.IPServiceCreateRequest{
				Network: "Internet",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Machine: new("57cd8678-9ff0-4f8c-a34a-43d8f16caacf"),
			},
			wantErr: false,
		},
		{
			name: "IPServiceCreateRequest name too short",
			msg: &apiv1.IPServiceCreateRequest{
				Network: "Internet",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Name:    new("a"),
			},
			wantErr:          true,
			wantErrorMessage: `validation error: name: must be within 2 and 128 characters`,
		},
	}

	validateProtos(t, tests)
}
