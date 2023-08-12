import grpc
import rpc.notifications_pb2_grpc as pb2_grpc

from drpc.settings import GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK
from rpc.notifications_pb2 import UserCountNotificationResponse
from notifications.modules.gRPC import BaseRPCRequestClass
from notifications.modules.typing.rpc.request import UserNotificationCount


class CountNotificationsRPC(BaseRPCRequestClass):
    def __init__(self, user_id):
        self.user_id = user_id

    def perform_request(self) -> 'UserCountNotificationResponse':
        channel = grpc.insecure_channel(GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK)
        stub = pb2_grpc.CreateNotificationsStub(channel)

        request = UserNotificationCount(
            user_id=self.user_id,
        ).as_grpc_request()

        response = stub.CountNotifications(request)
        self.process_response(response)
        return response

    def process_response(self, response: 'UserCountNotificationResponse') -> bool:
        return True
