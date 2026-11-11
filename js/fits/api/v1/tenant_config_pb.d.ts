import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta, UpdateMeta } from "./common_pb.js";
import type { JsonObject, Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/tenant_config.proto.
 */
export declare const file_fits_api_v1_tenant_config: GenFile;
/**
 * TenantConfig holds an arbitrary configuration payload for a tenant, keyed by name.
 * The name identifies the owning service and the config purpose, e.g. "mvm.fits.cloud/accounting".
 *
 * @generated from message fits.api.v1.TenantConfig
 */
export type TenantConfig = Message<"fits.api.v1.TenantConfig"> & {
    /**
     * Uuid is the unique identifier of this config.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Tenant is the login of the tenant this config belongs to.
     *
     * @generated from field: string tenant = 2;
     */
    tenant: string;
    /**
     * Namespace is the subdomain part of the config name, identifying the owning product or team,
     * e.g. "mvm", "fce", "fcn".
     *
     * @generated from field: string namespace = 3;
     */
    namespace: string;
    /**
     * Key is the config identifier within the namespace, e.g. "accounting".
     *
     * @generated from field: string key = 4;
     */
    key: string;
    /**
     * Name is the derived composite name in the form "{namespace}.fits.cloud/{key}",
     * e.g. "mvm.fits.cloud/accounting". It is read-only and computed by the server.
     *
     * @generated from field: string name = 5;
     */
    name: string;
    /**
     * Meta holds timestamps, labels, and the optimistic-locking generation counter.
     *
     * @generated from field: fits.api.v1.Meta meta = 6;
     */
    meta?: Meta | undefined;
    /**
     * Data is the configuration payload as a JSON object.
     * The structure is defined by the consuming service.
     *
     * @generated from field: google.protobuf.Struct data = 7;
     */
    data?: JsonObject | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfig.
 * Use `create(TenantConfigSchema)` to create a new message.
 */
export declare const TenantConfigSchema: GenMessage<TenantConfig>;
/**
 * TenantConfigServiceListRequest is the request payload for listing tenant configs.
 *
 * @generated from message fits.api.v1.TenantConfigServiceListRequest
 */
export type TenantConfigServiceListRequest = Message<"fits.api.v1.TenantConfigServiceListRequest"> & {
    /**
     * Tenant is the login of the tenant whose configs to list.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Namespace optionally filters by namespace, e.g. "mvm".
     *
     * @generated from field: optional string namespace = 2;
     */
    namespace?: string | undefined;
    /**
     * Key optionally filters by key within the namespace, e.g. "accounting".
     *
     * @generated from field: optional string key = 3;
     */
    key?: string | undefined;
    /**
     * Labels optionally filters configs that carry all of the given labels.
     *
     * @generated from field: optional fits.api.v1.Labels labels = 4;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceListRequest.
 * Use `create(TenantConfigServiceListRequestSchema)` to create a new message.
 */
export declare const TenantConfigServiceListRequestSchema: GenMessage<TenantConfigServiceListRequest>;
/**
 * TenantConfigServiceListResponse is the response payload for listing tenant configs.
 *
 * @generated from message fits.api.v1.TenantConfigServiceListResponse
 */
export type TenantConfigServiceListResponse = Message<"fits.api.v1.TenantConfigServiceListResponse"> & {
    /**
     * Configs is the list of matching tenant configs.
     *
     * @generated from field: repeated fits.api.v1.TenantConfig configs = 1;
     */
    configs: TenantConfig[];
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceListResponse.
 * Use `create(TenantConfigServiceListResponseSchema)` to create a new message.
 */
export declare const TenantConfigServiceListResponseSchema: GenMessage<TenantConfigServiceListResponse>;
/**
 * TenantConfigServiceGetRequest is the request payload for retrieving a single tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceGetRequest
 */
export type TenantConfigServiceGetRequest = Message<"fits.api.v1.TenantConfigServiceGetRequest"> & {
    /**
     * Tenant is the login of the tenant.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Uuid is the unique identifier of the config to retrieve.
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceGetRequest.
 * Use `create(TenantConfigServiceGetRequestSchema)` to create a new message.
 */
export declare const TenantConfigServiceGetRequestSchema: GenMessage<TenantConfigServiceGetRequest>;
/**
 * TenantConfigServiceGetResponse is the response payload for retrieving a single tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceGetResponse
 */
export type TenantConfigServiceGetResponse = Message<"fits.api.v1.TenantConfigServiceGetResponse"> & {
    /**
     * Config is the retrieved tenant config.
     *
     * @generated from field: fits.api.v1.TenantConfig config = 1;
     */
    config?: TenantConfig | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceGetResponse.
 * Use `create(TenantConfigServiceGetResponseSchema)` to create a new message.
 */
export declare const TenantConfigServiceGetResponseSchema: GenMessage<TenantConfigServiceGetResponse>;
/**
 * TenantConfigServiceCreateRequest is the request payload for creating a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceCreateRequest
 */
export type TenantConfigServiceCreateRequest = Message<"fits.api.v1.TenantConfigServiceCreateRequest"> & {
    /**
     * Tenant is the login of the tenant.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Namespace is the subdomain identifying the owning product or team, e.g. "mvm".
     * Must consist of lowercase letters only.
     *
     * @generated from field: string namespace = 2;
     */
    namespace: string;
    /**
     * Key is the config identifier within the namespace, e.g. "accounting", "accountingPg".
     * Must consist of alphanumeric characters only.
     *
     * @generated from field: string key = 3;
     */
    key: string;
    /**
     * Data is the configuration payload as a JSON object.
     *
     * @generated from field: google.protobuf.Struct data = 4;
     */
    data?: JsonObject | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceCreateRequest.
 * Use `create(TenantConfigServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantConfigServiceCreateRequestSchema: GenMessage<TenantConfigServiceCreateRequest>;
/**
 * TenantConfigServiceCreateResponse is the response payload after creating a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceCreateResponse
 */
export type TenantConfigServiceCreateResponse = Message<"fits.api.v1.TenantConfigServiceCreateResponse"> & {
    /**
     * Config is the created tenant config.
     *
     * @generated from field: fits.api.v1.TenantConfig config = 1;
     */
    config?: TenantConfig | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceCreateResponse.
 * Use `create(TenantConfigServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantConfigServiceCreateResponseSchema: GenMessage<TenantConfigServiceCreateResponse>;
/**
 * TenantConfigServiceUpdateRequest is the request payload for updating a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceUpdateRequest
 */
export type TenantConfigServiceUpdateRequest = Message<"fits.api.v1.TenantConfigServiceUpdateRequest"> & {
    /**
     * Tenant is the login of the tenant.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Uuid is the unique identifier of the config to update.
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
    /**
     * UpdateMeta carries the optimistic-locking timestamp and strategy.
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 3;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Data is the new configuration payload as a JSON object.
     *
     * @generated from field: google.protobuf.Struct data = 4;
     */
    data?: JsonObject | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceUpdateRequest.
 * Use `create(TenantConfigServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TenantConfigServiceUpdateRequestSchema: GenMessage<TenantConfigServiceUpdateRequest>;
/**
 * TenantConfigServiceUpdateResponse is the response payload after updating a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceUpdateResponse
 */
export type TenantConfigServiceUpdateResponse = Message<"fits.api.v1.TenantConfigServiceUpdateResponse"> & {
    /**
     * Config is the updated tenant config.
     *
     * @generated from field: fits.api.v1.TenantConfig config = 1;
     */
    config?: TenantConfig | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceUpdateResponse.
 * Use `create(TenantConfigServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TenantConfigServiceUpdateResponseSchema: GenMessage<TenantConfigServiceUpdateResponse>;
/**
 * TenantConfigServiceDeleteRequest is the request payload for deleting a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceDeleteRequest
 */
export type TenantConfigServiceDeleteRequest = Message<"fits.api.v1.TenantConfigServiceDeleteRequest"> & {
    /**
     * Tenant is the login of the tenant.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Uuid is the unique identifier of the config to delete.
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceDeleteRequest.
 * Use `create(TenantConfigServiceDeleteRequestSchema)` to create a new message.
 */
export declare const TenantConfigServiceDeleteRequestSchema: GenMessage<TenantConfigServiceDeleteRequest>;
/**
 * TenantConfigServiceDeleteResponse is the response payload after deleting a tenant config.
 *
 * @generated from message fits.api.v1.TenantConfigServiceDeleteResponse
 */
export type TenantConfigServiceDeleteResponse = Message<"fits.api.v1.TenantConfigServiceDeleteResponse"> & {
    /**
     * Config is the deleted tenant config.
     *
     * @generated from field: fits.api.v1.TenantConfig config = 1;
     */
    config?: TenantConfig | undefined;
};
/**
 * Describes the message fits.api.v1.TenantConfigServiceDeleteResponse.
 * Use `create(TenantConfigServiceDeleteResponseSchema)` to create a new message.
 */
export declare const TenantConfigServiceDeleteResponseSchema: GenMessage<TenantConfigServiceDeleteResponse>;
/**
 * TenantConfigService provides CRUD operations for tenant-scoped configuration entries.
 * Configs are for technical use only (e.g. service-specific settings per tenant) and are not meant to be exposed directly to end users.
 *
 * @generated from service fits.api.v1.TenantConfigService
 */
export declare const TenantConfigService: GenService<{
    /**
     * List returns the configs belonging to a tenant.
     *
     * @generated from rpc fits.api.v1.TenantConfigService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantConfigServiceListRequestSchema;
        output: typeof TenantConfigServiceListResponseSchema;
    };
    /**
     * Get returns a single tenant config by uuid.
     *
     * @generated from rpc fits.api.v1.TenantConfigService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TenantConfigServiceGetRequestSchema;
        output: typeof TenantConfigServiceGetResponseSchema;
    };
    /**
     * Create creates a new config for a tenant.
     * The config name must match the format <subdomain>.fits.cloud/<path>.
     *
     * @generated from rpc fits.api.v1.TenantConfigService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantConfigServiceCreateRequestSchema;
        output: typeof TenantConfigServiceCreateResponseSchema;
    };
    /**
     * Update replaces the data of an existing tenant config.
     *
     * @generated from rpc fits.api.v1.TenantConfigService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TenantConfigServiceUpdateRequestSchema;
        output: typeof TenantConfigServiceUpdateResponseSchema;
    };
    /**
     * Delete removes a tenant config by uuid.
     *
     * @generated from rpc fits.api.v1.TenantConfigService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof TenantConfigServiceDeleteRequestSchema;
        output: typeof TenantConfigServiceDeleteResponseSchema;
    };
}>;
