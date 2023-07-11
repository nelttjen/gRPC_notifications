from rest_framework.views import APIView
from rest_framework.response import Response
import pymongo


class TestView(APIView):

    def get(self, request):
        client = pymongo.MongoClient(
            host='127.0.0.1',
            port=27017,
            username='admin',
            password='adminpass123',
            authSource='development'
        )
        db = client['admin']
        collection = db["user_notification_settings"]
        collection.insert_one({
            "user_id": 1,
            "notifications": True
        })

        print(collection.find({}))

        return Response('ok')