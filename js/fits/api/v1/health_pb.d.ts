import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/health.proto.
 */
export declare const file_fits_api_v1_health: GenFile;
/**
 * Health reports the health status of all services.
 *
 * @generated from message fits.api.v1.Health
 */
export type Health = Message<"fits.api.v1.Health"> & {
    /**
     * Services contains the health of all individual services
     *
     * @generated from field: repeated fits.api.v1.HealthStatus services = 1;
     */
    services: HealthStatus[];
};
/**
 * Describes the message fits.api.v1.Health.
 * Use `create(HealthSchema)` to create a new message.
 */
export declare const HealthSchema: GenMessage<Health>;
/**
 * HealthStatus represents the health status of one service.
 *
 * @generated from message fits.api.v1.HealthStatus
 */
export type HealthStatus = Message<"fits.api.v1.HealthStatus"> & {
    /**
     * Service is the service this health status is reported for
     *
     * @generated from field: fits.api.v1.Service service = 1;
     */
    service: Service;
    /**
     * Status is the status of this service
     *
     * @generated from field: fits.api.v1.ServiceStatus status = 2;
     */
    status: ServiceStatus;
    /**
     * Message describes the reason for the unhealthy status if possible
     *
     * @generated from field: string message = 3;
     */
    message: string;
};
/**
 * Describes the message fits.api.v1.HealthStatus.
 * Use `create(HealthStatusSchema)` to create a new message.
 */
export declare const HealthStatusSchema: GenMessage<HealthStatus>;
/**
 * Dependencies reports the health status of all dependencies.
 *
 * @generated from message fits.api.v1.Dependencies
 */
export type Dependencies = Message<"fits.api.v1.Dependencies"> & {
    /**
     * Dependencies contains the health of all individual dependencies
     *
     * @generated from field: repeated fits.api.v1.DependencyHealth dependencies = 1;
     */
    dependencies: DependencyHealth[];
};
/**
 * Describes the message fits.api.v1.Dependencies.
 * Use `create(DependenciesSchema)` to create a new message.
 */
export declare const DependenciesSchema: GenMessage<Dependencies>;
/**
 * DependencyHealth represents the health status of one dependency.
 *
 * @generated from message fits.api.v1.DependencyHealth
 */
export type DependencyHealth = Message<"fits.api.v1.DependencyHealth"> & {
    /**
     * Dependency is the dependency this health status is reported for
     *
     * @generated from field: fits.api.v1.Dependency dependency = 1;
     */
    dependency: Dependency;
    /**
     * Status is the status of this dependency
     *
     * @generated from field: fits.api.v1.ServiceStatus status = 2;
     */
    status: ServiceStatus;
    /**
     * Message describes the reason for the unhealthy status if possible
     *
     * @generated from field: string message = 3;
     */
    message: string;
};
/**
 * Describes the message fits.api.v1.DependencyHealth.
 * Use `create(DependencyHealthSchema)` to create a new message.
 */
export declare const DependencyHealthSchema: GenMessage<DependencyHealth>;
/**
 * HealthServiceGetRequest is the request payload for getting the health of the system.
 *
 * @generated from message fits.api.v1.HealthServiceGetRequest
 */
export type HealthServiceGetRequest = Message<"fits.api.v1.HealthServiceGetRequest"> & {};
/**
 * Describes the message fits.api.v1.HealthServiceGetRequest.
 * Use `create(HealthServiceGetRequestSchema)` to create a new message.
 */
export declare const HealthServiceGetRequestSchema: GenMessage<HealthServiceGetRequest>;
/**
 * HealthServiceGetResponse is the response payload for getting the health of the system.
 *
 * @generated from message fits.api.v1.HealthServiceGetResponse
 */
export type HealthServiceGetResponse = Message<"fits.api.v1.HealthServiceGetResponse"> & {
    /**
     * Health contains the overall health of the system
     *
     * @generated from field: fits.api.v1.Health health = 1;
     */
    health?: Health | undefined;
};
/**
 * Describes the message fits.api.v1.HealthServiceGetResponse.
 * Use `create(HealthServiceGetResponseSchema)` to create a new message.
 */
export declare const HealthServiceGetResponseSchema: GenMessage<HealthServiceGetResponse>;
/**
 * HealthServiceGetDependenciesRequest is the request payload for getting the health of the dependencies.
 *
 * @generated from message fits.api.v1.HealthServiceGetDependenciesRequest
 */
export type HealthServiceGetDependenciesRequest = Message<"fits.api.v1.HealthServiceGetDependenciesRequest"> & {};
/**
 * Describes the message fits.api.v1.HealthServiceGetDependenciesRequest.
 * Use `create(HealthServiceGetDependenciesRequestSchema)` to create a new message.
 */
export declare const HealthServiceGetDependenciesRequestSchema: GenMessage<HealthServiceGetDependenciesRequest>;
/**
 * HealthServiceGetDependenciesResponse is the response payload for getting the health of the dependencies.
 *
 * @generated from message fits.api.v1.HealthServiceGetDependenciesResponse
 */
export type HealthServiceGetDependenciesResponse = Message<"fits.api.v1.HealthServiceGetDependenciesResponse"> & {
    /**
     * Dependencies contains the health of all dependencies
     *
     * @generated from field: fits.api.v1.Dependencies dependencies = 1;
     */
    dependencies?: Dependencies | undefined;
};
/**
 * Describes the message fits.api.v1.HealthServiceGetDependenciesResponse.
 * Use `create(HealthServiceGetDependenciesResponseSchema)` to create a new message.
 */
export declare const HealthServiceGetDependenciesResponseSchema: GenMessage<HealthServiceGetDependenciesResponse>;
/**
 * ServiceStatus defines the status of a service.
 *
 * @generated from enum fits.api.v1.ServiceStatus
 */
export declare enum ServiceStatus {
    /**
     * SERVICE_STATUS_UNSPECIFIED service status is not known or unspecified
     *
     * @generated from enum value: SERVICE_STATUS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SERVICE_STATUS_DEGRADED the service is in degraded status, not the whole functionality is available
     *
     * @generated from enum value: SERVICE_STATUS_DEGRADED = 1;
     */
    DEGRADED = 1,
    /**
     * SERVICE_STATUS_UNHEALTHY the service is in unhealthy status, serious impact is expected
     *
     * @generated from enum value: SERVICE_STATUS_UNHEALTHY = 2;
     */
    UNHEALTHY = 2,
    /**
     * SERVICE_STATUS_HEALTHY the service is in healthy status, for example fully functional
     *
     * @generated from enum value: SERVICE_STATUS_HEALTHY = 3;
     */
    HEALTHY = 3
}
/**
 * Describes the enum fits.api.v1.ServiceStatus.
 */
export declare const ServiceStatusSchema: GenEnum<ServiceStatus>;
/**
 * Service defines the service for which the healthiness is reported.
 *
 * @generated from enum fits.api.v1.Service
 */
export declare enum Service {
    /**
     * SERVICE_UNSPECIFIED is a unknown service
     *
     * @generated from enum value: SERVICE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SERVICE_MVM the managed vm service
     *
     * @generated from enum value: SERVICE_MVM = 1;
     */
    MVM = 1,
    /**
     * SERVICE_PCVM the public cloud vm service
     *
     * @generated from enum value: SERVICE_PCVM = 2;
     */
    PCVM = 2
}
/**
 * Describes the enum fits.api.v1.Service.
 */
export declare const ServiceSchema: GenEnum<Service>;
/**
 * Dependency defines the dependency for which the healthiness is reported.
 *
 * @generated from enum fits.api.v1.Dependency
 */
export declare enum Dependency {
    /**
     * DEPENDENCY_UNSPECIFIED is a unknown dependency
     *
     * @generated from enum value: DEPENDENCY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * DEPENDENCY_VALKEY the valkey service
     *
     * @generated from enum value: DEPENDENCY_VALKEY = 1;
     */
    VALKEY = 1,
    /**
     * DEPENDENCY_TENANT_APISERVER the tenant apiserver service
     *
     * @generated from enum value: DEPENDENCY_TENANT_APISERVER = 2;
     */
    TENANT_APISERVER = 2
}
/**
 * Describes the enum fits.api.v1.Dependency.
 */
export declare const DependencySchema: GenEnum<Dependency>;
/**
 * HealthService provides health status monitoring operations.
 *
 * @generated from service fits.api.v1.HealthService
 */
export declare const HealthService: GenService<{
    /**
     * Returns the health status of the platform.
     *
     * @generated from rpc fits.api.v1.HealthService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof HealthServiceGetRequestSchema;
        output: typeof HealthServiceGetResponseSchema;
    };
    /**
     * Returns the health status of the platform's dependencies.
     *
     * @generated from rpc fits.api.v1.HealthService.GetDependencies
     */
    getDependencies: {
        methodKind: "unary";
        input: typeof HealthServiceGetDependenciesRequestSchema;
        output: typeof HealthServiceGetDependenciesResponseSchema;
    };
}>;
