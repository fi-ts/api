package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
)

func TestValidateNetworkInterface(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.NetworkInterface{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - vlan_uuid: value is empty, which is not a valid UUID
 - ip_address: must be a valid IP address or hostname
 - mac_address: this string must be a valid mac address`,
		},
		{
			name: "valid",
			msg: &mvmv1.NetworkInterface{
				Uuid:       testUUID,
				VlanUuid:   testUUID,
				IpAddress:  "10.0.0.1",
				MacAddress: "aa:bb:cc:dd:ee:ff",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateInitialNetworkInterface(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &mvmv1.InitialNetworkInterface{},
			wantErr:          true,
			wantErrorMessage: `validation error: vlan_uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &mvmv1.InitialNetworkInterface{VlanUuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateManagedVM(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.ManagedVM{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - details: exactly one field is required in oneof
 - uuid: value is empty, which is not a valid UUID
 - fqdn: must be a valid FQDN, e.g. host.example.com
 - tenant: must be within 2 and 128 characters
 - project_uuid: value is empty, which is not a valid UUID
 - os_uuid: value is empty, which is not a valid UUID
 - location_uuid: value is empty, which is not a valid UUID
 - contact_uuid: value is empty, which is not a valid UUID
 - order_number: must be within 2 and 128 characters and not start or end with whitespace`,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.MVMServiceGetRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &mvmv1.MVMServiceGetRequest{Uuid: testUUID, ProjectUuid: testUUID},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceCreateWindowsRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.MVMServiceCreateWindowsRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - domain_uuid: value is empty, which is not a valid UUID
 - disks: must contain at least 1 item(s)`,
		},
	}

	validateProtos(t, tests)
}
