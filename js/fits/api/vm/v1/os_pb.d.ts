import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/vm/v1/os.proto.
 */
export declare const file_fits_api_vm_v1_os: GenFile;
/**
 * Os is the definition of a potential OS of a VM instance
 * TODO create a dedicated service
 *
 * @generated from message fits.api.vm.v1.Os
 */
export type Os = Message<"fits.api.vm.v1.Os"> & {
    /**
     * Uuid of this os
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.vm.v1.Os.
 * Use `create(OsSchema)` to create a new message.
 */
export declare const OsSchema: GenMessage<Os>;
