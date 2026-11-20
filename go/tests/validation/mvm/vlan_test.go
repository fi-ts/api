package mvm

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	"github.com/fi-ts/api/go/tests/validation"
)

func TestValidateVLAN(t *testing.T) {
	vlan := func(id int32) *mvmv1.VLAN {
		return &mvmv1.VLAN{
			Uuid:         validation.TestUUID,
			VlanId:       id,
			SubnetTitle:  "pod01-vlan2001 (Production)",
			SubnetCidr:   "10.0.0.0/24",
			LocationUuid: validation.TestUUID,
			Tenant:       "tenant",
		}
	}

	const msgVlanID = "validation error: vlan_id: must be greater than or equal to 1 and less than or equal to 4094"

	tests := validation.Prototests{
		{
			Name:    "min vlan id",
			Msg:     vlan(1),
			WantErr: false,
		},
		{
			Name:    "max vlan id",
			Msg:     vlan(4094),
			WantErr: false,
		},
		{
			Name:             "vlan id zero",
			Msg:              vlan(0),
			WantErr:          true,
			WantErrorMessage: msgVlanID,
		},
		{
			Name:             "vlan id too large",
			Msg:              vlan(4095),
			WantErr:          true,
			WantErrorMessage: msgVlanID,
		},
	}

	validation.ValidateProtos(t, tests)
}
