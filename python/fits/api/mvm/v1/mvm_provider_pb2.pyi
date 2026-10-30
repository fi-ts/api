import datetime

from fits.api.mvm.v1 import mvm_pb2 as _mvm_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProviderMVMServiceCreateRequest(_message.Message):
    __slots__ = ("request", "billable", "billable_from")
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FROM_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceCreateRequest
    billable: bool
    billable_from: _timestamp_pb2.Timestamp
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceCreateRequest, _Mapping]] = ..., billable: _Optional[bool] = ..., billable_from: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ProviderMVMServiceCreateResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _mvm_pb2.MVMServiceCreateResponse
    def __init__(self, response: _Optional[_Union[_mvm_pb2.MVMServiceCreateResponse, _Mapping]] = ...) -> None: ...
