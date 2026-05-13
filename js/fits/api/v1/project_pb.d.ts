import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/project.proto.
 */
export declare const file_fits_api_v1_project: GenFile;
/**
 * Project defines a group of resources belonging to a tenant
 * a tenant can have multiple projects
 *
 * @generated from message fits.api.v1.Project
 */
export type Project = Message<"fits.api.v1.Project"> & {
    /**
     * Uuid of this project
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Meta for this project
     *
     * @generated from field: fits.api.v1.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Name of this project must be unique per tenant
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Description of this project
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Tenant this project belongs to
     *
     * @generated from field: string tenant = 5;
     */
    tenant: string;
    /**
     * AvatarUrl of the Project
     *
     * @generated from field: optional string avatar_url = 6;
     */
    avatarUrl?: string | undefined;
};
/**
 * Describes the message fits.api.v1.Project.
 * Use `create(ProjectSchema)` to create a new message.
 */
export declare const ProjectSchema: GenMessage<Project>;
/**
 * ProjectServiceListRequest is the request payload to list all projects
 *
 * @generated from message fits.api.v1.ProjectServiceListRequest
 */
export type ProjectServiceListRequest = Message<"fits.api.v1.ProjectServiceListRequest"> & {
    /**
     * Id lists only projects with this id
     *
     * @generated from field: optional string id = 1;
     */
    id?: string | undefined;
    /**
     * Name lists only projects with this name
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Tenant lists only projects of this tenant
     *
     * @generated from field: optional string tenant = 3;
     */
    tenant?: string | undefined;
    /**
     * Labels lists only projects containing the given labels
     *
     * @generated from field: optional fits.api.v1.Labels labels = 4;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceListRequest.
 * Use `create(ProjectServiceListRequestSchema)` to create a new message.
 */
export declare const ProjectServiceListRequestSchema: GenMessage<ProjectServiceListRequest>;
/**
 * ProjectServiceListResponse is the response payload to list all projects
 *
 * @generated from message fits.api.v1.ProjectServiceListResponse
 */
export type ProjectServiceListResponse = Message<"fits.api.v1.ProjectServiceListResponse"> & {
    /**
     * Projects is a list of all your projects
     *
     * @generated from field: repeated fits.api.v1.Project projects = 1;
     */
    projects: Project[];
};
/**
 * Describes the message fits.api.v1.ProjectServiceListResponse.
 * Use `create(ProjectServiceListResponseSchema)` to create a new message.
 */
export declare const ProjectServiceListResponseSchema: GenMessage<ProjectServiceListResponse>;
/**
 * ProjectServiceGetRequest is the request payload to get a project
 *
 * @generated from message fits.api.v1.ProjectServiceGetRequest
 */
export type ProjectServiceGetRequest = Message<"fits.api.v1.ProjectServiceGetRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message fits.api.v1.ProjectServiceGetRequest.
 * Use `create(ProjectServiceGetRequestSchema)` to create a new message.
 */
export declare const ProjectServiceGetRequestSchema: GenMessage<ProjectServiceGetRequest>;
/**
 * ProjectServiceGetResponse is the response payload to get a projects
 *
 * @generated from message fits.api.v1.ProjectServiceGetResponse
 */
export type ProjectServiceGetResponse = Message<"fits.api.v1.ProjectServiceGetResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: fits.api.v1.Project project = 1;
     */
    project?: Project | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceGetResponse.
 * Use `create(ProjectServiceGetResponseSchema)` to create a new message.
 */
export declare const ProjectServiceGetResponseSchema: GenMessage<ProjectServiceGetResponse>;
/**
 * ProjectServiceCreateRequest is the request payload to Create a project
 *
 * @generated from message fits.api.v1.ProjectServiceCreateRequest
 */
export type ProjectServiceCreateRequest = Message<"fits.api.v1.ProjectServiceCreateRequest"> & {
    /**
     * Login is the tenant of this project
     * TODO: is login really a good name?
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Name of this project, unique per tenant
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Description of this project
     *
     * @generated from field: string description = 3;
     */
    description: string;
    /**
     * Avatar URL of the project
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the project
     *
     * @generated from field: fits.api.v1.Labels labels = 5;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceCreateRequest.
 * Use `create(ProjectServiceCreateRequestSchema)` to create a new message.
 */
export declare const ProjectServiceCreateRequestSchema: GenMessage<ProjectServiceCreateRequest>;
/**
 * ProjectServiceCreateResponse is the response payload of creation of a project
 *
 * @generated from message fits.api.v1.ProjectServiceCreateResponse
 */
export type ProjectServiceCreateResponse = Message<"fits.api.v1.ProjectServiceCreateResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: fits.api.v1.Project project = 1;
     */
    project?: Project | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceCreateResponse.
 * Use `create(ProjectServiceCreateResponseSchema)` to create a new message.
 */
export declare const ProjectServiceCreateResponseSchema: GenMessage<ProjectServiceCreateResponse>;
/**
 * ProjectServiceDeleteRequest is the request payload to delete a project
 *
 * @generated from message fits.api.v1.ProjectServiceDeleteRequest
 */
export type ProjectServiceDeleteRequest = Message<"fits.api.v1.ProjectServiceDeleteRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
};
/**
 * Describes the message fits.api.v1.ProjectServiceDeleteRequest.
 * Use `create(ProjectServiceDeleteRequestSchema)` to create a new message.
 */
export declare const ProjectServiceDeleteRequestSchema: GenMessage<ProjectServiceDeleteRequest>;
/**
 * ProjectServiceDeleteResponse is the response payload to delete a project
 *
 * @generated from message fits.api.v1.ProjectServiceDeleteResponse
 */
export type ProjectServiceDeleteResponse = Message<"fits.api.v1.ProjectServiceDeleteResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: fits.api.v1.Project project = 1;
     */
    project?: Project | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceDeleteResponse.
 * Use `create(ProjectServiceDeleteResponseSchema)` to create a new message.
 */
export declare const ProjectServiceDeleteResponseSchema: GenMessage<ProjectServiceDeleteResponse>;
/**
 * ProjectServiceUpdateRequest is the request payload to update a project
 *
 * @generated from message fits.api.v1.ProjectServiceUpdateRequest
 */
export type ProjectServiceUpdateRequest = Message<"fits.api.v1.ProjectServiceUpdateRequest"> & {
    /**
     * Project is the uuid of the project to get
     *
     * @generated from field: string project = 1;
     */
    project: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: fits.api.v1.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Name of this project unique per tenant
     *
     * @generated from field: optional string name = 3;
     */
    name?: string | undefined;
    /**
     * Description of this project
     *
     * @generated from field: optional string description = 4;
     */
    description?: string | undefined;
    /**
     * Avatar URL of the project
     *
     * @generated from field: optional string avatar_url = 5;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on this project
     *
     * @generated from field: optional fits.api.v1.UpdateLabels labels = 6;
     */
    labels?: UpdateLabels | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceUpdateRequest.
 * Use `create(ProjectServiceUpdateRequestSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateRequestSchema: GenMessage<ProjectServiceUpdateRequest>;
/**
 * ProjectServiceUpdateResponse is the response payload to update a project
 *
 * @generated from message fits.api.v1.ProjectServiceUpdateResponse
 */
export type ProjectServiceUpdateResponse = Message<"fits.api.v1.ProjectServiceUpdateResponse"> & {
    /**
     * Project is the project
     *
     * @generated from field: fits.api.v1.Project project = 1;
     */
    project?: Project | undefined;
};
/**
 * Describes the message fits.api.v1.ProjectServiceUpdateResponse.
 * Use `create(ProjectServiceUpdateResponseSchema)` to create a new message.
 */
export declare const ProjectServiceUpdateResponseSchema: GenMessage<ProjectServiceUpdateResponse>;
/**
 * ProjectService provides project management operations.
 *
 * @generated from service fits.api.v1.ProjectService
 */
export declare const ProjectService: GenService<{
    /**
     * Returns the list of all accessible projects.
     *
     * @generated from rpc fits.api.v1.ProjectService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ProjectServiceListRequestSchema;
        output: typeof ProjectServiceListResponseSchema;
    };
    /**
     * Returns the project with the specified UUID.
     *
     * @generated from rpc fits.api.v1.ProjectService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof ProjectServiceGetRequestSchema;
        output: typeof ProjectServiceGetResponseSchema;
    };
    /**
     * Create a project
     *
     * @generated from rpc fits.api.v1.ProjectService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof ProjectServiceCreateRequestSchema;
        output: typeof ProjectServiceCreateResponseSchema;
    };
    /**
     * Delete a project
     *
     * @generated from rpc fits.api.v1.ProjectService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof ProjectServiceDeleteRequestSchema;
        output: typeof ProjectServiceDeleteResponseSchema;
    };
    /**
     * Update a project
     *
     * @generated from rpc fits.api.v1.ProjectService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof ProjectServiceUpdateRequestSchema;
        output: typeof ProjectServiceUpdateResponseSchema;
    };
}>;
