import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { ManagedVM, MVMServiceCreateRequest, MVMServiceCreateResponse, MVMServiceDeleteRequest, MVMServiceDeleteResponse, MVMServiceGetRequest, MVMServiceListRequest } from "./mvm_pb.js";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/mvm_on_behalf.proto.
 */
export declare const file_fits_api_mvm_v1_mvm_on_behalf: GenFile;
/**
 * An MVM with additional fields beyond the ManagedVM.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVM
 */
export type OnBehalfMVM = Message<"fits.api.mvm.v1.OnBehalfMVM"> & {
    /**
     * The public MVM.
     *
     * @generated from field: fits.api.mvm.v1.ManagedVM mvm = 1;
     */
    mvm?: ManagedVM | undefined;
    /**
     * Determines if the MVM is billable.
     *
     * @generated from field: bool billable = 2;
     */
    billable: boolean;
    /**
     * The start date when the MVM will be billable.
     *
     * @generated from field: optional google.protobuf.Timestamp billable_from = 3;
     */
    billableFrom?: Timestamp | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVM.
 * Use `create(OnBehalfMVMSchema)` to create a new message.
 */
export declare const OnBehalfMVMSchema: GenMessage<OnBehalfMVM>;
/**
 * Request to get an MVM on behalf of a customer tenant.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceGetRequest
 */
export type OnBehalfMVMServiceGetRequest = Message<"fits.api.mvm.v1.OnBehalfMVMServiceGetRequest"> & {
    /**
     * The public get request.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceGetRequest request = 1;
     */
    request?: MVMServiceGetRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceGetRequest.
 * Use `create(OnBehalfMVMServiceGetRequestSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceGetRequestSchema: GenMessage<OnBehalfMVMServiceGetRequest>;
/**
 * Response for an on-behalf MVM get.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceGetResponse
 */
export type OnBehalfMVMServiceGetResponse = Message<"fits.api.mvm.v1.OnBehalfMVMServiceGetResponse"> & {
    /**
     * The MVM with additional fields.
     *
     * @generated from field: fits.api.mvm.v1.OnBehalfMVM mvm = 1;
     */
    mvm?: OnBehalfMVM | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceGetResponse.
 * Use `create(OnBehalfMVMServiceGetResponseSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceGetResponseSchema: GenMessage<OnBehalfMVMServiceGetResponse>;
/**
 * Request to list MVMs on behalf of a customer tenant.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceListRequest
 */
export type OnBehalfMVMServiceListRequest = Message<"fits.api.mvm.v1.OnBehalfMVMServiceListRequest"> & {
    /**
     * The public list request.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceListRequest request = 1;
     */
    request?: MVMServiceListRequest | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceListRequest.
 * Use `create(OnBehalfMVMServiceListRequestSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceListRequestSchema: GenMessage<OnBehalfMVMServiceListRequest>;
/**
 * Response for an on-behalf MVM list.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceListResponse
 */
export type OnBehalfMVMServiceListResponse = Message<"fits.api.mvm.v1.OnBehalfMVMServiceListResponse"> & {
    /**
     * The MVMs with additional fields.
     *
     * @generated from field: repeated fits.api.mvm.v1.OnBehalfMVM mvms = 1;
     */
    mvms: OnBehalfMVM[];
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceListResponse.
 * Use `create(OnBehalfMVMServiceListResponseSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceListResponseSchema: GenMessage<OnBehalfMVMServiceListResponse>;
/**
 * Request to create an MVM on behalf of a customer tenant.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceCreateRequest
 */
export type OnBehalfMVMServiceCreateRequest = Message<"fits.api.mvm.v1.OnBehalfMVMServiceCreateRequest"> & {
    /**
     * The public create request.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceCreateRequest request = 1;
     */
    request?: MVMServiceCreateRequest | undefined;
    /**
     * Determines if the MVM is billable.
     *
     * @generated from field: bool billable = 2;
     */
    billable: boolean;
    /**
     * The start date when the MVM will be billable.
     *
     * @generated from field: optional google.protobuf.Timestamp billable_from = 3;
     */
    billableFrom?: Timestamp | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceCreateRequest.
 * Use `create(OnBehalfMVMServiceCreateRequestSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceCreateRequestSchema: GenMessage<OnBehalfMVMServiceCreateRequest>;
/**
 * Response for an on-behalf MVM create.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceCreateResponse
 */
export type OnBehalfMVMServiceCreateResponse = Message<"fits.api.mvm.v1.OnBehalfMVMServiceCreateResponse"> & {
    /**
     * The public create response.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceCreateResponse response = 1;
     */
    response?: MVMServiceCreateResponse | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceCreateResponse.
 * Use `create(OnBehalfMVMServiceCreateResponseSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceCreateResponseSchema: GenMessage<OnBehalfMVMServiceCreateResponse>;
/**
 * Request to delete an MVM on behalf of a customer tenant.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceDeleteRequest
 */
export type OnBehalfMVMServiceDeleteRequest = Message<"fits.api.mvm.v1.OnBehalfMVMServiceDeleteRequest"> & {
    /**
     * The public delete request.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceDeleteRequest request = 1;
     */
    request?: MVMServiceDeleteRequest | undefined;
    /**
     * The future date when the MVM will be deleted, or immediately if unset.
     *
     * @generated from field: optional google.protobuf.Timestamp execution_date = 2;
     */
    executionDate?: Timestamp | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceDeleteRequest.
 * Use `create(OnBehalfMVMServiceDeleteRequestSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceDeleteRequestSchema: GenMessage<OnBehalfMVMServiceDeleteRequest>;
/**
 * Response for an on-behalf MVM delete.
 *
 * @generated from message fits.api.mvm.v1.OnBehalfMVMServiceDeleteResponse
 */
export type OnBehalfMVMServiceDeleteResponse = Message<"fits.api.mvm.v1.OnBehalfMVMServiceDeleteResponse"> & {
    /**
     * The public delete response.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceDeleteResponse response = 1;
     */
    response?: MVMServiceDeleteResponse | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.OnBehalfMVMServiceDeleteResponse.
 * Use `create(OnBehalfMVMServiceDeleteResponseSchema)` to create a new message.
 */
export declare const OnBehalfMVMServiceDeleteResponseSchema: GenMessage<OnBehalfMVMServiceDeleteResponse>;
/**
 * OnBehalfMVMService offers MVM operations on behalf of a customer tenant.
 *
 * @generated from service fits.api.mvm.v1.OnBehalfMVMService
 */
export declare const OnBehalfMVMService: GenService<{
    /**
     * Get returns an MVM with additional fields.
     *
     * @generated from rpc fits.api.mvm.v1.OnBehalfMVMService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof OnBehalfMVMServiceGetRequestSchema;
        output: typeof OnBehalfMVMServiceGetResponseSchema;
    };
    /**
     * List returns all MVMs with additional fields.
     *
     * @generated from rpc fits.api.mvm.v1.OnBehalfMVMService.List
     */
    list: {
        methodKind: "unary";
        input: typeof OnBehalfMVMServiceListRequestSchema;
        output: typeof OnBehalfMVMServiceListResponseSchema;
    };
    /**
     * Create makes a new MVM and lets the operator set additional fields.
     *
     * @generated from rpc fits.api.mvm.v1.OnBehalfMVMService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof OnBehalfMVMServiceCreateRequestSchema;
        output: typeof OnBehalfMVMServiceCreateResponseSchema;
    };
    /**
     * Delete removes an MVM on behalf of a customer tenant.
     *
     * @generated from rpc fits.api.mvm.v1.OnBehalfMVMService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof OnBehalfMVMServiceDeleteRequestSchema;
        output: typeof OnBehalfMVMServiceDeleteResponseSchema;
    };
}>;
