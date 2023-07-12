from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.notification_action_call import NotificationActionRPC


class ActionEmit(APIView):

    def _process_request(self, request, data_source):
        data = getattr(request, data_source)

        if (action := data.get('action')) not in ['title_new_name', 'title_new_chapter', 'site_notification']:
            return Response({"msg": "action forbidden"}, status=400)

        if action == 'title_new_name':
            grpc_client = NotificationActionRPC(
                action='title_new_name',
                target_id=data.get('target_id'),
                target_type=1,
                not_type=0,
            )
        elif action == 'title_new_chapter':
            grpc_client = NotificationActionRPC(
                action='title_new_chapter',
                target_id=data.get('target_id'),
                target_type=2,
                not_type=0
            )
        elif action == 'site_notification':
            grpc_client = NotificationActionRPC(
                action='site_notification',
                target_id=None,
                target_type=None,
                not_type=2,
                text=data.get('text', 'Site notification'),
                image=data.get('image', None),
                link=data.get('link', None),
            )

        response = grpc_client.perform_request()

        if not response.is_created:
            return Response({"msg": "rpc returns not created"}, status=400)
        return Response({"msg": "ok"}, status=200)

    def get(self, request):
        return self._process_request(request, data_source='GET')

    def post(self, request):
        return self._process_request(request, data_source='data')