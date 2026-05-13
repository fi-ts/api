import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { AdminRole, ProjectRole, TenantRole } from "./common_pb";
import type { MethodPermission } from "./token_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/method.proto.
 */
export declare const file_fits_api_v1_method: GenFile;
/**
 * MethodServiceListRequest is the request payload for listing all public methods.
 *
 * @generated from message fits.api.v1.MethodServiceListRequest
 */
export type MethodServiceListRequest = Message<"fits.api.v1.MethodServiceListRequest"> & {};
/**
 * Describes the message fits.api.v1.MethodServiceListRequest.
 * Use `create(MethodServiceListRequestSchema)` to create a new message.
 */
export declare const MethodServiceListRequestSchema: GenMessage<MethodServiceListRequest>;
/**
 * MethodServiceListResponse is the response payload for listing all public methods.
 *
 * @generated from message fits.api.v1.MethodServiceListResponse
 */
export type MethodServiceListResponse = Message<"fits.api.v1.MethodServiceListResponse"> & {
    /**
     * Methods is a list of methods public callable
     *
     * @generated from field: repeated string methods = 1;
     */
    methods: string[];
};
/**
 * Describes the message fits.api.v1.MethodServiceListResponse.
 * Use `create(MethodServiceListResponseSchema)` to create a new message.
 */
export declare const MethodServiceListResponseSchema: GenMessage<MethodServiceListResponse>;
/**
 * MethodServiceTokenScopedListRequest is the request payload for listing all methods callable with the token present in the request.
 *
 * @generated from message fits.api.v1.MethodServiceTokenScopedListRequest
 */
export type MethodServiceTokenScopedListRequest = Message<"fits.api.v1.MethodServiceTokenScopedListRequest"> & {};
/**
 * Describes the message fits.api.v1.MethodServiceTokenScopedListRequest.
 * Use `create(MethodServiceTokenScopedListRequestSchema)` to create a new message.
 */
export declare const MethodServiceTokenScopedListRequestSchema: GenMessage<MethodServiceTokenScopedListRequest>;
/**
 * MethodServiceTokenScopedListResponse is the response payload which contains all methods which are callable with the given token.
 *
 * @generated from message fits.api.v1.MethodServiceTokenScopedListResponse
 */
export type MethodServiceTokenScopedListResponse = Message<"fits.api.v1.MethodServiceTokenScopedListResponse"> & {
    /**
     * Permissions contains a list of methods which can be called
     *
     * @generated from field: repeated fits.api.v1.MethodPermission permissions = 1;
     */
    permissions: MethodPermission[];
    /**
     * ProjectRoles associates a project ID with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.ProjectRole> project_roles = 2;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant ID with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.TenantRole> tenant_roles = 3;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional fits.api.v1.AdminRole admin_role = 4;
     */
    adminRole?: AdminRole | undefined;
};
/**
 * Describes the message fits.api.v1.MethodServiceTokenScopedListResponse.
 * Use `create(MethodServiceTokenScopedListResponseSchema)` to create a new message.
 */
export declare const MethodServiceTokenScopedListResponseSchema: GenMessage<MethodServiceTokenScopedListResponse>;
/**
 * MethodService provides method discovery operations.
 * Methods are functions in services.
 *
 * @generated from service fits.api.v1.MethodService
 */
export declare const MethodService: GenService<{
    /**
     * Returns the list of all public visible methods.
     *
     * @generated from rpc fits.api.v1.MethodService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MethodServiceListRequestSchema;
        output: typeof MethodServiceListResponseSchema;
    };
    /**
     * TokenScopedList returns all methods callable with the token present in the request.
     *
     * @generated from rpc fits.api.v1.MethodService.TokenScopedList
     */
    tokenScopedList: {
        methodKind: "unary";
        input: typeof MethodServiceTokenScopedListRequestSchema;
        output: typeof MethodServiceTokenScopedListResponseSchema;
    };
}>;
