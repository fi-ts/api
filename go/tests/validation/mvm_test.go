package validation

import (
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
)

const (
	msgCpuCount       = "must be within 1 and 64 CPU cores"
	msgRamSize        = "must be within 4 and 512 GiB"
	msgDiskSize       = "must be within 1 and 8192 GiB"
	msgDiskCount      = "must contain at most 8 disks"
	msgOrderReference = "must be within 2 and 128 characters and not start or end with whitespace"
	msgMacAddress     = "this string must be a valid mac address"
	msgMountPoint     = "must be an absolute path without whitespace, e.g. /data/app"
	msgDriveLetter    = "must be a single uppercase drive letter, e.g. G"
	msgAddDriveLetter = "drive letter of a Windows data disk must be within F and Z, except R"
	msgFqdn           = "must be a valid FQDN, e.g. host.example.com"
	msgNameCode       = "must be a 4-character code consisting of digits and letters"
	msgSnapshotName   = "must be within 1 and 128 alphanumeric characters"
	msgNotUnspecified = "must not be in list [0]"
)

func validLinuxDisk() *mvmv1.LinuxDisk {
	return &mvmv1.LinuxDisk{
		Name:       "appdata",
		SizeGb:     50,
		MountPoint: proto.String("/appdata"),
		DiskType:   mvmv1.DiskType_DISK_TYPE_SYSTEM.Enum(),
	}
}

func validWindowsDisk() *mvmv1.WindowsDisk {
	return &mvmv1.WindowsDisk{
		Name:        "system",
		SizeGb:      50,
		DriveLetter: proto.String("F"),
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
		ProjectUuid:      testUUID,
		NameCode:         "ab12",
		OsUuid:           testUUID,
		LocationUuid:     testUUID,
		ContactUuid:      testUUID,
		Cpu:              2,
		RamGb:            8,
		OrderReference:   "TICKET-1",
		Backup:           &mvmv1.Backup{Enabled: true},
		Encryption:       &mvmv1.Encryption{Enabled: false},
		Availability:     mvmv1.Availability_AVAILABILITY_V1,
		ServiceClass:     mvmv1.ServiceClass_SERVICE_CLASS_SZ1,
		InitialIpAddress: &mvmv1.InitialIPAddress{VlanUuid: testUUID},
		MvmType: &mvmv1.MVMServiceCreateRequest_Linux{
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
		Uuid:           testUUID,
		ProjectUuid:    testUUID,
		UpdateMeta:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
		OrderReference: "TICKET-1",
	}
	for _, m := range modify {
		m(req)
	}
	return req
}

func TestValidateIPAddress(t *testing.T) {
	ni := &mvmv1.NetworkInterface{VlanUuid: testUUID, MacAddress: proto.String("AA:BB:CC:DD:EE:01")}
	tests := prototests{
		{
			name:    "empty",
			msg:     &mvmv1.IPAddress{},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - ip_address: value is empty, which is not a valid IPv4 address
 - network_interface: value is required`,
		},
		{
			name:             "hostname instead of ip address",
			msg:              &mvmv1.IPAddress{Uuid: testUUID, IpAddress: "host.example.com", NetworkInterface: ni},
			wantErr:          true,
			wantErrorMessage: "validation error: ip_address: must be a valid IPv4 address",
		},
		{
			name:             "ipv6 address",
			msg:              &mvmv1.IPAddress{Uuid: testUUID, IpAddress: "2001:db8::1", NetworkInterface: ni},
			wantErr:          true,
			wantErrorMessage: "validation error: ip_address: must be a valid IPv4 address",
		},
		{
			name: "missing network interface",
			msg: &mvmv1.IPAddress{
				Uuid:      testUUID,
				IpAddress: "10.0.0.1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: network_interface: value is required",
		},
		{
			name: "valid with network interface",
			msg: &mvmv1.IPAddress{
				Uuid:      testUUID,
				IpAddress: "10.0.0.1",
				NetworkInterface: &mvmv1.NetworkInterface{
					VlanUuid:   testUUID,
					MacAddress: proto.String("AA:BB:CC:DD:EE:01"),
				},
			},
			wantErr: false,
		},
		{
			name: "valid additional ip address without mac address",
			msg: &mvmv1.IPAddress{
				Uuid:      testUUID,
				IpAddress: "10.0.0.2",
				NetworkInterface: &mvmv1.NetworkInterface{
					VlanUuid: testUUID,
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
			name:             "empty",
			msg:              &mvmv1.NetworkInterface{},
			wantErr:          true,
			wantErrorMessage: "validation error: vlan_uuid: value is empty, which is not a valid UUID",
		},
		{
			name: "invalid mac address",
			msg: &mvmv1.NetworkInterface{
				VlanUuid:   testUUID,
				MacAddress: proto.String("AA:BB:CC:DD:EE"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac_address: " + msgMacAddress,
		},
		{
			name: "empty mac address",
			msg: &mvmv1.NetworkInterface{
				VlanUuid:   testUUID,
				MacAddress: proto.String(""),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac_address: " + msgMacAddress,
		},
		{
			name: "valid",
			msg: &mvmv1.NetworkInterface{
				VlanUuid:   testUUID,
				MacAddress: proto.String("AA:BB:CC:DD:EE:01"),
			},
			wantErr: false,
		},
		{
			name: "valid without mac address",
			msg: &mvmv1.NetworkInterface{
				VlanUuid: testUUID,
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
 - cpu: must be within 1 and 64 CPU cores
 - ram_gb: must be within 4 and 512 GiB
 - availability: ` + msgNotUnspecified + `
 - service_class: ` + msgNotUnspecified + `
 - ip_addresses: must contain at least one IP address with a MAC address
 - order_reference: must be within 2 and 128 characters and not start or end with whitespace`,
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
 - name: must be within 2 and 128 characters
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
 - name: must be within 2 and 128 characters
 - size_gb: ` + msgDiskSize,
		},
		{
			name:    "valid system disk",
			msg:     validWindowsDisk(),
			wantErr: false,
		},
		{
			name:             "lowercase drive letter",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("g") }),
			wantErr:          true,
			wantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			name:    "reserved drive letter is valid in responses",
			msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("R") }),
			wantErr: false,
		},
		{
			name:    "valid os disk drive letter",
			msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("C") }),
			wantErr: false,
		},
		{
			name:             "two drive letters",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("FG") }),
			wantErr:          true,
			wantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			name:             "drive letter with colon",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("F:") }),
			wantErr:          true,
			wantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			name:             "digit as drive letter",
			msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("1") }),
			wantErr:          true,
			wantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
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
					Name:        "system",
					SizeGb:      50,
					DriveLetter: proto.String("FG"),
				}},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: disks[0].drive_letter: " + msgDriveLetter,
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
				r.MvmType = &mvmv1.MVMServiceCreateRequest_Windows{
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
			name:             "missing order reference",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderReference = "" }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
		{
			name:             "untrimmed order reference",
			msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderReference = " TICKET-1 " }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_reference: " + msgOrderReference,
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
			name: "missing initial ip address",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.InitialIpAddress = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: initial_ip_address: value is required",
		},
		{
			name: "missing backup configuration",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Backup = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: backup: value is required",
		},
		{
			name: "missing encryption configuration",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Encryption = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: encryption: value is required",
		},
		{
			name: "unspecified availability",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Availability = mvmv1.Availability_AVAILABILITY_UNSPECIFIED
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: availability: " + msgNotUnspecified,
		},
		{
			name: "unspecified service class",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.ServiceClass = mvmv1.ServiceClass_SERVICE_CLASS_UNSPECIFIED
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: service_class: " + msgNotUnspecified,
		},
		{
			name: "missing os type",
			msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.MvmType = nil
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: mvm_type: exactly one field is required in oneof",
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
			name:             "missing order reference",
			msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.OrderReference = "" }),
			wantErr:          true,
			wantErrorMessage: "validation error: order_reference: " + msgOrderReference,
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
	windowsDisk := func(letter string) []*mvmv1.Disk {
		d := validWindowsDisk()
		d.DiskType = mvmv1.DiskType_DISK_TYPE_DATA.Enum()
		d.DriveLetter = nil
		if letter != "" {
			d.DriveLetter = proto.String(letter)
		}
		return []*mvmv1.Disk{{Disk: &mvmv1.Disk_Windows{Windows: d}}}
	}
	request := func(d []*mvmv1.Disk) *mvmv1.MVMServiceAddDisksRequest {
		return &mvmv1.MVMServiceAddDisksRequest{
			Uuid:           testUUID,
			ProjectUuid:    testUUID,
			OrderReference: "TICKET-1",
			Disks:          d,
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
				Name:       "data",
				SizeGb:     10,
				MountPoint: proto.String("data"),
			}}}}),
			wantErr:          true,
			wantErrorMessage: "validation error: disks[0].linux.mount_point: " + msgMountPoint,
		},
		{
			name:    "valid windows drive letter",
			msg:     request(windowsDisk("G")),
			wantErr: false,
		},
		{
			name:    "windows disk without drive letter",
			msg:     request(windowsDisk("")),
			wantErr: false,
		},
		{
			name:             "windows drive letter before F",
			msg:              request(windowsDisk("E")),
			wantErr:          true,
			wantErrorMessage: "validation error: disks: " + msgAddDriveLetter,
		},
		{
			name:             "windows drive letter R",
			msg:              request(windowsDisk("R")),
			wantErr:          true,
			wantErrorMessage: "validation error: disks: " + msgAddDriveLetter,
		},
		{
			name:    "windows lowercase drive letter",
			msg:     request(windowsDisk("g")),
			wantErr: true,
			wantErrorMessage: `validation errors:
 - disks: ` + msgAddDriveLetter + `
 - disks[0].windows.drive_letter: ` + msgDriveLetter,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceUpdateDiskRequest(t *testing.T) {
	request := func(size *uint64) *mvmv1.MVMServiceUpdateDiskRequest {
		return &mvmv1.MVMServiceUpdateDiskRequest{
			Uuid:           testUUID,
			DiskUuid:       testUUID,
			ProjectUuid:    testUUID,
			OrderReference: "TICKET-1",
			SizeGb:         size,
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
 - order_reference: ` + msgOrderReference,
		},
		{
			name: "order reference too short",
			msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:           testUUID,
				ProjectUuid:    testUUID,
				OrderReference: "x",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
		{
			name: "valid",
			msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:           testUUID,
				ProjectUuid:    testUUID,
				OrderReference: "TICKET-1",
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceAddSnapshotRequest(t *testing.T) {
	request := func(name string) *mvmv1.MVMServiceAddSnapshotRequest {
		return &mvmv1.MVMServiceAddSnapshotRequest{
			Uuid:           testUUID,
			ProjectUuid:    testUUID,
			OrderReference: "TICKET-1",
			Name:           name,
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
			wantErrorMessage: "validation error: name: " + msgSnapshotName,
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
			wantErrorMessage: "validation error: name: " + msgSnapshotName,
		},
	}

	validateProtos(t, tests)
}

func TestValidateMVMServiceReinstallRequest(t *testing.T) {
	request := func(modify ...func(*mvmv1.MVMServiceReinstallRequest)) *mvmv1.MVMServiceReinstallRequest {
		req := &mvmv1.MVMServiceReinstallRequest{
			Uuid:           testUUID,
			ProjectUuid:    testUUID,
			OrderReference: "TICKET-1",
		}
		for _, m := range modify {
			m(req)
		}
		return req
	}

	tests := prototests{
		{
			name:    "valid without changes keeps current configuration",
			msg:     request(),
			wantErr: false,
		},
		{
			name: "valid with disabled backup and enabled encryption",
			msg: request(func(r *mvmv1.MVMServiceReinstallRequest) {
				r.OsUuid = proto.String(testUUID)
				r.Backup = &mvmv1.Backup{Enabled: false}
				r.Encryption = &mvmv1.Encryption{Enabled: true}
			}),
			wantErr: false,
		},
		{
			name: "invalid os uuid",
			msg: request(func(r *mvmv1.MVMServiceReinstallRequest) {
				r.OsUuid = proto.String("not-a-uuid")
			}),
			wantErr:          true,
			wantErrorMessage: "validation error: os_uuid: must be a valid UUID",
		},
	}

	validateProtos(t, tests)
}
