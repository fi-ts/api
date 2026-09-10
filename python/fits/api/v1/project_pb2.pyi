from buf.validate import validate_pb2 as _validate_pb2
from fits.api.v1 import common_pb2 as _common_pb2
from fits.api.v1 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Project(_message.Message):
    __slots__ = ("uuid", "meta", "name", "description", "tenant", "avatar_url")
    UUID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    AVATAR_URL_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    meta: _common_pb2.Meta
    name: str
    description: str
    tenant: str
    avatar_url: str
    def __init__(self, uuid: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., tenant: _Optional[str] = ..., avatar_url: _Optional[str] = ...) -> None: ...

class ProjectServiceListRequest(_message.Message):
    __slots__ = ("id", "name", "tenant", "labels")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    tenant: str
    labels: _common_pb2.Labels
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., tenant: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class ProjectServiceListResponse(_message.Message):
    __slots__ = ("projects",)
    PROJECTS_FIELD_NUMBER: _ClassVar[int]
    projects: _containers.RepeatedCompositeFieldContainer[Project]
    def __init__(self, projects: _Optional[_Iterable[_Union[Project, _Mapping]]] = ...) -> None: ...

class ProjectServiceGetRequest(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: str
    def __init__(self, project: _Optional[str] = ...) -> None: ...

class ProjectServiceGetResponse(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...

class ProjectServiceCreateRequest(_message.Message):
    __slots__ = ("login", "name", "description", "avatar_url", "labels")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    AVATAR_URL_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    login: str
    name: str
    description: str
    avatar_url: str
    labels: _common_pb2.Labels
    def __init__(self, login: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., avatar_url: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class ProjectServiceCreateResponse(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...

class ProjectServiceDeleteRequest(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: str
    def __init__(self, project: _Optional[str] = ...) -> None: ...

class ProjectServiceDeleteResponse(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...

class ProjectServiceUpdateRequest(_message.Message):
    __slots__ = ("project", "update_meta", "name", "description", "avatar_url", "labels")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    AVATAR_URL_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    project: str
    update_meta: _common_pb2.UpdateMeta
    name: str
    description: str
    avatar_url: str
    labels: _common_pb2.UpdateLabels
    def __init__(self, project: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., avatar_url: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ...) -> None: ...

class ProjectServiceUpdateResponse(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...
