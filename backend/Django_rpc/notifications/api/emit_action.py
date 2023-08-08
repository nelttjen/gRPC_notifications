from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.notification_action_call import NotificationActionRPC


class ActionEmitView(APIView):

    def _process_request(self, request, data_source):
        data = getattr(request, data_source)

        if (action := data.get('action')) not in ['title_new_name', 'title_new_chapter', 'site_notification', 'title_chapter_free']:
            return Response({"msg": "action forbidden"}, status=400)

        if action == 'title_new_name':
            grpc_client = NotificationActionRPC(
                action='title_new_name',
                target_id=data.get('target_id'),
                target_type=1,
                not_type=1,
                important=bool(int(data.get('important', 0))),
                text=data.get('prev_name', 'Предыдущее название'),
            )
        elif action in ('title_new_chapter', 'title_chapter_free'):
            grpc_client = NotificationActionRPC(
                action=action,
                target_id=data.get('target_id'),
                target_type=2,
                not_type=1,
                important=bool(int(data.get('important', 0))),
            )
        elif action == 'site_notification':
            grpc_client = NotificationActionRPC(
                action='site_notification',
                target_id=None,
                target_type=None,
                not_type=3,
                text=data.get('text', 'Site notification'),
                image=data.get('image', ''),
                link=data.get('link', ''),
                important=bool(int(data.get('important', 0))),
            )

        response = grpc_client.perform_request()

        if not response.is_created:
            return Response({"msg": "rpc returns not created"}, status=400)
        return Response({"msg": "ok"}, status=200)

    def get(self, request):
        return self._process_request(request, data_source='GET')

    def post(self, request):
        return self._process_request(request, data_source='data')