from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Contact(_message.Message):
    __slots__ = ("uuid", "first_name", "last_name", "department", "phone", "email", "tenant")
    UUID_FIELD_NUMBER: _ClassVar[int]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    DEPARTMENT_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    first_name: str
    last_name: str
    department: str
    phone: str
    email: str
    tenant: str
    def __init__(self, uuid: _Optional[str] = ..., first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., department: _Optional[str] = ..., phone: _Optional[str] = ..., email: _Optional[str] = ..., tenant: _Optional[str] = ...) -> None: ...

class ContactServiceListRequest(_message.Message):
    __slots__ = ("tenant",)
    TENANT_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    def __init__(self, tenant: _Optional[str] = ...) -> None: ...

class ContactServiceListResponse(_message.Message):
    __slots__ = ("contacts",)
    CONTACTS_FIELD_NUMBER: _ClassVar[int]
    contacts: _containers.RepeatedCompositeFieldContainer[Contact]
    def __init__(self, contacts: _Optional[_Iterable[_Union[Contact, _Mapping]]] = ...) -> None: ...

class ContactServiceCreateRequest(_message.Message):
    __slots__ = ("tenant", "first_name", "last_name", "department", "phone", "email")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    DEPARTMENT_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    first_name: str
    last_name: str
    department: str
    phone: str
    email: str
    def __init__(self, tenant: _Optional[str] = ..., first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., department: _Optional[str] = ..., phone: _Optional[str] = ..., email: _Optional[str] = ...) -> None: ...

class ContactServiceCreateResponse(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class ContactServiceUpdateRequest(_message.Message):
    __slots__ = ("tenant", "uuid", "update_meta", "first_name", "last_name", "department", "phone", "email")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    DEPARTMENT_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    uuid: str
    update_meta: _common_pb2.UpdateMeta
    first_name: str
    last_name: str
    department: str
    phone: str
    email: str
    def __init__(self, tenant: _Optional[str] = ..., uuid: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., department: _Optional[str] = ..., phone: _Optional[str] = ..., email: _Optional[str] = ...) -> None: ...

class ContactServiceUpdateResponse(_message.Message):
    __slots__ = ("contact",)
    CONTACT_FIELD_NUMBER: _ClassVar[int]
    contact: Contact
    def __init__(self, contact: _Optional[_Union[Contact, _Mapping]] = ...) -> None: ...

class ContactServiceDeleteRequest(_message.Message):
    __slots__ = ("tenant", "uuid")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    uuid: str
    def __init__(self, tenant: _Optional[str] = ..., uuid: _Optional[str] = ...) -> None: ...

class ContactServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
