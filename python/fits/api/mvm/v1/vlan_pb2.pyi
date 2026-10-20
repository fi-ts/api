from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VLAN(_message.Message):
    __slots__ = ("uuid", "vlan_id", "subnet_title", "subnet_cidr", "location_uuid", "tenant", "domain_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    VLAN_ID_FIELD_NUMBER: _ClassVar[int]
    SUBNET_TITLE_FIELD_NUMBER: _ClassVar[int]
    SUBNET_CIDR_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    DOMAIN_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    vlan_id: int
    subnet_title: str
    subnet_cidr: str
    location_uuid: str
    tenant: str
    domain_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., vlan_id: _Optional[int] = ..., subnet_title: _Optional[str] = ..., subnet_cidr: _Optional[str] = ..., location_uuid: _Optional[str] = ..., tenant: _Optional[str] = ..., domain_uuid: _Optional[str] = ...) -> None: ...

class VLANQuery(_message.Message):
    __slots__ = ("domain_uuid",)
    DOMAIN_UUID_FIELD_NUMBER: _ClassVar[int]
    domain_uuid: str
    def __init__(self, domain_uuid: _Optional[str] = ...) -> None: ...

class VLANServiceListRequest(_message.Message):
    __slots__ = ("tenant", "location_uuid", "query")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    location_uuid: str
    query: VLANQuery
    def __init__(self, tenant: _Optional[str] = ..., location_uuid: _Optional[str] = ..., query: _Optional[_Union[VLANQuery, _Mapping]] = ...) -> None: ...

class VLANServiceListResponse(_message.Message):
    __slots__ = ("vlans",)
    VLANS_FIELD_NUMBER: _ClassVar[int]
    vlans: _containers.RepeatedCompositeFieldContainer[VLAN]
    def __init__(self, vlans: _Optional[_Iterable[_Union[VLAN, _Mapping]]] = ...) -> None: ...
