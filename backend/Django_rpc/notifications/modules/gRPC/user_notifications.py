import grpc
import rpc.notifications_pb2 as pb2
import rpc.notifications_pb2_grpc as pb2_grpc

from drpc.settings import GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK
from rpc.notifications_pb2 import UserCountNotificationResponse, UserNotificationsResponse
from notifications.modules.gRPC import BaseRPCRequestClass
from notifications.modules.typing.rpc import UserNotification


class UserNotificationsRPC(BaseRPCRequestClass):
    def __init__(self, user_id, page, count, only_important=False, read=False):
        self.user_id = user_id
        self.page = page
        self.count = count
        self.only_important = only_important
        self.read = read

    def perform_request(self) -> 'UserNotificationsResponse':
        channel = grpc.insecure_channel(GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK)
        stub = pb2_grpc.CreateNotificationsStub(channel)

        request = UserNotification(
            user_id=self.user_id,
            page=self.page,
            count=self.count,
            only_important=self.only_important,
            read=self.read
        ).as_grpc_request()

        response = stub.GetNotifications(request)
        self.process_response(response)
        return response

    def process_response(self, response: 'UserNotificationsResponse') -> bool:
        return True
