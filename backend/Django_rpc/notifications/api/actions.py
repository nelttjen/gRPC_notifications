from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.user_notification_actions import UserNotificationsActionRPC


class UserNotificationActionView(APIView):

    def post(self, request, user_id):
        # ignore authorization in this project
        ids = request.data.get('notifications')
        action = request.data.get('action')
        not_type = request.data.get('type')

        conds = [
            isinstance(ids, list),
            action in ('read', 'unread', 'delete'),
            not_type in ('user', 'mass')
        ]
        if not all(conds):
            return Response({"msg": 'bad parameters'}, status=400)

        grpc_client = UserNotificationsActionRPC(
            user_id=user_id,
            action=action,
            notification_type=not_type,
            notification_ids=ids
        )

        response = grpc_client.perform_request()

        status = 200 if response.success else 400

        return Response({"msg": f"rpc returned {response.success}"}, status=status)