from django.db import models
from django.contrib.auth.models import User
from django.db.models.signals import post_save
from django.dispatch import receiver
from pymongo import MongoClient

from notifications.modules.typing.user_not_setts import UserNotificationConfig
from drpc.settings import MONGODB_AUTHPARAMS


class NotificationBase:
    targets = (
        (1, 'Тайтл'),
        (2, 'Глава'),
        (3, 'Комментарий'),
        (4, 'Пополнение'),
        (5, 'Специальное предложение'),
        (6, 'Бейдж')
    )

    image = models.CharField(verbose_name='Картинка уведомления', max_length=1000, null=True, default=None)
    text = models.CharField(verbose_name='Текст уведомления', max_length=1000)
    link = models.CharField(verbose_name='Ссылка', max_length=1000, null=True, default=None)

    target_id = models.BigIntegerField(verbose_name='Таргет', null=True, default=None)
    target_type = models.PositiveSmallIntegerField(verbose_name='Тип таргета', choices=targets, null=True, default=None)

    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)
    important = models.BooleanField(verbose_name='Важный увед?', default=False)


# Create your models here.
class UserNotifications(models.Model, NotificationBase):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)

    confirmation = models.BooleanField(verbose_name='Нужно подтвердить?', default=False)
    read = models.BooleanField(verbose_name='Прочитано?', default=False)

    class Meta:
        db_table = 'user_notifications'
        verbose_name = 'Увед'
        verbose_name_plural = 'Уведы'


class UserMassNotifications(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    notification = models.ForeignKey(verbose_name='Уведомление', to='notifications.MassNotifications', on_delete=models.CASCADE)

    read = models.BooleanField(verbose_name='Прочитано?', default=False)


class MassNotifications(models.Model, NotificationBase):
    types = (
        (1, 'Обновление'),
        (2, 'Социальное'),
        (3, 'Важное'),
    )
    actions = (
        (1, 'Глава залита'),
        (2, 'Глава стала бесплатной'),
        (3, 'Уведомление сайта'),
    )

    type = models.IntegerField(verbose_name='Тип списка', choices=types)
    action = models.IntegerField(verbose_name='Действие', choices=actions, default=1)


class CommentMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    reply_to = models.ForeignKey(verbose_name='Комментарий', to='notifications.CommentMock', on_delete=models.CASCADE,
                                 null=True, blank=True, default=None)
    text = models.CharField(verbose_name='Текст комментария', max_length=1000)

    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)


class TitleMock(models.Model):
    name = models.CharField(verbose_name='Название', max_length=128)


class ChapterMock(models.Model):
    index = models.IntegerField(verbose_name='Номер главы')
    title = models.ForeignKey(verbose_name='Тайтл', to='notifications.TitleMock', on_delete=models.CASCADE)


class BillingMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    sum = models.DecimalField(verbose_name='Сумма', max_digits=10, decimal_places=2)


class SpecialOfferMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    need_sum = models.DecimalField(verbose_name='Нужная сумма', max_digits=10, decimal_places=2)
    reward = models.DecimalField(verbose_name='Награда', max_digits=10, decimal_places=2)


class BadgeMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    name = models.CharField(verbose_name='Название', max_length=128)


@receiver(post_save, sender=User)
def create_user_notifications(sender, instance, created, **kwargs):
    if created:
        client = MongoClient(**MONGODB_AUTHPARAMS)
        db = client.get_database('admin')
        collection = db.get_collection("user_notification_settings")
        collection.insert_one(UserNotificationConfig(user_id=instance.id).model_dump())

        client.close()
