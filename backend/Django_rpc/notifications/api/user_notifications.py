from rest_framework.views import APIView
from rest_framework.response import Response
from google.protobuf.json_format import MessageToDict

from notifications.modules.gRPC.user_notifications import UserNotificationsRPC


class UserNotificationsAPI(APIView):

    def get(self, request, user_id):

        grpc_client = UserNotificationsRPC(
            user_id=user_id,
            page=request.GET.get('page', 1),
            count=request.GET.get('count', 10),
            only_important=request.GET.get('only_important', False),
            read=request.GET.get('read', False),
        )

        response = grpc_client.perform_request()

        return Response({'msg': 'ok', 'notifications': list(response.notifications)})