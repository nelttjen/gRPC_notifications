import grpc
import rpc.notifications_pb2 as pb2
import rpc.notifications_pb2_grpc as pb2_grpc

from drpc.settings import GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK
from notifications.modules.typing.rpc import NotificationForUsers
from notifications.modules.gRPC import BaseRPCRequestClass


class NotificationForUsersRPC(BaseRPCRequestClass):
    def __init__(self, user_ids, sets_key, text,
                 text_as_model=True, important=False, confirmation=False,
                 target_id=None, target_type=None, link=None, image=None):
        self.user_ids = user_ids
        self.sets_key = sets_key
        self.text = text
        self.text_as_model = text_as_model
        self.important = important
        self.confirmation = confirmation
        self.target_id = target_id
        self.target_type = target_type
        self.link = link
        self.image = image

    def perform_request(self) -> 'pb2.NotificationCreateResponse':
        channel = grpc.insecure_channel(GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK)
        stub = pb2_grpc.CreateNotificationsStub(channel)

        request = NotificationForUsers(
            user_ids=self.user_ids,
            sets_key=self.sets_key,
            text=self.text,
            text_as_model=self.text_as_model,
            important=self.important,
            confirmation=self.confirmation,
            target_id=self.target_id,
            target_type=self.target_type,
            link=self.link,
            image=self.image
        ).as_grpc_request()

        response = stub.CreateNotificationForUsers(request)
        self.process_response(response)
        return response

    def process_response(self, response: pb2.NotificationCreateResponse) -> bool:
        if not response.is_created:
            print('Error creating notifications for users')
        return response.is_created
