from django.db import models
from django.contrib.auth.models import User
from django.db.models.signals import post_save, pre_delete
from django.dispatch import receiver
from pymongo import MongoClient

from notifications.modules.typing.user_not_setts import UserNotificationConfig
from drpc.settings import MONGODB_AUTHPARAMS

targets = (
    (1, 'Тайтл'),
    (2, 'Глава'),
    (3, 'Комментарий'),
    (4, 'Пополнение'),
    (5, 'Специальное предложение'),
    (6, 'Бейдж')
)


# Create your models here.
class UserNotifications(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    action = models.CharField(verbose_name='Действие', max_length=100, null=True, default=None)

    confirmation = models.BooleanField(verbose_name='Нужно подтвердить?', default=False)
    read = models.BooleanField(verbose_name='Прочитано?', default=False)

    image = models.CharField(verbose_name='Картинка уведомления', max_length=1000, null=True, default=None)

    text = models.CharField(verbose_name='Текст уведомления', max_length=1000, null=True, default=None)
    text_key = models.ForeignKey(verbose_name='Ключ текста', to='notifications.NotificationText', on_delete=models.SET_NULL, null=True, default=None)

    link = models.CharField(verbose_name='Ссылка', max_length=1000, null=True, default=None)

    target_id = models.BigIntegerField(verbose_name='Таргет', null=True, default=None)
    target_type = models.PositiveSmallIntegerField(verbose_name='Тип таргета', choices=targets, null=True, default=None)

    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)
    important = models.BooleanField(verbose_name='Важный увед?', default=False)

    class Meta:
        db_table = 'user_notifications'
        verbose_name = 'Уведомление пользователя'
        verbose_name_plural = 'Уведомления пользователей'


class UserMassNotifications(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    notification = models.ForeignKey(verbose_name='Уведомление', to='notifications.MassNotifications', on_delete=models.CASCADE)

    read = models.BooleanField(verbose_name='Прочитано?', default=False)

    class Meta:
        db_table = 'user_mass_notifications'
        verbose_name = 'Рассылка пользователя'
        verbose_name_plural = 'Рассылки пользователю'


class NotificationText(models.Model):
    text = models.TextField(verbose_name='Текст уведомления')

    class Meta:
        db_table = 'notification_texts'
        verbose_name = 'Текст уведомления'
        verbose_name_plural = 'Тексты уведомлений'


class MassNotifications(models.Model):
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

    image = models.CharField(verbose_name='Картинка уведомления', max_length=1000, null=True, default=None)
    text = models.CharField(verbose_name='Текст уведомления', max_length=1000)
    link = models.CharField(verbose_name='Ссылка', max_length=1000, null=True, default=None)

    target_id = models.BigIntegerField(verbose_name='Таргет', null=True, default=None)
    target_type = models.PositiveSmallIntegerField(verbose_name='Тип таргета', choices=targets, null=True, default=None)

    important = models.BooleanField(verbose_name='Важный увед?', default=False)
    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)

    class Meta:
        db_table = 'mass_notifications'
        verbose_name = 'Рассылка'
        verbose_name_plural = 'Рассылки'


class CommentMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    reply_to = models.ForeignKey(verbose_name='Комментарий', to='notifications.CommentMock', on_delete=models.CASCADE,
                                 null=True, blank=True, default=None)
    text = models.CharField(verbose_name='Текст комментария', max_length=1000)

    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)

    class Meta:
        db_table = 'comments'
        verbose_name = 'Комментарий'
        verbose_name_plural = 'Комментарии'


class TitleMock(models.Model):
    name = models.CharField(verbose_name='Название', max_length=128)

    class Meta:
        db_table = 'titles'
        verbose_name = 'Тайтл'
        verbose_name_plural = 'Тайтлы'


class ChapterMock(models.Model):
    index = models.IntegerField(verbose_name='Номер главы')
    title = models.ForeignKey(verbose_name='Тайтл', to='notifications.TitleMock', on_delete=models.CASCADE)
    is_paid = models.BooleanField(verbose_name='Платный?', default=False)

    class Meta:
        db_table = 'chapters'
        verbose_name = 'Глава'
        verbose_name_plural = 'Главы'


class BillingMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    sum = models.DecimalField(verbose_name='Сумма', max_digits=10, decimal_places=2)

    class Meta:
        db_table = 'billings'
        verbose_name = 'Пополнение'
        verbose_name_plural = 'Пополнения'


class SpecialOfferMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    need_sum = models.DecimalField(verbose_name='Нужная сумма', max_digits=10, decimal_places=2)
    reward = models.DecimalField(verbose_name='Награда', max_digits=10, decimal_places=2)

    class Meta:
        db_table = 'special_offers'
        verbose_name = 'Специльное предложение'
        verbose_name_plural = 'Специальные предложения'


class BadgeMock(models.Model):
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    name = models.CharField(verbose_name='Название', max_length=128)

    class Meta:
        db_table = 'badges'
        verbose_name = 'Бейдж'
        verbose_name_plural = 'Бейджи'


class TitleBookmark(models.Model):
    categories = (
        (1, 'Категория 1'),
        (2, 'Категория 2'),
        (3, 'Категория 3'),
    )
    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    title = models.ForeignKey(verbose_name='Тайтл', to='notifications.TitleMock', on_delete=models.CASCADE)

    category = models.IntegerField(verbose_name='Вкладка', choices=categories)

    class Meta:
        db_table = 'title_bookmarks'
        verbose_name = 'Вкладка'
        verbose_name_plural = 'Вкладки'


@receiver(post_save, sender=User)
def create_user_notifications(sender, instance, created, **kwargs):
    if created:
        client = MongoClient(**MONGODB_AUTHPARAMS)
        db = client.get_database('admin')
        collection = db.get_collection("user_notification_settings")
        collection.insert_one(UserNotificationConfig(user_id=instance.id).model_dump())

        client.close()


@receiver(pre_delete, sender=User)
def delete_user_notifications(sender, instance, **kwargs):
    client = MongoClient(**MONGODB_AUTHPARAMS)
    db = client.get_database('admin')
    collection = db.get_collection("user_notification_settings")
    collection.delete_one({"user_id": instance.id})

    client.close()