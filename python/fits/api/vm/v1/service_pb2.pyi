from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Service(_message.Message):
    __slots__ = ("uuid", "title")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    title: str
    def __init__(self, uuid: _Optional[str] = ..., title: _Optional[str] = ...) -> None: ...

class ServiceServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ServiceServiceListResponse(_message.Message):
    __slots__ = ("services",)
    SERVICES_FIELD_NUMBER: _ClassVar[int]
    services: _containers.RepeatedCompositeFieldContainer[Service]
    def __init__(self, services: _Optional[_Iterable[_Union[Service, _Mapping]]] = ...) -> None: ...
