from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StageType(_message.Message):
    __slots__ = ("uuid", "title")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    title: str
    def __init__(self, uuid: _Optional[str] = ..., title: _Optional[str] = ...) -> None: ...

class StageTypeServiceListRequest(_message.Message):
    __slots__ = ("tenant_uuid",)
    TENANT_UUID_FIELD_NUMBER: _ClassVar[int]
    tenant_uuid: str
    def __init__(self, tenant_uuid: _Optional[str] = ...) -> None: ...

class StageTypeServiceListResponse(_message.Message):
    __slots__ = ("stagetypes", "validation_errors")
    STAGETYPES_FIELD_NUMBER: _ClassVar[int]
    VALIDATION_ERRORS_FIELD_NUMBER: _ClassVar[int]
    stagetypes: _containers.RepeatedCompositeFieldContainer[StageType]
    validation_errors: _containers.RepeatedCompositeFieldContainer[_common_pb2.ValidationError]
    def __init__(self, stagetypes: _Optional[_Iterable[_Union[StageType, _Mapping]]] = ..., validation_errors: _Optional[_Iterable[_Union[_common_pb2.ValidationError, _Mapping]]] = ...) -> None: ...
