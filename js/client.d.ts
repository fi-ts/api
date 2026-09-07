import { Interceptor } from "@connectrpc/connect";
import type { Client as ConnectClient } from "@connectrpc/connect";
import { HealthService as Apiv1HealthService } from "./fits/api/v1/health_pb";
import { IPService as Apiv1IPService } from "./fits/api/v1/ip_pb";
import { MethodService as Apiv1MethodService } from "./fits/api/v1/method_pb";
import { ProjectService as Apiv1ProjectService } from "./fits/api/v1/project_pb";
import { TenantService as Apiv1TenantService } from "./fits/api/v1/tenant_pb";
import { TokenService as Apiv1TokenService } from "./fits/api/v1/token_pb";
import { VersionService as Apiv1VersionService } from "./fits/api/v1/version_pb";
export interface ClientConfig {
    baseUrl: string;
    token?: string;
    interceptors?: Interceptor[];
}
export interface Client {
    apiv1(): Apiv1;
}
export interface Apiv1 {
    health(): ConnectClient<typeof Apiv1HealthService>;
    ip(): ConnectClient<typeof Apiv1IPService>;
    method(): ConnectClient<typeof Apiv1MethodService>;
    project(): ConnectClient<typeof Apiv1ProjectService>;
    tenant(): ConnectClient<typeof Apiv1TenantService>;
    token(): ConnectClient<typeof Apiv1TokenService>;
    version(): ConnectClient<typeof Apiv1VersionService>;
}
export declare function newClient(config: ClientConfig): Client;
