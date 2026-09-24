from buf.validate import validate_pb2 as _validate_pb2
from fits.api.mvm.v1 import vlan_pb2 as _vlan_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Availability(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    AVAILABILITY_UNSPECIFIED: _ClassVar[Availability]
    AVAILABILITY_V0: _ClassVar[Availability]
    AVAILABILITY_V1: _ClassVar[Availability]
    AVAILABILITY_V2: _ClassVar[Availability]
    AVAILABILITY_V3: _ClassVar[Availability]

class ServiceClass(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SERVICE_CLASS_UNSPECIFIED: _ClassVar[ServiceClass]
    SERVICE_CLASS_SZ1: _ClassVar[ServiceClass]
    SERVICE_CLASS_SZ2: _ClassVar[ServiceClass]
    SERVICE_CLASS_SZ3: _ClassVar[ServiceClass]

class DiskType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DISK_TYPE_UNSPECIFIED: _ClassVar[DiskType]
    DISK_TYPE_OS: _ClassVar[DiskType]
    DISK_TYPE_SYSTEM: _ClassVar[DiskType]
    DISK_TYPE_DATA: _ClassVar[DiskType]
AVAILABILITY_UNSPECIFIED: Availability
AVAILABILITY_V0: Availability
AVAILABILITY_V1: Availability
AVAILABILITY_V2: Availability
AVAILABILITY_V3: Availability
SERVICE_CLASS_UNSPECIFIED: ServiceClass
SERVICE_CLASS_SZ1: ServiceClass
SERVICE_CLASS_SZ2: ServiceClass
SERVICE_CLASS_SZ3: ServiceClass
DISK_TYPE_UNSPECIFIED: DiskType
DISK_TYPE_OS: DiskType
DISK_TYPE_SYSTEM: DiskType
DISK_TYPE_DATA: DiskType

class LinuxDetails(_message.Message):
    __slots__ = ("disks", "ldap_uuid", "ldap_fqdn")
    DISKS_FIELD_NUMBER: _ClassVar[int]
    LDAP_UUID_FIELD_NUMBER: _ClassVar[int]
    LDAP_FQDN_FIELD_NUMBER: _ClassVar[int]
    disks: _containers.RepeatedCompositeFieldContainer[LinuxDisk]
    ldap_uuid: str
    ldap_fqdn: str
    def __init__(self, disks: _Optional[_Iterable[_Union[LinuxDisk, _Mapping]]] = ..., ldap_uuid: _Optional[str] = ..., ldap_fqdn: _Optional[str] = ...) -> None: ...

class WindowsDetails(_message.Message):
    __slots__ = ("disks", "domain_uuid", "domain_fqdn")
    DISKS_FIELD_NUMBER: _ClassVar[int]
    DOMAIN_UUID_FIELD_NUMBER: _ClassVar[int]
    DOMAIN_FQDN_FIELD_NUMBER: _ClassVar[int]
    disks: _containers.RepeatedCompositeFieldContainer[WindowsDisk]
    domain_uuid: str
    domain_fqdn: str
    def __init__(self, disks: _Optional[_Iterable[_Union[WindowsDisk, _Mapping]]] = ..., domain_uuid: _Optional[str] = ..., domain_fqdn: _Optional[str] = ...) -> None: ...

class NetworkInterface(_message.Message):
    __slots__ = ("uuid", "ipaddress", "macaddress")
    UUID_FIELD_NUMBER: _ClassVar[int]
    IPADDRESS_FIELD_NUMBER: _ClassVar[int]
    MACADDRESS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    ipaddress: str
    macaddress: str
    def __init__(self, uuid: _Optional[str] = ..., ipaddress: _Optional[str] = ..., macaddress: _Optional[str] = ...) -> None: ...

class MVMInstance(_message.Message):
    __slots__ = ("uuid", "meta", "fqdn", "tenant", "project_uuid", "os_uuid", "location_uuid", "contact_uuid", "cpu", "ram", "backup", "status", "status_info", "availability", "serviceclass", "windows_details", "linux_details", "interfaces", "vlan", "order_number", "contract")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    FQDN_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    STATUS_INFO_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    SERVICECLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_DETAILS_FIELD_NUMBER: _ClassVar[int]
    LINUX_DETAILS_FIELD_NUMBER: _ClassVar[int]
    INTERFACES_FIELD_NUMBER: _ClassVar[int]
    VLAN_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CONTRACT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    fqdn: str
    tenant: str
    project_uuid: str
    os_uuid: str
    location_uuid: str
    contact_uuid: str
    cpu: int
    ram: int
    backup: bool
    status: str
    status_info: str
    availability: Availability
    serviceclass: ServiceClass
    windows_details: WindowsDetails
    linux_details: LinuxDetails
    interfaces: _containers.RepeatedCompositeFieldContainer[NetworkInterface]
    vlan: _vlan_pb2.VLAN
    order_number: str
    contract: bool
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., fqdn: _Optional[str] = ..., tenant: _Optional[str] = ..., project_uuid: _Optional[str] = ..., os_uuid: _Optional[str] = ..., location_uuid: _Optional[str] = ..., contact_uuid: _Optional[str] = ..., cpu: _Optional[int] = ..., ram: _Optional[int] = ..., backup: _Optional[bool] = ..., status: _Optional[str] = ..., status_info: _Optional[str] = ..., availability: _Optional[_Union[Availability, str]] = ..., serviceclass: _Optional[_Union[ServiceClass, str]] = ..., windows_details: _Optional[_Union[WindowsDetails, _Mapping]] = ..., linux_details: _Optional[_Union[LinuxDetails, _Mapping]] = ..., interfaces: _Optional[_Iterable[_Union[NetworkInterface, _Mapping]]] = ..., vlan: _Optional[_Union[_vlan_pb2.VLAN, _Mapping]] = ..., order_number: _Optional[str] = ..., contract: _Optional[bool] = ...) -> None: ...

class MVMServiceGetRequest(_message.Message):
    __slots__ = ("uuid", "tenant")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    tenant: str
    def __init__(self, uuid: _Optional[str] = ..., tenant: _Optional[str] = ...) -> None: ...

class MVMServiceGetResponse(_message.Message):
    __slots__ = ("mvm",)
    MVM_FIELD_NUMBER: _ClassVar[int]
    mvm: MVMInstance
    def __init__(self, mvm: _Optional[_Union[MVMInstance, _Mapping]] = ...) -> None: ...

class MVMServiceCreateWindowsRequest(_message.Message):
    __slots__ = ("domain_uuid", "disks")
    DOMAIN_UUID_FIELD_NUMBER: _ClassVar[int]
    DISKS_FIELD_NUMBER: _ClassVar[int]
    domain_uuid: str
    disks: _containers.RepeatedCompositeFieldContainer[WindowsDisk]
    def __init__(self, domain_uuid: _Optional[str] = ..., disks: _Optional[_Iterable[_Union[WindowsDisk, _Mapping]]] = ...) -> None: ...

class MVMServiceCreateLinuxRequest(_message.Message):
    __slots__ = ("ldap_uuid", "disks")
    LDAP_UUID_FIELD_NUMBER: _ClassVar[int]
    DISKS_FIELD_NUMBER: _ClassVar[int]
    ldap_uuid: str
    disks: _containers.RepeatedCompositeFieldContainer[LinuxDisk]
    def __init__(self, ldap_uuid: _Optional[str] = ..., disks: _Optional[_Iterable[_Union[LinuxDisk, _Mapping]]] = ...) -> None: ...

class MVMServiceCreateRequest(_message.Message):
    __slots__ = ("project_uuid", "name", "os_uuid", "vlan_uuid", "location_uuid", "contact_uuid", "cpu", "ram", "order_number", "labels", "backup", "availability", "serviceclass", "windows", "linux")
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    VLAN_UUID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    SERVICECLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_FIELD_NUMBER: _ClassVar[int]
    LINUX_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    name: str
    os_uuid: str
    vlan_uuid: str
    location_uuid: str
    contact_uuid: str
    cpu: int
    ram: int
    order_number: str
    labels: _common_pb2.Labels
    backup: bool
    availability: Availability
    serviceclass: ServiceClass
    windows: MVMServiceCreateWindowsRequest
    linux: MVMServiceCreateLinuxRequest
    def __init__(self, project_uuid: _Optional[str] = ..., name: _Optional[str] = ..., os_uuid: _Optional[str] = ..., vlan_uuid: _Optional[str] = ..., location_uuid: _Optional[str] = ..., contact_uuid: _Optional[str] = ..., cpu: _Optional[int] = ..., ram: _Optional[int] = ..., order_number: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., backup: _Optional[bool] = ..., availability: _Optional[_Union[Availability, str]] = ..., serviceclass: _Optional[_Union[ServiceClass, str]] = ..., windows: _Optional[_Union[MVMServiceCreateWindowsRequest, _Mapping]] = ..., linux: _Optional[_Union[MVMServiceCreateLinuxRequest, _Mapping]] = ...) -> None: ...

class LinuxDisk(_message.Message):
    __slots__ = ("uuid", "label", "auto_extend", "size", "mount_point", "disk_type")
    UUID_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    MOUNT_POINT_FIELD_NUMBER: _ClassVar[int]
    DISK_TYPE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    label: str
    auto_extend: bool
    size: int
    mount_point: str
    disk_type: DiskType
    def __init__(self, uuid: _Optional[str] = ..., label: _Optional[str] = ..., auto_extend: _Optional[bool] = ..., size: _Optional[int] = ..., mount_point: _Optional[str] = ..., disk_type: _Optional[_Union[DiskType, str]] = ...) -> None: ...

class WindowsDisk(_message.Message):
    __slots__ = ("uuid", "label", "auto_extend", "size", "driveletter", "disk_type")
    UUID_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    DRIVELETTER_FIELD_NUMBER: _ClassVar[int]
    DISK_TYPE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    label: str
    auto_extend: bool
    size: int
    driveletter: str
    disk_type: DiskType
    def __init__(self, uuid: _Optional[str] = ..., label: _Optional[str] = ..., auto_extend: _Optional[bool] = ..., size: _Optional[int] = ..., driveletter: _Optional[str] = ..., disk_type: _Optional[_Union[DiskType, str]] = ...) -> None: ...

class MVMServiceCreateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceUpdateRequest(_message.Message):
    __slots__ = ("uuid", "project", "order_number", "update_meta", "performance_class", "service_class", "contact")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    PERFORMANCE_CLASS_FIELD_NUMBER: _ClassVar[int]
    SERVICE_CLASS_FIELD_NUMBER: _ClassVar[int]
    CONTACT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project: str
    order_number: str
    update_meta: _common_pb2.UpdateMeta
    performance_class: PerformanceClassChange
    service_class: ServiceClassChange
    contact: ContactChange
    def __init__(self, uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., performance_class: _Optional[_Union[PerformanceClassChange, _Mapping]] = ..., service_class: _Optional[_Union[ServiceClassChange, _Mapping]] = ..., contact: _Optional[_Union[ContactChange, _Mapping]] = ...) -> None: ...

class PerformanceClassChange(_message.Message):
    __slots__ = ("cpu", "ram")
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_FIELD_NUMBER: _ClassVar[int]
    cpu: int
    ram: int
    def __init__(self, cpu: _Optional[int] = ..., ram: _Optional[int] = ...) -> None: ...

class ServiceClassChange(_message.Message):
    __slots__ = ("serviceclass",)
    SERVICECLASS_FIELD_NUMBER: _ClassVar[int]
    serviceclass: ServiceClass
    def __init__(self, serviceclass: _Optional[_Union[ServiceClass, str]] = ...) -> None: ...

class ContactChange(_message.Message):
    __slots__ = ("contact_uuid",)
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    contact_uuid: str
    def __init__(self, contact_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceUpdateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceListRequest(_message.Message):
    __slots__ = ("tenant",)
    TENANT_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    def __init__(self, tenant: _Optional[str] = ...) -> None: ...

class MVMServiceListResponse(_message.Message):
    __slots__ = ("mvms",)
    MVMS_FIELD_NUMBER: _ClassVar[int]
    mvms: _containers.RepeatedCompositeFieldContainer[MVMInstance]
    def __init__(self, mvms: _Optional[_Iterable[_Union[MVMInstance, _Mapping]]] = ...) -> None: ...

class MVMServiceDeleteRequest(_message.Message):
    __slots__ = ("project", "uuid", "order_number")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    project: str
    uuid: str
    order_number: str
    def __init__(self, project: _Optional[str] = ..., uuid: _Optional[str] = ..., order_number: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateCreateRequest(_message.Message):
    __slots__ = ("create",)
    CREATE_FIELD_NUMBER: _ClassVar[int]
    create: MVMServiceCreateRequest
    def __init__(self, create: _Optional[_Union[MVMServiceCreateRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateCreateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateAddDiskRequest(_message.Message):
    __slots__ = ("add_disk",)
    ADD_DISK_FIELD_NUMBER: _ClassVar[int]
    add_disk: MVMServiceAddDiskRequest
    def __init__(self, add_disk: _Optional[_Union[MVMServiceAddDiskRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateAddDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateUpdateDiskRequest(_message.Message):
    __slots__ = ("update_disk",)
    UPDATE_DISK_FIELD_NUMBER: _ClassVar[int]
    update_disk: MVMServiceUpdateDiskRequest
    def __init__(self, update_disk: _Optional[_Union[MVMServiceUpdateDiskRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateUpdateDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceUpdateDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceDeleteDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddDiskRequest(_message.Message):
    __slots__ = ("uuid", "project", "order_number", "linux", "windows")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    LINUX_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project: str
    order_number: str
    linux: LinuxDisk
    windows: WindowsDisk
    def __init__(self, uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ..., linux: _Optional[_Union[LinuxDisk, _Mapping]] = ..., windows: _Optional[_Union[WindowsDisk, _Mapping]] = ...) -> None: ...

class MVMServiceUpdateDiskRequest(_message.Message):
    __slots__ = ("uuid", "disk_uuid", "project", "order_number", "auto_extend", "size")
    UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    disk_uuid: str
    project: str
    order_number: str
    auto_extend: bool
    size: int
    def __init__(self, uuid: _Optional[str] = ..., disk_uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ..., auto_extend: _Optional[bool] = ..., size: _Optional[int] = ...) -> None: ...

class MVMServiceDeleteDiskRequest(_message.Message):
    __slots__ = ("uuid", "disk_uuid", "project", "order_number")
    UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    disk_uuid: str
    project: str
    order_number: str
    def __init__(self, uuid: _Optional[str] = ..., disk_uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ...) -> None: ...

class MVMServiceAddNetworkInterfaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceMoveNetworkInterfaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceDeleteNetworkInterfaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddNetworkInterfaceRequest(_message.Message):
    __slots__ = ("uuid", "project", "order_number")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project: str
    order_number: str
    def __init__(self, uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ...) -> None: ...

class MVMServiceMoveNetworkInterfaceRequest(_message.Message):
    __slots__ = ("uuid", "interface_uuid", "project", "order_number", "target_mvm_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    INTERFACE_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    TARGET_MVM_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    interface_uuid: str
    project: str
    order_number: str
    target_mvm_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., interface_uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ..., target_mvm_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteNetworkInterfaceRequest(_message.Message):
    __slots__ = ("uuid", "interface_uuid", "project", "order_number")
    UUID_FIELD_NUMBER: _ClassVar[int]
    INTERFACE_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    ORDER_NUMBER_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    interface_uuid: str
    project: str
    order_number: str
    def __init__(self, uuid: _Optional[str] = ..., interface_uuid: _Optional[str] = ..., project: _Optional[str] = ..., order_number: _Optional[str] = ...) -> None: ...

class MVMServiceValidateAddNetworkInterfaceRequest(_message.Message):
    __slots__ = ("add_network_interface",)
    ADD_NETWORK_INTERFACE_FIELD_NUMBER: _ClassVar[int]
    add_network_interface: MVMServiceAddNetworkInterfaceRequest
    def __init__(self, add_network_interface: _Optional[_Union[MVMServiceAddNetworkInterfaceRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateAddNetworkInterfaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
