import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
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
     * Uuid of this VLAN
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Id of this VLAN
     * This is not a foreign key, this is the actual ID of the VLAN the network interface is attached to.
     *
     * @generated from field: int32 vlan_id = 2;
     */
    vlanId: number;
    /**
     * Title of the subnet this VLAN belongs to
     *
     * @generated from field: string subnet_title = 3;
     */
    subnetTitle: string;
    /**
     * CIDR of the subnet this VLAN belongs to
     *
     * @generated from field: string subnet_cidr = 4;
     */
    subnetCidr: string;
    /**
     * TODO instead of the pod title, could we have the location_uuid?
     *
     * @generated from field: optional string location_uuid = 5;
     */
    locationUuid?: string | undefined;
    /**
     * Stage type of this VLAN
     *
     * @generated from field: fits.api.mvm.v1.StageType stage_type = 6;
     */
    stageType: StageType;
    /**
     * Tenant this VLAN belongs to.
     *
     * @generated from field: string tenant = 7;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.VLAN.
 * Use `create(VLANSchema)` to create a new message.
 */
export declare const VLANSchema: GenMessage<VLAN>;
/**
 * VLANServiceListRequest is the request payload for a VLAN list request.
 *
 * @generated from message fits.api.mvm.v1.VLANServiceListRequest
 */
export type VLANServiceListRequest = Message<"fits.api.mvm.v1.VLANServiceListRequest"> & {
    /**
     * Tenant to list available VLANs for (the tenant login)
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * StageType filters the listed VLANs by stage type
     *
     * @generated from field: optional fits.api.mvm.v1.StageType stage_type = 2;
     */
    stageType?: StageType | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.VLANServiceListRequest.
 * Use `create(VLANServiceListRequestSchema)` to create a new message.
 */
export declare const VLANServiceListRequestSchema: GenMessage<VLANServiceListRequest>;
/**
 * VLANServiceListResponse is the response payload for a VLAN list request
 *
 * @generated from message fits.api.mvm.v1.VLANServiceListResponse
 */
export type VLANServiceListResponse = Message<"fits.api.mvm.v1.VLANServiceListResponse"> & {
    /**
     * The available VLANs
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
 * StageType specifies the fixed stage of a managed VM (MVM) instance or VLAN.
 * The set of stages is fixed upstream, so it is modeled as an enum instead of
 * a listable resource.
 *
 * @generated from enum fits.api.mvm.v1.StageType
 */
export declare enum StageType {
    /**
     * STAGE_TYPE_UNSPECIFIED is not specified.
     *
     * @generated from enum value: STAGE_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * STAGE_TYPE_DEVELOPMENT is the development stage.
     *
     * @generated from enum value: STAGE_TYPE_DEVELOPMENT = 1;
     */
    DEVELOPMENT = 1,
    /**
     * STAGE_TYPE_TEST is the test stage.
     *
     * @generated from enum value: STAGE_TYPE_TEST = 2;
     */
    TEST = 2,
    /**
     * STAGE_TYPE_INTEGRATION is the integration stage.
     *
     * @generated from enum value: STAGE_TYPE_INTEGRATION = 3;
     */
    INTEGRATION = 3,
    /**
     * STAGE_TYPE_PRODUCTION is the production stage.
     *
     * @generated from enum value: STAGE_TYPE_PRODUCTION = 4;
     */
    PRODUCTION = 4
}
/**
 * Describes the enum fits.api.mvm.v1.StageType.
 */
export declare const StageTypeSchema: GenEnum<StageType>;
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
