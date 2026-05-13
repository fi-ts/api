import datetime

from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TenantRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TENANT_ROLE_UNSPECIFIED: _ClassVar[TenantRole]
    TENANT_ROLE_OWNER: _ClassVar[TenantRole]
    TENANT_ROLE_EDITOR: _ClassVar[TenantRole]
    TENANT_ROLE_VIEWER: _ClassVar[TenantRole]
    TENANT_ROLE_GUEST: _ClassVar[TenantRole]

class ProjectRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PROJECT_ROLE_UNSPECIFIED: _ClassVar[ProjectRole]
    PROJECT_ROLE_OWNER: _ClassVar[ProjectRole]
    PROJECT_ROLE_EDITOR: _ClassVar[ProjectRole]
    PROJECT_ROLE_VIEWER: _ClassVar[ProjectRole]

class AdminRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ADMIN_ROLE_UNSPECIFIED: _ClassVar[AdminRole]
    ADMIN_ROLE_EDITOR: _ClassVar[AdminRole]
    ADMIN_ROLE_VIEWER: _ClassVar[AdminRole]

class InfraRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INFRA_ROLE_UNSPECIFIED: _ClassVar[InfraRole]
    INFRA_ROLE_EDITOR: _ClassVar[InfraRole]
    INFRA_ROLE_VIEWER: _ClassVar[InfraRole]

class MachineRole(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    MACHINE_ROLE_UNSPECIFIED: _ClassVar[MachineRole]
    MACHINE_ROLE_EDITOR: _ClassVar[MachineRole]
    MACHINE_ROLE_VIEWER: _ClassVar[MachineRole]

class Visibility(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    VISIBILITY_UNSPECIFIED: _ClassVar[Visibility]
    VISIBILITY_PUBLIC: _ClassVar[Visibility]
    VISIBILITY_SELF: _ClassVar[Visibility]

class Auditing(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    AUDITING_UNSPECIFIED: _ClassVar[Auditing]
    AUDITING_INCLUDED: _ClassVar[Auditing]
    AUDITING_EXCLUDED: _ClassVar[Auditing]

class OptimisticLockingStrategy(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OPTIMISTIC_LOCKING_STRATEGY_UNSPECIFIED: _ClassVar[OptimisticLockingStrategy]
    OPTIMISTIC_LOCKING_STRATEGY_CLIENT: _ClassVar[OptimisticLockingStrategy]
    OPTIMISTIC_LOCKING_STRATEGY_SERVER: _ClassVar[OptimisticLockingStrategy]
TENANT_ROLE_UNSPECIFIED: TenantRole
TENANT_ROLE_OWNER: TenantRole
TENANT_ROLE_EDITOR: TenantRole
TENANT_ROLE_VIEWER: TenantRole
TENANT_ROLE_GUEST: TenantRole
PROJECT_ROLE_UNSPECIFIED: ProjectRole
PROJECT_ROLE_OWNER: ProjectRole
PROJECT_ROLE_EDITOR: ProjectRole
PROJECT_ROLE_VIEWER: ProjectRole
ADMIN_ROLE_UNSPECIFIED: AdminRole
ADMIN_ROLE_EDITOR: AdminRole
ADMIN_ROLE_VIEWER: AdminRole
INFRA_ROLE_UNSPECIFIED: InfraRole
INFRA_ROLE_EDITOR: InfraRole
INFRA_ROLE_VIEWER: InfraRole
MACHINE_ROLE_UNSPECIFIED: MachineRole
MACHINE_ROLE_EDITOR: MachineRole
MACHINE_ROLE_VIEWER: MachineRole
VISIBILITY_UNSPECIFIED: Visibility
VISIBILITY_PUBLIC: Visibility
VISIBILITY_SELF: Visibility
AUDITING_UNSPECIFIED: Auditing
AUDITING_INCLUDED: Auditing
AUDITING_EXCLUDED: Auditing
OPTIMISTIC_LOCKING_STRATEGY_UNSPECIFIED: OptimisticLockingStrategy
OPTIMISTIC_LOCKING_STRATEGY_CLIENT: OptimisticLockingStrategy
OPTIMISTIC_LOCKING_STRATEGY_SERVER: OptimisticLockingStrategy
TENANT_ROLES_FIELD_NUMBER: _ClassVar[int]
tenant_roles: _descriptor.FieldDescriptor
PROJECT_ROLES_FIELD_NUMBER: _ClassVar[int]
project_roles: _descriptor.FieldDescriptor
ADMIN_ROLES_FIELD_NUMBER: _ClassVar[int]
admin_roles: _descriptor.FieldDescriptor
VISIBILITY_FIELD_NUMBER: _ClassVar[int]
visibility: _descriptor.FieldDescriptor
AUDITING_FIELD_NUMBER: _ClassVar[int]
auditing: _descriptor.FieldDescriptor
ENUM_STRING_VALUE_FIELD_NUMBER: _ClassVar[int]
enum_string_value: _descriptor.FieldDescriptor

class Paging(_message.Message):
    __slots__ = ("page", "count")
    PAGE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    page: int
    count: int
    def __init__(self, page: _Optional[int] = ..., count: _Optional[int] = ...) -> None: ...

class Labels(_message.Message):
    __slots__ = ("labels",)
    class LabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    LABELS_FIELD_NUMBER: _ClassVar[int]
    labels: _containers.ScalarMap[str, str]
    def __init__(self, labels: _Optional[_Mapping[str, str]] = ...) -> None: ...

class Meta(_message.Message):
    __slots__ = ("labels", "created_at", "updated_at", "generation")
    LABELS_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    GENERATION_FIELD_NUMBER: _ClassVar[int]
    labels: Labels
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    generation: int
    def __init__(self, labels: _Optional[_Union[Labels, _Mapping]] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., generation: _Optional[int] = ...) -> None: ...

class UpdateLabels(_message.Message):
    __slots__ = ("update", "remove")
    UPDATE_FIELD_NUMBER: _ClassVar[int]
    REMOVE_FIELD_NUMBER: _ClassVar[int]
    update: Labels
    remove: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, update: _Optional[_Union[Labels, _Mapping]] = ..., remove: _Optional[_Iterable[str]] = ...) -> None: ...

class UpdateMeta(_message.Message):
    __slots__ = ("updated_at", "locking_strategy")
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    LOCKING_STRATEGY_FIELD_NUMBER: _ClassVar[int]
    updated_at: _timestamp_pb2.Timestamp
    locking_strategy: OptimisticLockingStrategy
    def __init__(self, updated_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., locking_strategy: _Optional[_Union[OptimisticLockingStrategy, str]] = ...) -> None: ...
