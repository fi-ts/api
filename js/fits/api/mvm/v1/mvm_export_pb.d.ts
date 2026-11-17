import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/mvm_export.proto.
 */
export declare const file_fits_api_mvm_v1_mvm_export: GenFile;
/**
 * TODO: remove together with TodoRemoveDummy.
 *
 * @generated from message fits.api.mvm.v1.TodoRemoveDummyRequest
 */
export type TodoRemoveDummyRequest = Message<"fits.api.mvm.v1.TodoRemoveDummyRequest"> & {
    /**
     * Tenant
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.TodoRemoveDummyRequest.
 * Use `create(TodoRemoveDummyRequestSchema)` to create a new message.
 */
export declare const TodoRemoveDummyRequestSchema: GenMessage<TodoRemoveDummyRequest>;
/**
 * TODO: remove together with TodoRemoveDummy.
 *
 * @generated from message fits.api.mvm.v1.TodoRemoveDummyResponse
 */
export type TodoRemoveDummyResponse = Message<"fits.api.mvm.v1.TodoRemoveDummyResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.TodoRemoveDummyResponse.
 * Use `create(TodoRemoveDummyResponseSchema)` to create a new message.
 */
export declare const TodoRemoveDummyResponseSchema: GenMessage<TodoRemoveDummyResponse>;
/**
 * Request to export MVM accounting information as a CSV file.
 *
 * @generated from message fits.api.mvm.v1.CsvExportRequest
 */
export type CsvExportRequest = Message<"fits.api.mvm.v1.CsvExportRequest"> & {
    /**
     * Tenant
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Start of the time range (inclusive).
     *
     * @generated from field: google.protobuf.Timestamp from = 2;
     */
    from?: Timestamp | undefined;
    /**
     * End of the time range (inclusive).
     *
     * @generated from field: google.protobuf.Timestamp to = 3;
     */
    to?: Timestamp | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.CsvExportRequest.
 * Use `create(CsvExportRequestSchema)` to create a new message.
 */
export declare const CsvExportRequestSchema: GenMessage<CsvExportRequest>;
/**
 * CSV export of the mapped MVM accounting information.
 *
 * @generated from message fits.api.mvm.v1.CsvExportResponse
 */
export type CsvExportResponse = Message<"fits.api.mvm.v1.CsvExportResponse"> & {
    /**
     * One chunk of the CSV.
     *
     * @generated from field: string chunk = 1;
     */
    chunk: string;
};
/**
 * Describes the message fits.api.mvm.v1.CsvExportResponse.
 * Use `create(CsvExportResponseSchema)` to create a new message.
 */
export declare const CsvExportResponseSchema: GenMessage<CsvExportResponse>;
/**
 * CsvExportService offers MVM operations in the provider context.
 *
 * @generated from service fits.api.mvm.v1.CsvExportService
 */
export declare const CsvExportService: GenService<{
    /**
     * Create makes a new MVM and lets the operator set its billing fields.
     *
     * @generated from rpc fits.api.mvm.v1.CsvExportService.CsvExport
     */
    csvExport: {
        methodKind: "server_streaming";
        input: typeof CsvExportRequestSchema;
        output: typeof CsvExportResponseSchema;
    };
    /**
     * TODO: remove this placeholder method.
     * protoc-gen-go-mcp skips streaming methods but still emits its fixed import block,
     * so a service with only streaming RPCs generates an mcp file with unused imports
     * that does not compile. This unary method only exists to make the generated code compile.
     * It will be removed once CsvExport moves to a different service.
     *
     * @generated from rpc fits.api.mvm.v1.CsvExportService.TodoRemoveDummy
     */
    todoRemoveDummy: {
        methodKind: "unary";
        input: typeof TodoRemoveDummyRequestSchema;
        output: typeof TodoRemoveDummyResponseSchema;
    };
}>;
