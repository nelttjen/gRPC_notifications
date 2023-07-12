from django.contrib import admin

from .models import *

# Register your models here.
admin.site.register(TitleMock)
admin.site.register(ChapterMock)
admin.site.register(CommentMock)
admin.site.register(BadgeMock)
admin.site.register(BillingMock)
admin.site.register(SpecialOfferMock)
admin.site.register(TitleBookmark)

admin.site.register(UserNotifications)
admin.site.register(UserMassNotifications)
admin.site.register(MassNotifications)