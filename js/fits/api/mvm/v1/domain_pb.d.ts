import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/domain.proto.
 */
export declare const file_fits_api_mvm_v1_domain: GenFile;
/**
 * Domain is a Windows domain a Windows MVM can join.
 *
 * @generated from message fits.api.mvm.v1.Domain
 */
export type Domain = Message<"fits.api.mvm.v1.Domain"> & {
    /**
     * UUID of this domain.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * FQDN of this domain.
     *
     * @generated from field: string fqdn = 2;
     */
    fqdn: string;
    /**
     * Tenant this domain belongs to.
     *
     * @generated from field: string tenant = 3;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.Domain.
 * Use `create(DomainSchema)` to create a new message.
 */
export declare const DomainSchema: GenMessage<Domain>;
/**
 * DomainServiceListRequest is the request payload for a domain list request.
 *
 * @generated from message fits.api.mvm.v1.DomainServiceListRequest
 */
export type DomainServiceListRequest = Message<"fits.api.mvm.v1.DomainServiceListRequest"> & {
    /**
     * Tenant to list available domains for.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Filter the listed domains by location (the FCE datacenter ID). A domain is only available in some datacenters.
     *
     * @generated from field: string location_uuid = 2;
     */
    locationUuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.DomainServiceListRequest.
 * Use `create(DomainServiceListRequestSchema)` to create a new message.
 */
export declare const DomainServiceListRequestSchema: GenMessage<DomainServiceListRequest>;
/**
 * DomainServiceListResponse is the response payload for a domain list request.
 *
 * @generated from message fits.api.mvm.v1.DomainServiceListResponse
 */
export type DomainServiceListResponse = Message<"fits.api.mvm.v1.DomainServiceListResponse"> & {
    /**
     * The available domains.
     *
     * @generated from field: repeated fits.api.mvm.v1.Domain domains = 1;
     */
    domains: Domain[];
};
/**
 * Describes the message fits.api.mvm.v1.DomainServiceListResponse.
 * Use `create(DomainServiceListResponseSchema)` to create a new message.
 */
export declare const DomainServiceListResponseSchema: GenMessage<DomainServiceListResponse>;
/**
 * DomainService lists Windows domains available for managed VM (MVM) instances.
 *
 * @generated from service fits.api.mvm.v1.DomainService
 */
export declare const DomainService: GenService<{
    /**
     * Returns a list of all Windows domains available for a tenant in a location.
     *
     * @generated from rpc fits.api.mvm.v1.DomainService.List
     */
    list: {
        methodKind: "unary";
        input: typeof DomainServiceListRequestSchema;
        output: typeof DomainServiceListResponseSchema;
    };
}>;
