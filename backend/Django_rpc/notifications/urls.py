from django.urls import path

from .api.create_notifications import CreateNotificationsView
from .api.test import TestView

urlpatterns = [
    path("create/", CreateNotificationsView.as_view(), name='api-notifications-create'),
    path('test/', TestView.as_view(), name='api-notifications-test'),
]