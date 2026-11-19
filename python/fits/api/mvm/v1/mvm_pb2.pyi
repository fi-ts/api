from buf.validate import validate_pb2 as _validate_pb2
from fits.api.mvm.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2_1
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

class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    STATE_UNSPECIFIED: _ClassVar[State]
    STATE_CREATING: _ClassVar[State]
    STATE_ACTIVE: _ClassVar[State]
    STATE_UPDATING: _ClassVar[State]
    STATE_DELETED: _ClassVar[State]
    STATE_ERROR: _ClassVar[State]

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
STATE_UNSPECIFIED: State
STATE_CREATING: State
STATE_ACTIVE: State
STATE_UPDATING: State
STATE_DELETED: State
STATE_ERROR: State
DISK_TYPE_UNSPECIFIED: DiskType
DISK_TYPE_OS: DiskType
DISK_TYPE_SYSTEM: DiskType
DISK_TYPE_DATA: DiskType

class LinuxDisk(_message.Message):
    __slots__ = ("uuid", "label", "size_gb", "auto_extend", "mount_point", "disk_type")
    UUID_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    SIZE_GB_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    MOUNT_POINT_FIELD_NUMBER: _ClassVar[int]
    DISK_TYPE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    label: str
    size_gb: int
    auto_extend: bool
    mount_point: str
    disk_type: DiskType
    def __init__(self, uuid: _Optional[str] = ..., label: _Optional[str] = ..., size_gb: _Optional[int] = ..., auto_extend: _Optional[bool] = ..., mount_point: _Optional[str] = ..., disk_type: _Optional[_Union[DiskType, str]] = ...) -> None: ...

class WindowsDisk(_message.Message):
    __slots__ = ("uuid", "label", "size_gb", "auto_extend", "drive_letter", "disk_type")
    UUID_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    SIZE_GB_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    DRIVE_LETTER_FIELD_NUMBER: _ClassVar[int]
    DISK_TYPE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    label: str
    size_gb: int
    auto_extend: bool
    drive_letter: str
    disk_type: DiskType
    def __init__(self, uuid: _Optional[str] = ..., label: _Optional[str] = ..., size_gb: _Optional[int] = ..., auto_extend: _Optional[bool] = ..., drive_letter: _Optional[str] = ..., disk_type: _Optional[_Union[DiskType, str]] = ...) -> None: ...

class Disk(_message.Message):
    __slots__ = ("linux", "windows")
    LINUX_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_FIELD_NUMBER: _ClassVar[int]
    linux: LinuxDisk
    windows: WindowsDisk
    def __init__(self, linux: _Optional[_Union[LinuxDisk, _Mapping]] = ..., windows: _Optional[_Union[WindowsDisk, _Mapping]] = ...) -> None: ...

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

class IPAddress(_message.Message):
    __slots__ = ("uuid", "ip_address", "network_interface")
    UUID_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    NETWORK_INTERFACE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    ip_address: str
    network_interface: NetworkInterface
    def __init__(self, uuid: _Optional[str] = ..., ip_address: _Optional[str] = ..., network_interface: _Optional[_Union[NetworkInterface, _Mapping]] = ...) -> None: ...

class NetworkInterface(_message.Message):
    __slots__ = ("vlan_uuid", "mac_address")
    VLAN_UUID_FIELD_NUMBER: _ClassVar[int]
    MAC_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    vlan_uuid: str
    mac_address: str
    def __init__(self, vlan_uuid: _Optional[str] = ..., mac_address: _Optional[str] = ...) -> None: ...

class InitialIPAddress(_message.Message):
    __slots__ = ("vlan_uuid",)
    VLAN_UUID_FIELD_NUMBER: _ClassVar[int]
    vlan_uuid: str
    def __init__(self, vlan_uuid: _Optional[str] = ...) -> None: ...

class Status(_message.Message):
    __slots__ = ("state", "info")
    STATE_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    state: State
    info: str
    def __init__(self, state: _Optional[_Union[State, str]] = ..., info: _Optional[str] = ...) -> None: ...

class Backup(_message.Message):
    __slots__ = ("enabled",)
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    def __init__(self, enabled: _Optional[bool] = ...) -> None: ...

class Encryption(_message.Message):
    __slots__ = ("enabled",)
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    def __init__(self, enabled: _Optional[bool] = ...) -> None: ...

class ManagedVM(_message.Message):
    __slots__ = ("uuid", "meta", "fqdn", "tenant", "project_uuid", "os_uuid", "location_uuid", "contact_uuid", "cpu", "ram_gb", "backup", "encryption", "status", "availability", "service_class", "windows_details", "linux_details", "ip_addresses", "order_reference", "snapshots")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    FQDN_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_GB_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    ENCRYPTION_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    SERVICE_CLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_DETAILS_FIELD_NUMBER: _ClassVar[int]
    LINUX_DETAILS_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESSES_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    SNAPSHOTS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    fqdn: str
    tenant: str
    project_uuid: str
    os_uuid: str
    location_uuid: str
    contact_uuid: str
    cpu: int
    ram_gb: int
    backup: Backup
    encryption: Encryption
    status: Status
    availability: Availability
    service_class: ServiceClass
    windows_details: WindowsDetails
    linux_details: LinuxDetails
    ip_addresses: _containers.RepeatedCompositeFieldContainer[IPAddress]
    order_reference: str
    snapshots: _containers.RepeatedCompositeFieldContainer[Snapshot]
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., fqdn: _Optional[str] = ..., tenant: _Optional[str] = ..., project_uuid: _Optional[str] = ..., os_uuid: _Optional[str] = ..., location_uuid: _Optional[str] = ..., contact_uuid: _Optional[str] = ..., cpu: _Optional[int] = ..., ram_gb: _Optional[int] = ..., backup: _Optional[_Union[Backup, _Mapping]] = ..., encryption: _Optional[_Union[Encryption, _Mapping]] = ..., status: _Optional[_Union[Status, _Mapping]] = ..., availability: _Optional[_Union[Availability, str]] = ..., service_class: _Optional[_Union[ServiceClass, str]] = ..., windows_details: _Optional[_Union[WindowsDetails, _Mapping]] = ..., linux_details: _Optional[_Union[LinuxDetails, _Mapping]] = ..., ip_addresses: _Optional[_Iterable[_Union[IPAddress, _Mapping]]] = ..., order_reference: _Optional[str] = ..., snapshots: _Optional[_Iterable[_Union[Snapshot, _Mapping]]] = ...) -> None: ...

class Snapshot(_message.Message):
    __slots__ = ("uuid", "name")
    UUID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    name: str
    def __init__(self, uuid: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class MVMServiceGetRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceGetResponse(_message.Message):
    __slots__ = ("mvm",)
    MVM_FIELD_NUMBER: _ClassVar[int]
    mvm: ManagedVM
    def __init__(self, mvm: _Optional[_Union[ManagedVM, _Mapping]] = ...) -> None: ...

class MVMServiceListRequest(_message.Message):
    __slots__ = ("project_uuid",)
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    def __init__(self, project_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceListResponse(_message.Message):
    __slots__ = ("mvms",)
    MVMS_FIELD_NUMBER: _ClassVar[int]
    mvms: _containers.RepeatedCompositeFieldContainer[ManagedVM]
    def __init__(self, mvms: _Optional[_Iterable[_Union[ManagedVM, _Mapping]]] = ...) -> None: ...

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
    __slots__ = ("project_uuid", "name_code", "os_uuid", "location_uuid", "contact_uuid", "cpu", "ram_gb", "order_reference", "labels", "backup", "encryption", "availability", "service_class", "windows", "linux", "initial_ip_address")
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    NAME_CODE_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_GB_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    ENCRYPTION_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    SERVICE_CLASS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_FIELD_NUMBER: _ClassVar[int]
    LINUX_FIELD_NUMBER: _ClassVar[int]
    INITIAL_IP_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    project_uuid: str
    name_code: str
    os_uuid: str
    location_uuid: str
    contact_uuid: str
    cpu: int
    ram_gb: int
    order_reference: str
    labels: _common_pb2.Labels
    backup: Backup
    encryption: Encryption
    availability: Availability
    service_class: ServiceClass
    windows: MVMServiceCreateWindowsRequest
    linux: MVMServiceCreateLinuxRequest
    initial_ip_address: InitialIPAddress
    def __init__(self, project_uuid: _Optional[str] = ..., name_code: _Optional[str] = ..., os_uuid: _Optional[str] = ..., location_uuid: _Optional[str] = ..., contact_uuid: _Optional[str] = ..., cpu: _Optional[int] = ..., ram_gb: _Optional[int] = ..., order_reference: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., backup: _Optional[_Union[Backup, _Mapping]] = ..., encryption: _Optional[_Union[Encryption, _Mapping]] = ..., availability: _Optional[_Union[Availability, str]] = ..., service_class: _Optional[_Union[ServiceClass, str]] = ..., windows: _Optional[_Union[MVMServiceCreateWindowsRequest, _Mapping]] = ..., linux: _Optional[_Union[MVMServiceCreateLinuxRequest, _Mapping]] = ..., initial_ip_address: _Optional[_Union[InitialIPAddress, _Mapping]] = ...) -> None: ...

class MVMServiceCreateResponse(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class MVMServiceValidateCreateRequest(_message.Message):
    __slots__ = ("create",)
    CREATE_FIELD_NUMBER: _ClassVar[int]
    create: MVMServiceCreateRequest
    def __init__(self, create: _Optional[_Union[MVMServiceCreateRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateCreateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceUpdateRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "update_meta", "order_reference", "labels", "contact_uuid", "cpu", "ram_gb")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    CONTACT_UUID_FIELD_NUMBER: _ClassVar[int]
    CPU_FIELD_NUMBER: _ClassVar[int]
    RAM_GB_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    update_meta: _common_pb2.UpdateMeta
    order_reference: str
    labels: _common_pb2.UpdateLabels
    contact_uuid: str
    cpu: int
    ram_gb: int
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., order_reference: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ..., contact_uuid: _Optional[str] = ..., cpu: _Optional[int] = ..., ram_gb: _Optional[int] = ...) -> None: ...

class MVMServiceUpdateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateUpdateRequest(_message.Message):
    __slots__ = ("update",)
    UPDATE_FIELD_NUMBER: _ClassVar[int]
    update: MVMServiceUpdateRequest
    def __init__(self, update: _Optional[_Union[MVMServiceUpdateRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateUpdateResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceUpdateServiceClassRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "service_class", "availability")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    SERVICE_CLASS_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    service_class: ServiceClass
    availability: Availability
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., service_class: _Optional[_Union[ServiceClass, str]] = ..., availability: _Optional[_Union[Availability, str]] = ...) -> None: ...

class MVMServiceUpdateServiceClassResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateUpdateServiceClassRequest(_message.Message):
    __slots__ = ("update_service_class",)
    UPDATE_SERVICE_CLASS_FIELD_NUMBER: _ClassVar[int]
    update_service_class: MVMServiceUpdateServiceClassRequest
    def __init__(self, update_service_class: _Optional[_Union[MVMServiceUpdateServiceClassRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateUpdateServiceClassResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceDeleteRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddDisksRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference", "disks")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    DISKS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    disks: _containers.RepeatedCompositeFieldContainer[Disk]
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., disks: _Optional[_Iterable[_Union[Disk, _Mapping]]] = ...) -> None: ...

class MVMServiceAddDisksResponse(_message.Message):
    __slots__ = ("uuids",)
    UUIDS_FIELD_NUMBER: _ClassVar[int]
    uuids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, uuids: _Optional[_Iterable[str]] = ...) -> None: ...

class MVMServiceValidateAddDisksRequest(_message.Message):
    __slots__ = ("add_disks",)
    ADD_DISKS_FIELD_NUMBER: _ClassVar[int]
    add_disks: MVMServiceAddDisksRequest
    def __init__(self, add_disks: _Optional[_Union[MVMServiceAddDisksRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateAddDisksResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceUpdateDiskRequest(_message.Message):
    __slots__ = ("uuid", "disk_uuid", "project_uuid", "order_reference", "size_gb", "auto_extend")
    UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    SIZE_GB_FIELD_NUMBER: _ClassVar[int]
    AUTO_EXTEND_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    disk_uuid: str
    project_uuid: str
    order_reference: str
    size_gb: int
    auto_extend: bool
    def __init__(self, uuid: _Optional[str] = ..., disk_uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., size_gb: _Optional[int] = ..., auto_extend: _Optional[bool] = ...) -> None: ...

class MVMServiceUpdateDiskResponse(_message.Message):
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

class MVMServiceDeleteDiskRequest(_message.Message):
    __slots__ = ("uuid", "disk_uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    DISK_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    disk_uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., disk_uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteDiskResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddIPAddressRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceAddIPAddressResponse(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class MVMServiceMoveIPAddressRequest(_message.Message):
    __slots__ = ("uuid", "ip_address_uuid", "project_uuid", "order_reference", "target_mvm_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESS_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    TARGET_MVM_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    ip_address_uuid: str
    project_uuid: str
    order_reference: str
    target_mvm_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., ip_address_uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., target_mvm_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceMoveIPAddressResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceDeleteIPAddressRequest(_message.Message):
    __slots__ = ("uuid", "ip_address_uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESS_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    ip_address_uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., ip_address_uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteIPAddressResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceRestartRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceRestartResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceReinstallRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference", "os_uuid", "backup", "encryption")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    OS_UUID_FIELD_NUMBER: _ClassVar[int]
    BACKUP_FIELD_NUMBER: _ClassVar[int]
    ENCRYPTION_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    os_uuid: str
    backup: Backup
    encryption: Encryption
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., os_uuid: _Optional[str] = ..., backup: _Optional[_Union[Backup, _Mapping]] = ..., encryption: _Optional[_Union[Encryption, _Mapping]] = ...) -> None: ...

class MVMServiceReinstallResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceValidateReinstallRequest(_message.Message):
    __slots__ = ("reinstall",)
    REINSTALL_FIELD_NUMBER: _ClassVar[int]
    reinstall: MVMServiceReinstallRequest
    def __init__(self, reinstall: _Optional[_Union[MVMServiceReinstallRequest, _Mapping]] = ...) -> None: ...

class MVMServiceValidateReinstallResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceAddSnapshotRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference", "name")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    name: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class MVMServiceAddSnapshotResponse(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteSnapshotRequest(_message.Message):
    __slots__ = ("uuid", "snapshot_uuid", "project_uuid", "order_reference")
    UUID_FIELD_NUMBER: _ClassVar[int]
    SNAPSHOT_UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    snapshot_uuid: str
    project_uuid: str
    order_reference: str
    def __init__(self, uuid: _Optional[str] = ..., snapshot_uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ...) -> None: ...

class MVMServiceDeleteSnapshotResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MVMServiceMoveRequest(_message.Message):
    __slots__ = ("uuid", "project_uuid", "order_reference", "location_uuid")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_UUID_FIELD_NUMBER: _ClassVar[int]
    ORDER_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    LOCATION_UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    project_uuid: str
    order_reference: str
    location_uuid: str
    def __init__(self, uuid: _Optional[str] = ..., project_uuid: _Optional[str] = ..., order_reference: _Optional[str] = ..., location_uuid: _Optional[str] = ...) -> None: ...

class MVMServiceMoveResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
