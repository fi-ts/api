package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
)

func TestValidateVLAN(t *testing.T) {
	vlan := func(id int32) *mvmv1.VLAN {
		return &mvmv1.VLAN{
			Uuid:         testUUID,
			VlanId:       id,
			SubnetTitle:  "pod01-vlan2001 (Production)",
			SubnetCidr:   "10.0.0.0/24",
			LocationUuid: testUUID,
			Tenant:       "tenant",
		}
	}

	const msgVlanID = "validation error: vlan_id: must be greater than or equal to 1 and less than or equal to 4094"

	tests := prototests{
		{
			name:    "min vlan id",
			msg:     vlan(1),
			wantErr: false,
		},
		{
			name:    "max vlan id",
			msg:     vlan(4094),
			wantErr: false,
		},
		{
			name:             "vlan id zero",
			msg:              vlan(0),
			wantErr:          true,
			wantErrorMessage: msgVlanID,
		},
		{
			name:             "vlan id too large",
			msg:              vlan(4095),
			wantErr:          true,
			wantErrorMessage: msgVlanID,
		},
	}

	validateProtos(t, tests)
}
