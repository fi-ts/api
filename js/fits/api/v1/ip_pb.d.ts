import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/ip.proto.
 */
export declare const file_fits_api_v1_ip: GenFile;
/**
 * IP is a ip address which can be used as loadbalancer addresses
 *
 * @generated from message fits.api.v1.IP
 */
export type IP = Message<"fits.api.v1.IP"> & {
    /**
     * Uuid of this ip
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this ip
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Ip is either ipv4 or ipv6 address
     *
     * @generated from field: string ip = 3;
     */
    ip: string;
    /**
     * Name of this ip
     *
     * @generated from field: string name = 4;
     */
    name: string;
    /**
     * Description of this ip
     *
     * @generated from field: string description = 5;
     */
    description: string;
    /**
     * Network is the network this ip belongs to
     *
     * @generated from field: string network = 6;
     */
    network: string;
    /**
     * Project where this ip address belongs to
     *
     * @generated from field: string project = 7;
     */
    project: string;
    /**
     * Type of this ip
     *
     * @generated from field: fits.api.v1.IPType type = 8;
     */
    type: IPType;
    /**
     * Namespace if specified this ip is from a namespaced network and can therefore overlap with others
     * Will be equal with project most of the time
     *
     * @generated from field: optional string namespace = 9;
     */
    namespace?: string | undefined;
};
/**
 * Describes the message fits.api.v1.IP.
 * Use `create(IPSchema)` to create a new message.
 */
export declare const IPSchema: GenMessage<IP>;
/**
 * IPServiceGetRequest is the request payload for a ip get request
 *
 * @generated from message fits.api.v1.IPServiceGetRequest
 */
export type IPServiceGetRequest = Message<"fits.api.v1.IPServiceGetRequest"> & {
    /**
     * IP of the ip to get
     *
     * @generated from field: string ip = 1;
     */
    ip: string;
    /**
     * Project of the ip
     *
     * @generated from field: string project = 2;
     */
    project: string;
    /**
     * Namespace can be specified to get the ip of a namespace.
     *
     * @generated from field: optional string namespace = 3;
     */
    namespace?: string | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceGetRequest.
 * Use `create(IPServiceGetRequestSchema)` to create a new message.
 */
export declare const IPServiceGetRequestSchema: GenMessage<IPServiceGetRequest>;
/**
 * IPServiceCreateRequest is the request payload for a ip create request
 *
 * @generated from message fits.api.v1.IPServiceCreateRequest
 */
export type IPServiceCreateRequest = Message<"fits.api.v1.IPServiceCreateRequest"> & {
    /**
     * Network from which the IP should be created
     *
     * @generated from field: string network = 1;
     */
    network: string;
    /**
     * Project of the ip
     *
     * @generated from field: string project = 2;
     */
    project: string;
    /**
     * Name of the ip
     *
     * @generated from field: optional string name = 3;
     */
    name?: string | undefined;
    /**
     * Description of the ip
     *
     * @generated from field: optional string description = 4;
     */
    description?: string | undefined;
    /**
     * IP if given try to create this ip if still available
     *
     * @generated from field: optional string ip = 5;
     */
    ip?: string | undefined;
    /**
     * Machine for which this ip should get created
     *
     * @generated from field: optional string machine = 6;
     */
    machine?: string | undefined;
    /**
     * Labels to put onto the ip
     *
     * @generated from field: optional fits.api.v1.Labels labels = 7;
     */
    labels?: Labels | undefined;
    /**
     * Type of the IP, ether ephemeral (default), or static
     *
     * @generated from field: optional fits.api.v1.IPType type = 8;
     */
    type?: IPType | undefined;
    /**
     * Addressfamily of the IP to create, defaults to ipv4
     *
     * @generated from field: optional fits.api.v1.IPAddressFamily address_family = 9;
     */
    addressFamily?: IPAddressFamily | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceCreateRequest.
 * Use `create(IPServiceCreateRequestSchema)` to create a new message.
 */
export declare const IPServiceCreateRequestSchema: GenMessage<IPServiceCreateRequest>;
/**
 * IPServiceUpdateRequest is the request payload for a ip update request
 *
 * @generated from message fits.api.v1.IPServiceUpdateRequest
 */
export type IPServiceUpdateRequest = Message<"fits.api.v1.IPServiceUpdateRequest"> & {
    /**
     * Ip the ip to update
     *
     * @generated from field: string ip = 1;
     */
    ip: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Project id of the ip
     *
     * @generated from field: string project = 3;
     */
    project: string;
    /**
     * Name of this ip
     *
     * @generated from field: optional string name = 4;
     */
    name?: string | undefined;
    /**
     * Description of this ip
     *
     * @generated from field: optional string description = 5;
     */
    description?: string | undefined;
    /**
     * Type of this ip
     *
     * @generated from field: optional fits.api.v1.IPType type = 6;
     */
    type?: IPType | undefined;
    /**
     * Labels on this ip
     *
     * @generated from field: optional fits.api.v1.UpdateLabels labels = 7;
     */
    labels?: UpdateLabels | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceUpdateRequest.
 * Use `create(IPServiceUpdateRequestSchema)` to create a new message.
 */
export declare const IPServiceUpdateRequestSchema: GenMessage<IPServiceUpdateRequest>;
/**
 * IPServiceListRequest is the request payload for a ip list request
 *
 * @generated from message fits.api.v1.IPServiceListRequest
 */
export type IPServiceListRequest = Message<"fits.api.v1.IPServiceListRequest"> & {
    /**
     * Project of the ips to list
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * Query to list one ore more ips
     *
     * @generated from field: fits.api.v1.IPQuery query = 2;
     */
    query?: IPQuery | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceListRequest.
 * Use `create(IPServiceListRequestSchema)` to create a new message.
 */
export declare const IPServiceListRequestSchema: GenMessage<IPServiceListRequest>;
/**
 * IPQuery can be used to query a IP or a list of IP
 *
 * @generated from message fits.api.v1.IPQuery
 */
export type IPQuery = Message<"fits.api.v1.IPQuery"> & {
    /**
     * Ip the ip to list
     *
     * @generated from field: optional string ip = 1;
     */
    ip?: string | undefined;
    /**
     * Network from where the ips to list
     *
     * @generated from field: optional string network = 2;
     */
    network?: string | undefined;
    /**
     * Project of the ips to list
     *
     * @generated from field: optional string project = 3;
     */
    project?: string | undefined;
    /**
     * Name of this ip
     *
     * @generated from field: optional string name = 4;
     */
    name?: string | undefined;
    /**
     * Uuid for which this ips should get filtered
     *
     * @generated from field: optional string uuid = 5;
     */
    uuid?: string | undefined;
    /**
     * Machine for which this ips should get filtered
     *
     * @generated from field: optional string machine = 6;
     */
    machine?: string | undefined;
    /**
     * ParentPrefixCidr for which this ips should get filtered
     *
     * @generated from field: optional string parent_prefix_cidr = 7;
     */
    parentPrefixCidr?: string | undefined;
    /**
     * Labels for which this ips should get filtered
     *
     * @generated from field: optional fits.api.v1.Labels labels = 8;
     */
    labels?: Labels | undefined;
    /**
     * Static if set to true, this will be a Static ip
     *
     * @generated from field: optional fits.api.v1.IPType type = 9;
     */
    type?: IPType | undefined;
    /**
     * Addressfamily of the IPs to list, defaults to all addressfamilies
     *
     * @generated from field: optional fits.api.v1.IPAddressFamily address_family = 10;
     */
    addressFamily?: IPAddressFamily | undefined;
};
/**
 * Describes the message fits.api.v1.IPQuery.
 * Use `create(IPQuerySchema)` to create a new message.
 */
export declare const IPQuerySchema: GenMessage<IPQuery>;
/**
 * IPServiceDeleteRequest is the request payload for a ip delete request
 *
 * @generated from message fits.api.v1.IPServiceDeleteRequest
 */
export type IPServiceDeleteRequest = Message<"fits.api.v1.IPServiceDeleteRequest"> & {
    /**
     * IP of the ip to delete
     *
     * @generated from field: string ip = 1;
     */
    ip: string;
    /**
     * Project of the ip
     *
     * @generated from field: string project = 2;
     */
    project: string;
};
/**
 * Describes the message fits.api.v1.IPServiceDeleteRequest.
 * Use `create(IPServiceDeleteRequestSchema)` to create a new message.
 */
export declare const IPServiceDeleteRequestSchema: GenMessage<IPServiceDeleteRequest>;
/**
 * IPServiceGetResponse is the response payload for a ip get request
 *
 * @generated from message fits.api.v1.IPServiceGetResponse
 */
export type IPServiceGetResponse = Message<"fits.api.v1.IPServiceGetResponse"> & {
    /**
     * Ip the ip
     *
     * @generated from field: fits.api.v1.IP ip = 1;
     */
    ip?: IP | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceGetResponse.
 * Use `create(IPServiceGetResponseSchema)` to create a new message.
 */
export declare const IPServiceGetResponseSchema: GenMessage<IPServiceGetResponse>;
/**
 * IPServiceUpdateResponse is the response payload for a ip update request
 *
 * @generated from message fits.api.v1.IPServiceUpdateResponse
 */
export type IPServiceUpdateResponse = Message<"fits.api.v1.IPServiceUpdateResponse"> & {
    /**
     * Ip the ip
     *
     * @generated from field: fits.api.v1.IP ip = 1;
     */
    ip?: IP | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceUpdateResponse.
 * Use `create(IPServiceUpdateResponseSchema)` to create a new message.
 */
export declare const IPServiceUpdateResponseSchema: GenMessage<IPServiceUpdateResponse>;
/**
 * IPServiceCreateResponse is the response payload for a ip create request
 *
 * @generated from message fits.api.v1.IPServiceCreateResponse
 */
export type IPServiceCreateResponse = Message<"fits.api.v1.IPServiceCreateResponse"> & {
    /**
     * Ip the ip
     *
     * @generated from field: fits.api.v1.IP ip = 1;
     */
    ip?: IP | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceCreateResponse.
 * Use `create(IPServiceCreateResponseSchema)` to create a new message.
 */
export declare const IPServiceCreateResponseSchema: GenMessage<IPServiceCreateResponse>;
/**
 * IPServiceListResponse is the response payload for a ip list request
 *
 * @generated from message fits.api.v1.IPServiceListResponse
 */
export type IPServiceListResponse = Message<"fits.api.v1.IPServiceListResponse"> & {
    /**
     * Ips the ips
     *
     * @generated from field: repeated fits.api.v1.IP ips = 1;
     */
    ips: IP[];
};
/**
 * Describes the message fits.api.v1.IPServiceListResponse.
 * Use `create(IPServiceListResponseSchema)` to create a new message.
 */
export declare const IPServiceListResponseSchema: GenMessage<IPServiceListResponse>;
/**
 * IPServiceDeleteResponse is the response payload for a ip delete request
 *
 * @generated from message fits.api.v1.IPServiceDeleteResponse
 */
export type IPServiceDeleteResponse = Message<"fits.api.v1.IPServiceDeleteResponse"> & {
    /**
     * Ip the ip
     *
     * @generated from field: fits.api.v1.IP ip = 1;
     */
    ip?: IP | undefined;
};
/**
 * Describes the message fits.api.v1.IPServiceDeleteResponse.
 * Use `create(IPServiceDeleteResponseSchema)` to create a new message.
 */
export declare const IPServiceDeleteResponseSchema: GenMessage<IPServiceDeleteResponse>;
/**
 * IPType specifies different ip address types
 *
 * @generated from enum fits.api.v1.IPType
 */
export declare enum IPType {
    /**
     * IP_TYPE_UNSPECIFIED is not specified
     *
     * @generated from enum value: IP_TYPE_UNSPECIFIED = 0;
     */
    IP_TYPE_UNSPECIFIED = 0,
    /**
     * IP_TYPE_EPHEMERAL defines a ephemeral ip address which is freed/deleted after usage
     *
     * @generated from enum value: IP_TYPE_EPHEMERAL = 1;
     */
    IP_TYPE_EPHEMERAL = 1,
    /**
     * IP_TYPE_STATIC defines a static ip address which must be freed/deleted explicitly
     *
     * @generated from enum value: IP_TYPE_STATIC = 2;
     */
    IP_TYPE_STATIC = 2
}
/**
 * Describes the enum fits.api.v1.IPType.
 */
export declare const IPTypeSchema: GenEnum<IPType>;
/**
 * IPAddressFamily defines either IPv4 or IPv6 Addressfamily
 *
 * @generated from enum fits.api.v1.IPAddressFamily
 */
export declare enum IPAddressFamily {
    /**
     * IP_ADDRESS_FAMILY_UNSPECIFIED is not specified
     *
     * @generated from enum value: IP_ADDRESS_FAMILY_UNSPECIFIED = 0;
     */
    IP_ADDRESS_FAMILY_UNSPECIFIED = 0,
    /**
     * IP_ADDRESS_FAMILY_V4 defines a IPv4 address
     *
     * @generated from enum value: IP_ADDRESS_FAMILY_V4 = 1;
     */
    IP_ADDRESS_FAMILY_V4 = 1,
    /**
     * IP_ADDRESS_FAMILY_V6 defines a IPv6 address
     *
     * @generated from enum value: IP_ADDRESS_FAMILY_V6 = 2;
     */
    IP_ADDRESS_FAMILY_V6 = 2
}
/**
 * Describes the enum fits.api.v1.IPAddressFamily.
 */
export declare const IPAddressFamilySchema: GenEnum<IPAddressFamily>;
/**
 * IPService provides IP address management operations.
 *
 * @generated from service fits.api.v1.IPService
 */
export declare const IPService: GenService<{
    /**
     * Returns the IP address with the specified IP.
     *
     * @generated from rpc fits.api.v1.IPService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof IPServiceGetRequestSchema;
        output: typeof IPServiceGetResponseSchema;
    };
    /**
     * Creates a new IP address.
     *
     * @generated from rpc fits.api.v1.IPService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof IPServiceCreateRequestSchema;
        output: typeof IPServiceCreateResponseSchema;
    };
    /**
     * Updates an IP address.
     *
     * @generated from rpc fits.api.v1.IPService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof IPServiceUpdateRequestSchema;
        output: typeof IPServiceUpdateResponseSchema;
    };
    /**
     * Returns the list of all IP addresses.
     *
     * @generated from rpc fits.api.v1.IPService.List
     */
    list: {
        methodKind: "unary";
        input: typeof IPServiceListRequestSchema;
        output: typeof IPServiceListResponseSchema;
    };
    /**
     * Deletes an IP address.
     *
     * @generated from rpc fits.api.v1.IPService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof IPServiceDeleteRequestSchema;
        output: typeof IPServiceDeleteResponseSchema;
    };
}>;
