import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/vlan.proto.
 */
export declare const file_fits_api_mvm_v1_vlan: GenFile;
/**
 * VLAN is a VLAN that a MVM can be connected to.
 *
 * @generated from message fits.api.mvm.v1.VLAN
 */
export type VLAN = Message<"fits.api.mvm.v1.VLAN"> & {
    /**
     * UUID of this VLAN.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * ID of this VLAN. The actual VLAN ID a network interface is attached to.
     *
     * @generated from field: int32 vlan_id = 2;
     */
    vlanId: number;
    /**
     * Title of the subnet this VLAN belongs to, including the stage type, e.g. "pod01-vlan2001 (Production)".
     *
     * @generated from field: string subnet_title = 3;
     */
    subnetTitle: string;
    /**
     * CIDR of the subnet this VLAN belongs to.
     *
     * @generated from field: string subnet_cidr = 4;
     */
    subnetCidr: string;
    /**
     * UUID of the location this VLAN belongs to.
     *
     * @generated from field: string location_uuid = 5;
     */
    locationUuid: string;
    /**
     * Tenant this VLAN belongs to.
     *
     * @generated from field: string tenant = 6;
     */
    tenant: string;
    /**
     * UUID of the Windows domain this VLAN is associated with. Only set for Windows.
     *
     * @generated from field: optional string domain_uuid = 7;
     */
    domainUuid?: string | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.VLAN.
 * Use `create(VLANSchema)` to create a new message.
 */
export declare const VLANSchema: GenMessage<VLAN>;
/**
 * VLANQuery contains fields which can be specified to list specific VLANs.
 *
 * @generated from message fits.api.mvm.v1.VLANQuery
 */
export type VLANQuery = Message<"fits.api.mvm.v1.VLANQuery"> & {
    /**
     * Filter the listed VLANs by the Windows domain they are associated with.
     *
     * @generated from field: optional string domain_uuid = 1;
     */
    domainUuid?: string | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.VLANQuery.
 * Use `create(VLANQuerySchema)` to create a new message.
 */
export declare const VLANQuerySchema: GenMessage<VLANQuery>;
/**
 * VLANServiceListRequest is the request payload for a VLAN list request.
 *
 * @generated from message fits.api.mvm.v1.VLANServiceListRequest
 */
export type VLANServiceListRequest = Message<"fits.api.mvm.v1.VLANServiceListRequest"> & {
    /**
     * Tenant to list available VLANs for (the tenant login).
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Filter the listed VLANs by location.
     *
     * @generated from field: string location_uuid = 2;
     */
    locationUuid: string;
    /**
     * Query to list one or more VLANs.
     *
     * @generated from field: fits.api.mvm.v1.VLANQuery query = 3;
     */
    query?: VLANQuery | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.VLANServiceListRequest.
 * Use `create(VLANServiceListRequestSchema)` to create a new message.
 */
export declare const VLANServiceListRequestSchema: GenMessage<VLANServiceListRequest>;
/**
 * VLANServiceListResponse is the response payload for a VLAN list request.
 *
 * @generated from message fits.api.mvm.v1.VLANServiceListResponse
 */
export type VLANServiceListResponse = Message<"fits.api.mvm.v1.VLANServiceListResponse"> & {
    /**
     * The available VLANs.
     *
     * @generated from field: repeated fits.api.mvm.v1.VLAN vlans = 1;
     */
    vlans: VLAN[];
};
/**
 * Describes the message fits.api.mvm.v1.VLANServiceListResponse.
 * Use `create(VLANServiceListResponseSchema)` to create a new message.
 */
export declare const VLANServiceListResponseSchema: GenMessage<VLANServiceListResponse>;
/**
 * VLANService lists VLANs available for managed VM (MVM) instances.
 *
 * @generated from service fits.api.mvm.v1.VLANService
 */
export declare const VLANService: GenService<{
    /**
     * Returns a list of all VLANs available for a tenant.
     *
     * @generated from rpc fits.api.mvm.v1.VLANService.List
     */
    list: {
        methodKind: "unary";
        input: typeof VLANServiceListRequestSchema;
        output: typeof VLANServiceListResponseSchema;
    };
}>;
