import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { UpdateMeta } from "../../v1/common_pb.js";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/mvm/v1/contact.proto.
 */
export declare const file_fits_api_mvm_v1_contact: GenFile;
/**
 * Contact is a contact person an MVM instance can be assigned to.
 *
 * @generated from message fits.api.mvm.v1.Contact
 */
export type Contact = Message<"fits.api.mvm.v1.Contact"> & {
    /**
     * Uuid of this contact.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * First name of the contact.
     *
     * @generated from field: string first_name = 2;
     */
    firstName: string;
    /**
     * Last name of the contact.
     *
     * @generated from field: string last_name = 3;
     */
    lastName: string;
    /**
     * Department of the contact.
     *
     * @generated from field: string department = 4;
     */
    department: string;
    /**
     * Phone number of the contact.
     *
     * @generated from field: string phone = 5;
     */
    phone: string;
    /**
     * Email address of the contact.
     *
     * @generated from field: string email = 6;
     */
    email: string;
    /**
     * Tenant this contact belongs to.
     *
     * @generated from field: string tenant = 7;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.Contact.
 * Use `create(ContactSchema)` to create a new message.
 */
export declare const ContactSchema: GenMessage<Contact>;
/**
 * ContactServiceListRequest is the request payload for a contact list request.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceListRequest
 */
export type ContactServiceListRequest = Message<"fits.api.mvm.v1.ContactServiceListRequest"> & {
    /**
     * Tenant to list available contact persons for.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceListRequest.
 * Use `create(ContactServiceListRequestSchema)` to create a new message.
 */
export declare const ContactServiceListRequestSchema: GenMessage<ContactServiceListRequest>;
/**
 * ContactServiceListResponse is the response payload for a contact list request.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceListResponse
 */
export type ContactServiceListResponse = Message<"fits.api.mvm.v1.ContactServiceListResponse"> & {
    /**
     * The available contact persons.
     *
     * @generated from field: repeated fits.api.mvm.v1.Contact contacts = 1;
     */
    contacts: Contact[];
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceListResponse.
 * Use `create(ContactServiceListResponseSchema)` to create a new message.
 */
export declare const ContactServiceListResponseSchema: GenMessage<ContactServiceListResponse>;
/**
 * ContactServiceCreateRequest is the request payload for a contact create request.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceCreateRequest
 */
export type ContactServiceCreateRequest = Message<"fits.api.mvm.v1.ContactServiceCreateRequest"> & {
    /**
     * Tenant to create the contact person for.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * First name of the contact.
     *
     * @generated from field: string first_name = 2;
     */
    firstName: string;
    /**
     * Last name of the contact.
     *
     * @generated from field: string last_name = 3;
     */
    lastName: string;
    /**
     * Department of the contact.
     *
     * @generated from field: string department = 4;
     */
    department: string;
    /**
     * Phone number of the contact.
     *
     * @generated from field: string phone = 5;
     */
    phone: string;
    /**
     * Email address of the contact.
     *
     * @generated from field: string email = 6;
     */
    email: string;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceCreateRequest.
 * Use `create(ContactServiceCreateRequestSchema)` to create a new message.
 */
export declare const ContactServiceCreateRequestSchema: GenMessage<ContactServiceCreateRequest>;
/**
 * Response for contact create requests.
 * TODO: check whether create is really asynchronous; if it completes
 * synchronously, this response should return the created Contact instead
 * of just its UUID.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceCreateResponse
 */
export type ContactServiceCreateResponse = Message<"fits.api.mvm.v1.ContactServiceCreateResponse"> & {
    /**
     * UUID of the created contact.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceCreateResponse.
 * Use `create(ContactServiceCreateResponseSchema)` to create a new message.
 */
export declare const ContactServiceCreateResponseSchema: GenMessage<ContactServiceCreateResponse>;
/**
 * ContactServiceUpdateRequest is the request payload for a contact update request.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceUpdateRequest
 */
export type ContactServiceUpdateRequest = Message<"fits.api.mvm.v1.ContactServiceUpdateRequest"> & {
    /**
     * Tenant the contact belongs to.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * UUID of the contact to update.
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 3;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * First name of the contact.
     *
     * @generated from field: string first_name = 4;
     */
    firstName: string;
    /**
     * Last name of the contact.
     *
     * @generated from field: string last_name = 5;
     */
    lastName: string;
    /**
     * Department of the contact.
     *
     * @generated from field: string department = 6;
     */
    department: string;
    /**
     * Phone number of the contact.
     *
     * @generated from field: string phone = 7;
     */
    phone: string;
    /**
     * Email address of the contact.
     *
     * @generated from field: string email = 8;
     */
    email: string;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceUpdateRequest.
 * Use `create(ContactServiceUpdateRequestSchema)` to create a new message.
 */
export declare const ContactServiceUpdateRequestSchema: GenMessage<ContactServiceUpdateRequest>;
/**
 * Response for Update; an empty response means the change was accepted.
 * TODO: check whether update is really asynchronous. If it completes
 * synchronously, returning the updated Contact here is correct. If it is
 * asynchronous, this response should become a minimal ack instead, because
 * the returned entity would only be a snapshot at acceptance time and
 * clients must still poll List for the final state.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceUpdateResponse
 */
export type ContactServiceUpdateResponse = Message<"fits.api.mvm.v1.ContactServiceUpdateResponse"> & {
    /**
     * The updated Contact.
     *
     * @generated from field: fits.api.mvm.v1.Contact contact = 1;
     */
    contact?: Contact | undefined;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceUpdateResponse.
 * Use `create(ContactServiceUpdateResponseSchema)` to create a new message.
 */
export declare const ContactServiceUpdateResponseSchema: GenMessage<ContactServiceUpdateResponse>;
/**
 * ContactServiceDeleteRequest is the request payload for a contact delete request.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceDeleteRequest
 */
export type ContactServiceDeleteRequest = Message<"fits.api.mvm.v1.ContactServiceDeleteRequest"> & {
    /**
     * Tenant the contact belongs to.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * UUID of the contact to delete.
     *
     * @generated from field: string uuid = 2;
     */
    uuid: string;
};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceDeleteRequest.
 * Use `create(ContactServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ContactServiceDeleteRequestSchema: GenMessage<ContactServiceDeleteRequest>;
/**
 * Response for Delete; an empty response means the deletion was accepted.
 *
 * @generated from message fits.api.mvm.v1.ContactServiceDeleteResponse
 */
export type ContactServiceDeleteResponse = Message<"fits.api.mvm.v1.ContactServiceDeleteResponse"> & {};
/**
 * Describes the message fits.api.mvm.v1.ContactServiceDeleteResponse.
 * Use `create(ContactServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ContactServiceDeleteResponseSchema: GenMessage<ContactServiceDeleteResponse>;
/**
 * ContactService provides CRUD operations on the contact persons available for managed VM (MVM) instances, scoped per tenant.
 *
 * @generated from service fits.api.mvm.v1.ContactService
 */
export declare const ContactService: GenService<{
    /**
     * Returns a list of all contact persons available for a tenant.
     *
     * @generated from rpc fits.api.mvm.v1.ContactService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ContactServiceListRequestSchema;
        output: typeof ContactServiceListResponseSchema;
    };
    /**
     * Creates a new contact person.
     *
     * @generated from rpc fits.api.mvm.v1.ContactService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ContactServiceCreateRequestSchema;
        output: typeof ContactServiceCreateResponseSchema;
    };
    /**
     * Updates an existing contact person.
     *
     * @generated from rpc fits.api.mvm.v1.ContactService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof ContactServiceUpdateRequestSchema;
        output: typeof ContactServiceUpdateResponseSchema;
    };
    /**
     * Deletes a contact person.
     *
     * @generated from rpc fits.api.mvm.v1.ContactService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ContactServiceDeleteRequestSchema;
        output: typeof ContactServiceDeleteResponseSchema;
    };
}>;
