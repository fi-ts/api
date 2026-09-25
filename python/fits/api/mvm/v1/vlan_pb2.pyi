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

class StageType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    STAGE_TYPE_UNSPECIFIED: _ClassVar[StageType]
    STAGE_TYPE_DEVELOPMENT: _ClassVar[StageType]
    STAGE_TYPE_TEST: _ClassVar[StageType]
    STAGE_TYPE_INTEGRATION: _ClassVar[StageType]
    STAGE_TYPE_PRODUCTION: _ClassVar[StageType]
STAGE_TYPE_UNSPECIFIED: StageType
STAGE_TYPE_DEVELOPMENT: StageType
STAGE_TYPE_TEST: StageType
STAGE_TYPE_INTEGRATION: StageType
STAGE_TYPE_PRODUCTION: StageType

class VLAN(_message.Message):
    __slots__ = ("uuid", "vlan_id", "subnet_title", "subnet_cidr", "location_uuid", "stage_type", "tenant")
    UUID_FIELD_NUMBER: _ClassVar[int]
    VLAN_ID_FIELD_NUMBER: _ClassVar[int]
    SUBNET_TITLE_FIELD_NUMBER: _ClassVar[int]
    SUBNET_CIDR_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    STAGE_TYPE_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    vlan_id: int
    subnet_title: str
    subnet_cidr: str
    location_uuid: str
    stage_type: StageType
    tenant: str
    def __init__(self, uuid: _Optional[str] = ..., vlan_id: _Optional[int] = ..., subnet_title: _Optional[str] = ..., subnet_cidr: _Optional[str] = ..., location_uuid: _Optional[str] = ..., stage_type: _Optional[_Union[StageType, str]] = ..., tenant: _Optional[str] = ...) -> None: ...

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
