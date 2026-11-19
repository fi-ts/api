package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

const (
	msgCpuCount    = "must be within 1 and 64 CPU cores"
	msgRamSize     = "must be within 4 and 512 GiB"
	msgDiskSize    = "must be within 1 and 8192 GiB"
	msgDiskCount   = "must contain at most 8 disks"
	msgOrderNumber = "must be within 2 and 128 characters and not start or end with whitespace"
	msgMountPoint  = "must be an absolute path without whitespace, e.g. /data/app"
	msgDriveLetter = "must be a single drive letter, e.g. G"
	msgFqdn        = "must be a valid FQDN, e.g. host.example.com"
	msgNameCode    = "must be a 4-character code consisting of digits and letters"
)

func validLinuxDisk() *mvmv1.LinuxDisk {
	return &mvmv1.LinuxDisk{
		Label:      "appdata",
		SizeGb:     50,
		MountPoint: proto.String("/appdata"),
		DiskType:   mvmv1.DiskType_DISK_TYPE_SYSTEM.Enum(),
	}
}

func validWindowsDisk() *mvmv1.WindowsDisk {
	return &mvmv1.WindowsDisk{
		Label:       "system",
		SizeGb:      50,
		Driveletter: proto.String("F"),
		DiskType:    mvmv1.DiskType_DISK_TYPE_SYSTEM.Enum(),
	}
}

func linuxDisks(n int) []*mvmv1.LinuxDisk {
	disks := make([]*mvmv1.LinuxDisk, 0, n)
	for range n {
		disks = append(disks, validLinuxDisk())
	}
	return disks
}

func validLinuxCreateRequest(modify ...func(*mvmv1.MVMServiceCreateRequest)) *mvmv1.MVMServiceCreateRequest {
	req := &mvmv1.MVMServiceCreateRequest{
		ProjectUuid:  testUUID,
		NameCode:     "ab12",
		OsUuid:       testUUID,
		LocationUuid: testUUID,
		ContactUuid:  testUUID,
		Cpu:          2,
		RamGb:        8,
		OrderNumber:  "TICKET-1",
		Mvmtype: &mvmv1.MVMServiceCreateRequest_Linux{
			Linux: &mvmv1.MVMServiceCreateLinuxRequest{
				LdapUuid: testUUID,
				Disks:    []*mvmv1.LinuxDisk{validLinuxDisk()},
			},
		},
	}
	for _, m := range modify {
		m(req)
	}
	return req
}

func validUpdateRequest(modify ...func(*mvmv1.MVMServiceUpdateRequest)) *mvmv1.MVMServiceUpdateRequest {
	req := &mvmv1.MVMServiceUpdateRequest{
		Uuid:        testUUID,
		ProjectUuid: testUUID,
		UpdateMeta:  &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
		OrderNumber: "TICKET-1",
	}
	for _, m := range modify {
		m(req)
	}
	return req
}

func TestValidateIPAddress(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.IPAddress{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - ip_address: must be a valid IP address or hostname`,
		},
		{
			name: "valid",
			msg: &mvmv1.IPAddress{
				Uuid:      testUUID,
				IpAddress: "10.0.0.1",
			},
			wantErr: false,
		},
		{
			name: "valid with network interface",
			msg: &mvmv1.IPAddress{
				Uuid:      testUUID,
				IpAddress: "10.0.0.1",
				NetworkInterface: &mvmv1.NetworkInterface{
					VlanUuid:   testUUID,
					Macaddress: "AA:BB:CC:DD:EE:01",
				},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateNetworkInterface(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.NetworkInterface{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - vlan_uuid: value is empty, which is not a valid UUID
 - macaddress: this string must be a valid macaddress`,
		},
		{
			name: "invalid mac address",
			msg: &mvmv1.NetworkInterface{
				VlanUuid:   testUUID,
				Macaddress: "AA:BB:CC:DD:EE",
			},
			wantErr:          true,
			wantErrorMessage: `validation error: macaddress: this string must be a valid macaddress`,
		},
		{
			name: "valid",
			msg: &mvmv1.NetworkInterface{
				VlanUuid:   testUUID,
				Macaddress: "AA:BB:CC:DD:EE:01",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateInitialIPAddress(t *testing.T) {
	tests := prototests{
		{
			name:             "empty",
			msg:              &mvmv1.InitialIPAddress{},
			wantErr:          true,
			wantErrorMessage: `validation error: vlan_uuid: value is empty, which is not a valid UUID`,
		},
		{
			name:    "valid",
			msg:     &mvmv1.InitialIPAddress{VlanUuid: testUUID},
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

func TestValidateLinuxDetails(t *testing.T) {
	tests := prototests{
		{
			name:             "fqdn without domain",
			msg:              &mvmv1.LinuxDetails{LdapUuid: testUUID, LdapFqdn: "host"},
			wantErr:          true,
			wantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			name:             "fqdn with numeric tld",
			msg:              &mvmv1.LinuxDetails{LdapUuid: testUUID, LdapFqdn: "host.example.123"},
			wantErr:          true,
			wantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			name:    "valid",
			msg:     &mvmv1.LinuxDetails{LdapUuid: testUUID, LdapFqdn: "host.example.com"},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateWindowsDetails(t *testing.T) {
	tests := prototests{
		{
			name:             "invalid fqdn",
			msg:              &mvmv1.WindowsDetails{DomainUuid: testUUID, DomainFqdn: "not a fqdn"},
			wantErr:          true,
			wantErrorMessage: "validation error: domain_fqdn: " + msgFqdn,
		},
		{
			name:    "valid",
			msg:     &mvmv1.WindowsDetails{DomainUuid: testUUID, DomainFqdn: "host.domain.example.com"},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateLinuxDisk(t *testing.T) {
	withDisk := func(modify func(*mvmv1.LinuxDisk)) *mvmv1.LinuxDisk {
		d := validLinuxDisk()
		modify(d)
		return d
	}

	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.LinuxDisk{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - label: must be within 2 and 128 characters
 - size_gb: ` + msgDiskSize,
		},
		{
			name:    "valid system disk",
			msg:     validLinuxDisk(),
			wantErr: false,
		},
		{
			name: "valid data disk without mount point",
			msg: withDisk(func(d *mvmv1.LinuxDisk) {
				d.MountPoint = nil
				d.DiskType = mvmv1.DiskType_DISK_TYPE_DATA.Enum()
				d.AutoExtend = proto.Bool(true)
			}),
			wantErr: false,
		},
		{
			name:    "valid root mount point",
			msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/") }),
			wantErr: false,
		},
		{
			name:    "valid mount point with trailing slash and special characters",
			msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/data/my+dir/") }),
			wantErr: false,
		},
		{
			name:             "relative mount point",
			msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("data/app") }),
			wantErr:          true,
			wantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			name:             "empty mount point",
			msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("") }),
			wantErr:          true,
			wantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			name:             "mount point with whitespace",
			msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/data/my dir") }),
			wantErr:          true,
			wantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			name:    "min disk size",
			msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 1 }),
			wantErr: false,
		},
		{
			name:    "max disk size",
			msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 8192 }),
			wantErr: false,
		},
		{
			name:             "disk too large",
			msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 8193 }),
			wantErr:          true,
			wantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
		{
			name:             "invalid uuid",
			msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.Uuid = proto.String("not-a-uuid") }),
			wantErr:          true,
			wantErrorMessage: "validation error: uuid: must be a valid UUID",
		},
	}

	validateProtos(t, tests)
}

func TestValidateWindowsDisk(t *testing.T) {
	withDisk := func(modify func(*mvmv1.WindowsDisk)) *mvmv1.WindowsDisk {
		d := validWindowsDisk()
		modify(d)
		return d
	}

	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.WindowsDisk{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - label: must be within 2 and 128 characters
 - size_gb: ` + msgDiskSize,
		},
		{
			name:    "valid system disk",
			msg:     validWindowsDisk(),
			wantErr: false,
		},
		{
			name:    "valid lowercase drive letter",
			msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.Driveletter = proto.String("g") }),
			wantErr: false,
		},
		{
			name:    "valid os disk drive letter",
			msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.Driveletter = proto.String("C") }),
			wantErr: false,
		},
		{
			name:             "two drive letters",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.Driveletter = proto.String("FG") }),
			wantErr:          true,
			wantErrorMessage: "validation error: driveletter: " + msgDriveLetter,
		},
		{
			name:             "drive letter with colon",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.Driveletter = proto.String("F:") }),
			wantErr:          true,
			wantErrorMessage: "validation error: driveletter: " + msgDriveLetter,
		},
		{
			name:             "digit as drive letter",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.Driveletter = proto.String("1") }),
			wantErr:          true,
			wantErrorMessage: "validation error: driveletter: " + msgDriveLetter,
		},
		{
			name:             "disk too large",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.SizeGb = 8193 }),
			wantErr:          true,
			wantErrorMessage: "validation error: size_gb: " + msgDiskSize,
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
		{
			name: "invalid disk",
			msg: &mvmv1.MVMServiceCreateWindowsRequest{
				DomainUuid: testUUID,
				Disks: []*mvmv1.WindowsDisk{{
					Label:       "system",
					SizeGb:      50,
					Driveletter: proto.String("FG"),
				}},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: disks[0].driveletter: " + msgDriveLetter,
		},
		{
			name: "valid",
			msg: &mvmv1.MVMServiceCreateWindowsRequest{
				DomainUuid: testUUID,
				Disks:      []*mvmv1.WindowsDisk{validWindowsDisk()},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceCreateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "valid linux",
			msg:     validLinuxCreateRequest(),
			wantErr: false,
		},
		{
			name: "valid windows",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Mvmtype = &mvmv1.MVMServiceCreateRequest_Windows{
					Windows: &mvmv1.MVMServiceCreateWindowsRequest{
						DomainUuid: testUUID,
						Disks:      []*mvmv1.WindowsDisk{validWindowsDisk()},
					},
				}
			}),
			wantErr: false,
		},
		{
			name: "valid bounds",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Cpu = 64
				r.RamGb = 512
			}),
			wantErr: false,
		},
		{
			name:             "no cpu",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.Cpu = 0 }),
			wantErr:          true,
			wantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			name:             "too many cpus",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.Cpu = 65 }),
			wantErr:          true,
			wantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			name:             "too little ram",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.RamGb = 3 }),
			wantErr:          true,
			wantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			name:             "too much ram",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.RamGb = 513 }),
			wantErr:          true,
			wantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			name:             "invalid name code",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.NameCode = "ab-1" }),
			wantErr:          true,
			wantErrorMessage: "validation error: name_code: " + msgNameCode,
		},
		{
			name:             "missing order number",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderNumber = "" }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_number: " + msgOrderNumber,
		},
		{
			name:             "untrimmed order number",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderNumber = " TICKET-1 " }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_number: " + msgOrderNumber,
		},
		{
			name: "no disks",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: linux.disks: must contain at least 1 item(s)",
		},
		{
			name: "max disks",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = linuxDisks(8)
			}),
			wantErr: false,
		},
		{
			name: "too many disks",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = linuxDisks(9)
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: linux.disks: " + msgDiskCount,
		},
		{
			name: "invalid disk",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks[0].SizeGb = 0
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: linux.disks[0].size_gb: " + msgDiskSize,
		},
		{
			name: "missing os type",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Mvmtype = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: mvmtype: exactly one field is required in oneof",
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceUpdateRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "valid without changes",
			msg:     validUpdateRequest(),
			wantErr: false,
		},
		{
			name: "valid cpu and ram",
			msg: validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) {
				r.Cpu = proto.Uint32(4)
				r.RamGb = proto.Uint32(16)
			}),
			wantErr: false,
		},
		{
			name:             "too many cpus",
			msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.Cpu = proto.Uint32(65) }),
			wantErr:          true,
			wantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			name:             "zero ram",
			msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.RamGb = proto.Uint32(0) }),
			wantErr:          true,
			wantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			name:             "missing order number",
			msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.OrderNumber = "" }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_number: " + msgOrderNumber,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceAddDisksRequest(t *testing.T) {
	disks := func(n int) []*mvmv1.Disk {
		d := make([]*mvmv1.Disk, 0, n)
		for range n {
			d = append(d, &mvmv1.Disk{Disk: &mvmv1.Disk_Linux{Linux: validLinuxDisk()}})
		}
		return d
	}
	request := func(d []*mvmv1.Disk) *mvmv1.MVMServiceAddDisksRequest {
		return &mvmv1.MVMServiceAddDisksRequest{
			Uuid:        testUUID,
			ProjectUuid: testUUID,
			OrderNumber: "TICKET-1",
			Disks:       d,
		}
	}

	tests := prototests{
		{
			name:    "valid",
			msg:     request(disks(1)),
			wantErr: false,
		},
		{
			name:             "no disks",
			msg:              request(nil),
			wantErr:          true,
			wantErrorMessage: "validation error: disks: must contain at least 1 item(s)",
		},
		{
			name:             "too many disks",
			msg:              request(disks(9)),
			wantErr:          true,
			wantErrorMessage: "validation error: disks: " + msgDiskCount,
		},
		{
			name:             "disk without os type",
			msg:              request([]*mvmv1.Disk{{}}),
			wantErr:          true,
			wantErrorMessage: "validation error: disks[0].disk: exactly one field is required in oneof",
		},
		{
			name: "invalid disk",
			msg: request([]*mvmv1.Disk{{Disk: &mvmv1.Disk_Linux{Linux: &mvmv1.LinuxDisk{
				Label:      "data",
				SizeGb:     10,
				MountPoint: proto.String("data"),
			}}}}),
			wantErr:          true,
			wantErrorMessage: "validation error: disks[0].linux.mount_point: " + msgMountPoint,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceUpdateDiskRequest(t *testing.T) {
	request := func(size *uint64) *mvmv1.MVMServiceUpdateDiskRequest {
		return &mvmv1.MVMServiceUpdateDiskRequest{
			Uuid:        testUUID,
			DiskUuid:    testUUID,
			ProjectUuid: testUUID,
			OrderNumber: "TICKET-1",
			SizeGb:      size,
		}
	}

	tests := prototests{
		{
			name:    "valid without size",
			msg:     request(nil),
			wantErr: false,
		},
		{
			name:    "valid size",
			msg:     request(proto.Uint64(100)),
			wantErr: false,
		},
		{
			name:             "zero size",
			msg:              request(proto.Uint64(0)),
			wantErr:          true,
			wantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
		{
			name:             "too large",
			msg:              request(proto.Uint64(8193)),
			wantErr:          true,
			wantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceDeleteRequest(t *testing.T) {
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.MVMServiceDeleteRequest{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID
 - order_number: ` + msgOrderNumber,
		},
		{
			name: "order number too short",
			msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:        testUUID,
				ProjectUuid: testUUID,
				OrderNumber: "x",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: order_number: " + msgOrderNumber,
		},
		{
			name: "valid",
			msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:        testUUID,
				ProjectUuid: testUUID,
				OrderNumber: "TICKET-1",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceAddSnapshotRequest(t *testing.T) {
	request := func(name string) *mvmv1.MVMServiceAddSnapshotRequest {
		return &mvmv1.MVMServiceAddSnapshotRequest{
			Uuid:        testUUID,
			ProjectUuid: testUUID,
			OrderNumber: "TICKET-1",
			Name:        name,
		}
	}

	tests := prototests{
		{
			name:    "valid",
			msg:     request("beforeUpgrade1"),
			wantErr: false,
		},
		{
			name:             "name with special characters",
			msg:              request("before-upgrade"),
			wantErr:          true,
			wantErrorMessage: "validation error: name: does not match regex pattern `^[a-zA-Z0-9]+$`",
		},
	}

	validateProtos(t, tests)
}

func TestValidateSnapshot(t *testing.T) {
	tests := prototests{
		{
			name:    "valid",
			msg:     &mvmv1.Snapshot{Uuid: testUUID, Name: "snapshot1"},
			wantErr: false,
		},
		{
			name:             "name with whitespace",
			msg:              &mvmv1.Snapshot{Uuid: testUUID, Name: "my snapshot"},
			wantErr:          true,
			wantErrorMessage: "validation error: name: does not match regex pattern `^[a-zA-Z0-9]+$`",
		},
	}

	validateProtos(t, tests)
}
