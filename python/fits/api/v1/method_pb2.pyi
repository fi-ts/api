from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import token_pb2 as _token_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MethodServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MethodServiceListResponse(_message.Message):
    __slots__ = ("methods",)
    METHODS_FIELD_NUMBER: _ClassVar[int]
    methods: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, methods: _Optional[_Iterable[str]] = ...) -> None: ...

class MethodServiceTokenScopedListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MethodServiceTokenScopedListResponse(_message.Message):
    __slots__ = ("permissions", "project_roles", "tenant_roles", "admin_role")
    class ProjectRolesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _common_pb2.ProjectRole
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_common_pb2.ProjectRole, str]] = ...) -> None: ...
    class TenantRolesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _common_pb2.TenantRole
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_common_pb2.TenantRole, str]] = ...) -> None: ...
    PERMISSIONS_FIELD_NUMBER: _ClassVar[int]
    PROJECT_ROLES_FIELD_NUMBER: _ClassVar[int]
    TENANT_ROLES_FIELD_NUMBER: _ClassVar[int]
    ADMIN_ROLE_FIELD_NUMBER: _ClassVar[int]
    permissions: _containers.RepeatedCompositeFieldContainer[_token_pb2.MethodPermission]
    project_roles: _containers.ScalarMap[str, _common_pb2.ProjectRole]
    tenant_roles: _containers.ScalarMap[str, _common_pb2.TenantRole]
    admin_role: _common_pb2.AdminRole
    def __init__(self, permissions: _Optional[_Iterable[_Union[_token_pb2.MethodPermission, _Mapping]]] = ..., project_roles: _Optional[_Mapping[str, _common_pb2.ProjectRole]] = ..., tenant_roles: _Optional[_Mapping[str, _common_pb2.TenantRole]] = ..., admin_role: _Optional[_Union[_common_pb2.AdminRole, str]] = ...) -> None: ...
