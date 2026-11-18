import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/stage.proto.
 */
export declare const file_fits_api_mvm_v1_stage: GenFile;
/**
 * Stages a MVM can belong to.
 *
 * @generated from message fits.api.mvm.v1.Stage
 */
export type Stage = Message<"fits.api.mvm.v1.Stage"> & {
    /**
     * UUID of this Stage.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Title of the Stage.
     *
     * @generated from field: string title = 2;
     */
    title: string;
};
/**
 * Describes the message fits.api.mvm.v1.Stage.
 * Use `create(StageSchema)` to create a new message.
 */
export declare const StageSchema: GenMessage<Stage>;
/**
 * StageServiceListRequest is the request payload for a stage list request.
 *
 * @generated from message fits.api.mvm.v1.StageServiceListRequest
 */
export type StageServiceListRequest = Message<"fits.api.mvm.v1.StageServiceListRequest"> & {};
/**
 * Describes the message fits.api.mvm.v1.StageServiceListRequest.
 * Use `create(StageServiceListRequestSchema)` to create a new message.
 */
export declare const StageServiceListRequestSchema: GenMessage<StageServiceListRequest>;
/**
 * StageServiceListResponse is the response payload for a stage list request.
 *
 * @generated from message fits.api.mvm.v1.StageServiceListResponse
 */
export type StageServiceListResponse = Message<"fits.api.mvm.v1.StageServiceListResponse"> & {
    /**
     * The available stages.
     *
     * @generated from field: repeated fits.api.mvm.v1.Stage stages = 1;
     */
    stages: Stage[];
};
/**
 * Describes the message fits.api.mvm.v1.StageServiceListResponse.
 * Use `create(StageServiceListResponseSchema)` to create a new message.
 */
export declare const StageServiceListResponseSchema: GenMessage<StageServiceListResponse>;
/**
 * StageService lists available stages.
 *
 * @generated from service fits.api.mvm.v1.StageService
 */
export declare const StageService: GenService<{
    /**
     * Returns a list of all available stages.
     *
     * @generated from rpc fits.api.mvm.v1.StageService.List
     */
    list: {
        methodKind: "unary";
        input: typeof StageServiceListRequestSchema;
        output: typeof StageServiceListResponseSchema;
    };
}>;
