from rest_framework.views import APIView
from rest_framework.response import Response
import pymongo


class TestView(APIView):

    def get(self, request):
        return Response('ok')