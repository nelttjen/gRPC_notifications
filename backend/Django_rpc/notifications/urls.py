from django.urls import path

from .api.emit_event import ActionEmitView
from .api.emit_for_user import EmitForUserView
from .api.user_notifications import UserNotificationsView, MassUserNotificationsView
from .api.actions import UserNotificationActionView

urlpatterns = [
    path("emit/", ActionEmitView.as_view(), name='api-notifications-emit'),
    path("emit/user/", EmitForUserView.as_view(), name='api-notifications-emit-users'),
    path('<int:user_id>/list/', UserNotificationsView.as_view(), name='api-notifications-get-user-notifications'),
    path('<int:user_id>/mass/list/', MassUserNotificationsView.as_view(), name='api-notifications-get-mass-user-notifications'),
    path('<int:user_id>/action/', UserNotificationActionView.as_view(), name='api-notifications-user-notification-actions'),
]