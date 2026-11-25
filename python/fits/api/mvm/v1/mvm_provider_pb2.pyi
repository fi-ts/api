from fits.api.mvm.v1 import mvm_pb2 as _mvm_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProviderMVMServiceCreateRequest(_message.Message):
    __slots__ = ("request", "billable")
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    BILLABLE_FIELD_NUMBER: _ClassVar[int]
    request: _mvm_pb2.MVMServiceCreateRequest
    billable: bool
    def __init__(self, request: _Optional[_Union[_mvm_pb2.MVMServiceCreateRequest, _Mapping]] = ..., billable: _Optional[bool] = ...) -> None: ...

class ProviderMVMServiceCreateResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _mvm_pb2.MVMServiceCreateResponse
    def __init__(self, response: _Optional[_Union[_mvm_pb2.MVMServiceCreateResponse, _Mapping]] = ...) -> None: ...
