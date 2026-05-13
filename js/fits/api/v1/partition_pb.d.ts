import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/partition.proto.
 */
export declare const file_fits_api_v1_partition: GenFile;
/**
 * Partition is a failure domain with machines and switches
 *
 * @generated from message fits.api.v1.Partition
 */
export type Partition = Message<"fits.api.v1.Partition"> & {
    /**
     * ID of this partition
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this ip
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Description of this partition
     *
     * @generated from field: string description = 3;
     */
    description: string;
    /**
     * PartitionBootConfiguration defines how metal-hammer boots
     *
     * @generated from field: fits.api.v1.PartitionBootConfiguration boot_configuration = 4;
     */
    bootConfiguration?: PartitionBootConfiguration | undefined;
    /**
     * DNSServers for this partition
     *
     * @generated from field: repeated fits.api.v1.DNSServer dns_servers = 5;
     */
    dnsServers: DNSServer[];
    /**
     * NTPServers for this partition
     *
     * @generated from field: repeated fits.api.v1.NTPServer ntp_servers = 6;
     */
    ntpServers: NTPServer[];
    /**
     * ManagementServiceAddresses defines where the management is reachable
     * should be in the form <ip|host>:<port>
     *
     * @generated from field: repeated string mgmt_service_addresses = 7;
     */
    mgmtServiceAddresses: string[];
};
/**
 * Describes the message fits.api.v1.Partition.
 * Use `create(PartitionSchema)` to create a new message.
 */
export declare const PartitionSchema: GenMessage<Partition>;
/**
 * PartitionQuery is used to search partitions
 *
 * @generated from message fits.api.v1.PartitionQuery
 */
export type PartitionQuery = Message<"fits.api.v1.PartitionQuery"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string | undefined;
    /**
     * Labels lists only partitions containing the given labels
     *
     * @generated from field: optional fits.api.v1.Labels labels = 2;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.PartitionQuery.
 * Use `create(PartitionQuerySchema)` to create a new message.
 */
export declare const PartitionQuerySchema: GenMessage<PartitionQuery>;
/**
 * PartitionBootConfiguration defines how metal-hammer boots
 *
 * @generated from message fits.api.v1.PartitionBootConfiguration
 */
export type PartitionBootConfiguration = Message<"fits.api.v1.PartitionBootConfiguration"> & {
    /**
     * ImageURL the url to download the initrd for the boot image
     *
     * @generated from field: string image_url = 1;
     */
    imageUrl: string;
    /**
     * KernelURL the url to download the kernel for the boot image
     *
     * @generated from field: string kernel_url = 2;
     */
    kernelUrl: string;
    /**
     * Commandline the cmdline to the kernel for the boot image
     *
     * @generated from field: string commandline = 3;
     */
    commandline: string;
};
/**
 * Describes the message fits.api.v1.PartitionBootConfiguration.
 * Use `create(PartitionBootConfigurationSchema)` to create a new message.
 */
export declare const PartitionBootConfigurationSchema: GenMessage<PartitionBootConfiguration>;
/**
 * DNSServer
 *
 * @generated from message fits.api.v1.DNSServer
 */
export type DNSServer = Message<"fits.api.v1.DNSServer"> & {
    /**
     * IP address of this dns server
     *
     * @generated from field: string ip = 1;
     */
    ip: string;
};
/**
 * Describes the message fits.api.v1.DNSServer.
 * Use `create(DNSServerSchema)` to create a new message.
 */
export declare const DNSServerSchema: GenMessage<DNSServer>;
/**
 * NTPServer
 *
 * @generated from message fits.api.v1.NTPServer
 */
export type NTPServer = Message<"fits.api.v1.NTPServer"> & {
    /**
     * Address either as ip or hostname
     *
     * @generated from field: string address = 1;
     */
    address: string;
};
/**
 * Describes the message fits.api.v1.NTPServer.
 * Use `create(NTPServerSchema)` to create a new message.
 */
export declare const NTPServerSchema: GenMessage<NTPServer>;
/**
 * PartitionServiceGetRequest is the request payload for a partition get request
 *
 * @generated from message fits.api.v1.PartitionServiceGetRequest
 */
export type PartitionServiceGetRequest = Message<"fits.api.v1.PartitionServiceGetRequest"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message fits.api.v1.PartitionServiceGetRequest.
 * Use `create(PartitionServiceGetRequestSchema)` to create a new message.
 */
export declare const PartitionServiceGetRequestSchema: GenMessage<PartitionServiceGetRequest>;
/**
 * PartitionServiceListRequest is the request payload for a partition list request
 *
 * @generated from message fits.api.v1.PartitionServiceListRequest
 */
export type PartitionServiceListRequest = Message<"fits.api.v1.PartitionServiceListRequest"> & {
    /**
     * Query for partitions
     *
     * @generated from field: fits.api.v1.PartitionQuery query = 1;
     */
    query?: PartitionQuery | undefined;
};
/**
 * Describes the message fits.api.v1.PartitionServiceListRequest.
 * Use `create(PartitionServiceListRequestSchema)` to create a new message.
 */
export declare const PartitionServiceListRequestSchema: GenMessage<PartitionServiceListRequest>;
/**
 * PartitionServiceGetResponse is the response payload for a partition get request
 *
 * @generated from message fits.api.v1.PartitionServiceGetResponse
 */
export type PartitionServiceGetResponse = Message<"fits.api.v1.PartitionServiceGetResponse"> & {
    /**
     * Ip the partition
     *
     * @generated from field: fits.api.v1.Partition partition = 1;
     */
    partition?: Partition | undefined;
};
/**
 * Describes the message fits.api.v1.PartitionServiceGetResponse.
 * Use `create(PartitionServiceGetResponseSchema)` to create a new message.
 */
export declare const PartitionServiceGetResponseSchema: GenMessage<PartitionServiceGetResponse>;
/**
 * PartitionServiceListResponse is the response payload for a partition list request
 *
 * @generated from message fits.api.v1.PartitionServiceListResponse
 */
export type PartitionServiceListResponse = Message<"fits.api.v1.PartitionServiceListResponse"> & {
    /**
     * Ips the partitions
     *
     * @generated from field: repeated fits.api.v1.Partition partitions = 1;
     */
    partitions: Partition[];
};
/**
 * Describes the message fits.api.v1.PartitionServiceListResponse.
 * Use `create(PartitionServiceListResponseSchema)` to create a new message.
 */
export declare const PartitionServiceListResponseSchema: GenMessage<PartitionServiceListResponse>;
/**
 * PartitionService provides partition management operations.
 *
 * @generated from service fits.api.v1.PartitionService
 */
export declare const PartitionService: GenService<{
    /**
     * Returns the partition with the specified ID.
     *
     * @generated from rpc fits.api.v1.PartitionService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof PartitionServiceGetRequestSchema;
        output: typeof PartitionServiceGetResponseSchema;
    };
    /**
     * Returns the list of all partitions.
     *
     * @generated from rpc fits.api.v1.PartitionService.List
     */
    list: {
        methodKind: "unary";
        input: typeof PartitionServiceListRequestSchema;
        output: typeof PartitionServiceListResponseSchema;
    };
}>;
