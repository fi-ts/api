import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { MVMServiceCreateRequest, MVMServiceCreateResponse } from "./mvm_pb.js";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/mvm_provider.proto.
 */
export declare const file_fits_api_mvm_v1_mvm_provider: GenFile;
/**
 * Request to create an MVM in the provider context.
 *
 * @generated from message fits.api.mvm.v1.ProviderMVMServiceCreateRequest
 */
export type ProviderMVMServiceCreateRequest = Message<"fits.api.mvm.v1.ProviderMVMServiceCreateRequest"> & {
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
};
/**
 * Describes the message fits.api.mvm.v1.ProviderMVMServiceCreateRequest.
 * Use `create(ProviderMVMServiceCreateRequestSchema)` to create a new message.
 */
export declare const ProviderMVMServiceCreateRequestSchema: GenMessage<ProviderMVMServiceCreateRequest>;
/**
 * Response for a provider-context MVM create.
 *
 * @generated from message fits.api.mvm.v1.ProviderMVMServiceCreateResponse
 */
export type ProviderMVMServiceCreateResponse = Message<"fits.api.mvm.v1.ProviderMVMServiceCreateResponse"> & {
    /**
     * The public create response.
     *
     * @generated from field: fits.api.mvm.v1.MVMServiceCreateResponse response = 1;
     */
    response?: MVMServiceCreateResponse | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.ProviderMVMServiceCreateResponse.
 * Use `create(ProviderMVMServiceCreateResponseSchema)` to create a new message.
 */
export declare const ProviderMVMServiceCreateResponseSchema: GenMessage<ProviderMVMServiceCreateResponse>;
/**
 * ProviderMVMService offers MVM operations in the provider context.
 *
 * @generated from service fits.api.mvm.v1.ProviderMVMService
 */
export declare const ProviderMVMService: GenService<{
    /**
     * Create makes a new MVM and lets the operator set its billing fields.
     *
     * @generated from rpc fits.api.mvm.v1.ProviderMVMService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ProviderMVMServiceCreateRequestSchema;
        output: typeof ProviderMVMServiceCreateResponseSchema;
    };
}>;
