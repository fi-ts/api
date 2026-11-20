from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Stage(_message.Message):
    __slots__ = ("uuid", "title")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    title: str
    def __init__(self, uuid: _Optional[str] = ..., title: _Optional[str] = ...) -> None: ...

class StageServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class StageServiceListResponse(_message.Message):
    __slots__ = ("stages",)
    STAGES_FIELD_NUMBER: _ClassVar[int]
    stages: _containers.RepeatedCompositeFieldContainer[Stage]
    def __init__(self, stages: _Optional[_Iterable[_Union[Stage, _Mapping]]] = ...) -> None: ...
