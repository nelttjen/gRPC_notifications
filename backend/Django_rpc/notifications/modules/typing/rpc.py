import pydantic
import rpc.notifications_pb2 as pb2

from typing import Optional, List


class NotificationAction(pydantic.BaseModel):
    action: str
    target_id: Optional[int]
    target_type: Optional[int]
    important: bool = pydantic.Field(default=False)

    type: int

    link: Optional[str]
    image: Optional[str]
    text: Optional[str]

    def as_grpc_request(self) -> pb2.NotificationCreateRequest:
        data = {
            'action': self.action,
            'target_id': self.target_id,
            'target_type': self.target_type,
            'important': self.important,
            'type': self.type,
        }

        for item in ['link', 'image', 'text']:
            if getattr(self, item) is not None:
                data[item] = getattr(self, item)

        return pb2.NotificationCreateRequest(**data)


class NotificationForUsers(pydantic.BaseModel):
    user_ids: List[int]
    settings_key: str

    target_id: Optional[int]
    target_type: Optional[int]
    text: str
    text_as_model: bool = pydantic.Field(default=True)
    important: bool = pydantic.Field(default=False)
    confirmation: bool = pydantic.Field(default=False)

    link: Optional[str]
    image: Optional[str]

    def as_grpc_request(self) -> pb2.NotificationCreateManualRequest:
        data = {
            'user_ids': self.user_ids,
            'settings_key': self.settings_key,
            'text': self.text,
            'text_as_model': self.text_as_model,
            'important': self.important,
            'confirmation': self.confirmation,
        }

        for item in ['link', 'image', 'target_id', 'target_type']:
            if getattr(self, item) is not None:
                data[item] = getattr(self, item)

        return pb2.NotificationCreateManualRequest(**data)


class UserNotification(pydantic.BaseModel):
    user_id: int
    page: int
    count: int

    only_important: bool = pydantic.Field(default=False)
    read: bool = pydantic.Field(default=False)

    def as_grpc_request(self) -> pb2.UserNotificationsRequest:
        data = {
            'user_id': self.user_id,
            'page': self.page,
            'count': self.count,
            'only_important': self.only_important,
            'read': self.read,
        }

        return pb2.UserNotificationsRequest(**data)