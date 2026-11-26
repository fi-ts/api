import datetime

from fits.api.mvm.v1 import mvm_pb2 as _mvm_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class OnBehalfMVM(_message.Message):
    __slots__ = ("mvm", "billable", "billable_from")
    MVM_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FROM_FIELD_NUMBER: _ClassVar[int]
    mvm: _mvm_pb2.ManagedVM
    billable: bool
    billable_from: _timestamp_pb2.Timestamp
    def __init__(self, mvm: _Optional[_Union[_mvm_pb2.ManagedVM, _Mapping]] = ..., billable: _Optional[bool] = ..., billable_from: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceGetRequest(_message.Message):
    __slots__ = ("request",)
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceGetRequest
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceGetRequest, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceGetResponse(_message.Message):
    __slots__ = ("mvm",)
    MVM_FIELD_NUMBER: _ClassVar[int]
    mvm: OnBehalfMVM
    def __init__(self, mvm: _Optional[_Union[OnBehalfMVM, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceListRequest(_message.Message):
    __slots__ = ("request",)
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceListRequest
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceListRequest, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceListResponse(_message.Message):
    __slots__ = ("mvms",)
    MVMS_FIELD_NUMBER: _ClassVar[int]
    mvms: _containers.RepeatedCompositeFieldContainer[OnBehalfMVM]
    def __init__(self, mvms: _Optional[_Iterable[_Union[OnBehalfMVM, _Mapping]]] = ...) -> None: ...

class OnBehalfMVMServiceCreateRequest(_message.Message):
    __slots__ = ("request", "billable", "billable_from")
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FROM_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceCreateRequest
    billable: bool
    billable_from: _timestamp_pb2.Timestamp
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceCreateRequest, _Mapping]] = ..., billable: _Optional[bool] = ..., billable_from: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceCreateResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _mvm_pb2.MVMServiceCreateResponse
    def __init__(self, response: _Optional[_Union[_mvm_pb2.MVMServiceCreateResponse, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceDeleteRequest(_message.Message):
    __slots__ = ("request", "execution_date")
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_DATE_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceDeleteRequest
    execution_date: _timestamp_pb2.Timestamp
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceDeleteRequest, _Mapping]] = ..., execution_date: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class OnBehalfMVMServiceDeleteResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _mvm_pb2.MVMServiceDeleteResponse
    def __init__(self, response: _Optional[_Union[_mvm_pb2.MVMServiceDeleteResponse, _Mapping]] = ...) -> None: ...
