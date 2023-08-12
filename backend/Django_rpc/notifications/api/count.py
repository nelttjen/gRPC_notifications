from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.count_notifications import CountNotificationsRPC
from notifications.modules.typing.rpc.response import UserCountNotificationResponse


class UserCountNotificationsView(APIView):

    def get(self, request, user_id):
        request = CountNotificationsRPC(user_id=user_id)

        response = request.perform_request()

        return Response({"msg": "ok", "response": {"count": response.count, "has_important": response.has_important}})