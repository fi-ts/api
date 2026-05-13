# fi-ts.io api

[![Release](https://github.com/fi-ts/api/actions/workflows/main.yml/badge.svg)](https://github.com/fi-ts/api/actions/workflows/main.yml) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/fi-ts/api) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fi-ts/api) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/fi-ts/api) [![Go Report Card](https://goreportcard.com/badge/github.com/fi-ts/api)](https://goreportcard.com/report/github.com/fi-ts/api)

> [!IMPORTANT]
> Please note that this repository is still in alpha stage. For now, the primary API is still defined in the [metal-api](https://github.com/fi-ts/metal-api) project. We will make a special announcement when this repository becomes GA.

This is the V2 API of [fi-ts.io](https://fi-ts.io) to implement [MEP-4](https://fi-ts.io/docs/MEP-4-multi-tenancy-for-the-metal-api).

## Usage examples

Can be found in the [examples](examples/) folder.

## Conventions

Method options provide an intuitve and declarative way of providing annotations to service methods.
These are used for scoping api-methods, which are getting utilized for authentication, authorization and auditing (mainly in interceptors).

### Motivational Example

```proto
service IPService {
  // Get an ip
  rpc Get(IPServiceGetRequest) returns (IPServiceGetResponse) {
    option (project_roles) = PROJECT_ROLE_OWNER;
    option (project_roles) = PROJECT_ROLE_EDITOR;
    option (project_roles) = PROJECT_ROLE_VIEWER;
    option (auditing) = AUDITING_EXCLUDED;
  }
  // Create an ip
  rpc Create(IPServiceCreateRequest) returns (IPServiceCreateResponse) {
    option (project_roles) = PROJECT_ROLE_OWNER;
    option (project_roles) = PROJECT_ROLE_EDITOR;
  }
}

// IPServiceCreateRequest is the request payload for a ip create request
message IPServiceCreateRequest {
  // Network from which the IP should be created
  string network = 1 [(buf.validate.field).string = {
    min_len: 2
    max_len: 128
  }];
  // Project of the ip
  string project = 2 [(buf.validate.field).string.uuid = true];
  // Name of the ip
  optional string name = 3 [(buf.validate.field).string.(metalstack.api.v2.is_name) = true];
  // Description of the ip
  optional string description = 4 [(buf.validate.field).string.(metalstack.api.v2.is_description) = true];
  // IP if given try to create this ip if still available
  optional string ip = 5 [(buf.validate.field).string.ip = true];
  // Machine for which this ip should get created
  optional string machine = 6 [(buf.validate.field).string.uuid = true];
  // Labels to put onto the ip
  optional Labels labels = 7;
  // Type of the IP, ether ephemeral (default), or static
  optional IPType type = 8 [(buf.validate.field).enum.defined_only = true];
  // Addressfamily of the IP to create, defaults to ipv4
  optional IPAddressFamily address_family = 9 [(buf.validate.field).enum.defined_only = true];
}
```

In this example we can see the motivation behind the method options.

1. Get: can be issued by project owner, editor, viewer and is excluded from auditing
2. Allocate: can be used by project owner, editor
3. Both methods are project-scoped, since they are annotated by a project role -> Request object needs to have the **project** field in order to specify the target project of the service method

Further explanations are explained in the following.

### Auth

These options specify the RBAC of the api-endpoint.

| Option         | Description                                  | Values      | Explanation                                            |
| -------------- | -------------------------------------------- | ----------- | ------------------------------------------------------ |
| TENANT_ROLE\_  | Specifies the required tenant role           | UNSPECIFIED |                                                        |
|                |                                              | OWNER       | tenant owner                                           |
|                |                                              | EDITOR      | tenant editor                                          |
|                |                                              | VIEWER      | tenant viewer                                          |
|                |                                              | GUEST       | tenant guest                                           |
| PROJECT_ROLE\_ | Specifies the required project role          | UNSPECIFIED |                                                        |
|                |                                              | OWNER       | project owner                                          |
|                |                                              | EDITOR      | project editor                                         |
|                |                                              | VIEWER      | project viewer                                         |
| ADMIN_ROLE\_   | Specifies the required admin role            | UNSPECIFIED |                                                        |
|                |                                              | EDITOR      | admin editor                                           |
|                |                                              | VIEWER      | admin viewer                                           |
| VISIBILITY\_   | Specifies the visibility of the api-endpoint | UNSPECIFIED |                                                        |
|                |                                              | PUBLIC      | api-method is visible to public, a token is not needed |
|                |                                              | SELF        | api-method is scoped to owner resources                |

> [!IMPORTANT]
>
> Every operation needs at least an option, which references the scope of the request: **ROLE** or **VISIBILITY**

> [!CAUTION]
>
> If we use a Tenant or Project role, the request will be respectively scoped as Tenant or Project request.
> Tenant-Requests must have the field **login**, which is the tenant id and specifies the tenant on which the service-method is scoped.
> Project-Requests must have the field **project**, which is the project id and specifies the project on which the service-method is scoped.

### Auditing

For traceability we require to store audit-logs.

| Option     | Description                              | Values      | Explanation              |
| ---------- | ---------------------------------------- | ----------- | ------------------------ |
| AUDITING\_ | Specifies if the api-endpoint is audited | UNSPECIFIED | DEFAULT: included        |
|            |                                          | INCLUDED    | operation is audited     |
|            |                                          | EXCLUDED    | operation is not audited |
