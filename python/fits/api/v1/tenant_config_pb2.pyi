from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TenantConfig(_message.Message):
    __slots__ = ("uuid", "tenant", "namespace", "key", "name", "meta", "data")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    tenant: str
    namespace: str
    key: str
    name: str
    meta: _common_pb2.Meta
    data: _struct_pb2.Struct
    def __init__(self, uuid: _Optional[str] = ..., tenant: _Optional[str] = ..., namespace: _Optional[str] = ..., key: _Optional[str] = ..., name: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class TenantConfigServiceListRequest(_message.Message):
    __slots__ = ("tenant", "namespace", "key", "labels")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    namespace: str
    key: str
    labels: _common_pb2.Labels
    def __init__(self, tenant: _Optional[str] = ..., namespace: _Optional[str] = ..., key: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class TenantConfigServiceListResponse(_message.Message):
    __slots__ = ("configs",)
    CONFIGS_FIELD_NUMBER: _ClassVar[int]
    configs: _containers.RepeatedCompositeFieldContainer[TenantConfig]
    def __init__(self, configs: _Optional[_Iterable[_Union[TenantConfig, _Mapping]]] = ...) -> None: ...

class TenantConfigServiceGetRequest(_message.Message):
    __slots__ = ("tenant", "uuid")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    uuid: str
    def __init__(self, tenant: _Optional[str] = ..., uuid: _Optional[str] = ...) -> None: ...

class TenantConfigServiceGetResponse(_message.Message):
    __slots__ = ("config",)
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    config: TenantConfig
    def __init__(self, config: _Optional[_Union[TenantConfig, _Mapping]] = ...) -> None: ...

class TenantConfigServiceCreateRequest(_message.Message):
    __slots__ = ("tenant", "namespace", "key", "data")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    namespace: str
    key: str
    data: _struct_pb2.Struct
    def __init__(self, tenant: _Optional[str] = ..., namespace: _Optional[str] = ..., key: _Optional[str] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class TenantConfigServiceCreateResponse(_message.Message):
    __slots__ = ("config",)
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    config: TenantConfig
    def __init__(self, config: _Optional[_Union[TenantConfig, _Mapping]] = ...) -> None: ...

class TenantConfigServiceUpdateRequest(_message.Message):
    __slots__ = ("tenant", "uuid", "update_meta", "data")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    uuid: str
    update_meta: _common_pb2.UpdateMeta
    data: _struct_pb2.Struct
    def __init__(self, tenant: _Optional[str] = ..., uuid: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class TenantConfigServiceUpdateResponse(_message.Message):
    __slots__ = ("config",)
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    config: TenantConfig
    def __init__(self, config: _Optional[_Union[TenantConfig, _Mapping]] = ...) -> None: ...

class TenantConfigServiceDeleteRequest(_message.Message):
    __slots__ = ("tenant", "uuid")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    uuid: str
    def __init__(self, tenant: _Optional[str] = ..., uuid: _Optional[str] = ...) -> None: ...

class TenantConfigServiceDeleteResponse(_message.Message):
    __slots__ = ("config",)
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    config: TenantConfig
    def __init__(self, config: _Optional[_Union[TenantConfig, _Mapping]] = ...) -> None: ...
