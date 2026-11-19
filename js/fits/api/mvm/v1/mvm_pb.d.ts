import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta, UpdateLabels, UpdateMeta } from "../../v1/common_pb.js";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/mvm.proto.
 */
export declare const file_fits_api_mvm_v1_mvm: GenFile;
/**
 * Disk configuration for a Linux MVM. See DiskType for the kinds of disks
 * (os, system, data) and their lifecycle rules.
 *
 * @generated from message fits.api.mvm.v1.LinuxDisk
 */
export type LinuxDisk = Message<"fits.api.mvm.v1.LinuxDisk"> & {
    /**
     * UUID of the disk; set for existing disks, unset for new ones.
     *
     * @generated from field: optional string uuid = 1;
     */
    uuid?: string | undefined;
    /**
     * Label (name) of the disk.
     *
     * @generated from field: string label = 2;
     */
    label: string;
    /**
     * Whether the disk grows automatically when space runs low.
     *
     * @generated from field: bool auto_extend = 3;
     */
    autoExtend: boolean;
    /**
     * Size of the disk in GB; defaults to 50 when unset.
     *
     * @generated from field: optional uint64 size_gb = 4;
     */
    sizeGb?: bigint | undefined;
    /**
     * Mount point (e.g. "/data"). Fixed to "/appdata" for the system disk;
     * free-form for data disks.
     *
     * @generated from field: optional string mount_point = 5;
     */
    mountPoint?: string | undefined;
    /**
     * Kind of the disk; determines which operations are allowed on it.
     * Provided by the client in create/add requests and populated by the
     * API in responses.
     *
     * @generated from field: optional fits.api.mvm.v1.DiskType disk_type = 6;
     */
    diskType?: DiskType | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.LinuxDisk.
 * Use `create(LinuxDiskSchema)` to create a new message.
 */
export declare const LinuxDiskSchema: GenMessage<LinuxDisk>;
/**
 * Disk configuration for a Windows MVM. See DiskType for the kinds of
 * disks (os, system, data) and their lifecycle rules.
 *
 * @generated from message fits.api.mvm.v1.WindowsDisk
 */
export type WindowsDisk = Message<"fits.api.mvm.v1.WindowsDisk"> & {
    /**
     * UUID of the disk; set for existing disks, unset for new ones.
     *
     * @generated from field: optional string uuid = 1;
     */
    uuid?: string | undefined;
    /**
     * Label (name) of the disk.
     *
     * @generated from field: string label = 2;
     */
    label: string;
    /**
     * Whether the disk grows automatically when space runs low.
     *
     * @generated from field: bool auto_extend = 3;
     */
    autoExtend: boolean;
    /**
     * Size of the disk in GB; defaults to 50 when unset.
     *
     * @generated from field: optional uint64 size_gb = 4;
     */
    sizeGb?: bigint | undefined;
    /**
     * Drive letter (e.g. "D"). Fixed to "F" for the system disk; free-form
     * for data disks.
     *
     * @generated from field: optional string driveletter = 5;
     */
    driveletter?: string | undefined;
    /**
     * Kind of the disk; determines which operations are allowed on it.
     * Provided by the client in create/add requests and populated by the
     * API in responses.
     *
     * @generated from field: optional fits.api.mvm.v1.DiskType disk_type = 6;
     */
    diskType?: DiskType | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.WindowsDisk.
 * Use `create(WindowsDiskSchema)` to create a new message.
 */
export declare const WindowsDiskSchema: GenMessage<WindowsDisk>;
/**
 * Disk contains either a Windows or Linux disk.
 *
 * @generated from message fits.api.mvm.v1.Disk
 */
export type Disk = Message<"fits.api.mvm.v1.Disk"> & {
    /**
     * OS-specific disk specification; must match the MVM's OS type.
     *
     * @generated from oneof fits.api.mvm.v1.Disk.disk
     */
    disk: {
        /**
         * Linux data disks.
         *
         * @generated from field: fits.api.mvm.v1.LinuxDisk linux = 4;
         */
        value: LinuxDisk;
        case: "linux";
    } | {
        /**
         * Windows data disks.
         *
         * @generated from field: fits.api.mvm.v1.WindowsDisk windows = 5;
         */
        value: WindowsDisk;
        case: "windows";
    } | {
        case: undefined;
        value?: undefined;
    };
};
/**
 * Describes the message fits.api.mvm.v1.Disk.
 * Use `create(DiskSchema)` to create a new message.
 */
export declare const DiskSchema: GenMessage<Disk>;
/**
 * Linux-specific MVM details.
 *
 * @generated from message fits.api.mvm.v1.LinuxDetails
 */
export type LinuxDetails = Message<"fits.api.mvm.v1.LinuxDetails"> & {
    /**
     * Disks configured on the Linux MVM.
     *
     * @generated from field: repeated fits.api.mvm.v1.LinuxDisk disks = 1;
     */
    disks: LinuxDisk[];
    /**
     * UUID of the LDAP directory the MVM is integrated with.
     *
     * @generated from field: string ldap_uuid = 2;
     */
    ldapUuid: string;
    /**
     * Linux-specific FQDN, derived from the name at creation or generated by the backend.
     *
     * @generated from field: string ldap_fqdn = 3;
     */
    ldapFqdn: string;
};
/**
 * Describes the message fits.api.mvm.v1.LinuxDetails.
 * Use `create(LinuxDetailsSchema)` to create a new message.
 */
export declare const LinuxDetailsSchema: GenMessage<LinuxDetails>;
/**
 * Windows-specific MVM details.
 *
 * @generated from message fits.api.mvm.v1.WindowsDetails
 */
export type WindowsDetails = Message<"fits.api.mvm.v1.WindowsDetails"> & {
    /**
     * Disks configured on the Windows MVM.
     *
     * @generated from field: repeated fits.api.mvm.v1.WindowsDisk disks = 1;
     */
    disks: WindowsDisk[];
    /**
     * UUID of the Windows domain the MVM has joined.
     *
     * @generated from field: string domain_uuid = 2;
     */
    domainUuid: string;
    /**
     * Windows-specific FQDN, derived from the name at creation or generated by the backend.
     *
     * @generated from field: string domain_fqdn = 3;
     */
    domainFqdn: string;
};
/**
 * Describes the message fits.api.mvm.v1.WindowsDetails.
 * Use `create(WindowsDetailsSchema)` to create a new message.
 */
export declare const WindowsDetailsSchema: GenMessage<WindowsDetails>;
/**
 * An IPv4 address of a MVM. IPv4 only; IPv6 is not yet supported.
 *
 * @generated from message fits.api.mvm.v1.IPAddress
 */
export type IPAddress = Message<"fits.api.mvm.v1.IPAddress"> & {
    /**
     * UUID of this IP address.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * The IPv4 address.
     *
     * @generated from field: string ip_address = 2;
     */
    ipAddress: string;
    /**
     * The network interface this IP address is bound to.
     * Unset for additional IP addresses that share an existing interface such addresses can be deleted via DeleteIPAddress.
     *
     * @generated from field: optional fits.api.mvm.v1.NetworkInterface network_interface = 3;
     */
    networkInterface?: NetworkInterface | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.IPAddress.
 * Use `create(IPAddressSchema)` to create a new message.
 */
export declare const IPAddressSchema: GenMessage<IPAddress>;
/**
 * A network interface of a MVM. IPv4 only; IPv6 is not yet supported.
 *
 * @generated from message fits.api.mvm.v1.NetworkInterface
 */
export type NetworkInterface = Message<"fits.api.mvm.v1.NetworkInterface"> & {
    /**
     * UUID of the VLAN.
     *
     * @generated from field: string vlan_uuid = 1;
     */
    vlanUuid: string;
    /**
     * MAC address of this interface.
     *
     * TODO: add validation
     *
     * @generated from field: string macaddress = 2;
     */
    macaddress: string;
};
/**
 * Describes the message fits.api.mvm.v1.NetworkInterface.
 * Use `create(NetworkInterfaceSchema)` to create a new message.
 */
export declare const NetworkInterfaceSchema: GenMessage<NetworkInterface>;
/**
 * The InitialIPAddress needs to be provided at creation of a MVM.
 * It provides the UUID of the target VLAN for the MVM's first IP address.
 * The created MVM comes with an IPAddress on a network interface connected
 * to this VLAN.
 *
 * @generated from message fits.api.mvm.v1.InitialIPAddress
 */
export type InitialIPAddress = Message<"fits.api.mvm.v1.InitialIPAddress"> & {
    /**
     * UUID of the VLAN.
     *
     * @generated from field: string vlan_uuid = 1;
     */
    vlanUuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.InitialIPAddress.
 * Use `create(InitialIPAddressSchema)` to create a new message.
 */
export declare const InitialIPAddressSchema: GenMessage<InitialIPAddress>;
/**
 * The status of the MVM.
 *
 * @generated from message fits.api.mvm.v1.Status
 */
export type Status = Message<"fits.api.mvm.v1.Status"> & {
    /**
     * State of the MVM.
     *
     * @generated from field: fits.api.mvm.v1.State state = 1;
     */
    state: State;
    /**
     * Additional information about the current status.
     *
     * @generated from field: string info = 2;
     */
    info: string;
};
/**
 * Describes the message fits.api.mvm.v1.Status.
 * Use `create(StatusSchema)` to create a new message.
 */
export declare const StatusSchema: GenMessage<Status>;
/**
 * ManagedVM is a managed VM instance as returned by the API.
 *
 * @generated from message fits.api.mvm.v1.ManagedVM
 */
export type ManagedVM = Message<"fits.api.mvm.v1.ManagedVM"> & {
    /**
     * UUID of this MVM.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Metadata for this MVM.
     * TODO: dates are available in the nulink db but not passed through yet.
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * FQDN of the MVM, derived from the name at creation or generated by the backend.
     *
     * @generated from field: string fqdn = 3;
     */
    fqdn: string;
    /**
     * Tenant this MVM belongs to.
     *
     * @generated from field: string tenant = 4;
     */
    tenant: string;
    /**
     * UUID of the project this MVM belongs to.
     *
     * @generated from field: string project_uuid = 5;
     */
    projectUuid: string;
    /**
     * UUID of the OS image installed on the MVM; resolve the title via the OS endpoint.
     *
     * @generated from field: string os_uuid = 6;
     */
    osUuid: string;
    /**
     * UUID of the datacenter location.
     *
     * @generated from field: string location_uuid = 7;
     */
    locationUuid: string;
    /**
     * UUID of the contact person responsible for this MVM.
     *
     * @generated from field: string contact_uuid = 8;
     */
    contactUuid: string;
    /**
     * Number of CPU cores.
     *
     * @generated from field: uint32 cpu = 9;
     */
    cpu: number;
    /**
     * RAM in GB.
     *
     * @generated from field: uint32 ram_gb = 10;
     */
    ramGb: number;
    /**
     * Whether backup is enabled.
     *
     * @generated from field: bool backup = 11;
     */
    backup: boolean;
    /**
     * Whether the MVM is encrypted.
     *
     * @generated from field: bool encrypted = 12;
     */
    encrypted: boolean;
    /**
     * Status of the MVM.
     *
     * @generated from field: fits.api.mvm.v1.Status status = 13;
     */
    status?: Status | undefined;
    /**
     * Availability level.
     *
     * @generated from field: fits.api.mvm.v1.Availability availability = 14;
     */
    availability: Availability;
    /**
     * Service class.
     *
     * @generated from field: fits.api.mvm.v1.ServiceClass service_class = 15;
     */
    serviceClass: ServiceClass;
    /**
     * OS-specific configuration.
     *
     * @generated from oneof fits.api.mvm.v1.ManagedVM.details
     */
    details: {
        /**
         * Windows-specific settings (domain + disks).
         *
         * @generated from field: fits.api.mvm.v1.WindowsDetails windows_details = 16;
         */
        value: WindowsDetails;
        case: "windowsDetails";
    } | {
        /**
         * Linux-specific settings (LDAP + disks).
         *
         * @generated from field: fits.api.mvm.v1.LinuxDetails linux_details = 17;
         */
        value: LinuxDetails;
        case: "linuxDetails";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * IP addresses attached to the MVM (IPv4 only).
     *
     * @generated from field: repeated fits.api.mvm.v1.IPAddress ip_addresses = 18;
     */
    ipAddresses: IPAddress[];
    /**
     * Internal reference number provided by the user, e.g. a ticket ID.
     *
     * @generated from field: string order_number = 19;
     */
    orderNumber: string;
    /**
     * Snapshots of the MVM.
     *
     * @generated from field: repeated fits.api.mvm.v1.Snapshot snapshots = 20;
     */
    snapshots: Snapshot[];
};
/**
 * Describes the message fits.api.mvm.v1.ManagedVM.
 * Use `create(ManagedVMSchema)` to create a new message.
 */
export declare const ManagedVMSchema: GenMessage<ManagedVM>;
/**
 * A snapshot of a MVM.
 *
 * @generated from message fits.api.mvm.v1.Snapshot
 */
export type Snapshot = Message<"fits.api.mvm.v1.Snapshot"> & {
    /**
     * UUID of this snapshot.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Name of the snapshot.
     *
     * @generated from field: string name = 2;
     */
    name: string;
};
/**
 * Describes the message fits.api.mvm.v1.Snapshot.
 * Use `create(SnapshotSchema)` to create a new message.
 */
export declare const SnapshotSchema: GenMessage<Snapshot>;
/**
 * Request to get a MVM by UUID.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceGetRequest
 */
export type MVMServiceGetRequest = Message<"fits.api.mvm.v1.MVMServiceGetRequest"> & {
    /**
     * UUID of the MVM.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project this MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceGetRequest.
 * Use `create(MVMServiceGetRequestSchema)` to create a new message.
 */
export declare const MVMServiceGetRequestSchema: GenMessage<MVMServiceGetRequest>;
/**
 * Response containing the requested MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceGetResponse
 */
export type MVMServiceGetResponse = Message<"fits.api.mvm.v1.MVMServiceGetResponse"> & {
    /**
     * The MVM.
     *
     * @generated from field: fits.api.mvm.v1.ManagedVM mvm = 1;
     */
    mvm?: ManagedVM | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceGetResponse.
 * Use `create(MVMServiceGetResponseSchema)` to create a new message.
 */
export declare const MVMServiceGetResponseSchema: GenMessage<MVMServiceGetResponse>;
/**
 * Request to list MVMs.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceListRequest
 */
export type MVMServiceListRequest = Message<"fits.api.mvm.v1.MVMServiceListRequest"> & {
    /**
     * UUID of the project to list MVMs for.
     *
     * @generated from field: string project_uuid = 1;
     */
    projectUuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceListRequest.
 * Use `create(MVMServiceListRequestSchema)` to create a new message.
 */
export declare const MVMServiceListRequestSchema: GenMessage<MVMServiceListRequest>;
/**
 * Response containing the list of MVMs.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceListResponse
 */
export type MVMServiceListResponse = Message<"fits.api.mvm.v1.MVMServiceListResponse"> & {
    /**
     * The MVM instances.
     *
     * @generated from field: repeated fits.api.mvm.v1.ManagedVM mvms = 1;
     */
    mvms: ManagedVM[];
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceListResponse.
 * Use `create(MVMServiceListResponseSchema)` to create a new message.
 */
export declare const MVMServiceListResponseSchema: GenMessage<MVMServiceListResponse>;
/**
 * Windows-specific create request parts.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceCreateWindowsRequest
 */
export type MVMServiceCreateWindowsRequest = Message<"fits.api.mvm.v1.MVMServiceCreateWindowsRequest"> & {
    /**
     * UUID of the Windows domain the MVM should join.
     *
     * @generated from field: string domain_uuid = 1;
     */
    domainUuid: string;
    /**
     * Disks for the Windows MVM, excluding the OS disk; must contain exactly
     * one system disk plus any number of data disks (see DiskType).
     *
     * @generated from field: repeated fits.api.mvm.v1.WindowsDisk disks = 2;
     */
    disks: WindowsDisk[];
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceCreateWindowsRequest.
 * Use `create(MVMServiceCreateWindowsRequestSchema)` to create a new message.
 */
export declare const MVMServiceCreateWindowsRequestSchema: GenMessage<MVMServiceCreateWindowsRequest>;
/**
 * Linux-specific create request parts.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceCreateLinuxRequest
 */
export type MVMServiceCreateLinuxRequest = Message<"fits.api.mvm.v1.MVMServiceCreateLinuxRequest"> & {
    /**
     * UUID of the LDAP directory the MVM should integrate with.
     *
     * @generated from field: string ldap_uuid = 1;
     */
    ldapUuid: string;
    /**
     * Disks for the Linux MVM, excluding the OS disk; must contain exactly
     * one system disk plus any number of data disks (see DiskType).
     *
     * @generated from field: repeated fits.api.mvm.v1.LinuxDisk disks = 2;
     */
    disks: LinuxDisk[];
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceCreateLinuxRequest.
 * Use `create(MVMServiceCreateLinuxRequestSchema)` to create a new message.
 */
export declare const MVMServiceCreateLinuxRequestSchema: GenMessage<MVMServiceCreateLinuxRequest>;
/**
 * Request to create a new MVM instance.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceCreateRequest
 */
export type MVMServiceCreateRequest = Message<"fits.api.mvm.v1.MVMServiceCreateRequest"> & {
    /**
     * UUID of the project this MVM belongs to.
     *
     * @generated from field: string project_uuid = 1;
     */
    projectUuid: string;
    /**
     * Required 4-character code (digits and letters); the FQDN is derived from it.
     *
     * @generated from field: string name_code = 2;
     */
    nameCode: string;
    /**
     * UUID of the operating system image to install.
     *
     * @generated from field: string os_uuid = 3;
     */
    osUuid: string;
    /**
     * UUID of the datacenter location; the MVM is scheduled into one data center within it.
     *
     * @generated from field: string location_uuid = 4;
     */
    locationUuid: string;
    /**
     * UUID of the contact person responsible for this MVM.
     *
     * @generated from field: string contact_uuid = 5;
     */
    contactUuid: string;
    /**
     * Number of CPU cores.
     *
     * @generated from field: uint32 cpu = 6;
     */
    cpu: number;
    /**
     * RAM in GB. Must be a power of 2.
     *
     * @generated from field: uint32 ram_gb = 7;
     */
    ramGb: number;
    /**
     * Customer reference for this operation provided by the user, e.g. a ticket ID.
     *
     * @generated from field: string order_number = 8;
     */
    orderNumber: string;
    /**
     * Optional key-value labels for accounting purposes.
     * (e.g. contract number, transaction number, technical key, accounting key).
     *
     * @generated from field: fits.api.v1.Labels labels = 9;
     */
    labels?: Labels | undefined;
    /**
     * Whether backup is enabled.
     *
     * @generated from field: bool backup = 10;
     */
    backup: boolean;
    /**
     * Whether the MVM should be encrypted.
     *
     * @generated from field: bool encrypted = 11;
     */
    encrypted: boolean;
    /**
     * Availability level.
     *
     * @generated from field: fits.api.mvm.v1.Availability availability = 12;
     */
    availability: Availability;
    /**
     * Service class; only certain combinations with availability are valid.
     *
     * @generated from field: fits.api.mvm.v1.ServiceClass service_class = 13;
     */
    serviceClass: ServiceClass;
    /**
     * OS type and its configuration.
     *
     * @generated from oneof fits.api.mvm.v1.MVMServiceCreateRequest.mvmtype
     */
    mvmtype: {
        /**
         * Windows-specific settings (domain + disk).
         *
         * @generated from field: fits.api.mvm.v1.MVMServiceCreateWindowsRequest windows = 14;
         */
        value: MVMServiceCreateWindowsRequest;
        case: "windows";
    } | {
        /**
         * Linux-specific settings (LDAP + disk).
         *
         * @generated from field: fits.api.mvm.v1.MVMServiceCreateLinuxRequest linux = 15;
         */
        value: MVMServiceCreateLinuxRequest;
        case: "linux";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * Provides the target VLAN for the initial IP address.
     *
     * @generated from field: fits.api.mvm.v1.InitialIPAddress initial_ip_address = 16;
     */
    initialIpAddress?: InitialIPAddress | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceCreateRequest.
 * Use `create(MVMServiceCreateRequestSchema)` to create a new message.
 */
export declare const MVMServiceCreateRequestSchema: GenMessage<MVMServiceCreateRequest>;
/**
 * Response for MVM create requests.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceCreateResponse
 */
export type MVMServiceCreateResponse = Message<"fits.api.mvm.v1.MVMServiceCreateResponse"> & {
    /**
     * UUID of the created MVM.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceCreateResponse.
 * Use `create(MVMServiceCreateResponseSchema)` to create a new message.
 */
export declare const MVMServiceCreateResponseSchema: GenMessage<MVMServiceCreateResponse>;
/**
 * Request to validate a MVM create without creating anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateCreateRequest
 */
export type MVMServiceValidateCreateRequest = Message<"fits.api.mvm.v1.MVMServiceValidateCreateRequest"> & {
    /**
     * The create request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceCreateRequest create = 1;
     */
    create?: MVMServiceCreateRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateCreateRequest.
 * Use `create(MVMServiceValidateCreateRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateCreateRequestSchema: GenMessage<MVMServiceValidateCreateRequest>;
/**
 * Response for ValidateCreate; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateCreateResponse
 */
export type MVMServiceValidateCreateResponse = Message<"fits.api.mvm.v1.MVMServiceValidateCreateResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateCreateResponse.
 * Use `create(MVMServiceValidateCreateResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateCreateResponseSchema: GenMessage<MVMServiceValidateCreateResponse>;
/**
 * Request to trigger the update of an existing MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateRequest
 */
export type MVMServiceUpdateRequest = Message<"fits.api.mvm.v1.MVMServiceUpdateRequest"> & {
    /**
     * UUID of the MVM to update.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project this MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Timestamp and strategy for this update.
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 3;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Order reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
    /**
     * The labels to update.
     *
     * @generated from field: optional fits.api.v1.UpdateLabels labels = 5;
     */
    labels?: UpdateLabels | undefined;
    /**
     * UUID of the new contact; unset keeps the current value.
     *
     * @generated from field: optional string contact_uuid = 6;
     */
    contactUuid?: string | undefined;
    /**
     * Number of vCPUs; unset keeps the current value.
     *
     * @generated from field: optional uint32 cpu = 7;
     */
    cpu?: number | undefined;
    /**
     * RAM size in GB; unset keeps the current value.
     *
     * @generated from field: optional uint32 ram_gb = 8;
     */
    ramGb?: number | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateRequest.
 * Use `create(MVMServiceUpdateRequestSchema)` to create a new message.
 */
export declare const MVMServiceUpdateRequestSchema: GenMessage<MVMServiceUpdateRequest>;
/**
 * Response for Update; an empty response means the update was accepted.
 * The update is applied asynchronously; poll Get until the state settles.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateResponse
 */
export type MVMServiceUpdateResponse = Message<"fits.api.mvm.v1.MVMServiceUpdateResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateResponse.
 * Use `create(MVMServiceUpdateResponseSchema)` to create a new message.
 */
export declare const MVMServiceUpdateResponseSchema: GenMessage<MVMServiceUpdateResponse>;
/**
 * Request to validate a MVM update without updating anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateRequest
 */
export type MVMServiceValidateUpdateRequest = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateRequest"> & {
    /**
     * The update request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceUpdateRequest update = 1;
     */
    update?: MVMServiceUpdateRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateRequest.
 * Use `create(MVMServiceValidateUpdateRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateRequestSchema: GenMessage<MVMServiceValidateUpdateRequest>;
/**
 * Response for ValidateUpdate; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateResponse
 */
export type MVMServiceValidateUpdateResponse = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateResponse.
 * Use `create(MVMServiceValidateUpdateResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateResponseSchema: GenMessage<MVMServiceValidateUpdateResponse>;
/**
 * Deprecated: request for a service class update.
 * Will be optional fields of the MVMServiceUpdateRequest.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateServiceClassRequest
 */
export type MVMServiceUpdateServiceClassRequest = Message<"fits.api.mvm.v1.MVMServiceUpdateServiceClassRequest"> & {
    /**
     * UUID of the MVM to update service class for.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Target service class; unset keeps the current value.
     *
     * @generated from field: optional fits.api.mvm.v1.ServiceClass service_class = 3;
     */
    serviceClass?: ServiceClass | undefined;
    /**
     * Target availability level; unset keeps current value.
     *
     * @generated from field: optional fits.api.mvm.v1.Availability availability = 4;
     */
    availability?: Availability | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateServiceClassRequest.
 * Use `create(MVMServiceUpdateServiceClassRequestSchema)` to create a new message.
 */
export declare const MVMServiceUpdateServiceClassRequestSchema: GenMessage<MVMServiceUpdateServiceClassRequest>;
/**
 * Deprecated: response for a service class update.
 * An empty response means the update was accepted.
 * The update is applied asynchronously; poll Get until the state settles.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateServiceClassResponse
 */
export type MVMServiceUpdateServiceClassResponse = Message<"fits.api.mvm.v1.MVMServiceUpdateServiceClassResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateServiceClassResponse.
 * Use `create(MVMServiceUpdateServiceClassResponseSchema)` to create a new message.
 */
export declare const MVMServiceUpdateServiceClassResponseSchema: GenMessage<MVMServiceUpdateServiceClassResponse>;
/**
 * Deprecated: Request to validate a MVM service class update without updating anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassRequest
 */
export type MVMServiceValidateUpdateServiceClassRequest = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassRequest"> & {
    /**
     * The UpdateServiceClass request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceUpdateServiceClassRequest update_service_class = 1;
     */
    updateServiceClass?: MVMServiceUpdateServiceClassRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassRequest.
 * Use `create(MVMServiceValidateUpdateServiceClassRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateServiceClassRequestSchema: GenMessage<MVMServiceValidateUpdateServiceClassRequest>;
/**
 * Deprecated: Response for ValidateUpdateServiceClass; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassResponse
 */
export type MVMServiceValidateUpdateServiceClassResponse = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateServiceClassResponse.
 * Use `create(MVMServiceValidateUpdateServiceClassResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateServiceClassResponseSchema: GenMessage<MVMServiceValidateUpdateServiceClassResponse>;
/**
 * Request to cancel a MVM instance.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteRequest
 */
export type MVMServiceDeleteRequest = Message<"fits.api.mvm.v1.MVMServiceDeleteRequest"> & {
    /**
     * UUID of the MVM to delete.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteRequest.
 * Use `create(MVMServiceDeleteRequestSchema)` to create a new message.
 */
export declare const MVMServiceDeleteRequestSchema: GenMessage<MVMServiceDeleteRequest>;
/**
 * Response for Delete; an empty response means the cancellation was accepted.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteResponse
 */
export type MVMServiceDeleteResponse = Message<"fits.api.mvm.v1.MVMServiceDeleteResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteResponse.
 * Use `create(MVMServiceDeleteResponseSchema)` to create a new message.
 */
export declare const MVMServiceDeleteResponseSchema: GenMessage<MVMServiceDeleteResponse>;
/**
 * Request to order additional data disks for a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddDisksRequest
 */
export type MVMServiceAddDisksRequest = Message<"fits.api.mvm.v1.MVMServiceAddDisksRequest"> & {
    /**
     * UUID of the MVM to attach the disks to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
    /**
     * Disks to add.
     *
     * @generated from field: repeated fits.api.mvm.v1.Disk disks = 4;
     */
    disks: Disk[];
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddDisksRequest.
 * Use `create(MVMServiceAddDisksRequestSchema)` to create a new message.
 */
export declare const MVMServiceAddDisksRequestSchema: GenMessage<MVMServiceAddDisksRequest>;
/**
 * Response for AddDisks.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddDisksResponse
 */
export type MVMServiceAddDisksResponse = Message<"fits.api.mvm.v1.MVMServiceAddDisksResponse"> & {
    /**
     * UUIDs of the created disks; one per requested disk.
     *
     * @generated from field: repeated string uuids = 1;
     */
    uuids: string[];
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddDisksResponse.
 * Use `create(MVMServiceAddDisksResponseSchema)` to create a new message.
 */
export declare const MVMServiceAddDisksResponseSchema: GenMessage<MVMServiceAddDisksResponse>;
/**
 * Request to validate an AddDisks without ordering anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateAddDisksRequest
 */
export type MVMServiceValidateAddDisksRequest = Message<"fits.api.mvm.v1.MVMServiceValidateAddDisksRequest"> & {
    /**
     * The add disks request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceAddDisksRequest add_disks = 1;
     */
    addDisks?: MVMServiceAddDisksRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateAddDisksRequest.
 * Use `create(MVMServiceValidateAddDisksRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateAddDisksRequestSchema: GenMessage<MVMServiceValidateAddDisksRequest>;
/**
 * Response for ValidateAddDisks; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateAddDisksResponse
 */
export type MVMServiceValidateAddDisksResponse = Message<"fits.api.mvm.v1.MVMServiceValidateAddDisksResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateAddDisksResponse.
 * Use `create(MVMServiceValidateAddDisksResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateAddDisksResponseSchema: GenMessage<MVMServiceValidateAddDisksResponse>;
/**
 * Request to change an existing disk; only size and auto-extend are
 * mutable (see DiskType for applicability per disk kind).
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateDiskRequest
 */
export type MVMServiceUpdateDiskRequest = Message<"fits.api.mvm.v1.MVMServiceUpdateDiskRequest"> & {
    /**
     * UUID of the MVM the disk belongs to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the disk to change.
     *
     * @generated from field: string disk_uuid = 2;
     */
    diskUuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 3;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
    /**
     * New auto-extend behavior; unset keeps the current value.
     *
     * @generated from field: optional bool auto_extend = 5;
     */
    autoExtend?: boolean | undefined;
    /**
     * New size of the disk in GB; unset keeps the current value.
     *
     * @generated from field: optional uint64 size_gb = 6;
     */
    sizeGb?: bigint | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateDiskRequest.
 * Use `create(MVMServiceUpdateDiskRequestSchema)` to create a new message.
 */
export declare const MVMServiceUpdateDiskRequestSchema: GenMessage<MVMServiceUpdateDiskRequest>;
/**
 * Response for UpdateDisk.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceUpdateDiskResponse
 */
export type MVMServiceUpdateDiskResponse = Message<"fits.api.mvm.v1.MVMServiceUpdateDiskResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceUpdateDiskResponse.
 * Use `create(MVMServiceUpdateDiskResponseSchema)` to create a new message.
 */
export declare const MVMServiceUpdateDiskResponseSchema: GenMessage<MVMServiceUpdateDiskResponse>;
/**
 * Request to validate an UpdateDisk without changing anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateDiskRequest
 */
export type MVMServiceValidateUpdateDiskRequest = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateDiskRequest"> & {
    /**
     * The update disk request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceUpdateDiskRequest update_disk = 1;
     */
    updateDisk?: MVMServiceUpdateDiskRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateDiskRequest.
 * Use `create(MVMServiceValidateUpdateDiskRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateDiskRequestSchema: GenMessage<MVMServiceValidateUpdateDiskRequest>;
/**
 * Response for ValidateUpdateDisk; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateUpdateDiskResponse
 */
export type MVMServiceValidateUpdateDiskResponse = Message<"fits.api.mvm.v1.MVMServiceValidateUpdateDiskResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateUpdateDiskResponse.
 * Use `create(MVMServiceValidateUpdateDiskResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateUpdateDiskResponseSchema: GenMessage<MVMServiceValidateUpdateDiskResponse>;
/**
 * Request to cancel a data disk; OS and system disks are rejected by the
 * upstream.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteDiskRequest
 */
export type MVMServiceDeleteDiskRequest = Message<"fits.api.mvm.v1.MVMServiceDeleteDiskRequest"> & {
    /**
     * UUID of the MVM the disk belongs to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the disk to cancel.
     *
     * @generated from field: string disk_uuid = 2;
     */
    diskUuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 3;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteDiskRequest.
 * Use `create(MVMServiceDeleteDiskRequestSchema)` to create a new message.
 */
export declare const MVMServiceDeleteDiskRequestSchema: GenMessage<MVMServiceDeleteDiskRequest>;
/**
 * Response for DeleteDisk.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteDiskResponse
 */
export type MVMServiceDeleteDiskResponse = Message<"fits.api.mvm.v1.MVMServiceDeleteDiskResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteDiskResponse.
 * Use `create(MVMServiceDeleteDiskResponseSchema)` to create a new message.
 */
export declare const MVMServiceDeleteDiskResponseSchema: GenMessage<MVMServiceDeleteDiskResponse>;
/**
 * Request to order an additional IP address for a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddIPAddressRequest
 */
export type MVMServiceAddIPAddressRequest = Message<"fits.api.mvm.v1.MVMServiceAddIPAddressRequest"> & {
    /**
     * UUID of the MVM to attach the IP address to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddIPAddressRequest.
 * Use `create(MVMServiceAddIPAddressRequestSchema)` to create a new message.
 */
export declare const MVMServiceAddIPAddressRequestSchema: GenMessage<MVMServiceAddIPAddressRequest>;
/**
 * Response for AddIPAddress.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddIPAddressResponse
 */
export type MVMServiceAddIPAddressResponse = Message<"fits.api.mvm.v1.MVMServiceAddIPAddressResponse"> & {
    /**
     * UUID of the created IP address.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddIPAddressResponse.
 * Use `create(MVMServiceAddIPAddressResponseSchema)` to create a new message.
 */
export declare const MVMServiceAddIPAddressResponseSchema: GenMessage<MVMServiceAddIPAddressResponse>;
/**
 * Request to move an IP address from one MVM to another.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceMoveIPAddressRequest
 */
export type MVMServiceMoveIPAddressRequest = Message<"fits.api.mvm.v1.MVMServiceMoveIPAddressRequest"> & {
    /**
     * UUID of the MVM the IP address currently belongs to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the IP address to move.
     *
     * @generated from field: string ip_address_uuid = 2;
     */
    ipAddressUuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 3;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
    /**
     * UUID of the MVM to move the IP address to.
     *
     * @generated from field: string target_mvm_uuid = 5;
     */
    targetMvmUuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceMoveIPAddressRequest.
 * Use `create(MVMServiceMoveIPAddressRequestSchema)` to create a new message.
 */
export declare const MVMServiceMoveIPAddressRequestSchema: GenMessage<MVMServiceMoveIPAddressRequest>;
/**
 * Response for MoveIPAddress.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceMoveIPAddressResponse
 */
export type MVMServiceMoveIPAddressResponse = Message<"fits.api.mvm.v1.MVMServiceMoveIPAddressResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceMoveIPAddressResponse.
 * Use `create(MVMServiceMoveIPAddressResponseSchema)` to create a new message.
 */
export declare const MVMServiceMoveIPAddressResponseSchema: GenMessage<MVMServiceMoveIPAddressResponse>;
/**
 * Request to delete an IP address from a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteIPAddressRequest
 */
export type MVMServiceDeleteIPAddressRequest = Message<"fits.api.mvm.v1.MVMServiceDeleteIPAddressRequest"> & {
    /**
     * UUID of the MVM the IP address belongs to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the IP address to delete.
     *
     * @generated from field: string ip_address_uuid = 2;
     */
    ipAddressUuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 3;
     */
    projectUuid: string;
    /**
     * Order reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteIPAddressRequest.
 * Use `create(MVMServiceDeleteIPAddressRequestSchema)` to create a new message.
 */
export declare const MVMServiceDeleteIPAddressRequestSchema: GenMessage<MVMServiceDeleteIPAddressRequest>;
/**
 * Response for DeleteIPAddress.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteIPAddressResponse
 */
export type MVMServiceDeleteIPAddressResponse = Message<"fits.api.mvm.v1.MVMServiceDeleteIPAddressResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteIPAddressResponse.
 * Use `create(MVMServiceDeleteIPAddressResponseSchema)` to create a new message.
 */
export declare const MVMServiceDeleteIPAddressResponseSchema: GenMessage<MVMServiceDeleteIPAddressResponse>;
/**
 * Request to restart a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceRestartRequest
 */
export type MVMServiceRestartRequest = Message<"fits.api.mvm.v1.MVMServiceRestartRequest"> & {
    /**
     * UUID of the MVM to restart.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceRestartRequest.
 * Use `create(MVMServiceRestartRequestSchema)` to create a new message.
 */
export declare const MVMServiceRestartRequestSchema: GenMessage<MVMServiceRestartRequest>;
/**
 * Response for Restart; an empty response means the restart was accepted.
 * The restart is applied asynchronously; poll Get until the state settles.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceRestartResponse
 */
export type MVMServiceRestartResponse = Message<"fits.api.mvm.v1.MVMServiceRestartResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceRestartResponse.
 * Use `create(MVMServiceRestartResponseSchema)` to create a new message.
 */
export declare const MVMServiceRestartResponseSchema: GenMessage<MVMServiceRestartResponse>;
/**
 * Request to reinstall the operating system of a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceReinstallRequest
 */
export type MVMServiceReinstallRequest = Message<"fits.api.mvm.v1.MVMServiceReinstallRequest"> & {
    /**
     * UUID of the MVM to reinstall.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
    /**
     * UUID of the operating system image to install.
     *
     * @generated from field: optional string os_uuid = 4;
     */
    osUuid?: string | undefined;
    /**
     * Whether backup is enabled after the reinstall; unset keeps the current value.
     *
     * @generated from field: optional bool backup = 5;
     */
    backup?: boolean | undefined;
    /**
     * Whether the MVM is encrypted after the reinstall; unset keeps the current value.
     *
     * @generated from field: optional bool encrypted = 6;
     */
    encrypted?: boolean | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceReinstallRequest.
 * Use `create(MVMServiceReinstallRequestSchema)` to create a new message.
 */
export declare const MVMServiceReinstallRequestSchema: GenMessage<MVMServiceReinstallRequest>;
/**
 * Response for Reinstall; an empty response means the reinstall was accepted.
 * The reinstall is applied asynchronously; poll Get until the state settles.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceReinstallResponse
 */
export type MVMServiceReinstallResponse = Message<"fits.api.mvm.v1.MVMServiceReinstallResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceReinstallResponse.
 * Use `create(MVMServiceReinstallResponseSchema)` to create a new message.
 */
export declare const MVMServiceReinstallResponseSchema: GenMessage<MVMServiceReinstallResponse>;
/**
 * Request to validate a Reinstall without reinstalling anything.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateReinstallRequest
 */
export type MVMServiceValidateReinstallRequest = Message<"fits.api.mvm.v1.MVMServiceValidateReinstallRequest"> & {
    /**
     * The reinstall request to validate.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceReinstallRequest reinstall = 1;
     */
    reinstall?: MVMServiceReinstallRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateReinstallRequest.
 * Use `create(MVMServiceValidateReinstallRequestSchema)` to create a new message.
 */
export declare const MVMServiceValidateReinstallRequestSchema: GenMessage<MVMServiceValidateReinstallRequest>;
/**
 * Response for ValidateReinstall; violations are returned as an InvalidArgument error.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceValidateReinstallResponse
 */
export type MVMServiceValidateReinstallResponse = Message<"fits.api.mvm.v1.MVMServiceValidateReinstallResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceValidateReinstallResponse.
 * Use `create(MVMServiceValidateReinstallResponseSchema)` to create a new message.
 */
export declare const MVMServiceValidateReinstallResponseSchema: GenMessage<MVMServiceValidateReinstallResponse>;
/**
 * Request to create a snapshot of a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddSnapshotRequest
 */
export type MVMServiceAddSnapshotRequest = Message<"fits.api.mvm.v1.MVMServiceAddSnapshotRequest"> & {
    /**
     * UUID of the MVM to create the snapshot for.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
    /**
     * Name of the snapshot; alphanumeric only.
     *
     * @generated from field: string name = 4;
     */
    name: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddSnapshotRequest.
 * Use `create(MVMServiceAddSnapshotRequestSchema)` to create a new message.
 */
export declare const MVMServiceAddSnapshotRequestSchema: GenMessage<MVMServiceAddSnapshotRequest>;
/**
 * Response for AddSnapshot.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceAddSnapshotResponse
 */
export type MVMServiceAddSnapshotResponse = Message<"fits.api.mvm.v1.MVMServiceAddSnapshotResponse"> & {
    /**
     * UUID of the created snapshot.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceAddSnapshotResponse.
 * Use `create(MVMServiceAddSnapshotResponseSchema)` to create a new message.
 */
export declare const MVMServiceAddSnapshotResponseSchema: GenMessage<MVMServiceAddSnapshotResponse>;
/**
 * Request to delete a snapshot of a MVM.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteSnapshotRequest
 */
export type MVMServiceDeleteSnapshotRequest = Message<"fits.api.mvm.v1.MVMServiceDeleteSnapshotRequest"> & {
    /**
     * UUID of the MVM the snapshot belongs to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the snapshot to delete.
     *
     * @generated from field: string snapshot_uuid = 2;
     */
    snapshotUuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 3;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 4;
     */
    orderNumber: string;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteSnapshotRequest.
 * Use `create(MVMServiceDeleteSnapshotRequestSchema)` to create a new message.
 */
export declare const MVMServiceDeleteSnapshotRequestSchema: GenMessage<MVMServiceDeleteSnapshotRequest>;
/**
 * Response for DeleteSnapshot; an empty response means the deletion was accepted.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceDeleteSnapshotResponse
 */
export type MVMServiceDeleteSnapshotResponse = Message<"fits.api.mvm.v1.MVMServiceDeleteSnapshotResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceDeleteSnapshotResponse.
 * Use `create(MVMServiceDeleteSnapshotResponseSchema)` to create a new message.
 */
export declare const MVMServiceDeleteSnapshotResponseSchema: GenMessage<MVMServiceDeleteSnapshotResponse>;
/**
 * Request to move a MVM from one datacenter to another.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceMoveRequest
 */
export type MVMServiceMoveRequest = Message<"fits.api.mvm.v1.MVMServiceMoveRequest"> & {
    /**
     * UUID of the MVM to move.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UUID of the project the MVM belongs to.
     *
     * @generated from field: string project_uuid = 2;
     */
    projectUuid: string;
    /**
     * Customer reference for this operation.
     *
     * @generated from field: string order_number = 3;
     */
    orderNumber: string;
    /**
     * If not provided, another datacenter within the same location will be chosen automatically.
     *
     * @generated from field: optional string location_uuid = 4;
     */
    locationUuid?: string | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceMoveRequest.
 * Use `create(MVMServiceMoveRequestSchema)` to create a new message.
 */
export declare const MVMServiceMoveRequestSchema: GenMessage<MVMServiceMoveRequest>;
/**
 * Response for Move; an empty response means the move was accepted.
 * The move is applied asynchronously; poll Get until the state settles.
 *
 * @generated from message fits.api.mvm.v1.MVMServiceMoveResponse
 */
export type MVMServiceMoveResponse = Message<"fits.api.mvm.v1.MVMServiceMoveResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.MVMServiceMoveResponse.
 * Use `create(MVMServiceMoveResponseSchema)` to create a new message.
 */
export declare const MVMServiceMoveResponseSchema: GenMessage<MVMServiceMoveResponse>;
/**
 * Availability specifies the availability level of a MVM; higher value = higher availability.
 *
 * @generated from enum fits.api.mvm.v1.Availability
 */
export declare enum Availability {
    /**
     * AVAILABILITY_UNSPECIFIED is not specified.
     *
     * @generated from enum value: AVAILABILITY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * AVAILABILITY_V0 is availability level V0.
     *
     * @generated from enum value: AVAILABILITY_V0 = 1;
     */
    V0 = 1,
    /**
     * AVAILABILITY_V1 is availability level V1.
     *
     * @generated from enum value: AVAILABILITY_V1 = 2;
     */
    V1 = 2,
    /**
     * AVAILABILITY_V2 is availability level V2.
     *
     * @generated from enum value: AVAILABILITY_V2 = 3;
     */
    V2 = 3,
    /**
     * AVAILABILITY_V3 is availability level V3.
     *
     * @generated from enum value: AVAILABILITY_V3 = 4;
     */
    V3 = 4
}
/**
 * Describes the enum fits.api.mvm.v1.Availability.
 */
export declare const AvailabilitySchema: GenEnum<Availability>;
/**
 * ServiceClass specifies the service tier of a MVM; higher value = higher tier.
 * Only certain combinations with Availability are valid.
 *
 * @generated from enum fits.api.mvm.v1.ServiceClass
 */
export declare enum ServiceClass {
    /**
     * SERVICE_CLASS_UNSPECIFIED is not specified.
     *
     * @generated from enum value: SERVICE_CLASS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SERVICE_CLASS_SZ1 is service class SZ1.
     *
     * @generated from enum value: SERVICE_CLASS_SZ1 = 1;
     */
    SZ1 = 1,
    /**
     * SERVICE_CLASS_SZ2 is service class SZ2.
     *
     * @generated from enum value: SERVICE_CLASS_SZ2 = 2;
     */
    SZ2 = 2,
    /**
     * SERVICE_CLASS_SZ3 is service class SZ3.
     *
     * @generated from enum value: SERVICE_CLASS_SZ3 = 3;
     */
    SZ3 = 3
}
/**
 * Describes the enum fits.api.mvm.v1.ServiceClass.
 */
export declare const ServiceClassSchema: GenEnum<ServiceClass>;
/**
 * State specifies the current state of a MVM.
 *
 * @generated from enum fits.api.mvm.v1.State
 */
export declare enum State {
    /**
     * STATE_UNSPECIFIED is not specified.
     *
     * @generated from enum value: STATE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * STATE_CREATING: the MVM hostname is reserved and provisioning has started.
     *
     * @generated from enum value: STATE_CREATING = 1;
     */
    CREATING = 1,
    /**
     * STATE_ACTIVE: the MVM is active.
     *
     * @generated from enum value: STATE_ACTIVE = 2;
     */
    ACTIVE = 2,
    /**
     * STATE_UPDATING: there is an ongoing operation for the MVM.
     *
     * @generated from enum value: STATE_UPDATING = 3;
     */
    UPDATING = 3,
    /**
     * STATE_DELETED: the MVM is deleted.
     *
     * @generated from enum value: STATE_DELETED = 4;
     */
    DELETED = 4,
    /**
     * STATE_ERROR: the exact error status is derived from the request type.
     *
     * @generated from enum value: STATE_ERROR = 5;
     */
    ERROR = 5
}
/**
 * Describes the enum fits.api.mvm.v1.State.
 */
export declare const StateSchema: GenEnum<State>;
/**
 * DiskType classifies a MVM disk by its lifecycle rules.
 * The upstream (nulink) API returns every disk of a MVM, including the
 * backend-provisioned OS disk, so that oversized OS disks stay billable.
 *
 * Lifecycle per kind:
 *   - OS: provisioned by the backend and never part of create/add
 *     requests; resizable via UpdateDisk, never deletable.
 *   - SYSTEM: exactly one per MVM (drive letter F on Windows, mount point
 *     /appdata on Linux), whose letter/mount point is fixed; must be
 *     included in the create request; resizable via UpdateDisk, never
 *     deletable.
 *   - DATA: fully controllable (letter/mount point, size, auto-extend)
 *     and deletable; may be included at creation or added later via
 *     AddDisks.
 *
 * Whether auto-extend is supported for OS and SYSTEM disks is not yet
 * known; only DATA disks are documented as supporting it.
 *
 * These rules are enforced by the upstream (nulink) API at creation and
 * edit time (e.g. duplicate drive letters, deleting a system disk);
 * fco-api does not re-validate them locally and surfaces upstream
 * violations as InvalidArgument errors. Use the Validate* RPCs to
 * pre-check a request against the upstream rules.
 *
 * @generated from enum fits.api.mvm.v1.DiskType
 */
export declare enum DiskType {
    /**
     * DISK_TYPE_UNSPECIFIED is not specified.
     *
     * @generated from enum value: DISK_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * DISK_TYPE_OS is the operating system disk; provisioned by the backend,
     * resizable, never deletable.
     *
     * @generated from enum value: DISK_TYPE_OS = 1;
     */
    OS = 1,
    /**
     * DISK_TYPE_SYSTEM is the additional system disk (volume F on Windows,
     * /appdata on Linux); the letter/mount point is fixed, the size is
     * resizable, never deletable.
     *
     * @generated from enum value: DISK_TYPE_SYSTEM = 2;
     */
    SYSTEM = 2,
    /**
     * DISK_TYPE_DATA is a user-managed data disk; fully controllable and
     * deletable.
     *
     * @generated from enum value: DISK_TYPE_DATA = 3;
     */
    DATA = 3
}
/**
 * Describes the enum fits.api.mvm.v1.DiskType.
 */
export declare const DiskTypeSchema: GenEnum<DiskType>;
/**
 * MVMService provides managed VM (MVM) operations.
 *
 * @generated from service fits.api.mvm.v1.MVMService
 */
export declare const MVMService: GenService<{
    /**
     * Creates a new MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof MVMServiceCreateRequestSchema;
        output: typeof MVMServiceCreateResponseSchema;
    };
    /**
     * Validates a MVM create request without creating anything (dry-run).
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateCreate
     */
    validateCreate: {
        methodKind: "unary";
        input: typeof MVMServiceValidateCreateRequestSchema;
        output: typeof MVMServiceValidateCreateResponseSchema;
    };
    /**
     * Returns the MVM with the specified UUID.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof MVMServiceGetRequestSchema;
        output: typeof MVMServiceGetResponseSchema;
    };
    /**
     * Returns the list of all MVMs.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MVMServiceListRequestSchema;
        output: typeof MVMServiceListResponseSchema;
    };
    /**
     * Updates a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof MVMServiceUpdateRequestSchema;
        output: typeof MVMServiceUpdateResponseSchema;
    };
    /**
     * Validates an Update request without changing anything.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateUpdate
     */
    validateUpdate: {
        methodKind: "unary";
        input: typeof MVMServiceValidateUpdateRequestSchema;
        output: typeof MVMServiceValidateUpdateResponseSchema;
    };
    /**
     * Deprecated: Updates the service class of a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.UpdateServiceClass
     */
    updateServiceClass: {
        methodKind: "unary";
        input: typeof MVMServiceUpdateServiceClassRequestSchema;
        output: typeof MVMServiceUpdateServiceClassResponseSchema;
    };
    /**
     * Deprecated: Validates an UpdateServiceClass request without changing anything.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateUpdateServiceClass
     */
    validateUpdateServiceClass: {
        methodKind: "unary";
        input: typeof MVMServiceValidateUpdateServiceClassRequestSchema;
        output: typeof MVMServiceValidateUpdateServiceClassResponseSchema;
    };
    /**
     * Deletes a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof MVMServiceDeleteRequestSchema;
        output: typeof MVMServiceDeleteResponseSchema;
    };
    /**
     * Orders additional data disks for a MVM. System disks are ordered at
     * MVM creation and cannot be added later.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.AddDisks
     */
    addDisks: {
        methodKind: "unary";
        input: typeof MVMServiceAddDisksRequestSchema;
        output: typeof MVMServiceAddDisksResponseSchema;
    };
    /**
     * Validates an AddDisks request without ordering anything.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateAddDisks
     */
    validateAddDisks: {
        methodKind: "unary";
        input: typeof MVMServiceValidateAddDisksRequestSchema;
        output: typeof MVMServiceValidateAddDisksResponseSchema;
    };
    /**
     * Changes the size or auto-extend behavior of an existing disk; see
     * DiskType for which operations apply per disk kind.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.UpdateDisk
     */
    updateDisk: {
        methodKind: "unary";
        input: typeof MVMServiceUpdateDiskRequestSchema;
        output: typeof MVMServiceUpdateDiskResponseSchema;
    };
    /**
     * Validates an UpdateDisk request without changing anything.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateUpdateDisk
     */
    validateUpdateDisk: {
        methodKind: "unary";
        input: typeof MVMServiceValidateUpdateDiskRequestSchema;
        output: typeof MVMServiceValidateUpdateDiskResponseSchema;
    };
    /**
     * Cancels a data disk. OS and system disks cannot be deleted.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.DeleteDisk
     */
    deleteDisk: {
        methodKind: "unary";
        input: typeof MVMServiceDeleteDiskRequestSchema;
        output: typeof MVMServiceDeleteDiskResponseSchema;
    };
    /**
     * Orders an additional IP address for a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.AddIPAddress
     */
    addIPAddress: {
        methodKind: "unary";
        input: typeof MVMServiceAddIPAddressRequestSchema;
        output: typeof MVMServiceAddIPAddressResponseSchema;
    };
    /**
     * Moves an IP address from one MVM to another.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.MoveIPAddress
     */
    moveIPAddress: {
        methodKind: "unary";
        input: typeof MVMServiceMoveIPAddressRequestSchema;
        output: typeof MVMServiceMoveIPAddressResponseSchema;
    };
    /**
     * Deletes an IP address from a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.DeleteIPAddress
     */
    deleteIPAddress: {
        methodKind: "unary";
        input: typeof MVMServiceDeleteIPAddressRequestSchema;
        output: typeof MVMServiceDeleteIPAddressResponseSchema;
    };
    /**
     * Restarts a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Restart
     */
    restart: {
        methodKind: "unary";
        input: typeof MVMServiceRestartRequestSchema;
        output: typeof MVMServiceRestartResponseSchema;
    };
    /**
     * Reinstalls the operating system of a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Reinstall
     */
    reinstall: {
        methodKind: "unary";
        input: typeof MVMServiceReinstallRequestSchema;
        output: typeof MVMServiceReinstallResponseSchema;
    };
    /**
     * Validates a Reinstall request without reinstalling anything.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.ValidateReinstall
     */
    validateReinstall: {
        methodKind: "unary";
        input: typeof MVMServiceValidateReinstallRequestSchema;
        output: typeof MVMServiceValidateReinstallResponseSchema;
    };
    /**
     * Creates a snapshot of a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.AddSnapshot
     */
    addSnapshot: {
        methodKind: "unary";
        input: typeof MVMServiceAddSnapshotRequestSchema;
        output: typeof MVMServiceAddSnapshotResponseSchema;
    };
    /**
     * Deletes a snapshot of a MVM.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.DeleteSnapshot
     */
    deleteSnapshot: {
        methodKind: "unary";
        input: typeof MVMServiceDeleteSnapshotRequestSchema;
        output: typeof MVMServiceDeleteSnapshotResponseSchema;
    };
    /**
     * Moves a MVM from one datacenter to another.
     *
     * @generated from rpc fits.api.mvm.v1.MVMService.Move
     */
    move: {
        methodKind: "unary";
        input: typeof MVMServiceMoveRequestSchema;
        output: typeof MVMServiceMoveResponseSchema;
    };
}>;
