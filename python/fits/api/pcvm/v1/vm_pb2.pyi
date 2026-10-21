from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    STATE_UNSPECIFIED: _ClassVar[State]
    STATE_HEALTHY: _ClassVar[State]
    STATE_WARNING: _ClassVar[State]
    STATE_ERROR: _ClassVar[State]

class Operation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OPERATION_UNSPECIFIED: _ClassVar[Operation]
    OPERATION_CREATING: _ClassVar[Operation]
    OPERATION_ACTIVE: _ClassVar[Operation]
    OPERATION_UPDATING: _ClassVar[Operation]
    OPERATION_DELETED: _ClassVar[Operation]
    OPERATION_ERROR: _ClassVar[Operation]

class Size(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SIZE_UNSPECIFIED: _ClassVar[Size]
    SIZE_SMALL: _ClassVar[Size]
    SIZE_MEDIUM: _ClassVar[Size]
    SIZE_LARGE: _ClassVar[Size]
STATE_UNSPECIFIED: State
STATE_HEALTHY: State
STATE_WARNING: State
STATE_ERROR: State
OPERATION_UNSPECIFIED: Operation
OPERATION_CREATING: Operation
OPERATION_ACTIVE: Operation
OPERATION_UPDATING: Operation
OPERATION_DELETED: Operation
OPERATION_ERROR: Operation
SIZE_UNSPECIFIED: Size
SIZE_SMALL: Size
SIZE_MEDIUM: Size
SIZE_LARGE: Size

class VM(_message.Message):
    __slots__ = ("uuid", "meta", "name", "tenant", "project_uuid", "size", "status")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    name: str
    tenant: str
    project_uuid: str
    size: Size
    status: Status
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., tenant: _Optional[str] = ..., project_uuid: _Optional[str] = ..., size: _Optional[_Union[Size, str]] = ..., status: _Optional[_Union[Status, _Mapping]] = ...) -> None: ...

class Status(_message.Message):
    __slots__ = ("state", "info", "operation")
    STATE_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    OPERATION_FIELD_NUMBER: _ClassVar[int]
    state: State
    info: str
    operation: Operation
    def __init__(self, state: _Optional[_Union[State, str]] = ..., info: _Optional[str] = ..., operation: _Optional[_Union[Operation, str]] = ...) -> None: ...

class VMQuery(_message.Message):
    __slots__ = ("name", "labels")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    name: str
    labels: _common_pb2.Labels
    def __init__(self, name: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class VMServiceGetRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ...) -> None: ...

class VMServiceGetResponse(_message.Message):
    __slots__ = ("vm",)
    VM_FIELD_NUMBER: _ClassVar[int]
    vm: VM
    def __init__(self, vm: _Optional[_Union[VM, _Mapping]] = ...) -> None: ...

class VMServiceCreateRequest(_message.Message):
    __slots__ = ("project_uuid", "name", "labels", "size")
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    name: str
    labels: _common_pb2.Labels
    size: Size
    def __init__(self, project_uuid: _Optional[str] = ..., name: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., size: _Optional[_Union[Size, str]] = ...) -> None: ...

class VMServiceCreateResponse(_message.Message):
    __slots__ = ("vm",)
    VM_FIELD_NUMBER: _ClassVar[int]
    vm: VM
    def __init__(self, vm: _Optional[_Union[VM, _Mapping]] = ...) -> None: ...

class VMServiceListRequest(_message.Message):
    __slots__ = ("project_uuid", "paging", "query")
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    PAGING_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    paging: _common_pb2.Paging
    query: VMQuery
    def __init__(self, project_uuid: _Optional[str] = ..., paging: _Optional[_Union[_common_pb2.Paging, _Mapping]] = ..., query: _Optional[_Union[VMQuery, _Mapping]] = ...) -> None: ...

class VMServiceListResponse(_message.Message):
    __slots__ = ("vms",)
    VMS_FIELD_NUMBER: _ClassVar[int]
    vms: _containers.RepeatedCompositeFieldContainer[VM]
    def __init__(self, vms: _Optional[_Iterable[_Union[VM, _Mapping]]] = ...) -> None: ...

class VMServiceUpdateRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "update_meta", "labels", "name", "size")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    update_meta: _common_pb2.UpdateMeta
    labels: _common_pb2.UpdateLabels
    name: str
    size: Size
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ..., name: _Optional[str] = ..., size: _Optional[_Union[Size, str]] = ...) -> None: ...

class VMServiceUpdateResponse(_message.Message):
    __slots__ = ("vm",)
    VM_FIELD_NUMBER: _ClassVar[int]
    vm: VM
    def __init__(self, vm: _Optional[_Union[VM, _Mapping]] = ...) -> None: ...

class VMServiceDeleteRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ...) -> None: ...

class VMServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
