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
    __slots__ = ("uuid", "id", "subnet_title", "subnet", "subnetmask", "gateway", "ip_range_start", "ip_range_end", "use_dhcp", "location_uuid", "stage_type_uuid", "tenant")
    UUID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    SUBNET_TITLE_FIELD_NUMBER: _ClassVar[int]
    SUBNET_FIELD_NUMBER: _ClassVar[int]
    SUBNETMASK_FIELD_NUMBER: _ClassVar[int]
    GATEWAY_FIELD_NUMBER: _ClassVar[int]
    IP_RANGE_START_FIELD_NUMBER: _ClassVar[int]
    IP_RANGE_END_FIELD_NUMBER: _ClassVar[int]
    USE_DHCP_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    STAGE_TYPE_UUID_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    id: int
    subnet_title: str
    subnet: str
    subnetmask: str
    gateway: str
    ip_range_start: str
    ip_range_end: str
    use_dhcp: bool
    location_uuid: str
    stage_type_uuid: str
    tenant: str
    def __init__(self, uuid: _Optional[str] = ..., id: _Optional[int] = ..., subnet_title: _Optional[str] = ..., subnet: _Optional[str] = ..., subnetmask: _Optional[str] = ..., gateway: _Optional[str] = ..., ip_range_start: _Optional[str] = ..., ip_range_end: _Optional[str] = ..., use_dhcp: _Optional[bool] = ..., location_uuid: _Optional[str] = ..., stage_type_uuid: _Optional[str] = ..., tenant: _Optional[str] = ...) -> None: ...

class VLANServiceListRequest(_message.Message):
    __slots__ = ("tenant",)
    TENANT_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    def __init__(self, tenant: _Optional[str] = ...) -> None: ...

class VLANServiceListResponse(_message.Message):
    __slots__ = ("vlans",)
    VLANS_FIELD_NUMBER: _ClassVar[int]
    vlans: _containers.RepeatedCompositeFieldContainer[VLAN]
    def __init__(self, vlans: _Optional[_Iterable[_Union[VLAN, _Mapping]]] = ...) -> None: ...
