from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ServiceStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SERVICE_STATUS_UNSPECIFIED: _ClassVar[ServiceStatus]
    SERVICE_STATUS_DEGRADED: _ClassVar[ServiceStatus]
    SERVICE_STATUS_UNHEALTHY: _ClassVar[ServiceStatus]
    SERVICE_STATUS_HEALTHY: _ClassVar[ServiceStatus]

class Service(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SERVICE_UNSPECIFIED: _ClassVar[Service]
    SERVICE_K8S: _ClassVar[Service]
SERVICE_STATUS_UNSPECIFIED: ServiceStatus
SERVICE_STATUS_DEGRADED: ServiceStatus
SERVICE_STATUS_UNHEALTHY: ServiceStatus
SERVICE_STATUS_HEALTHY: ServiceStatus
SERVICE_UNSPECIFIED: Service
SERVICE_K8S: Service

class Health(_message.Message):
    __slots__ = ("services",)
    SERVICES_FIELD_NUMBER: _ClassVar[int]
    services: _containers.RepeatedCompositeFieldContainer[HealthStatus]
    def __init__(self, services: _Optional[_Iterable[_Union[HealthStatus, _Mapping]]] = ...) -> None: ...

class HealthStatus(_message.Message):
    __slots__ = ("name", "status", "message", "partitions")
    class PartitionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: PartitionHealth
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[PartitionHealth, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    PARTITIONS_FIELD_NUMBER: _ClassVar[int]
    name: Service
    status: ServiceStatus
    message: str
    partitions: _containers.MessageMap[str, PartitionHealth]
    def __init__(self, name: _Optional[_Union[Service, str]] = ..., status: _Optional[_Union[ServiceStatus, str]] = ..., message: _Optional[str] = ..., partitions: _Optional[_Mapping[str, PartitionHealth]] = ...) -> None: ...

class PartitionHealth(_message.Message):
    __slots__ = ("status", "message")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    status: ServiceStatus
    message: str
    def __init__(self, status: _Optional[_Union[ServiceStatus, str]] = ..., message: _Optional[str] = ...) -> None: ...

class HealthServiceGetRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HealthServiceGetResponse(_message.Message):
    __slots__ = ("health",)
    HEALTH_FIELD_NUMBER: _ClassVar[int]
    health: Health
    def __init__(self, health: _Optional[_Union[Health, _Mapping]] = ...) -> None: ...
