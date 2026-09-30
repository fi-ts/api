import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/location.proto.
 */
export declare const file_fits_api_mvm_v1_location: GenFile;
/**
 * Location is the definition of an available datacenter location for MVM instances.
 *
 * @generated from message fits.api.mvm.v1.Location
 */
export type Location = Message<"fits.api.mvm.v1.Location"> & {
    /**
     * Uuid of this location
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * title of the location
     *
     * @generated from field: string title = 2;
     */
    title: string;
};
/**
 * Describes the message fits.api.mvm.v1.Location.
 * Use `create(LocationSchema)` to create a new message.
 */
export declare const LocationSchema: GenMessage<Location>;
/**
 * LocationServiceListRequest is the request payload for a location list request.
 *
 * @generated from message fits.api.mvm.v1.LocationServiceListRequest
 */
export type LocationServiceListRequest = Message<"fits.api.mvm.v1.LocationServiceListRequest"> & {
    /**
     * Tenant to list available locations for
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.LocationServiceListRequest.
 * Use `create(LocationServiceListRequestSchema)` to create a new message.
 */
export declare const LocationServiceListRequestSchema: GenMessage<LocationServiceListRequest>;
/**
 * LocationServiceListResponse is the response payload for a location list request.
 *
 * @generated from message fits.api.mvm.v1.LocationServiceListResponse
 */
export type LocationServiceListResponse = Message<"fits.api.mvm.v1.LocationServiceListResponse"> & {
    /**
     * The available locations
     *
     * @generated from field: repeated fits.api.mvm.v1.Location locations = 1;
     */
    locations: Location[];
};
/**
 * Describes the message fits.api.mvm.v1.LocationServiceListResponse.
 * Use `create(LocationServiceListResponseSchema)` to create a new message.
 */
export declare const LocationServiceListResponseSchema: GenMessage<LocationServiceListResponse>;
/**
 * LocationService lists datacenter locations available for managed VM (MVM) instances, scoped per tenant.
 *
 * @generated from service fits.api.mvm.v1.LocationService
 */
export declare const LocationService: GenService<{
    /**
     * Returns a list of all datacenter locations available for a tenant.
     *
     * @generated from rpc fits.api.mvm.v1.LocationService.List
     */
    list: {
        methodKind: "unary";
        input: typeof LocationServiceListRequestSchema;
        output: typeof LocationServiceListResponseSchema;
    };
}>;
