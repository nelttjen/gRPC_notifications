from rest_framework.views import APIView
from rest_framework.response import Response

import grpc
import rpc.notifications_pb2 as pb2
import rpc.notifications_pb2_grpc as pb2_grpc


class CreateNotificationsView(APIView):

    def post(self, request):
        example_users = [1, 3, 5, 6, 7, 8]
        example_text = 'Это текст уведомления'

        channel = grpc.insecure_channel('127.0.0.1:55000')
        stub = pb2_grpc.CreateNotificationsStub(channel=channel)

        payload = {
            "user_ids": example_users,
            "text": example_text
        }

        stub_response: pb2.NotificationCreateResponse = stub.CreateNotificationsForUsers(
            pb2.NotificationCreateRequest(**payload)
        )

        print(stub_response)
        return Response("ok")
