from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.notification_action_call import NotificationActionRPC


class ActionEmit(APIView):

    def get(self, request):

        if (action := request.GET.get('action')) not in ['title_new_name', 'title_new_chapter', ]:
            return Response({"msg": "action forbidden"}, status=400)

        if action == 'title_new_name':
            grpc_client = NotificationActionRPC(
                notification_action='title_new_name',
                target_id=request.GET.get('target_id'),
                target_type=1,
                action=1,
                notification_type=0
            )
        elif action == 'title_new_chapter':
            grpc_client = NotificationActionRPC(
                notification_action='title_new_chapter',
                target_id=request.GET.get('target_id'),
                target_type=1,
                action=1,
                notification_type=0
            )

        response = grpc_client.perform_request()

        if not response.is_created:
            return Response({"msg": "rpc returns not created"}, status=400)
        return Response({"msg": "ok"}, status=200)