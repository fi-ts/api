package mvm

import (
	"strings"
	"testing"

	mvmv1 "github.com/fi-ts/api/go/fits/api/mvm/v1"
	v1 "github.com/fi-ts/api/go/fits/api/v1"
	"github.com/fi-ts/api/go/tests/validation"
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
		ProjectUuid:      validation.TestUUID,
		NameCode:         "ab12",
		OsUuid:           validation.TestUUID,
		LocationUuid:     validation.TestUUID,
		ContactUuid:      validation.TestUUID,
		Cpu:              2,
		RamGb:            8,
		OrderReference:   "TICKET-1",
		Backup:           &mvmv1.Backup{Enabled: true},
		Encryption:       &mvmv1.Encryption{Enabled: false},
		Availability:     mvmv1.Availability_AVAILABILITY_V1,
		ServiceClass:     mvmv1.ServiceClass_SERVICE_CLASS_SZ1,
		InitialIpAddress: &mvmv1.InitialIPAddress{VlanUuid: validation.TestUUID},
		MvmType: &mvmv1.MVMServiceCreateRequest_Linux{
			Linux: &mvmv1.MVMServiceCreateLinuxRequest{
				LdapUuid: validation.TestUUID,
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
		Uuid:           validation.TestUUID,
		ProjectUuid:    validation.TestUUID,
		UpdateMeta:     &v1.UpdateMeta{LockingStrategy: v1.OptimisticLockingStrategy_OPTIMISTIC_LOCKING_STRATEGY_SERVER},
		OrderReference: "TICKET-1",
	}
	for _, m := range modify {
		m(req)
	}
	return req
}

func TestValidateIPAddress(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.IPAddress{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - ip_address: value is empty, which is not a valid IPv4 address`,
		},
		{
			Name:             "hostname instead of ip address",
			Msg:              &mvmv1.IPAddress{Uuid: validation.TestUUID, IpAddress: "host.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ip_address: must be a valid IPv4 address",
		},
		{
			Name:             "ipv6 address",
			Msg:              &mvmv1.IPAddress{Uuid: validation.TestUUID, IpAddress: "2001:db8::1"},
			WantErr:          true,
			WantErrorMessage: "validation error: ip_address: must be a valid IPv4 address",
		},
		{
			Name: "valid",
			Msg: &mvmv1.IPAddress{
				Uuid:      validation.TestUUID,
				IpAddress: "10.0.0.1",
			},
			WantErr: false,
		},
		{
			Name: "valid with network interface",
			Msg: &mvmv1.IPAddress{
				Uuid:      validation.TestUUID,
				IpAddress: "10.0.0.1",
				NetworkInterface: &mvmv1.NetworkInterface{
					VlanUuid:   validation.TestUUID,
					MacAddress: "AA:BB:CC:DD:EE:01",
				},
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateNetworkInterface(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.NetworkInterface{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - vlan_uuid: value is empty, which is not a valid UUID
 - mac_address: ` + msgMacAddress,
		},
		{
			Name: "invalid mac address",
			Msg: &mvmv1.NetworkInterface{
				VlanUuid:   validation.TestUUID,
				MacAddress: "AA:BB:CC:DD:EE",
			},
			WantErr:          true,
			WantErrorMessage: "validation error: mac_address: " + msgMacAddress,
		},
		{
			Name: "valid",
			Msg: &mvmv1.NetworkInterface{
				VlanUuid:   validation.TestUUID,
				MacAddress: "AA:BB:CC:DD:EE:01",
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateInitialIPAddress(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "empty",
			Msg:              &mvmv1.InitialIPAddress{},
			WantErr:          true,
			WantErrorMessage: `validation error: vlan_uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.InitialIPAddress{VlanUuid: validation.TestUUID},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateManagedVM(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.ManagedVM{},
			WantErr: true,
			WantErrorMessage: `validation errors:
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
 - order_reference: must be within 2 and 128 characters and not start or end with whitespace`,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateLinuxDetails(t *testing.T) {
	// label63 is the longest allowed single label.
	label63 := strings.Repeat("a", 63)
	// fqdn253 is the longest allowed FQDN: 3*(63+1) + 57 + 4 (".com") = 253 characters.
	fqdn253 := strings.Repeat(label63+".", 3) + strings.Repeat("b", 57) + ".com"
	// fqdn254 exceeds the total length by one, while every label is still valid.
	fqdn254 := strings.Repeat(label63+".", 3) + strings.Repeat("b", 58) + ".com"
	if len(fqdn253) != 253 || len(fqdn254) != 254 {
		t.Fatalf("got lengths %d and %d, want 253 and 254", len(fqdn253), len(fqdn254))
	}

	tests := validation.Prototests{
		{
			Name:             "fqdn without domain",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "fqdn with numeric tld",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host.example.123"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host.example.com"},
			WantErr: false,
		},
		{
			Name:    "valid hyphen inside label",
			Msg:     &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "my-host.ex-ample.com"},
			WantErr: false,
		},
		{
			Name:    "valid single character label",
			Msg:     &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "a.example.com"},
			WantErr: false,
		},
		{
			Name:    "valid 63 character label",
			Msg:     &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: label63 + ".example.com"},
			WantErr: false,
		},
		{
			Name:    "valid 253 characters total",
			Msg:     &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: fqdn253},
			WantErr: false,
		},
		{
			Name:             "leading hyphen in first label",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "-host.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "trailing hyphen in first label",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host-.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "leading hyphen in inner label",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "a.-b.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "empty label",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host..example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "trailing dot",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "host.example.com."},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "64 character label",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: label63 + "a.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "254 characters total",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: fqdn254},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
		{
			Name:             "underscore",
			Msg:              &mvmv1.LinuxDetails{LdapUuid: validation.TestUUID, LdapFqdn: "my_host.example.com"},
			WantErr:          true,
			WantErrorMessage: "validation error: ldap_fqdn: " + msgFqdn,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateWindowsDetails(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:             "invalid fqdn",
			Msg:              &mvmv1.WindowsDetails{DomainUuid: validation.TestUUID, DomainFqdn: "not a fqdn"},
			WantErr:          true,
			WantErrorMessage: "validation error: domain_fqdn: " + msgFqdn,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.WindowsDetails{DomainUuid: validation.TestUUID, DomainFqdn: "host.domain.example.com"},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateLinuxDisk(t *testing.T) {
	withDisk := func(modify func(*mvmv1.LinuxDisk)) *mvmv1.LinuxDisk {
		d := validLinuxDisk()
		modify(d)
		return d
	}

	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.LinuxDisk{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - name: must be within 2 and 128 characters
 - size_gb: ` + msgDiskSize,
		},
		{
			Name:    "valid system disk",
			Msg:     validLinuxDisk(),
			WantErr: false,
		},
		{
			Name: "valid data disk without mount point",
			Msg: withDisk(func(d *mvmv1.LinuxDisk) {
				d.MountPoint = nil
				d.DiskType = mvmv1.DiskType_DISK_TYPE_DATA.Enum()
				d.AutoExtend = proto.Bool(true)
			}),
			WantErr: false,
		},
		{
			Name:    "valid root mount point",
			Msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/") }),
			WantErr: false,
		},
		{
			Name:    "valid mount point with trailing slash and special characters",
			Msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/data/my+dir/") }),
			WantErr: false,
		},
		{
			Name:             "relative mount point",
			Msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("data/app") }),
			WantErr:          true,
			WantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			Name:             "empty mount point",
			Msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("") }),
			WantErr:          true,
			WantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			Name:             "mount point with whitespace",
			Msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.MountPoint = proto.String("/data/my dir") }),
			WantErr:          true,
			WantErrorMessage: "validation error: mount_point: " + msgMountPoint,
		},
		{
			Name:    "min disk size",
			Msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 1 }),
			WantErr: false,
		},
		{
			Name:    "max disk size",
			Msg:     withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 8192 }),
			WantErr: false,
		},
		{
			Name:             "disk too large",
			Msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.SizeGb = 8193 }),
			WantErr:          true,
			WantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
		{
			Name:             "invalid uuid",
			Msg:              withDisk(func(d *mvmv1.LinuxDisk) { d.Uuid = proto.String("not-a-uuid") }),
			WantErr:          true,
			WantErrorMessage: "validation error: uuid: must be a valid UUID",
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateWindowsDisk(t *testing.T) {
	withDisk := func(modify func(*mvmv1.WindowsDisk)) *mvmv1.WindowsDisk {
		d := validWindowsDisk()
		modify(d)
		return d
	}

	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.WindowsDisk{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - name: must be within 2 and 128 characters
 - size_gb: ` + msgDiskSize,
		},
		{
			Name:    "valid system disk",
			Msg:     validWindowsDisk(),
			WantErr: false,
		},
		{
			Name:             "lowercase drive letter",
			Msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("g") }),
			WantErr:          true,
			WantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			Name:    "reserved drive letter is valid in responses",
			Msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("R") }),
			WantErr: false,
		},
		{
			Name:    "valid os disk drive letter",
			Msg:     withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("C") }),
			WantErr: false,
		},
		{
			Name:             "two drive letters",
			Msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("FG") }),
			WantErr:          true,
			WantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			Name:             "drive letter with colon",
			Msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("F:") }),
			WantErr:          true,
			WantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			Name:             "digit as drive letter",
			Msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.DriveLetter = proto.String("1") }),
			WantErr:          true,
			WantErrorMessage: "validation error: drive_letter: " + msgDriveLetter,
		},
		{
			Name:             "disk too large",
			Msg:              withDisk(func(d *mvmv1.WindowsDisk) { d.SizeGb = 8193 }),
			WantErr:          true,
			WantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceGetRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.MVMServiceGetRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID`,
		},
		{
			Name:    "valid",
			Msg:     &mvmv1.MVMServiceGetRequest{Uuid: validation.TestUUID, ProjectUuid: validation.TestUUID},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceCreateWindowsRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.MVMServiceCreateWindowsRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - domain_uuid: value is empty, which is not a valid UUID
 - disks: must contain at least 1 item(s)`,
		},
		{
			Name: "invalid disk",
			Msg: &mvmv1.MVMServiceCreateWindowsRequest{
				DomainUuid: validation.TestUUID,
				Disks: []*mvmv1.WindowsDisk{{
					Name:        "system",
					SizeGb:      50,
					DriveLetter: proto.String("FG"),
				}},
			},
			WantErr:          true,
			WantErrorMessage: "validation error: disks[0].drive_letter: " + msgDriveLetter,
		},
		{
			Name: "valid",
			Msg: &mvmv1.MVMServiceCreateWindowsRequest{
				DomainUuid: validation.TestUUID,
				Disks:      []*mvmv1.WindowsDisk{validWindowsDisk()},
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceCreateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "valid linux",
			Msg:     validLinuxCreateRequest(),
			WantErr: false,
		},
		{
			Name: "valid windows",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.MvmType = &mvmv1.MVMServiceCreateRequest_Windows{
					Windows: &mvmv1.MVMServiceCreateWindowsRequest{
						DomainUuid: validation.TestUUID,
						Disks:      []*mvmv1.WindowsDisk{validWindowsDisk()},
					},
				}
			}),
			WantErr: false,
		},
		{
			Name: "valid bounds",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Cpu = 64
				r.RamGb = 512
			}),
			WantErr: false,
		},
		{
			Name:             "no cpu",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.Cpu = 0 }),
			WantErr:          true,
			WantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			Name:             "too many cpus",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.Cpu = 65 }),
			WantErr:          true,
			WantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			Name:             "too little ram",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.RamGb = 3 }),
			WantErr:          true,
			WantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			Name:             "too much ram",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.RamGb = 513 }),
			WantErr:          true,
			WantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			Name:             "invalid name code",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.NameCode = "ab-1" }),
			WantErr:          true,
			WantErrorMessage: "validation error: name_code: " + msgNameCode,
		},
		{
			Name:             "missing order reference",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderReference = "" }),
			WantErr:          true,
			WantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
		{
			Name:             "untrimmed order reference",
			Msg:              validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) { r.OrderReference = " TICKET-1 " }),
			WantErr:          true,
			WantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
		{
			Name: "no disks",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = nil
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: linux.disks: must contain at least 1 item(s)",
		},
		{
			Name: "max disks",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = linuxDisks(8)
			}),
			WantErr: false,
		},
		{
			Name: "too many disks",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks = linuxDisks(9)
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: linux.disks: " + msgDiskCount,
		},
		{
			Name: "invalid disk",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.GetLinux().Disks[0].SizeGb = 0
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: linux.disks[0].size_gb: " + msgDiskSize,
		},
		{
			Name: "missing initial ip address",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.InitialIpAddress = nil
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: initial_ip_address: value is required",
		},
		{
			Name: "missing backup configuration",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Backup = nil
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: backup: value is required",
		},
		{
			Name: "missing encryption configuration",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Encryption = nil
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: encryption: value is required",
		},
		{
			Name: "unspecified availability",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.Availability = mvmv1.Availability_AVAILABILITY_UNSPECIFIED
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: availability: " + msgNotUnspecified,
		},
		{
			Name: "unspecified service class",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.ServiceClass = mvmv1.ServiceClass_SERVICE_CLASS_UNSPECIFIED
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: service_class: " + msgNotUnspecified,
		},
		{
			Name: "missing os type",
			Msg: validLinuxCreateRequest(func(r *mvmv1.MVMServiceCreateRequest) {
				r.MvmType = nil
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: mvm_type: exactly one field is required in oneof",
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceUpdateRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "valid without changes",
			Msg:     validUpdateRequest(),
			WantErr: false,
		},
		{
			Name: "valid cpu and ram",
			Msg: validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) {
				r.Cpu = proto.Uint32(4)
				r.RamGb = proto.Uint32(16)
			}),
			WantErr: false,
		},
		{
			Name:             "too many cpus",
			Msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.Cpu = proto.Uint32(65) }),
			WantErr:          true,
			WantErrorMessage: "validation error: cpu: " + msgCpuCount,
		},
		{
			Name:             "zero ram",
			Msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.RamGb = proto.Uint32(0) }),
			WantErr:          true,
			WantErrorMessage: "validation error: ram_gb: " + msgRamSize,
		},
		{
			Name:             "missing order reference",
			Msg:              validUpdateRequest(func(r *mvmv1.MVMServiceUpdateRequest) { r.OrderReference = "" }),
			WantErr:          true,
			WantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
	}

	validation.ValidateProtos(t, tests)
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
			Uuid:           validation.TestUUID,
			ProjectUuid:    validation.TestUUID,
			OrderReference: "TICKET-1",
			Disks:          d,
		}
	}

	tests := validation.Prototests{
		{
			Name:    "valid",
			Msg:     request(disks(1)),
			WantErr: false,
		},
		{
			Name:             "no disks",
			Msg:              request(nil),
			WantErr:          true,
			WantErrorMessage: "validation error: disks: must contain at least 1 item(s)",
		},
		{
			Name:             "too many disks",
			Msg:              request(disks(9)),
			WantErr:          true,
			WantErrorMessage: "validation error: disks: " + msgDiskCount,
		},
		{
			Name:             "disk without os type",
			Msg:              request([]*mvmv1.Disk{{}}),
			WantErr:          true,
			WantErrorMessage: "validation error: disks[0].disk: exactly one field is required in oneof",
		},
		{
			Name: "invalid disk",
			Msg: request([]*mvmv1.Disk{{Disk: &mvmv1.Disk_Linux{Linux: &mvmv1.LinuxDisk{
				Name:       "data",
				SizeGb:     10,
				MountPoint: proto.String("data"),
			}}}}),
			WantErr:          true,
			WantErrorMessage: "validation error: disks[0].linux.mount_point: " + msgMountPoint,
		},
		{
			Name:    "valid windows drive letter",
			Msg:     request(windowsDisk("G")),
			WantErr: false,
		},
		{
			Name:    "windows disk without drive letter",
			Msg:     request(windowsDisk("")),
			WantErr: false,
		},
		{
			Name:             "windows drive letter before F",
			Msg:              request(windowsDisk("E")),
			WantErr:          true,
			WantErrorMessage: "validation error: disks: " + msgAddDriveLetter,
		},
		{
			Name:             "windows drive letter R",
			Msg:              request(windowsDisk("R")),
			WantErr:          true,
			WantErrorMessage: "validation error: disks: " + msgAddDriveLetter,
		},
		{
			Name:    "windows lowercase drive letter",
			Msg:     request(windowsDisk("g")),
			WantErr: true,
			WantErrorMessage: `validation errors:
 - disks: ` + msgAddDriveLetter + `
 - disks[0].windows.drive_letter: ` + msgDriveLetter,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceUpdateDiskRequest(t *testing.T) {
	request := func(size *uint64) *mvmv1.MVMServiceUpdateDiskRequest {
		return &mvmv1.MVMServiceUpdateDiskRequest{
			Uuid:           validation.TestUUID,
			DiskUuid:       validation.TestUUID,
			ProjectUuid:    validation.TestUUID,
			OrderReference: "TICKET-1",
			SizeGb:         size,
		}
	}

	tests := validation.Prototests{
		{
			Name:    "valid without size",
			Msg:     request(nil),
			WantErr: false,
		},
		{
			Name:    "valid size",
			Msg:     request(proto.Uint64(100)),
			WantErr: false,
		},
		{
			Name:             "zero size",
			Msg:              request(proto.Uint64(0)),
			WantErr:          true,
			WantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
		{
			Name:             "too large",
			Msg:              request(proto.Uint64(8193)),
			WantErr:          true,
			WantErrorMessage: "validation error: size_gb: " + msgDiskSize,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceDeleteRequest(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "empty",
			Msg:     &mvmv1.MVMServiceDeleteRequest{},
			WantErr: true,
			WantErrorMessage: `validation errors:
 - uuid: value is empty, which is not a valid UUID
 - project_uuid: value is empty, which is not a valid UUID
 - order_reference: ` + msgOrderReference,
		},
		{
			Name: "order reference too short",
			Msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:           validation.TestUUID,
				ProjectUuid:    validation.TestUUID,
				OrderReference: "x",
			},
			WantErr:          true,
			WantErrorMessage: "validation error: order_reference: " + msgOrderReference,
		},
		{
			Name: "valid",
			Msg: &mvmv1.MVMServiceDeleteRequest{
				Uuid:           validation.TestUUID,
				ProjectUuid:    validation.TestUUID,
				OrderReference: "TICKET-1",
			},
			WantErr: false,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceAddSnapshotRequest(t *testing.T) {
	request := func(name string) *mvmv1.MVMServiceAddSnapshotRequest {
		return &mvmv1.MVMServiceAddSnapshotRequest{
			Uuid:           validation.TestUUID,
			ProjectUuid:    validation.TestUUID,
			OrderReference: "TICKET-1",
			Name:           name,
		}
	}

	tests := validation.Prototests{
		{
			Name:    "valid",
			Msg:     request("beforeUpgrade1"),
			WantErr: false,
		},
		{
			Name:             "name with special characters",
			Msg:              request("before-upgrade"),
			WantErr:          true,
			WantErrorMessage: "validation error: name: " + msgSnapshotName,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateSnapshot(t *testing.T) {
	tests := validation.Prototests{
		{
			Name:    "valid",
			Msg:     &mvmv1.Snapshot{Uuid: validation.TestUUID, Name: "snapshot1"},
			WantErr: false,
		},
		{
			Name:             "name with whitespace",
			Msg:              &mvmv1.Snapshot{Uuid: validation.TestUUID, Name: "my snapshot"},
			WantErr:          true,
			WantErrorMessage: "validation error: name: " + msgSnapshotName,
		},
	}

	validation.ValidateProtos(t, tests)
}

func TestValidateMVMServiceReinstallRequest(t *testing.T) {
	request := func(modify ...func(*mvmv1.MVMServiceReinstallRequest)) *mvmv1.MVMServiceReinstallRequest {
		req := &mvmv1.MVMServiceReinstallRequest{
			Uuid:           validation.TestUUID,
			ProjectUuid:    validation.TestUUID,
			OrderReference: "TICKET-1",
		}
		for _, m := range modify {
			m(req)
		}
		return req
	}

	tests := validation.Prototests{
		{
			Name:    "valid without changes keeps current configuration",
			Msg:     request(),
			WantErr: false,
		},
		{
			Name: "valid with disabled backup and enabled encryption",
			Msg: request(func(r *mvmv1.MVMServiceReinstallRequest) {
				r.OsUuid = proto.String(validation.TestUUID)
				r.Backup = &mvmv1.Backup{Enabled: false}
				r.Encryption = &mvmv1.Encryption{Enabled: true}
			}),
			WantErr: false,
		},
		{
			Name: "invalid os uuid",
			Msg: request(func(r *mvmv1.MVMServiceReinstallRequest) {
				r.OsUuid = proto.String("not-a-uuid")
			}),
			WantErr:          true,
			WantErrorMessage: "validation error: os_uuid: must be a valid UUID",
		},
	}

	validation.ValidateProtos(t, tests)
}
