import type { GenEnum, GenExtension, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { EnumValueOptions, MethodOptions, Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/common.proto.
 */
export declare const file_fits_api_v1_common: GenFile;
/**
 * Paging defines paging for methods with a lot of results
 *
 * @generated from message fits.api.v1.Paging
 */
export type Paging = Message<"fits.api.v1.Paging"> & {
    /**
     * Page is used for pagination, if unset only the first page is returned,
     * the list response contains then the page number for the next page.
     *
     * @generated from field: optional uint64 page = 1;
     */
    page?: bigint | undefined;
    /**
     * Count is the number of results returned per page, if not given server side defaults apply
     *
     * @generated from field: optional uint64 count = 2;
     */
    count?: bigint | undefined;
};
/**
 * Describes the message fits.api.v1.Paging.
 * Use `create(PagingSchema)` to create a new message.
 */
export declare const PagingSchema: GenMessage<Paging>;
/**
 * Labels define additional information to a entity
 *
 * @generated from message fits.api.v1.Labels
 */
export type Labels = Message<"fits.api.v1.Labels"> & {
    /**
     * Labels consists labels
     *
     * @generated from field: map<string, string> labels = 1;
     */
    labels: {
        [key: string]: string;
    };
};
/**
 * Describes the message fits.api.v1.Labels.
 * Use `create(LabelsSchema)` to create a new message.
 */
export declare const LabelsSchema: GenMessage<Labels>;
/**
 * Meta of a message
 *
 * @generated from message fits.api.v1.Meta
 */
export type Meta = Message<"fits.api.v1.Meta"> & {
    /**
     * Tags on this entity
     *
     * @generated from field: optional fits.api.v1.Labels labels = 1;
     */
    labels?: Labels | undefined;
    /**
     * CreatedAt is the date when this entity was created
     *
     * @generated from field: google.protobuf.Timestamp created_at = 2;
     */
    createdAt?: Timestamp | undefined;
    /**
     * UpdatedAt is the date when this entity was updated
     * must be part of the update request to ensure optimistic locking
     *
     * @generated from field: google.protobuf.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp | undefined;
    /**
     * Generation identifies how often this entity was modified since creation.
     *
     * @generated from field: uint64 generation = 4;
     */
    generation: bigint;
};
/**
 * Describes the message fits.api.v1.Meta.
 * Use `create(MetaSchema)` to create a new message.
 */
export declare const MetaSchema: GenMessage<Meta>;
/**
 * UpdateLabels is a message to update labels
 *
 * @generated from message fits.api.v1.UpdateLabels
 */
export type UpdateLabels = Message<"fits.api.v1.UpdateLabels"> & {
    /**
     * Update labels. New ones will be added, existing ones overwritten
     *
     * @generated from field: fits.api.v1.Labels update = 1;
     */
    update?: Labels | undefined;
    /**
     * Remove labels by key
     *
     * @generated from field: repeated string remove = 2;
     */
    remove: string[];
};
/**
 * Describes the message fits.api.v1.UpdateLabels.
 * Use `create(UpdateLabelsSchema)` to create a new message.
 */
export declare const UpdateLabelsSchema: GenMessage<UpdateLabels>;
/**
 * UpdateMeta must be provided with every UpdateRequest to define how optimistic locking should be handled
 *
 * @generated from message fits.api.v1.UpdateMeta
 */
export type UpdateMeta = Message<"fits.api.v1.UpdateMeta"> & {
    /**
     * UpdatedAt is the date when this entity was updated
     * must be part of the update request to ensure optimistic locking
     *
     * @generated from field: google.protobuf.Timestamp updated_at = 1;
     */
    updatedAt?: Timestamp | undefined;
    /**
     * LockingStrategy to be used for this update request
     *
     * @generated from field: fits.api.v1.OptimisticLockingStrategy locking_strategy = 2;
     */
    lockingStrategy: OptimisticLockingStrategy;
};
/**
 * Describes the message fits.api.v1.UpdateMeta.
 * Use `create(UpdateMetaSchema)` to create a new message.
 */
export declare const UpdateMetaSchema: GenMessage<UpdateMeta>;
/**
 * TenantRole specifies what role a logged in user needs to call this tenant scoped service
 *
 * @generated from enum fits.api.v1.TenantRole
 */
export declare enum TenantRole {
    /**
     * TENANT_ROLE_UNSPECIFIED is not specified
     *
     * @generated from enum value: TENANT_ROLE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * TENANT_ROLE_OWNER the logged in user needs at least owner role to call this method
     *
     * @generated from enum value: TENANT_ROLE_OWNER = 1;
     */
    OWNER = 1,
    /**
     * TENANT_ROLE_EDITOR the logged in user needs at least editor role to call this method
     *
     * @generated from enum value: TENANT_ROLE_EDITOR = 2;
     */
    EDITOR = 2,
    /**
     * TENANT_ROLE_VIEWER the logged in user needs at least viewer role to call this method
     *
     * @generated from enum value: TENANT_ROLE_VIEWER = 3;
     */
    VIEWER = 3,
    /**
     * TENANT_ROLE_GUEST the logged in user needs at least guest role to call this method
     * The guest role is assumed by users who are invited to a tenant's project without them
     * having a direct membership within the tenant.
     *
     * @generated from enum value: TENANT_ROLE_GUEST = 4;
     */
    GUEST = 4
}
/**
 * Describes the enum fits.api.v1.TenantRole.
 */
export declare const TenantRoleSchema: GenEnum<TenantRole>;
/**
 * ProjectRole specifies what role a logged in user needs to call this project scoped service
 *
 * @generated from enum fits.api.v1.ProjectRole
 */
export declare enum ProjectRole {
    /**
     * PROJECT_ROLE_UNSPECIFIED is not specified
     *
     * @generated from enum value: PROJECT_ROLE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * PROJECT_ROLE_OWNER the logged in user needs at least owner role to call this method
     *
     * @generated from enum value: PROJECT_ROLE_OWNER = 1;
     */
    OWNER = 1,
    /**
     * PROJECT_ROLE_EDITOR the logged in user needs at least editor role to call this method
     *
     * @generated from enum value: PROJECT_ROLE_EDITOR = 2;
     */
    EDITOR = 2,
    /**
     * PROJECT_ROLE_VIEWER the logged in user needs at least viewer role to call this method
     *
     * @generated from enum value: PROJECT_ROLE_VIEWER = 3;
     */
    VIEWER = 3
}
/**
 * Describes the enum fits.api.v1.ProjectRole.
 */
export declare const ProjectRoleSchema: GenEnum<ProjectRole>;
/**
 * AdminRole specifies what role a logged in user needs to call this admin service
 *
 * @generated from enum fits.api.v1.AdminRole
 */
export declare enum AdminRole {
    /**
     * ADMIN_ROLE_UNSPECIFIED is not specified
     *
     * @generated from enum value: ADMIN_ROLE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * ADMIN_ROLE_EDITOR the logged in user needs at least editor role to call this method
     *
     * @generated from enum value: ADMIN_ROLE_EDITOR = 1;
     */
    EDITOR = 1,
    /**
     * ADMIN_ROLE_VIEWER the logged in user needs at least viewer role to call this method
     *
     * @generated from enum value: ADMIN_ROLE_VIEWER = 2;
     */
    VIEWER = 2
}
/**
 * Describes the enum fits.api.v1.AdminRole.
 */
export declare const AdminRoleSchema: GenEnum<AdminRole>;
/**
 * InfraRole specifies what role a microservice needs to call this infra service
 *
 * @generated from enum fits.api.v1.InfraRole
 */
export declare enum InfraRole {
    /**
     * INFRA_ROLE_UNSPECIFIED is not specified
     *
     * @generated from enum value: INFRA_ROLE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * INFRA_ROLE_EDITOR a microservice needs at least editor role to call this method
     *
     * @generated from enum value: INFRA_ROLE_EDITOR = 1;
     */
    EDITOR = 1,
    /**
     * INFRA_ROLE_VIEWER a microservice needs at least viewer role to call this method
     *
     * @generated from enum value: INFRA_ROLE_VIEWER = 2;
     */
    VIEWER = 2
}
/**
 * Describes the enum fits.api.v1.InfraRole.
 */
export declare const InfraRoleSchema: GenEnum<InfraRole>;
/**
 * MachineRole specifies what role a microservice needs to call this machine service
 *
 * @generated from enum fits.api.v1.MachineRole
 */
export declare enum MachineRole {
    /**
     * MACHINE_ROLE_UNSPECIFIED is not specified
     *
     * @generated from enum value: MACHINE_ROLE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_ROLE_EDITOR a microservice needs at least editor role to call this method
     *
     * @generated from enum value: MACHINE_ROLE_EDITOR = 1;
     */
    EDITOR = 1,
    /**
     * MACHINE_ROLE_VIEWER a microservice needs at least viewer role to call this method
     *
     * @generated from enum value: MACHINE_ROLE_VIEWER = 2;
     */
    VIEWER = 2
}
/**
 * Describes the enum fits.api.v1.MachineRole.
 */
export declare const MachineRoleSchema: GenEnum<MachineRole>;
/**
 * Visibility of a method
 *
 * @generated from enum fits.api.v1.Visibility
 */
export declare enum Visibility {
    /**
     * VISIBILITY_UNSPECIFIED is not defined
     *
     * @generated from enum value: VISIBILITY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * VISIBILITY_PUBLIC specifies that this service is accessible without authentication
     *
     * @generated from enum value: VISIBILITY_PUBLIC = 1;
     */
    PUBLIC = 1,
    /**
     * VISIBILITY_SELF enable call this endpoint from the authenticated user only
     *
     * @generated from enum value: VISIBILITY_SELF = 2;
     */
    SELF = 2
}
/**
 * Describes the enum fits.api.v1.Visibility.
 */
export declare const VisibilitySchema: GenEnum<Visibility>;
/**
 * Auditing option specified per service method
 * by default all service methods are included
 * add the auditing option if you want to exclude a method in auditing
 *
 * @generated from enum fits.api.v1.Auditing
 */
export declare enum Auditing {
    /**
     * AUDITING_UNSPECIFIED is not specified
     *
     * @generated from enum value: AUDITING_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * AUDITING_INCLUDED if a method is annotated with this, all calls are audited
     *
     * @generated from enum value: AUDITING_INCLUDED = 1;
     */
    INCLUDED = 1,
    /**
     * AUDITING_EXCLUDED if a method is annotated with this, no calls are audited
     *
     * @generated from enum value: AUDITING_EXCLUDED = 2;
     */
    EXCLUDED = 2
}
/**
 * Describes the enum fits.api.v1.Auditing.
 */
export declare const AuditingSchema: GenEnum<Auditing>;
/**
 * OptimisticLockingStrategy defines how optimistic locking should be handled.
 * It defaults to client side, which requires the UpdatedAt timestamp to be provided
 *
 * @generated from enum fits.api.v1.OptimisticLockingStrategy
 */
export declare enum OptimisticLockingStrategy {
    /**
     * OPTIMISTIC_LOCKING_STRATEGY_UNSPECIFIED same as client side
     *
     * @generated from enum value: OPTIMISTIC_LOCKING_STRATEGY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * OPTIMISTIC_LOCKING_STRATEGY_CLIENT requires UpdatedAt to be specified
     *
     * @generated from enum value: OPTIMISTIC_LOCKING_STRATEGY_CLIENT = 1;
     */
    CLIENT = 1,
    /**
     * OPTIMISTIC_LOCKING_STRATEGY_SERVER allows to omit UpdatedAt
     * and fetches the most recent entity before updating with the given values with the update request
     *
     * @generated from enum value: OPTIMISTIC_LOCKING_STRATEGY_SERVER = 2;
     */
    SERVER = 2
}
/**
 * Describes the enum fits.api.v1.OptimisticLockingStrategy.
 */
export declare const OptimisticLockingStrategySchema: GenEnum<OptimisticLockingStrategy>;
/**
 * TenantRoles are used to define which tenant role a logged in user must provide to call this method
 *
 * @generated from extension: repeated fits.api.v1.TenantRole tenant_roles = 51000;
 */
export declare const tenant_roles: GenExtension<MethodOptions, TenantRole[]>;
/**
 * ProjectRoles are used to define which project role a logged in user must provide to call this method
 *
 * @generated from extension: repeated fits.api.v1.ProjectRole project_roles = 51001;
 */
export declare const project_roles: GenExtension<MethodOptions, ProjectRole[]>;
/**
 * AdminRoles are used to define which admin role a logged in user must provide to call this method
 *
 * @generated from extension: repeated fits.api.v1.AdminRole admin_roles = 51002;
 */
export declare const admin_roles: GenExtension<MethodOptions, AdminRole[]>;
/**
 * Visibility defines the visibility of this method, this is used to have public or self visible methods
 *
 * @generated from extension: fits.api.v1.Visibility visibility = 51003;
 */
export declare const visibility: GenExtension<MethodOptions, Visibility>;
/**
 * Auditing defines if calls to this method should be audited or not
 *
 * @generated from extension: fits.api.v1.Auditing auditing = 51004;
 */
export declare const auditing: GenExtension<MethodOptions, Auditing>;
/**
 * StringValue which can be set to a enum
 *
 * @generated from extension: string enum_string_value = 52000;
 */
export declare const enum_string_value: GenExtension<EnumValueOptions, string>;
