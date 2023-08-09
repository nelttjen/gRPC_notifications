import grpc
import rpc.notifications_pb2_grpc as pb2_grpc

from drpc.settings import GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK
from rpc.notifications_pb2 import NotificationManageResponse
from notifications.modules.gRPC import BaseRPCRequestClass
from notifications.modules.typing.rpc.request import UserNotificationAction

import logging
logger = logging.getLogger('error')


class UserNotificationsActionRPC(BaseRPCRequestClass):
    def __init__(self, user_id, action, notification_ids, notification_type):
        self.user_id = user_id
        self.action = action
        self.notification_ids = notification_ids
        self.notification_type = notification_type

    def perform_request(self) -> 'NotificationManageResponse':
        channel = grpc.insecure_channel(GRPC_NOTIFICATIONS_SERVICE_CONNECT_LINK)
        stub = pb2_grpc.CreateNotificationsStub(channel=channel)

        request = UserNotificationAction(
            user_id=self.user_id,
            action=self.action,
            notification_ids=self.notification_ids,
            notification_type=self.notification_type
        ).as_grpc_request()

        response = stub.ManageNotifications(request)
        self.process_response(response)
        return response

    def process_response(self, response: 'NotificationManageResponse') -> bool:
        if not response.success:
            logger.error(f'Failed to {self.action} notifications for user {self.user_id}')
        return True
