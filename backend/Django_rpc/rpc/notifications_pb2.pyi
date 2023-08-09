import rpc.struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NotificationCreateManualRequest(_message.Message):
    __slots__ = ["confirmation", "image", "important", "link", "settings_key", "target_id", "target_type", "text", "text_as_model", "user_ids"]
    CONFIRMATION_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    IMPORTANT_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    SETTINGS_KEY_FIELD_NUMBER: _ClassVar[int]
    TARGET_ID_FIELD_NUMBER: _ClassVar[int]
    TARGET_TYPE_FIELD_NUMBER: _ClassVar[int]
    TEXT_AS_MODEL_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    USER_IDS_FIELD_NUMBER: _ClassVar[int]
    confirmation: bool
    image: str
    important: bool
    link: str
    settings_key: str
    target_id: int
    target_type: int
    text: str
    text_as_model: bool
    user_ids: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, user_ids: _Optional[_Iterable[int]] = ..., settings_key: _Optional[str] = ..., target_id: _Optional[int] = ..., target_type: _Optional[int] = ..., text: _Optional[str] = ..., text_as_model: bool = ..., important: bool = ..., confirmation: bool = ..., link: _Optional[str] = ..., image: _Optional[str] = ...) -> None: ...

class NotificationCreateRequest(_message.Message):
    __slots__ = ["action", "image", "important", "link", "target_id", "target_type", "text", "type"]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    IMPORTANT_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    TARGET_ID_FIELD_NUMBER: _ClassVar[int]
    TARGET_TYPE_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    action: str
    image: str
    important: bool
    link: str
    target_id: int
    target_type: int
    text: str
    type: int
    def __init__(self, action: _Optional[str] = ..., target_id: _Optional[int] = ..., target_type: _Optional[int] = ..., important: bool = ..., type: _Optional[int] = ..., link: _Optional[str] = ..., image: _Optional[str] = ..., text: _Optional[str] = ...) -> None: ...

class NotificationCreateResponse(_message.Message):
    __slots__ = ["is_created"]
    IS_CREATED_FIELD_NUMBER: _ClassVar[int]
    is_created: bool
    def __init__(self, is_created: bool = ...) -> None: ...

class NotificationManageRequest(_message.Message):
    __slots__ = ["action", "notification_ids", "notification_type", "user_id"]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_IDS_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_TYPE_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    action: str
    notification_ids: _containers.RepeatedScalarFieldContainer[int]
    notification_type: str
    user_id: int
    def __init__(self, notification_ids: _Optional[_Iterable[int]] = ..., user_id: _Optional[int] = ..., action: _Optional[str] = ..., notification_type: _Optional[str] = ...) -> None: ...

class NotificationManageResponse(_message.Message):
    __slots__ = ["success"]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class UserCountNotificationRequest(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    def __init__(self, user_id: _Optional[int] = ...) -> None: ...

class UserCountNotificationResponse(_message.Message):
    __slots__ = ["count", "has_important"]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    HAS_IMPORTANT_FIELD_NUMBER: _ClassVar[int]
    count: int
    has_important: bool
    def __init__(self, count: _Optional[int] = ..., has_important: bool = ...) -> None: ...

class UserMassNotification(_message.Message):
    __slots__ = ["notification", "read"]
    NOTIFICATION_FIELD_NUMBER: _ClassVar[int]
    READ_FIELD_NUMBER: _ClassVar[int]
    notification: _struct_pb2.Struct
    read: bool
    def __init__(self, notification: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., read: bool = ...) -> None: ...

class UserMassNotificationRequest(_message.Message):
    __slots__ = ["action", "count", "only_important", "page", "read", "type", "user_id"]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    ONLY_IMPORTANT_FIELD_NUMBER: _ClassVar[int]
    PAGE_FIELD_NUMBER: _ClassVar[int]
    READ_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    action: str
    count: int
    only_important: bool
    page: int
    read: bool
    type: int
    user_id: int
    def __init__(self, type: _Optional[int] = ..., user_id: _Optional[int] = ..., action: _Optional[str] = ..., only_important: bool = ..., read: bool = ..., page: _Optional[int] = ..., count: _Optional[int] = ...) -> None: ...

class UserMassNotificationResponse(_message.Message):
    __slots__ = ["notifications"]
    NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    notifications: _containers.RepeatedCompositeFieldContainer[_struct_pb2.Struct]
    def __init__(self, notifications: _Optional[_Iterable[_Union[_struct_pb2.Struct, _Mapping]]] = ...) -> None: ...

class UserNotificationsRequest(_message.Message):
    __slots__ = ["count", "only_important", "page", "read", "type", "user_id"]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    ONLY_IMPORTANT_FIELD_NUMBER: _ClassVar[int]
    PAGE_FIELD_NUMBER: _ClassVar[int]
    READ_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    count: int
    only_important: bool
    page: int
    read: bool
    type: int
    user_id: int
    def __init__(self, user_id: _Optional[int] = ..., type: _Optional[int] = ..., only_important: bool = ..., read: bool = ..., page: _Optional[int] = ..., count: _Optional[int] = ...) -> None: ...

class UserNotificationsResponse(_message.Message):
    __slots__ = ["notifications"]
    NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    notifications: _containers.RepeatedCompositeFieldContainer[_struct_pb2.Struct]
    def __init__(self, notifications: _Optional[_Iterable[_Union[_struct_pb2.Struct, _Mapping]]] = ...) -> None: ...
