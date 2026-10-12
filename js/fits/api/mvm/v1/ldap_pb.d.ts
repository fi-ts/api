import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/ldap.proto.
 */
export declare const file_fits_api_mvm_v1_ldap: GenFile;
/**
 * LDAP is an LDAP directory a Linux MVM can be integrated with.
 *
 * @generated from message fits.api.mvm.v1.LDAP
 */
export type LDAP = Message<"fits.api.mvm.v1.LDAP"> & {
    /**
     * UUID of this LDAP directory.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * FQDN of this LDAP directory.
     *
     * @generated from field: string fqdn = 2;
     */
    fqdn: string;
    /**
     * Tenant this LDAP directory belongs to.
     *
     * @generated from field: string tenant = 3;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.LDAP.
 * Use `create(LDAPSchema)` to create a new message.
 */
export declare const LDAPSchema: GenMessage<LDAP>;
/**
 * LDAPServiceListRequest is the request payload for an LDAP list request.
 *
 * @generated from message fits.api.mvm.v1.LDAPServiceListRequest
 */
export type LDAPServiceListRequest = Message<"fits.api.mvm.v1.LDAPServiceListRequest"> & {
    /**
     * Tenant to list available LDAP directories for.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.LDAPServiceListRequest.
 * Use `create(LDAPServiceListRequestSchema)` to create a new message.
 */
export declare const LDAPServiceListRequestSchema: GenMessage<LDAPServiceListRequest>;
/**
 * LDAPServiceListResponse is the response payload for an LDAP list request.
 *
 * @generated from message fits.api.mvm.v1.LDAPServiceListResponse
 */
export type LDAPServiceListResponse = Message<"fits.api.mvm.v1.LDAPServiceListResponse"> & {
    /**
     * The available LDAP directories.
     *
     * @generated from field: repeated fits.api.mvm.v1.LDAP ldaps = 1;
     */
    ldaps: LDAP[];
};
/**
 * Describes the message fits.api.mvm.v1.LDAPServiceListResponse.
 * Use `create(LDAPServiceListResponseSchema)` to create a new message.
 */
export declare const LDAPServiceListResponseSchema: GenMessage<LDAPServiceListResponse>;
/**
 * LDAPService lists LDAP directories available for managed VM (MVM) instances.
 *
 * @generated from service fits.api.mvm.v1.LDAPService
 */
export declare const LDAPService: GenService<{
    /**
     * Returns a list of all LDAP directories available for a tenant.
     *
     * @generated from rpc fits.api.mvm.v1.LDAPService.List
     */
    list: {
        methodKind: "unary";
        input: typeof LDAPServiceListRequestSchema;
        output: typeof LDAPServiceListResponseSchema;
    };
}>;
