import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { PermissionsByVisibility } from "./token_pb.js";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/apikey.proto.
 */
export declare const file_fits_api_v1_apikey: GenFile;
/**
 * ApiKey is a personal access token which authenticates api requests on behalf of a user.
 *
 * Secrets have the format fco_v1_<random>, the version selects the generation and
 * hash scheme of the secret, v1 is hashed with sha-256.
 * The plaintext secret is only known at creation time and is returned exactly once,
 * only the sha-256 hash of the secret is persisted, never the secret itself.
 *
 * @generated from message fits.api.v1.ApiKey
 */
export type ApiKey = Message<"fits.api.v1.ApiKey"> & {
    /**
     * Id is the unique identifier of this api key
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UserId is the id of the user this api key belongs to
     *
     * @generated from field: string user_id = 2;
     */
    userId: string;
    /**
     * Realm is the identity realm of the user this api key belongs to
     *
     * @generated from field: string realm = 3;
     */
    realm: string;
    /**
     * Name is a user given display name of this api key
     *
     * @generated from field: string name = 4;
     */
    name: string;
    /**
     * Hash is the sha-256 hash of the api key secret, the secret itself is never stored.
     * The api server leaves this field unset in api responses, it is only persisted internally.
     *
     * @generated from field: string hash = 5;
     */
    hash: string;
    /**
     * Prefix is the first 12 bytes of the api key secret, used to display and identify
     * the api key in the ui, e.g. fco_v1_abc12...
     *
     * @generated from field: string prefix = 6;
     */
    prefix: string;
    /**
     * CreatedAt is the date when this api key was created
     *
     * @generated from field: google.protobuf.Timestamp created_at = 7;
     */
    createdAt?: Timestamp | undefined;
    /**
     * ExpiresAt is the date after which this api key can not be used anymore.
     * If unset the api key does not expire.
     *
     * @generated from field: optional google.protobuf.Timestamp expires_at = 8;
     */
    expiresAt?: Timestamp | undefined;
    /**
     * RevokedAt is the date when this api key was revoked.
     * If unset the api key is not revoked.
     *
     * @generated from field: optional google.protobuf.Timestamp revoked_at = 9;
     */
    revokedAt?: Timestamp | undefined;
    /**
     * LastUsedAt is the date when this api key was last successfully used for authentication.
     * If unset the api key was never used.
     *
     * @generated from field: optional google.protobuf.Timestamp last_used_at = 10;
     */
    lastUsedAt?: Timestamp | undefined;
    /**
     * Permissions limit the methods this api key can call.
     * If unset the api key inherits the full permissions of the user it belongs to.
     * If set, the api key can only call the intersection of these permissions and
     * the permissions of the user it belongs to, see ApiKeyServiceCreateRequest.permissions.
     *
     * @generated from field: repeated fits.api.v1.PermissionsByVisibility permissions = 11;
     */
    permissions: PermissionsByVisibility[];
};
/**
 * Describes the message fits.api.v1.ApiKey.
 * Use `create(ApiKeySchema)` to create a new message.
 */
export declare const ApiKeySchema: GenMessage<ApiKey>;
/**
 * ApiKeyServiceCreateRequest is the request payload to create an api key
 *
 * @generated from message fits.api.v1.ApiKeyServiceCreateRequest
 */
export type ApiKeyServiceCreateRequest = Message<"fits.api.v1.ApiKeyServiceCreateRequest"> & {
    /**
     * Name is a user given display name of this api key
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * ExpiresAt is the date after which this api key can not be used anymore.
     * If unset the api key does not expire.
     *
     * @generated from field: optional google.protobuf.Timestamp expires_at = 2;
     */
    expiresAt?: Timestamp | undefined;
    /**
     * Permissions limit the methods this api key can call.
     * If unset the api key inherits the full permissions of the user it belongs to.
     * If set, the api key can only call the intersection of these permissions and
     * the permissions of the user it belongs to. Operator permissions are rejected,
     * a user can never obtain them.
     *
     * @generated from field: repeated fits.api.v1.PermissionsByVisibility permissions = 3;
     */
    permissions: PermissionsByVisibility[];
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceCreateRequest.
 * Use `create(ApiKeyServiceCreateRequestSchema)` to create a new message.
 */
export declare const ApiKeyServiceCreateRequestSchema: GenMessage<ApiKeyServiceCreateRequest>;
/**
 * ApiKeyServiceCreateResponse is the response payload of an api key create request
 *
 * @generated from message fits.api.v1.ApiKeyServiceCreateResponse
 */
export type ApiKeyServiceCreateResponse = Message<"fits.api.v1.ApiKeyServiceCreateResponse"> & {
    /**
     * Secret is the plaintext api key secret with format fco_v1_<random>, it is
     * returned exactly once and never stored by the api server.
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
    /**
     * ApiKey is the created api key
     *
     * @generated from field: fits.api.v1.ApiKey api_key = 2;
     */
    apiKey?: ApiKey | undefined;
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceCreateResponse.
 * Use `create(ApiKeyServiceCreateResponseSchema)` to create a new message.
 */
export declare const ApiKeyServiceCreateResponseSchema: GenMessage<ApiKeyServiceCreateResponse>;
/**
 * ApiKeyServiceListRequest is the request payload to list api keys
 *
 * @generated from message fits.api.v1.ApiKeyServiceListRequest
 */
export type ApiKeyServiceListRequest = Message<"fits.api.v1.ApiKeyServiceListRequest"> & {};
/**
 * Describes the message fits.api.v1.ApiKeyServiceListRequest.
 * Use `create(ApiKeyServiceListRequestSchema)` to create a new message.
 */
export declare const ApiKeyServiceListRequestSchema: GenMessage<ApiKeyServiceListRequest>;
/**
 * ApiKeyServiceListResponse is the response payload of an api key list request
 *
 * @generated from message fits.api.v1.ApiKeyServiceListResponse
 */
export type ApiKeyServiceListResponse = Message<"fits.api.v1.ApiKeyServiceListResponse"> & {
    /**
     * ApiKeys is a list of api keys without the secrets
     *
     * @generated from field: repeated fits.api.v1.ApiKey api_keys = 1;
     */
    apiKeys: ApiKey[];
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceListResponse.
 * Use `create(ApiKeyServiceListResponseSchema)` to create a new message.
 */
export declare const ApiKeyServiceListResponseSchema: GenMessage<ApiKeyServiceListResponse>;
/**
 * ApiKeyServiceGetRequest is the request payload of an api key get request
 *
 * @generated from message fits.api.v1.ApiKeyServiceGetRequest
 */
export type ApiKeyServiceGetRequest = Message<"fits.api.v1.ApiKeyServiceGetRequest"> & {
    /**
     * Id of the api key to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceGetRequest.
 * Use `create(ApiKeyServiceGetRequestSchema)` to create a new message.
 */
export declare const ApiKeyServiceGetRequestSchema: GenMessage<ApiKeyServiceGetRequest>;
/**
 * ApiKeyServiceGetResponse is the response payload of an api key get request
 *
 * @generated from message fits.api.v1.ApiKeyServiceGetResponse
 */
export type ApiKeyServiceGetResponse = Message<"fits.api.v1.ApiKeyServiceGetResponse"> & {
    /**
     * ApiKey is the api key
     *
     * @generated from field: fits.api.v1.ApiKey api_key = 1;
     */
    apiKey?: ApiKey | undefined;
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceGetResponse.
 * Use `create(ApiKeyServiceGetResponseSchema)` to create a new message.
 */
export declare const ApiKeyServiceGetResponseSchema: GenMessage<ApiKeyServiceGetResponse>;
/**
 * ApiKeyServiceRevokeRequest is the request payload of an api key revoke request
 *
 * @generated from message fits.api.v1.ApiKeyServiceRevokeRequest
 */
export type ApiKeyServiceRevokeRequest = Message<"fits.api.v1.ApiKeyServiceRevokeRequest"> & {
    /**
     * Id of the api key to revoke
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceRevokeRequest.
 * Use `create(ApiKeyServiceRevokeRequestSchema)` to create a new message.
 */
export declare const ApiKeyServiceRevokeRequestSchema: GenMessage<ApiKeyServiceRevokeRequest>;
/**
 * ApiKeyServiceRevokeResponse is the response payload of an api key revoke request
 *
 * @generated from message fits.api.v1.ApiKeyServiceRevokeResponse
 */
export type ApiKeyServiceRevokeResponse = Message<"fits.api.v1.ApiKeyServiceRevokeResponse"> & {};
/**
 * Describes the message fits.api.v1.ApiKeyServiceRevokeResponse.
 * Use `create(ApiKeyServiceRevokeResponseSchema)` to create a new message.
 */
export declare const ApiKeyServiceRevokeResponseSchema: GenMessage<ApiKeyServiceRevokeResponse>;
/**
 * ApiKeyServiceDeleteRequest is the request payload of an api key delete request
 *
 * @generated from message fits.api.v1.ApiKeyServiceDeleteRequest
 */
export type ApiKeyServiceDeleteRequest = Message<"fits.api.v1.ApiKeyServiceDeleteRequest"> & {
    /**
     * Id of the api key to delete
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message fits.api.v1.ApiKeyServiceDeleteRequest.
 * Use `create(ApiKeyServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ApiKeyServiceDeleteRequestSchema: GenMessage<ApiKeyServiceDeleteRequest>;
/**
 * ApiKeyServiceDeleteResponse is the response payload of an api key delete request
 *
 * @generated from message fits.api.v1.ApiKeyServiceDeleteResponse
 */
export type ApiKeyServiceDeleteResponse = Message<"fits.api.v1.ApiKeyServiceDeleteResponse"> & {};
/**
 * Describes the message fits.api.v1.ApiKeyServiceDeleteResponse.
 * Use `create(ApiKeyServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ApiKeyServiceDeleteResponseSchema: GenMessage<ApiKeyServiceDeleteResponse>;
/**
 * ApiKeyService provides api key (personal access token) management operations.
 *
 * An api key authenticates requests on behalf of the user it belongs to.
 * Clients send the secret of the api key in the authorization header:
 *
 * 	Authorization: Bearer fco_v1_abcd12...
 *
 * The api server resolves the secret to the user stored with the api key
 * and acquires a jwt for this user from the identity provider
 * (jwt authorization grant). The jwt is cached until it expires and is
 * reused for subsequent requests authenticated with the same api key.
 * Revoking or deleting an api key immediately invalidates its cached jwt.
 *
 * @generated from service fits.api.v1.ApiKeyService
 */
export declare const ApiKeyService: GenService<{
    /**
     * Creates a new api key, the plaintext secret is returned exactly once and never stored.
     *
     * @generated from rpc fits.api.v1.ApiKeyService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ApiKeyServiceCreateRequestSchema;
        output: typeof ApiKeyServiceCreateResponseSchema;
    };
    /**
     * Returns the list of all api keys of the authenticated user.
     *
     * @generated from rpc fits.api.v1.ApiKeyService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ApiKeyServiceListRequestSchema;
        output: typeof ApiKeyServiceListResponseSchema;
    };
    /**
     * Returns the api key with the specified id.
     *
     * @generated from rpc fits.api.v1.ApiKeyService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ApiKeyServiceGetRequestSchema;
        output: typeof ApiKeyServiceGetResponseSchema;
    };
    /**
     * Revokes an api key, no further usage is possible afterwards.
     *
     * @generated from rpc fits.api.v1.ApiKeyService.Revoke
     */
    revoke: {
        methodKind: "unary";
        input: typeof ApiKeyServiceRevokeRequestSchema;
        output: typeof ApiKeyServiceRevokeResponseSchema;
    };
    /**
     * Deletes an api key irreversibly.
     *
     * @generated from rpc fits.api.v1.ApiKeyService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ApiKeyServiceDeleteRequestSchema;
        output: typeof ApiKeyServiceDeleteResponseSchema;
    };
}>;
