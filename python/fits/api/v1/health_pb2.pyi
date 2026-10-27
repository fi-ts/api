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
    SERVICE_MVM: _ClassVar[Service]
    SERVICE_PCVM: _ClassVar[Service]

class Dependency(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DEPENDENCY_UNSPECIFIED: _ClassVar[Dependency]
    DEPENDENCY_VALKEY: _ClassVar[Dependency]
    DEPENDENCY_TENANT_APISERVER: _ClassVar[Dependency]
SERVICE_STATUS_UNSPECIFIED: ServiceStatus
SERVICE_STATUS_DEGRADED: ServiceStatus
SERVICE_STATUS_UNHEALTHY: ServiceStatus
SERVICE_STATUS_HEALTHY: ServiceStatus
SERVICE_UNSPECIFIED: Service
SERVICE_MVM: Service
SERVICE_PCVM: Service
DEPENDENCY_UNSPECIFIED: Dependency
DEPENDENCY_VALKEY: Dependency
DEPENDENCY_TENANT_APISERVER: Dependency

class Health(_message.Message):
    __slots__ = ("services",)
    SERVICES_FIELD_NUMBER: _ClassVar[int]
    services: _containers.RepeatedCompositeFieldContainer[HealthStatus]
    def __init__(self, services: _Optional[_Iterable[_Union[HealthStatus, _Mapping]]] = ...) -> None: ...

class HealthStatus(_message.Message):
    __slots__ = ("service", "status", "message")
    SERVICE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    service: Service
    status: ServiceStatus
    message: str
    def __init__(self, service: _Optional[_Union[Service, str]] = ..., status: _Optional[_Union[ServiceStatus, str]] = ..., message: _Optional[str] = ...) -> None: ...

class Dependencies(_message.Message):
    __slots__ = ("dependencies",)
    DEPENDENCIES_FIELD_NUMBER: _ClassVar[int]
    dependencies: _containers.RepeatedCompositeFieldContainer[DependencyHealth]
    def __init__(self, dependencies: _Optional[_Iterable[_Union[DependencyHealth, _Mapping]]] = ...) -> None: ...

class DependencyHealth(_message.Message):
    __slots__ = ("dependency", "status", "message")
    DEPENDENCY_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    dependency: Dependency
    status: ServiceStatus
    message: str
    def __init__(self, dependency: _Optional[_Union[Dependency, str]] = ..., status: _Optional[_Union[ServiceStatus, str]] = ..., message: _Optional[str] = ...) -> None: ...

class HealthServiceGetRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HealthServiceGetResponse(_message.Message):
    __slots__ = ("health",)
    HEALTH_FIELD_NUMBER: _ClassVar[int]
    health: Health
    def __init__(self, health: _Optional[_Union[Health, _Mapping]] = ...) -> None: ...

class HealthServiceGetDependenciesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HealthServiceGetDependenciesResponse(_message.Message):
    __slots__ = ("dependencies",)
    DEPENDENCIES_FIELD_NUMBER: _ClassVar[int]
    dependencies: Dependencies
    def __init__(self, dependencies: _Optional[_Union[Dependencies, _Mapping]] = ...) -> None: ...
