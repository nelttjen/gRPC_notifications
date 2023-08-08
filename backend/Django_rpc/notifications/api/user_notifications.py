from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.user_notifications import UserNotificationsRPC, MassUserNotificationsRPC
from notifications.modules.typing.rpc.response import UserNotificationResponse, UserMassNotificationResponse


class UserNotificationsView(APIView):

    def get(self, request, user_id):
        grpc_client = UserNotificationsRPC(
            user_id=user_id,
            page=request.GET.get('page', 1),
            count=request.GET.get('count', 10),
            only_important=request.GET.get('only_important', False),
            read=request.GET.get('read', False),
        )

        response = grpc_client.perform_request()
        notifications = []

        for notif in list(response.notifications):

            model = UserNotificationResponse.model_validate(notif)
            model.target_validate()
            notifications.append(model.model_dump())

        return Response({'msg': 'ok', 'notifications': notifications})


class MassUserNotificationsView(APIView):
    def get(self, request, user_id):
        grpc_client = MassUserNotificationsRPC(
            user_id=user_id,
            page=request.GET.get('page', 1),
            count=request.GET.get('count', 10),
            not_type=request.GET.get('type', 1),
            only_important=request.GET.get('only_important', False),
            read=request.GET.get('read', False),
            action=request.GET.get('action', None),
        )

        response = grpc_client.perform_request()

        notifications = []
        for notif in list(response.notifications):

            model = UserMassNotificationResponse.model_validate(notif)
            model.target_validate()
            notifications.append(model.model_dump())
        return Response({'msg': 'ok', 'notifications': notifications})