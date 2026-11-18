import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { MVMServiceCreateRequest, MVMServiceCreateResponse } from "./mvm_pb.js";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/mvm_on_behalf.proto.
 */
export declare const file_fits_api_mvm_v1_mvm_on_behalf: GenFile;
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
 * OnBehalfMVMService offers MVM operations on behalf of a customer tenant.
 *
 * @generated from service fits.api.mvm.v1.OnBehalfMVMService
 */
export declare const OnBehalfMVMService: GenService<{
    /**
     * Create makes a new MVM and lets the operator set its billing fields.
     *
     * @generated from rpc fits.api.mvm.v1.OnBehalfMVMService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof OnBehalfMVMServiceCreateRequestSchema;
        output: typeof OnBehalfMVMServiceCreateResponseSchema;
    };
}>;
