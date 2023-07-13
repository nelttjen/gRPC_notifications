from django.urls import path

from .api.test import TestView
from .api.emit_action import ActionEmit

urlpatterns = [
    path("emit/", ActionEmit.as_view(), name='api-notifications-emit'),
]