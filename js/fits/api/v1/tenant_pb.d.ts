import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/tenant.proto.
 */
export declare const file_fits_api_v1_tenant: GenFile;
/**
 * Tenant is a customer of the platform
 *
 * @generated from message fits.api.v1.Tenant
 */
export type Tenant = Message<"fits.api.v1.Tenant"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Meta for this tenant
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Name of the tenant
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Email of the tenant
     *
     * @generated from field: string email = 4;
     */
    email: string;
    /**
     * Description of this tenant
     *
     * @generated from field: string description = 5;
     */
    description: string;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: string avatar_url = 6;
     */
    avatarUrl: string;
    /**
     * CreatedBy stores who created this tenant
     *
     * @generated from field: string created_by = 7;
     */
    createdBy: string;
};
/**
 * Describes the message fits.api.v1.Tenant.
 * Use `create(TenantSchema)` to create a new message.
 */
export declare const TenantSchema: GenMessage<Tenant>;
/**
 * TenantServiceListRequest is the request payload of the tenant list request
 *
 * @generated from message fits.api.v1.TenantServiceListRequest
 */
export type TenantServiceListRequest = Message<"fits.api.v1.TenantServiceListRequest"> & {
    /**
     * Id filters tenants by id
     *
     * @generated from field: optional string id = 1;
     */
    id?: string | undefined;
    /**
     * Name filters tenants by name
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Labels lists only projects containing the given labels
     *
     * @generated from field: optional fits.api.v1.Labels labels = 3;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceListRequest.
 * Use `create(TenantServiceListRequestSchema)` to create a new message.
 */
export declare const TenantServiceListRequestSchema: GenMessage<TenantServiceListRequest>;
/**
 * TenantServiceGetRequest is the request payload of the tenant get request
 *
 * @generated from message fits.api.v1.TenantServiceGetRequest
 */
export type TenantServiceGetRequest = Message<"fits.api.v1.TenantServiceGetRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message fits.api.v1.TenantServiceGetRequest.
 * Use `create(TenantServiceGetRequestSchema)` to create a new message.
 */
export declare const TenantServiceGetRequestSchema: GenMessage<TenantServiceGetRequest>;
/**
 * TenantServiceCreateRequest is the request payload of the tenant create request
 *
 * @generated from message fits.api.v1.TenantServiceCreateRequest
 */
export type TenantServiceCreateRequest = Message<"fits.api.v1.TenantServiceCreateRequest"> & {
    /**
     * Name of this tenant
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Description of this tenant
     *
     * @generated from field: optional string description = 2;
     */
    description?: string | undefined;
    /**
     * Email of the tenant, if not set will be inherited from the creator
     *
     * @generated from field: optional string email = 3;
     */
    email?: string | undefined;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the tenant
     *
     * @generated from field: fits.api.v1.Labels labels = 5;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceCreateRequest.
 * Use `create(TenantServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantServiceCreateRequestSchema: GenMessage<TenantServiceCreateRequest>;
/**
 * TenantServiceUpdateRequest is the request payload of the tenant update request
 *
 * @generated from message fits.api.v1.TenantServiceUpdateRequest
 */
export type TenantServiceUpdateRequest = Message<"fits.api.v1.TenantServiceUpdateRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Name of the tenant
     *
     * @generated from field: optional string name = 3;
     */
    name?: string | undefined;
    /**
     * Email of the tenant
     *
     * @generated from field: optional string email = 4;
     */
    email?: string | undefined;
    /**
     * Description of this tenant
     *
     * @generated from field: optional string description = 5;
     */
    description?: string | undefined;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: optional string avatar_url = 6;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the tenant
     *
     * @generated from field: optional fits.api.v1.UpdateLabels labels = 7;
     */
    labels?: UpdateLabels | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceUpdateRequest.
 * Use `create(TenantServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TenantServiceUpdateRequestSchema: GenMessage<TenantServiceUpdateRequest>;
/**
 * TenantServiceDeleteRequest is the request payload of the tenant delete request
 *
 * @generated from message fits.api.v1.TenantServiceDeleteRequest
 */
export type TenantServiceDeleteRequest = Message<"fits.api.v1.TenantServiceDeleteRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message fits.api.v1.TenantServiceDeleteRequest.
 * Use `create(TenantServiceDeleteRequestSchema)` to create a new message.
 */
export declare const TenantServiceDeleteRequestSchema: GenMessage<TenantServiceDeleteRequest>;
/**
 * TenantServiceGetResponse is the response payload of the tenant get request
 *
 * @generated from message fits.api.v1.TenantServiceGetResponse
 */
export type TenantServiceGetResponse = Message<"fits.api.v1.TenantServiceGetResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: fits.api.v1.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceGetResponse.
 * Use `create(TenantServiceGetResponseSchema)` to create a new message.
 */
export declare const TenantServiceGetResponseSchema: GenMessage<TenantServiceGetResponse>;
/**
 * TenantServiceListResponse is the response payload of the tenant list request
 *
 * @generated from message fits.api.v1.TenantServiceListResponse
 */
export type TenantServiceListResponse = Message<"fits.api.v1.TenantServiceListResponse"> & {
    /**
     * Tenants is the list of tenants
     *
     * @generated from field: repeated fits.api.v1.Tenant tenants = 1;
     */
    tenants: Tenant[];
};
/**
 * Describes the message fits.api.v1.TenantServiceListResponse.
 * Use `create(TenantServiceListResponseSchema)` to create a new message.
 */
export declare const TenantServiceListResponseSchema: GenMessage<TenantServiceListResponse>;
/**
 * TenantServiceCreateResponse is the response payload of the tenant create request
 *
 * @generated from message fits.api.v1.TenantServiceCreateResponse
 */
export type TenantServiceCreateResponse = Message<"fits.api.v1.TenantServiceCreateResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: fits.api.v1.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceCreateResponse.
 * Use `create(TenantServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantServiceCreateResponseSchema: GenMessage<TenantServiceCreateResponse>;
/**
 * TenantServiceUpdateResponse is the response payload of the tenant update request
 *
 * @generated from message fits.api.v1.TenantServiceUpdateResponse
 */
export type TenantServiceUpdateResponse = Message<"fits.api.v1.TenantServiceUpdateResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: fits.api.v1.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceUpdateResponse.
 * Use `create(TenantServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TenantServiceUpdateResponseSchema: GenMessage<TenantServiceUpdateResponse>;
/**
 * TenantServiceDeleteResponse is the response payload of the tenant delete request
 *
 * @generated from message fits.api.v1.TenantServiceDeleteResponse
 */
export type TenantServiceDeleteResponse = Message<"fits.api.v1.TenantServiceDeleteResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: fits.api.v1.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message fits.api.v1.TenantServiceDeleteResponse.
 * Use `create(TenantServiceDeleteResponseSchema)` to create a new message.
 */
export declare const TenantServiceDeleteResponseSchema: GenMessage<TenantServiceDeleteResponse>;
/**
 * TenantService provides tenant management operations.
 *
 * @generated from service fits.api.v1.TenantService
 */
export declare const TenantService: GenService<{
    /**
     * Creates a new tenant.
     *
     * @generated from rpc fits.api.v1.TenantService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantServiceCreateRequestSchema;
        output: typeof TenantServiceCreateResponseSchema;
    };
    /**
     * Returns the list of tenants.
     *
     * @generated from rpc fits.api.v1.TenantService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantServiceListRequestSchema;
        output: typeof TenantServiceListResponseSchema;
    };
    /**
     * Get a tenant
     *
     * @generated from rpc fits.api.v1.TenantService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TenantServiceGetRequestSchema;
        output: typeof TenantServiceGetResponseSchema;
    };
    /**
     * Update a tenant
     *
     * @generated from rpc fits.api.v1.TenantService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TenantServiceUpdateRequestSchema;
        output: typeof TenantServiceUpdateResponseSchema;
    };
    /**
     * Delete a tenant
     *
     * @generated from rpc fits.api.v1.TenantService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof TenantServiceDeleteRequestSchema;
        output: typeof TenantServiceDeleteResponseSchema;
    };
}>;
