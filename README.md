# fi-ts.io api

[![Release](https://github.com/fi-ts/api/actions/workflows/main.yml/badge.svg)](https://github.com/fi-ts/api/actions/workflows/main.yml) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/fi-ts/api) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fi-ts/api) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/fi-ts/api) [![Go Report Card](https://goreportcard.com/badge/github.com/fi-ts/api)](https://goreportcard.com/report/github.com/fi-ts/api)

> [!IMPORTANT]
> Please note that this repository is still in alpha stage. We will make a special announcement when this repository becomes GA.

This is the API of fi-ts.io, defined in the `fits.api.v1` proto package.

## Usage examples

The API ships generated clients for Go, TypeScript and Python.

**Go** ([`go/client`](go/client)):

```go
import (
	"github.com/fi-ts/api/go/client"
	apiv1 "github.com/fi-ts/api/go/fits/api/v1"
)

c, err := client.New(&client.DialConfig{
	BaseURL: "https://api.fi-ts.io",
	Token:   token,
})
version, err := c.Apiv1().Version().Get(ctx, &apiv1.VersionServiceGetRequest{})
```

**TypeScript** ([`js/client.ts`](js/client.ts)):

```ts
import { newClient } from "@fi-ts/api/client";

const c = newClient({ baseUrl: "https://api.fi-ts.io", token });
const version = await c.apiv1().version().get({});
```

**Python** ([`python/fits/client`](python/fits/client)):

```python
from fits.api.v1 import version_pb2
from fits.client.client import Client

c = Client(baseurl="https://api.fi-ts.io")
version = c.apiv1().version().get(version_pb2.VersionServiceGetRequest())
```

## Conventions

Method options provide an intuitive and declarative way of providing annotations to service methods.
These are used for scoping api-methods, which are getting utilized for authentication, authorization and auditing (mainly in interceptors).

### Motivational Example

```proto
service ProjectService {
  // Returns the list of all accessible projects.
  rpc List(ProjectServiceListRequest) returns (ProjectServiceListResponse) {
    option (visibility) = VISIBILITY_SELF;
    option (auditing) = AUDITING_EXCLUDED;
  }
  // Returns the project with the specified UUID.
  rpc Get(ProjectServiceGetRequest) returns (ProjectServiceGetResponse) {
    option (project_roles) = PROJECT_ROLE_OWNER;
    option (project_roles) = PROJECT_ROLE_EDITOR;
    option (project_roles) = PROJECT_ROLE_VIEWER;
    option (auditing) = AUDITING_EXCLUDED;
  }
  // Create a project
  rpc Create(ProjectServiceCreateRequest) returns (ProjectServiceCreateResponse) {
    option (tenant_roles) = TENANT_ROLE_OWNER;
    option (tenant_roles) = TENANT_ROLE_EDITOR;
    option (auditing) = AUDITING_INCLUDED;
  }
}

// ProjectServiceCreateRequest is the request payload to create a project
message ProjectServiceCreateRequest {
  // Tenant is the tenant of this project (the tenant login)
  string tenant = 1 [(buf.validate.field).string.(fits.api.v1.is_name) = true];
  // Name of this project, unique per tenant
  string name = 2 [(buf.validate.field).string.(fits.api.v1.is_name) = true];
  // Description of this project
  string description = 3 [(buf.validate.field).string.(fits.api.v1.is_description) = true];
  // Avatar URL of the project
  optional string avatar_url = 4;
  // Labels on the project
  Labels labels = 5;
}
```

In this example we can see the motivation behind the method options.

1. List: is visible to the authenticated user only (`VISIBILITY_SELF`) and is excluded from auditing
2. Get: is project-scoped, can be issued by project owner, editor and viewer, and is excluded from auditing
3. Create: is tenant-scoped, can be issued by tenant owner and editor, and is audited
4. Methods annotated with a project role are project-scoped, so the request object needs to have the **project** field to specify the target project of the service method; methods annotated with a tenant role are tenant-scoped and need the **tenant** field instead

Further explanations are explained in the following.

### Auth

These options specify the RBAC of the api-endpoint.

| Option          | Description                                      | Values      | Explanation                                            |
| --------------- | ------------------------------------------------ | ----------- | ------------------------------------------------------ |
| `tenant_roles`  | Specifies the required tenant role (repeatable)  | UNSPECIFIED |                                                        |
|                 |                                                  | OWNER       | tenant owner                                           |
|                 |                                                  | EDITOR      | tenant editor                                          |
|                 |                                                  | VIEWER      | tenant viewer                                          |
|                 |                                                  | GUEST       | tenant guest                                           |
| `project_roles` | Specifies the required project role (repeatable) | UNSPECIFIED |                                                        |
|                 |                                                  | OWNER       | project owner                                          |
|                 |                                                  | EDITOR      | project editor                                         |
|                 |                                                  | VIEWER      | project viewer                                         |
| `admin_roles`   | Specifies the required admin role (repeatable)   | UNSPECIFIED |                                                        |
|                 |                                                  | EDITOR      | admin editor                                           |
|                 |                                                  | VIEWER      | admin viewer                                           |
| `visibility`    | Specifies the visibility of the api-endpoint     | UNSPECIFIED |                                                        |
|                 |                                                  | PUBLIC      | api-method is visible to public, a token is not needed |
|                 |                                                  | SELF        | api-method is scoped to owner resources                |

> [!IMPORTANT]
>
> Every operation needs at least an option, which references the scope of the request: **ROLE** or **VISIBILITY**

> [!CAUTION]
>
> If we use a Tenant or Project role, the request will be respectively scoped as Tenant or Project request.
> Tenant-Requests must have the field **tenant**, which is the tenant id and specifies the tenant on which the service-method is scoped.
> Project-Requests must have the field **project**, which is the project id and specifies the project on which the service-method is scoped.

### Auditing

For traceability we require to store audit-logs.

| Option     | Description                              | Values      | Explanation              |
| ---------- | ---------------------------------------- | ----------- | ------------------------ |
| AUDITING\_ | Specifies if the api-endpoint is audited | UNSPECIFIED | DEFAULT: included        |
|            |                                          | INCLUDED    | operation is audited     |
|            |                                          | EXCLUDED    | operation is not audited |
