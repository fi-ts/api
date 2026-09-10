import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Meta, UpdateMeta } from "../../v1/common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/vm/v1/vm.proto.
 */
export declare const file_fits_api_vm_v1_vm: GenFile;
/**
 * VM Instance Messages
 *
 * @generated from message fits.api.vm.v1.VMInstance
 */
export type VMInstance = Message<"fits.api.vm.v1.VMInstance"> & {
    /**
     * Uuid of this vm
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this vm
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Name of this vm
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Project where this vm address belongs to
     *
     * @generated from field: string project = 4;
     */
    project: string;
};
/**
 * Describes the message fits.api.vm.v1.VMInstance.
 * Use `create(VMInstanceSchema)` to create a new message.
 */
export declare const VMInstanceSchema: GenMessage<VMInstance>;
/**
 * VMServiceGetRequest TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceGetRequest
 */
export type VMServiceGetRequest = Message<"fits.api.vm.v1.VMServiceGetRequest"> & {
    /**
     * Uuid of this vm
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Project where this vm address belongs to
     *
     * @generated from field: string project = 2;
     */
    project: string;
};
/**
 * Describes the message fits.api.vm.v1.VMServiceGetRequest.
 * Use `create(VMServiceGetRequestSchema)` to create a new message.
 */
export declare const VMServiceGetRequestSchema: GenMessage<VMServiceGetRequest>;
/**
 * VMServiceGetResponse TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceGetResponse
 */
export type VMServiceGetResponse = Message<"fits.api.vm.v1.VMServiceGetResponse"> & {
    /**
     * The vm
     *
     * @generated from field: fits.api.vm.v1.VMInstance vm = 1;
     */
    vm?: VMInstance | undefined;
};
/**
 * Describes the message fits.api.vm.v1.VMServiceGetResponse.
 * Use `create(VMServiceGetResponseSchema)` to create a new message.
 */
export declare const VMServiceGetResponseSchema: GenMessage<VMServiceGetResponse>;
/**
 * VMServiceCreateRequest TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceCreateRequest
 */
export type VMServiceCreateRequest = Message<"fits.api.vm.v1.VMServiceCreateRequest"> & {
    /**
     * Project of the vm
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Name of the vm
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Number of CPUs of this vm
     *
     * @generated from field: uint32 cpu = 3;
     */
    cpu: number;
    /**
     * Ram of the vm in GB
     *
     * @generated from field: uint32 ram = 4;
     */
    ram: number;
    /**
     * OS Uuid of the OS to install in the vm instance
     *
     * @generated from field: string os_uuid = 5;
     */
    osUuid: string;
    /**
     * VLAN Uuid of the VLAN to install in the vm instance into
     *
     * @generated from field: string vlan_uuid = 6;
     */
    vlanUuid: string;
    /**
     * Location Uuid of the datacenter location to install in the vm instance
     *
     * @generated from field: string location_uuid = 7;
     */
    locationUuid: string;
    /**
     * Contact Uuid of who is the responsible contact of the vm instance
     *
     * @generated from field: string contact_uuid = 8;
     */
    contactUuid: string;
    /**
     * List of disks for this vm
     *
     * Disks             []DiskInputModel
     * RequestedBy       *string // TODO is inherited from the token making the create request -> stored in meta.createdBy
     *
     * @generated from field: repeated fits.api.vm.v1.Disk disks = 9;
     */
    disks: Disk[];
};
/**
 * Describes the message fits.api.vm.v1.VMServiceCreateRequest.
 * Use `create(VMServiceCreateRequestSchema)` to create a new message.
 */
export declare const VMServiceCreateRequestSchema: GenMessage<VMServiceCreateRequest>;
/**
 * Disk
 *
 * @generated from message fits.api.vm.v1.Disk
 */
export type Disk = Message<"fits.api.vm.v1.Disk"> & {
    /**
     * AutoExtend if set to true the disk grows automatically
     *
     * @generated from field: bool auto_extend = 1;
     */
    autoExtend: boolean;
    /**
     * Size if the disk in GB
     *
     * TODO discuss how this could be optional
     *
     * @generated from field: optional uint64 size_in_gb = 2;
     */
    sizeInGb?: bigint | undefined;
    /**
     * DriveLetter where this disk should be assigned
     *
     * @generated from field: optional string drive_letter = 3;
     */
    driveLetter?: string | undefined;
    /**
     * Label of the disk
     *
     * @generated from field: string label = 4;
     */
    label: string;
    /**
     * MountPoint where this disk should be mounted
     *
     * TODO this requires a Filesystem on the disk ?
     *
     * @generated from field: optional string mount_point = 5;
     */
    mountPoint?: string | undefined;
};
/**
 * Describes the message fits.api.vm.v1.Disk.
 * Use `create(DiskSchema)` to create a new message.
 */
export declare const DiskSchema: GenMessage<Disk>;
/**
 * VMServiceCreateResponse TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceCreateResponse
 */
export type VMServiceCreateResponse = Message<"fits.api.vm.v1.VMServiceCreateResponse"> & {};
/**
 * Describes the message fits.api.vm.v1.VMServiceCreateResponse.
 * Use `create(VMServiceCreateResponseSchema)` to create a new message.
 */
export declare const VMServiceCreateResponseSchema: GenMessage<VMServiceCreateResponse>;
/**
 * VMServiceUpdateRequest TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceUpdateRequest
 */
export type VMServiceUpdateRequest = Message<"fits.api.vm.v1.VMServiceUpdateRequest"> & {
    /**
     * Project of the vm
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
};
/**
 * Describes the message fits.api.vm.v1.VMServiceUpdateRequest.
 * Use `create(VMServiceUpdateRequestSchema)` to create a new message.
 */
export declare const VMServiceUpdateRequestSchema: GenMessage<VMServiceUpdateRequest>;
/**
 * VMServiceUpdateResponse TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceUpdateResponse
 */
export type VMServiceUpdateResponse = Message<"fits.api.vm.v1.VMServiceUpdateResponse"> & {};
/**
 * Describes the message fits.api.vm.v1.VMServiceUpdateResponse.
 * Use `create(VMServiceUpdateResponseSchema)` to create a new message.
 */
export declare const VMServiceUpdateResponseSchema: GenMessage<VMServiceUpdateResponse>;
/**
 * VMServiceListRequest TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceListRequest
 */
export type VMServiceListRequest = Message<"fits.api.vm.v1.VMServiceListRequest"> & {
    /**
     * Project of the vm
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message fits.api.vm.v1.VMServiceListRequest.
 * Use `create(VMServiceListRequestSchema)` to create a new message.
 */
export declare const VMServiceListRequestSchema: GenMessage<VMServiceListRequest>;
/**
 * VMServiceListResponse TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceListResponse
 */
export type VMServiceListResponse = Message<"fits.api.vm.v1.VMServiceListResponse"> & {};
/**
 * Describes the message fits.api.vm.v1.VMServiceListResponse.
 * Use `create(VMServiceListResponseSchema)` to create a new message.
 */
export declare const VMServiceListResponseSchema: GenMessage<VMServiceListResponse>;
/**
 * VMServiceDeleteRequest TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceDeleteRequest
 */
export type VMServiceDeleteRequest = Message<"fits.api.vm.v1.VMServiceDeleteRequest"> & {
    /**
     * Project of the vm
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message fits.api.vm.v1.VMServiceDeleteRequest.
 * Use `create(VMServiceDeleteRequestSchema)` to create a new message.
 */
export declare const VMServiceDeleteRequestSchema: GenMessage<VMServiceDeleteRequest>;
/**
 * VMServiceDeleteResponse TODO
 *
 * @generated from message fits.api.vm.v1.VMServiceDeleteResponse
 */
export type VMServiceDeleteResponse = Message<"fits.api.vm.v1.VMServiceDeleteResponse"> & {};
/**
 * Describes the message fits.api.vm.v1.VMServiceDeleteResponse.
 * Use `create(VMServiceDeleteResponseSchema)` to create a new message.
 */
export declare const VMServiceDeleteResponseSchema: GenMessage<VMServiceDeleteResponse>;
/**
 * VMService provides VM address management operations.
 *
 * @generated from service fits.api.vm.v1.VMService
 */
export declare const VMService: GenService<{
    /**
     * Returns the VM address with the specified VM.
     *
     * @generated from rpc fits.api.vm.v1.VMService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof VMServiceGetRequestSchema;
        output: typeof VMServiceGetResponseSchema;
    };
    /**
     * Creates a new VM address.
     *
     * @generated from rpc fits.api.vm.v1.VMService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof VMServiceCreateRequestSchema;
        output: typeof VMServiceCreateResponseSchema;
    };
    /**
     * Updates an VM address.
     *
     * @generated from rpc fits.api.vm.v1.VMService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof VMServiceUpdateRequestSchema;
        output: typeof VMServiceUpdateResponseSchema;
    };
    /**
     * Returns the list of all VM addresses.
     *
     * @generated from rpc fits.api.vm.v1.VMService.List
     */
    list: {
        methodKind: "unary";
        input: typeof VMServiceListRequestSchema;
        output: typeof VMServiceListResponseSchema;
    };
    /**
     * Deletes an VM address.
     *
     * @generated from rpc fits.api.vm.v1.VMService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof VMServiceDeleteRequestSchema;
        output: typeof VMServiceDeleteResponseSchema;
    };
}>;
