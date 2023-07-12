import grpc
import rpc.notifications_pb2 as pb2
import rpc.notifications_pb2_grpc as pb2_grpc

from drpc.settings import GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK
from notifications.modules.typing.rpc import NotificationAction
from notifications.modules.gRPC import BaseRPCRequestClass


class NotificationActionRPC(BaseRPCRequestClass):
    def __init__(self, action, target_id, target_type, not_type,
                 important=False, link=None, image=None, text=None):
        self.action = action
        self.target_id = target_id
        self.target_type = target_type
        self.type = not_type
        self.important = important
        self.link = link
        self.image = image
        self.text = text

    def perform_request(self) -> pb2.NotificationCreateResponse:
        channel = grpc.insecure_channel(GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK)
        stub = pb2_grpc.CreateNotificationsStub(channel=channel)

        request = NotificationAction(
            action=self.action,
            target_id=self.target_id,
            target_type=self.target_type,
            type=self.type,
            important=self.important,
            link=self.link,
            image=self.image,
            text=self.text,
        ).as_grpc_request()

        response = stub.CreateNotificationsAction(request)
        self.process_response(response)
        return response

    def process_response(self, response: pb2.NotificationCreateResponse) -> bool:
        return True
