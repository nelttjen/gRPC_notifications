import pydantic
import rpc.notifications_pb2 as pb2


from typing import Optional


class NotificationAction(pydantic.BaseModel):
    action: str
    target_id: Optional[int]
    target_type: Optional[str]
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