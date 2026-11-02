import datetime

from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ApiKey(_message.Message):
    __slots__ = ("id", "user_id", "realm", "name", "hash", "prefix", "created_at", "expires_at", "revoked_at", "last_used_at")
    ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    REALM_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    PREFIX_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    REVOKED_AT_FIELD_NUMBER: _ClassVar[int]
    LAST_USED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    user_id: str
    realm: str
    name: str
    hash: str
    prefix: str
    created_at: _timestamp_pb2.Timestamp
    expires_at: _timestamp_pb2.Timestamp
    revoked_at: _timestamp_pb2.Timestamp
    last_used_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., user_id: _Optional[str] = ..., realm: _Optional[str] = ..., name: _Optional[str] = ..., hash: _Optional[str] = ..., prefix: _Optional[str] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., expires_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., revoked_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., last_used_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ApiKeyServiceCreateRequest(_message.Message):
    __slots__ = ("name", "expires_at")
    NAME_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    name: str
    expires_at: _timestamp_pb2.Timestamp
    def __init__(self, name: _Optional[str] = ..., expires_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ApiKeyServiceCreateResponse(_message.Message):
    __slots__ = ("secret", "api_key")
    SECRET_FIELD_NUMBER: _ClassVar[int]
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    secret: str
    api_key: ApiKey
    def __init__(self, secret: _Optional[str] = ..., api_key: _Optional[_Union[ApiKey, _Mapping]] = ...) -> None: ...

class ApiKeyServiceListRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ApiKeyServiceListResponse(_message.Message):
    __slots__ = ("api_keys",)
    API_KEYS_FIELD_NUMBER: _ClassVar[int]
    api_keys: _containers.RepeatedCompositeFieldContainer[ApiKey]
    def __init__(self, api_keys: _Optional[_Iterable[_Union[ApiKey, _Mapping]]] = ...) -> None: ...

class ApiKeyServiceGetRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ApiKeyServiceGetResponse(_message.Message):
    __slots__ = ("api_key",)
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    api_key: ApiKey
    def __init__(self, api_key: _Optional[_Union[ApiKey, _Mapping]] = ...) -> None: ...

class ApiKeyServiceRevokeRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ApiKeyServiceRevokeResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ApiKeyServiceDeleteRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ApiKeyServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
