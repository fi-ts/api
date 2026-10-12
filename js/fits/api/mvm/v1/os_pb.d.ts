import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/os.proto.
 */
export declare const file_fits_api_mvm_v1_os: GenFile;
/**
 * OS is the definition of an available OS for MVM instances.
 *
 * @generated from message fits.api.mvm.v1.OS
 */
export type OS = Message<"fits.api.mvm.v1.OS"> & {
    /**
     * Uuid of this OS.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Title of the OS.
     *
     * @generated from field: string title = 2;
     */
    title: string;
    /**
     * Type of the OS.
     *
     * @generated from field: fits.api.mvm.v1.OSType type = 3;
     */
    type: OSType;
};
/**
 * Describes the message fits.api.mvm.v1.OS.
 * Use `create(OSSchema)` to create a new message.
 */
export declare const OSSchema: GenMessage<OS>;
/**
 * OSServiceListRequest is the request payload for a OS list request.
 *
 * @generated from message fits.api.mvm.v1.OSServiceListRequest
 */
export type OSServiceListRequest = Message<"fits.api.mvm.v1.OSServiceListRequest"> & {};
/**
 * Describes the message fits.api.mvm.v1.OSServiceListRequest.
 * Use `create(OSServiceListRequestSchema)` to create a new message.
 */
export declare const OSServiceListRequestSchema: GenMessage<OSServiceListRequest>;
/**
 * OSServiceListResponse is the response payload for a OS list request
 *
 * @generated from message fits.api.mvm.v1.OSServiceListResponse
 */
export type OSServiceListResponse = Message<"fits.api.mvm.v1.OSServiceListResponse"> & {
    /**
     * The available operating systems.
     *
     * @generated from field: repeated fits.api.mvm.v1.OS operating_systems = 1;
     */
    operatingSystems: OS[];
};
/**
 * Describes the message fits.api.mvm.v1.OSServiceListResponse.
 * Use `create(OSServiceListResponseSchema)` to create a new message.
 */
export declare const OSServiceListResponseSchema: GenMessage<OSServiceListResponse>;
/**
 * OSType specifies the type of an operating system.
 *
 * @generated from enum fits.api.mvm.v1.OSType
 */
export declare enum OSType {
    /**
     * OS_TYPE_UNSPECIFIED is not specified.
     *
     * @generated from enum value: OS_TYPE_UNSPECIFIED = 0;
     */
    OS_TYPE_UNSPECIFIED = 0,
    /**
     * OS_TYPE_LINUX is a Linux operating system.
     *
     * @generated from enum value: OS_TYPE_LINUX = 1;
     */
    OS_TYPE_LINUX = 1,
    /**
     * OS_TYPE_WINDOWS is a Windows operating system.
     *
     * @generated from enum value: OS_TYPE_WINDOWS = 2;
     */
    OS_TYPE_WINDOWS = 2
}
/**
 * Describes the enum fits.api.mvm.v1.OSType.
 */
export declare const OSTypeSchema: GenEnum<OSType>;
/**
 * OSService lists operating systems available for managed VM (MVM) instances.
 *
 * @generated from service fits.api.mvm.v1.OSService
 */
export declare const OSService: GenService<{
    /**
     * Returns a list of all operating systems.
     *
     * @generated from rpc fits.api.mvm.v1.OSService.List
     */
    list: {
        methodKind: "unary";
        input: typeof OSServiceListRequestSchema;
        output: typeof OSServiceListResponseSchema;
    };
}>;
