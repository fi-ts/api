from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VMInstance(_message.Message):
    __slots__ = ("uuid", "meta", "name", "project")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    name: str
    project: str
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., project: _Optional[str] = ...) -> None: ...

class VMServiceGetRequest(_message.Message):
    __slots__ = ("uuid", "project")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project: str
    def __init__(self, uuid: _Optional[str] = ..., project: _Optional[str] = ...) -> None: ...

class VMServiceGetResponse(_message.Message):
    __slots__ = ("vm",)
    VM_FIELD_NUMBER: _ClassVar[int]
    vm: VMInstance
    def __init__(self, vm: _Optional[_Union[VMInstance, _Mapping]] = ...) -> None: ...

class VMServiceCreateWindowsRequest(_message.Message):
    __slots__ = ("domain_uuid", "disk")
    DOMAIN_UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_FIELD_NUMBER: _ClassVar[int]
    domain_uuid: str
    disk: WindowsDisk
    def __init__(self, domain_uuid: _Optional[str] = ..., disk: _Optional[_Union[WindowsDisk, _Mapping]] = ...) -> None: ...

class VMServiceCreateLinuxRequest(_message.Message):
    __slots__ = ("ldap_uuid", "disk")
    LDAP_UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_FIELD_NUMBER: _ClassVar[int]
    ldap_uuid: str
    disk: LinuxDisk
    def __init__(self, ldap_uuid: _Optional[str] = ..., disk: _Optional[_Union[LinuxDisk, _Mapping]] = ...) -> None: ...

class VMServiceCreateRequest(_message.Message):
    __slots__ = ("project_uuid", "name", "cpu", "ram", "os_uuid", "vlan_uuid", "location_uuid", "contact_uuid", "labels", "order_number", "backup", "availability", "serviceclass", "windows", "linux")
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    VLAN_UUID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    SERVICECLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_FIELD_NUMBER: _ClassVar[int]
    LINUX_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    name: str
    cpu: int
    ram: int
    os_uuid: str
    vlan_uuid: str
    location_uuid: str
    contact_uuid: str
    labels: _common_pb2.Labels
    order_number: str
    backup: bool
    availability: str
    serviceclass: str
    windows: VMServiceCreateWindowsRequest
    linux: VMServiceCreateLinuxRequest
    def __init__(self, project_uuid: _Optional[str] = ..., name: _Optional[str] = ..., cpu: _Optional[int] = ..., ram: _Optional[int] = ..., os_uuid: _Optional[str] = ..., vlan_uuid: _Optional[str] = ..., location_uuid: _Optional[str] = ..., contact_uuid: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., order_number: _Optional[str] = ..., backup: _Optional[bool] = ..., availability: _Optional[str] = ..., serviceclass: _Optional[str] = ..., windows: _Optional[_Union[VMServiceCreateWindowsRequest, _Mapping]] = ..., linux: _Optional[_Union[VMServiceCreateLinuxRequest, _Mapping]] = ...) -> None: ...

class LinuxDisk(_message.Message):
    __slots__ = ("auto_extend", "size_in_gb", "label", "mount_point")
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    SIZE_IN_GB_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    MOUNT_POINT_FIELD_NUMBER: _ClassVar[int]
    auto_extend: bool
    size_in_gb: int
    label: str
    mount_point: str
    def __init__(self, auto_extend: _Optional[bool] = ..., size_in_gb: _Optional[int] = ..., label: _Optional[str] = ..., mount_point: _Optional[str] = ...) -> None: ...

class WindowsDisk(_message.Message):
    __slots__ = ("auto_extend", "size_in_gb", "driveletter", "label")
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    SIZE_IN_GB_FIELD_NUMBER: _ClassVar[int]
    DRIVELETTER_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    auto_extend: bool
    size_in_gb: int
    driveletter: str
    label: str
    def __init__(self, auto_extend: _Optional[bool] = ..., size_in_gb: _Optional[int] = ..., driveletter: _Optional[str] = ..., label: _Optional[str] = ...) -> None: ...

class VMServiceCreateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class VMServiceUpdateRequest(_message.Message):
    __slots__ = ("project", "update_meta")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    project: str
    update_meta: _common_pb2.UpdateMeta
    def __init__(self, project: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ...) -> None: ...

class VMServiceUpdateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class VMServiceListRequest(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: str
    def __init__(self, project: _Optional[str] = ...) -> None: ...

class VMServiceListResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class VMServiceDeleteRequest(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: str
    def __init__(self, project: _Optional[str] = ...) -> None: ...

class VMServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
