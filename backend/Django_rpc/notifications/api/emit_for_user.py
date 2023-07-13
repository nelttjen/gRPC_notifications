from rest_framework.views import APIView
from rest_framework.response import Response

from notifications.modules.gRPC.notification_for_users import NotificationForUsersRPC


class EmitForUserView(APIView):

    def _process_request(self, request, source):
        data = getattr(request, source)

        actions = []

        if action := data.get('action') not in actions:
            return Response(status=400, data={'msg': 'action forbidden'})

        if not isinstance((ids := data.get('user_ids')), list):
            return Response(status=400, data={'msg': 'user_ids must be a list'})

        required = ['sets_key', 'text']


        if not all([data.get(key) for key in ['sets_key', 'text']]):
            return Response(status=400, data={'msg': f'required keys {", ".join(required)} not found'})

        grpc_client = NotificationForUsersRPC(
            user_ids=ids,
            sets_key=data.get('sets_key'),
            text=data.get('text'),
            text_as_model=data.get('text_as_model', True),
            important=data.get('important', False),
            confirmation=data.get('confirmation', False),
        )

        for key, mapper in [('target_id', int), ('target_type', int), ('link', str), ('image', str)]:
            if val := data.get(key):
                setattr(grpc_client, key, mapper(val))

        response = grpc_client.perform_request()

        if not response.is_created:
            return Response({"msg": "rpc returns not created"}, status=400)
        return Response({"msg": "ok"}, status=200)

    def get(self, request):
        return self._process_request(request, 'GET')

    def post(self, request):
        return self._process_request(request, 'data')
