from django.urls import path

from .api.test import TestView
from .api.emit_action import ActionEmit
from .api.emit_for_user import EmitForUserView
from .api.user_notifications import UserNotificationsAPI

urlpatterns = [
    path("emit/", ActionEmit.as_view(), name='api-notifications-emit'),
    path("emit/user/", EmitForUserView.as_view(), name='api-notifications-emit-users'),
    path('<int:user_id>/list/', UserNotificationsAPI.as_view(), name='api-notifications-get-user-notifications'),
]