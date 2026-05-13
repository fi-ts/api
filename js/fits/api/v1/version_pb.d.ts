import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/version.proto.
 */
export declare const file_fits_api_v1_version: GenFile;
/**
 * Version represents the version of the application.
 *
 * @generated from message fits.api.v1.Version
 */
export type Version = Message<"fits.api.v1.Version"> & {
    /**
     * Version of the application
     *
     * @generated from field: string version = 1;
     */
    version: string;
    /**
     * Revision of the application
     *
     * @generated from field: string revision = 2;
     */
    revision: string;
    /**
     * GitSHA1 of the application
     *
     * @generated from field: string git_sha1 = 3;
     */
    gitSha1: string;
    /**
     * BuildDate of the application
     *
     * @generated from field: string build_date = 4;
     */
    buildDate: string;
};
/**
 * Describes the message fits.api.v1.Version.
 * Use `create(VersionSchema)` to create a new message.
 */
export declare const VersionSchema: GenMessage<Version>;
/**
 * VersionServiceGetRequest is the request payload for getting the version.
 *
 * @generated from message fits.api.v1.VersionServiceGetRequest
 */
export type VersionServiceGetRequest = Message<"fits.api.v1.VersionServiceGetRequest"> & {};
/**
 * Describes the message fits.api.v1.VersionServiceGetRequest.
 * Use `create(VersionServiceGetRequestSchema)` to create a new message.
 */
export declare const VersionServiceGetRequestSchema: GenMessage<VersionServiceGetRequest>;
/**
 * VersionServiceGetResponse is the response payload for getting the version.
 *
 * @generated from message fits.api.v1.VersionServiceGetResponse
 */
export type VersionServiceGetResponse = Message<"fits.api.v1.VersionServiceGetResponse"> & {
    /**
     * Version contains the version of the application
     *
     * @generated from field: fits.api.v1.Version version = 1;
     */
    version?: Version | undefined;
};
/**
 * Describes the message fits.api.v1.VersionServiceGetResponse.
 * Use `create(VersionServiceGetResponseSchema)` to create a new message.
 */
export declare const VersionServiceGetResponseSchema: GenMessage<VersionServiceGetResponse>;
/**
 * VersionService provides version information operations.
 *
 * @generated from service fits.api.v1.VersionService
 */
export declare const VersionService: GenService<{
    /**
     * Returns the version of the application.
     *
     * @generated from rpc fits.api.v1.VersionService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof VersionServiceGetRequestSchema;
        output: typeof VersionServiceGetResponseSchema;
    };
}>;
