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

class OSType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OS_TYPE_UNSPECIFIED: _ClassVar[OSType]
    OS_TYPE_LINUX: _ClassVar[OSType]
    OS_TYPE_WINDOWS: _ClassVar[OSType]
OS_TYPE_UNSPECIFIED: OSType
OS_TYPE_LINUX: OSType
OS_TYPE_WINDOWS: OSType

class OS(_message.Message):
    __slots__ = ("uuid", "title", "type")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    title: str
    type: OSType
    def __init__(self, uuid: _Optional[str] = ..., title: _Optional[str] = ..., type: _Optional[_Union[OSType, str]] = ...) -> None: ...

class OSServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class OSServiceListResponse(_message.Message):
    __slots__ = ("operating_systems",)
    OPERATING_SYSTEMS_FIELD_NUMBER: _ClassVar[int]
    operating_systems: _containers.RepeatedCompositeFieldContainer[OS]
    def __init__(self, operating_systems: _Optional[_Iterable[_Union[OS, _Mapping]]] = ...) -> None: ...
