from django.urls import path

from .api.create_notifications import CreateNotificationsView


urlpatterns = [
    path("create/", CreateNotificationsView.as_view(), name='api-notifications-create'),
]