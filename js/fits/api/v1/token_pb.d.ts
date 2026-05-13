import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { AdminRole, Labels, Meta, ProjectRole, TenantRole, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/token.proto.
 */
export declare const file_fits_api_v1_token: GenFile;
/**
 * Token generates a jwt authentication token to access the api
 *
 * There are two different types of tokens, api- and user- tokens
 *
 * A user token is used to authenticate end user requests for example from a cli.
 * The configured roles in a user token are expanded in the api server
 * based on the memberships in other projects and tenants based on the role granted there.
 * User tokens will never contain permissions.
 * Permissions are always derived from the tenant and project roles and memberships.
 *
 * The api token should be used to authenticate services.
 * In contrast to a user token, the api token permissions and roles apply as configured during the token create process.
 *
 * @generated from message fits.api.v1.Token
 */
export type Token = Message<"fits.api.v1.Token"> & {
    /**
     * Uuid of the jwt token, used to reference it by revoke
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * User who created this token
     *
     * @generated from field: string user = 2;
     */
    user: string;
    /**
     * Meta for this token
     *
     * @generated from field: fits.api.v1.Meta meta = 3;
     */
    meta?: Meta | undefined;
    /**
     * Description is a user given description of this token.
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated fits.api.v1.MethodPermission permissions = 5;
     */
    permissions: MethodPermission[];
    /**
     * Expires gives the date in the future after which this token can not be used anymore
     *
     * @generated from field: google.protobuf.Timestamp expires = 6;
     */
    expires?: Timestamp | undefined;
    /**
     * IssuedAt gives the date when this token was created
     *
     * @generated from field: google.protobuf.Timestamp issued_at = 7;
     */
    issuedAt?: Timestamp | undefined;
    /**
     * TokenType describes the type of this token
     *
     * @generated from field: fits.api.v1.TokenType token_type = 8;
     */
    tokenType: TokenType;
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.ProjectRole> project_roles = 9;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.TenantRole> tenant_roles = 10;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional fits.api.v1.AdminRole admin_role = 11;
     */
    adminRole?: AdminRole | undefined;
};
/**
 * Describes the message fits.api.v1.Token.
 * Use `create(TokenSchema)` to create a new message.
 */
export declare const TokenSchema: GenMessage<Token>;
/**
 * TokenServiceCreateRequest is the request payload to create a token
 *
 * @generated from message fits.api.v1.TokenServiceCreateRequest
 */
export type TokenServiceCreateRequest = Message<"fits.api.v1.TokenServiceCreateRequest"> & {
    /**
     * Description of the token
     *
     * @generated from field: string description = 1;
     */
    description: string;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated fits.api.v1.MethodPermission permissions = 2;
     */
    permissions: MethodPermission[];
    /**
     * Expires gives the duration since now, after which this token can not be used anymore
     *
     * @generated from field: google.protobuf.Duration expires = 3;
     */
    expires?: Duration | undefined;
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.ProjectRole> project_roles = 4;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles_associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.TenantRole> tenant_roles = 5;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional fits.api.v1.AdminRole admin_role = 6;
     */
    adminRole?: AdminRole | undefined;
    /**
     * Labels on this token
     *
     * @generated from field: fits.api.v1.Labels labels = 7;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.TokenServiceCreateRequest.
 * Use `create(TokenServiceCreateRequestSchema)` to create a new message.
 */
export declare const TokenServiceCreateRequestSchema: GenMessage<TokenServiceCreateRequest>;
/**
 * MethodPermission is a mapping from a subject/project to a service method
 *
 * @generated from message fits.api.v1.MethodPermission
 */
export type MethodPermission = Message<"fits.api.v1.MethodPermission"> & {
    /**
     * Subject maybe either the project or the tenant
     * for which the methods should be allowed
     *
     * asterisk (*) can be specified to match any subject
     * empty string ("") can be specified for requests that do not require a subject, e.g. partition list
     * otherwise either a projectid or a tenant login should be specified
     *
     * @generated from field: string subject = 1;
     */
    subject: string;
    /**
     * Methods which should be accessible
     *
     * @generated from field: repeated string methods = 2;
     */
    methods: string[];
};
/**
 * Describes the message fits.api.v1.MethodPermission.
 * Use `create(MethodPermissionSchema)` to create a new message.
 */
export declare const MethodPermissionSchema: GenMessage<MethodPermission>;
/**
 * TokenServiceCreateResponse is the response payload of a token create request
 *
 * @generated from message fits.api.v1.TokenServiceCreateResponse
 */
export type TokenServiceCreateResponse = Message<"fits.api.v1.TokenServiceCreateResponse"> & {
    /**
     * Token which was created
     *
     * @generated from field: fits.api.v1.Token token = 1;
     */
    token?: Token | undefined;
    /**
     * Secret is the body if the jwt token, should be used in api requests as bearer token
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message fits.api.v1.TokenServiceCreateResponse.
 * Use `create(TokenServiceCreateResponseSchema)` to create a new message.
 */
export declare const TokenServiceCreateResponseSchema: GenMessage<TokenServiceCreateResponse>;
/**
 * TokenServiceListRequest is the request payload to list tokens
 *
 * @generated from message fits.api.v1.TokenServiceListRequest
 */
export type TokenServiceListRequest = Message<"fits.api.v1.TokenServiceListRequest"> & {};
/**
 * Describes the message fits.api.v1.TokenServiceListRequest.
 * Use `create(TokenServiceListRequestSchema)` to create a new message.
 */
export declare const TokenServiceListRequestSchema: GenMessage<TokenServiceListRequest>;
/**
 * TokenServiceListResponse is the response payload of a token list request
 *
 * @generated from message fits.api.v1.TokenServiceListResponse
 */
export type TokenServiceListResponse = Message<"fits.api.v1.TokenServiceListResponse"> & {
    /**
     * Tokens is a list of tokens without the secrets
     *
     * @generated from field: repeated fits.api.v1.Token tokens = 1;
     */
    tokens: Token[];
};
/**
 * Describes the message fits.api.v1.TokenServiceListResponse.
 * Use `create(TokenServiceListResponseSchema)` to create a new message.
 */
export declare const TokenServiceListResponseSchema: GenMessage<TokenServiceListResponse>;
/**
 * TokenServiceRevokeRequest is the request payload of a token revoke request
 *
 * @generated from message fits.api.v1.TokenServiceRevokeRequest
 */
export type TokenServiceRevokeRequest = Message<"fits.api.v1.TokenServiceRevokeRequest"> & {
    /**
     * Uuid of the token to revoke
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.v1.TokenServiceRevokeRequest.
 * Use `create(TokenServiceRevokeRequestSchema)` to create a new message.
 */
export declare const TokenServiceRevokeRequestSchema: GenMessage<TokenServiceRevokeRequest>;
/**
 * TokenServiceRevokeResponse is the response payload of a token revoke request
 *
 * @generated from message fits.api.v1.TokenServiceRevokeResponse
 */
export type TokenServiceRevokeResponse = Message<"fits.api.v1.TokenServiceRevokeResponse"> & {};
/**
 * Describes the message fits.api.v1.TokenServiceRevokeResponse.
 * Use `create(TokenServiceRevokeResponseSchema)` to create a new message.
 */
export declare const TokenServiceRevokeResponseSchema: GenMessage<TokenServiceRevokeResponse>;
/**
 * TokenServiceUpdateRequest is the request payload of a token update request
 *
 * @generated from message fits.api.v1.TokenServiceUpdateRequest
 */
export type TokenServiceUpdateRequest = Message<"fits.api.v1.TokenServiceUpdateRequest"> & {
    /**
     * Uuid of the token to update
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     * TokenUpdate is not guarded with optlock in the backend
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Description is a user given description of this token.
     *
     * @generated from field: optional string description = 3;
     */
    description?: string | undefined;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated fits.api.v1.MethodPermission permissions = 4;
     */
    permissions: MethodPermission[];
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.ProjectRole> project_roles = 5;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, fits.api.v1.TenantRole> tenant_roles = 6;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional fits.api.v1.AdminRole admin_role = 7;
     */
    adminRole?: AdminRole | undefined;
    /**
     * Labels on this token
     *
     * @generated from field: fits.api.v1.UpdateLabels labels = 8;
     */
    labels?: UpdateLabels | undefined;
};
/**
 * Describes the message fits.api.v1.TokenServiceUpdateRequest.
 * Use `create(TokenServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TokenServiceUpdateRequestSchema: GenMessage<TokenServiceUpdateRequest>;
/**
 * TokenServiceUpdateResponse is the response payload of a token update request
 *
 * @generated from message fits.api.v1.TokenServiceUpdateResponse
 */
export type TokenServiceUpdateResponse = Message<"fits.api.v1.TokenServiceUpdateResponse"> & {
    /**
     * Token is the updated token
     *
     * @generated from field: fits.api.v1.Token token = 1;
     */
    token?: Token | undefined;
};
/**
 * Describes the message fits.api.v1.TokenServiceUpdateResponse.
 * Use `create(TokenServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TokenServiceUpdateResponseSchema: GenMessage<TokenServiceUpdateResponse>;
/**
 * TokenServiceGetRequest is the request payload of a token get request
 *
 * @generated from message fits.api.v1.TokenServiceGetRequest
 */
export type TokenServiceGetRequest = Message<"fits.api.v1.TokenServiceGetRequest"> & {
    /**
     * Uuid of the token to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.v1.TokenServiceGetRequest.
 * Use `create(TokenServiceGetRequestSchema)` to create a new message.
 */
export declare const TokenServiceGetRequestSchema: GenMessage<TokenServiceGetRequest>;
/**
 * TokenServiceGetResponse is the response payload of a token get request
 *
 * @generated from message fits.api.v1.TokenServiceGetResponse
 */
export type TokenServiceGetResponse = Message<"fits.api.v1.TokenServiceGetResponse"> & {
    /**
     * Token is the token
     *
     * @generated from field: fits.api.v1.Token token = 1;
     */
    token?: Token | undefined;
};
/**
 * Describes the message fits.api.v1.TokenServiceGetResponse.
 * Use `create(TokenServiceGetResponseSchema)` to create a new message.
 */
export declare const TokenServiceGetResponseSchema: GenMessage<TokenServiceGetResponse>;
/**
 * TokenServiceRefreshRequest is the request payload of a token refresh request
 * Permissions, Roles and Expiration duration and all other properties are inherited from the calling token.
 * The expiration duration will be calculated from the existing token (exp - iat)
 *
 * @generated from message fits.api.v1.TokenServiceRefreshRequest
 */
export type TokenServiceRefreshRequest = Message<"fits.api.v1.TokenServiceRefreshRequest"> & {};
/**
 * Describes the message fits.api.v1.TokenServiceRefreshRequest.
 * Use `create(TokenServiceRefreshRequestSchema)` to create a new message.
 */
export declare const TokenServiceRefreshRequestSchema: GenMessage<TokenServiceRefreshRequest>;
/**
 * TokenServiceRefreshResponse is the response payload of a token refresh request
 *
 * @generated from message fits.api.v1.TokenServiceRefreshResponse
 */
export type TokenServiceRefreshResponse = Message<"fits.api.v1.TokenServiceRefreshResponse"> & {
    /**
     * Token which was refreshed
     *
     * @generated from field: fits.api.v1.Token token = 1;
     */
    token?: Token | undefined;
    /**
     * Secret is the body if the jwt token, should be used in api requests as bearer token
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message fits.api.v1.TokenServiceRefreshResponse.
 * Use `create(TokenServiceRefreshResponseSchema)` to create a new message.
 */
export declare const TokenServiceRefreshResponseSchema: GenMessage<TokenServiceRefreshResponse>;
/**
 * TokenType specifies different use cases of tokens
 *
 * @generated from enum fits.api.v1.TokenType
 */
export declare enum TokenType {
    /**
     * TOKEN_TYPE_UNSPECIFIED is not specified
     *
     * @generated from enum value: TOKEN_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * TOKEN_TYPE_API is a token for api usage
     *
     * @generated from enum value: TOKEN_TYPE_API = 1;
     */
    API = 1,
    /**
     * TOKEN_TYPE_USER is a token to access the api with cli, a web application or other user induced actions.
     *
     * @generated from enum value: TOKEN_TYPE_USER = 2;
     */
    USER = 2
}
/**
 * Describes the enum fits.api.v1.TokenType.
 */
export declare const TokenTypeSchema: GenEnum<TokenType>;
/**
 * TokenService provides token management operations.
 *
 * @generated from service fits.api.v1.TokenService
 */
export declare const TokenService: GenService<{
    /**
     * Returns the token with the specified UUID.
     *
     * @generated from rpc fits.api.v1.TokenService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TokenServiceGetRequestSchema;
        output: typeof TokenServiceGetResponseSchema;
    };
    /**
     * Creates a token to authenticate against the platform, the secret will be only visible in the response.
     *
     * @generated from rpc fits.api.v1.TokenService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TokenServiceCreateRequestSchema;
        output: typeof TokenServiceCreateResponseSchema;
    };
    /**
     * Updates a token.
     *
     * @generated from rpc fits.api.v1.TokenService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TokenServiceUpdateRequestSchema;
        output: typeof TokenServiceUpdateResponseSchema;
    };
    /**
     * Returns the list of all user tokens.
     *
     * @generated from rpc fits.api.v1.TokenService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TokenServiceListRequestSchema;
        output: typeof TokenServiceListResponseSchema;
    };
    /**
     * Revokes a token, no further usage is possible afterwards.
     *
     * @generated from rpc fits.api.v1.TokenService.Revoke
     */
    revoke: {
        methodKind: "unary";
        input: typeof TokenServiceRevokeRequestSchema;
        output: typeof TokenServiceRevokeResponseSchema;
    };
    /**
     * Refreshes a token, this will create a new token with the exact same permissions as the calling token contains.
     *
     * @generated from rpc fits.api.v1.TokenService.Refresh
     */
    refresh: {
        methodKind: "unary";
        input: typeof TokenServiceRefreshRequestSchema;
        output: typeof TokenServiceRefreshResponseSchema;
    };
}>;
