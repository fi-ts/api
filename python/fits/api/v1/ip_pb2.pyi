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

class IPType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    IP_TYPE_UNSPECIFIED: _ClassVar[IPType]
    IP_TYPE_EPHEMERAL: _ClassVar[IPType]
    IP_TYPE_STATIC: _ClassVar[IPType]

class IPAddressFamily(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    IP_ADDRESS_FAMILY_UNSPECIFIED: _ClassVar[IPAddressFamily]
    IP_ADDRESS_FAMILY_V4: _ClassVar[IPAddressFamily]
    IP_ADDRESS_FAMILY_V6: _ClassVar[IPAddressFamily]
IP_TYPE_UNSPECIFIED: IPType
IP_TYPE_EPHEMERAL: IPType
IP_TYPE_STATIC: IPType
IP_ADDRESS_FAMILY_UNSPECIFIED: IPAddressFamily
IP_ADDRESS_FAMILY_V4: IPAddressFamily
IP_ADDRESS_FAMILY_V6: IPAddressFamily

class IP(_message.Message):
    __slots__ = ("uuid", "meta", "ip", "name", "description", "network", "project", "type", "namespace")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    IP_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    NETWORK_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    ip: str
    name: str
    description: str
    network: str
    project: str
    type: IPType
    namespace: str
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., ip: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., network: _Optional[str] = ..., project: _Optional[str] = ..., type: _Optional[_Union[IPType, str]] = ..., namespace: _Optional[str] = ...) -> None: ...

class IPServiceGetRequest(_message.Message):
    __slots__ = ("ip", "project", "namespace")
    IP_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    ip: str
    project: str
    namespace: str
    def __init__(self, ip: _Optional[str] = ..., project: _Optional[str] = ..., namespace: _Optional[str] = ...) -> None: ...

class IPServiceCreateRequest(_message.Message):
    __slots__ = ("network", "project", "name", "description", "ip", "machine", "labels", "type", "address_family")
    NETWORK_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IP_FIELD_NUMBER: _ClassVar[int]
    MACHINE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FAMILY_FIELD_NUMBER: _ClassVar[int]
    network: str
    project: str
    name: str
    description: str
    ip: str
    machine: str
    labels: _common_pb2.Labels
    type: IPType
    address_family: IPAddressFamily
    def __init__(self, network: _Optional[str] = ..., project: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., ip: _Optional[str] = ..., machine: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., type: _Optional[_Union[IPType, str]] = ..., address_family: _Optional[_Union[IPAddressFamily, str]] = ...) -> None: ...

class IPServiceUpdateRequest(_message.Message):
    __slots__ = ("ip", "update_meta", "project", "name", "description", "type", "labels")
    IP_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    ip: str
    update_meta: _common_pb2.UpdateMeta
    project: str
    name: str
    description: str
    type: IPType
    labels: _common_pb2.UpdateLabels
    def __init__(self, ip: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., project: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., type: _Optional[_Union[IPType, str]] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ...) -> None: ...

class IPServiceListRequest(_message.Message):
    __slots__ = ("project", "query")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    project: str
    query: IPQuery
    def __init__(self, project: _Optional[str] = ..., query: _Optional[_Union[IPQuery, _Mapping]] = ...) -> None: ...

class IPQuery(_message.Message):
    __slots__ = ("ip", "network", "project", "name", "uuid", "machine", "parent_prefix_cidr", "labels", "type", "address_family")
    IP_FIELD_NUMBER: _ClassVar[int]
    NETWORK_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    MACHINE_FIELD_NUMBER: _ClassVar[int]
    PARENT_PREFIX_CIDR_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FAMILY_FIELD_NUMBER: _ClassVar[int]
    ip: str
    network: str
    project: str
    name: str
    uuid: str
    machine: str
    parent_prefix_cidr: str
    labels: _common_pb2.Labels
    type: IPType
    address_family: IPAddressFamily
    def __init__(self, ip: _Optional[str] = ..., network: _Optional[str] = ..., project: _Optional[str] = ..., name: _Optional[str] = ..., uuid: _Optional[str] = ..., machine: _Optional[str] = ..., parent_prefix_cidr: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., type: _Optional[_Union[IPType, str]] = ..., address_family: _Optional[_Union[IPAddressFamily, str]] = ...) -> None: ...

class IPServiceDeleteRequest(_message.Message):
    __slots__ = ("ip", "project")
    IP_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ip: str
    project: str
    def __init__(self, ip: _Optional[str] = ..., project: _Optional[str] = ...) -> None: ...

class IPServiceGetResponse(_message.Message):
    __slots__ = ("ip",)
    IP_FIELD_NUMBER: _ClassVar[int]
    ip: IP
    def __init__(self, ip: _Optional[_Union[IP, _Mapping]] = ...) -> None: ...

class IPServiceUpdateResponse(_message.Message):
    __slots__ = ("ip",)
    IP_FIELD_NUMBER: _ClassVar[int]
    ip: IP
    def __init__(self, ip: _Optional[_Union[IP, _Mapping]] = ...) -> None: ...

class IPServiceCreateResponse(_message.Message):
    __slots__ = ("ip",)
    IP_FIELD_NUMBER: _ClassVar[int]
    ip: IP
    def __init__(self, ip: _Optional[_Union[IP, _Mapping]] = ...) -> None: ...

class IPServiceListResponse(_message.Message):
    __slots__ = ("ips",)
    IPS_FIELD_NUMBER: _ClassVar[int]
    ips: _containers.RepeatedCompositeFieldContainer[IP]
    def __init__(self, ips: _Optional[_Iterable[_Union[IP, _Mapping]]] = ...) -> None: ...

class IPServiceDeleteResponse(_message.Message):
    __slots__ = ("ip",)
    IP_FIELD_NUMBER: _ClassVar[int]
    ip: IP
    def __init__(self, ip: _Optional[_Union[IP, _Mapping]] = ...) -> None: ...
