import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Project } from "./project_pb";
import type { Tenant } from "./tenant_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/user.proto.
 */
export declare const file_fits_api_v1_user: GenFile;
/**
 * User represents an end user of the platform.
 *
 * @generated from message fits.api.v1.User
 */
export type User = Message<"fits.api.v1.User"> & {
    /**
     * Login is the login at the provider
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Name of the user
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Email of the user
     *
     * @generated from field: string email = 3;
     */
    email: string;
    /**
     * AvatarUrl of the user
     *
     * @generated from field: string avatar_url = 4;
     */
    avatarUrl: string;
    /**
     * Tenants the user belongs to
     *
     * @generated from field: repeated fits.api.v1.Tenant tenants = 5;
     */
    tenants: Tenant[];
    /**
     * Projects the user belongs to
     *
     * @generated from field: repeated fits.api.v1.Project projects = 6;
     */
    projects: Project[];
    /**
     * DefaultTenant this user belongs to
     *
     * @generated from field: fits.api.v1.Tenant default_tenant = 7;
     */
    defaultTenant?: Tenant | undefined;
};
/**
 * Describes the message fits.api.v1.User.
 * Use `create(UserSchema)` to create a new message.
 */
export declare const UserSchema: GenMessage<User>;
/**
 * UserServiceGetRequest is the request payload for getting the user.
 *
 * @generated from message fits.api.v1.UserServiceGetRequest
 */
export type UserServiceGetRequest = Message<"fits.api.v1.UserServiceGetRequest"> & {};
/**
 * Describes the message fits.api.v1.UserServiceGetRequest.
 * Use `create(UserServiceGetRequestSchema)` to create a new message.
 */
export declare const UserServiceGetRequestSchema: GenMessage<UserServiceGetRequest>;
/**
 * UserServiceGetResponse is the response payload for getting the user.
 *
 * @generated from message fits.api.v1.UserServiceGetResponse
 */
export type UserServiceGetResponse = Message<"fits.api.v1.UserServiceGetResponse"> & {
    /**
     * User contains the authenticated user
     *
     * @generated from field: fits.api.v1.User user = 1;
     */
    user?: User | undefined;
};
/**
 * Describes the message fits.api.v1.UserServiceGetResponse.
 * Use `create(UserServiceGetResponseSchema)` to create a new message.
 */
export declare const UserServiceGetResponseSchema: GenMessage<UserServiceGetResponse>;
/**
 * UserService provides user information operations.
 *
 * @generated from service fits.api.v1.UserService
 */
export declare const UserService: GenService<{
    /**
     * Returns the authenticated user.
     *
     * @generated from rpc fits.api.v1.UserService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof UserServiceGetRequestSchema;
        output: typeof UserServiceGetResponseSchema;
    };
}>;
