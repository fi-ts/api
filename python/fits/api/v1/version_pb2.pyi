from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Version(_message.Message):
    __slots__ = ("version", "revision", "git_sha1", "build_date")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    REVISION_FIELD_NUMBER: _ClassVar[int]
    GIT_SHA1_FIELD_NUMBER: _ClassVar[int]
    BUILD_DATE_FIELD_NUMBER: _ClassVar[int]
    version: str
    revision: str
    git_sha1: str
    build_date: str
    def __init__(self, version: _Optional[str] = ..., revision: _Optional[str] = ..., git_sha1: _Optional[str] = ..., build_date: _Optional[str] = ...) -> None: ...

class VersionServiceGetRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class VersionServiceGetResponse(_message.Message):
    __slots__ = ("version",)
    VERSION_FIELD_NUMBER: _ClassVar[int]
    version: Version
    def __init__(self, version: _Optional[_Union[Version, _Mapping]] = ...) -> None: ...
