from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class OS(_message.Message):
    __slots__ = ("uuid", "title", "type", "version")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    title: str
    type: str
    version: str
    def __init__(self, uuid: _Optional[str] = ..., title: _Optional[str] = ..., type: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...

class OSServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class OSServiceListResponse(_message.Message):
    __slots__ = ("oss",)
    OSS_FIELD_NUMBER: _ClassVar[int]
    oss: _containers.RepeatedCompositeFieldContainer[OS]
    def __init__(self, oss: _Optional[_Iterable[_Union[OS, _Mapping]]] = ...) -> None: ...
