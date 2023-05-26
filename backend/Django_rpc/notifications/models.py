from django.db import models
from django.contrib.auth.models import User


# Create your models here.
class UserNotifications(models.Model):
    types = (
        (0, 'Обновление'),
        (1, 'Социальное'),
        (2, 'Важное'),
    )
    action_choices = (
        (1, 'Ссылка'),
        (2, 'Подтверждение'),
        (3, 'Комментарий'),
    )

    user = models.ForeignKey(verbose_name='Пользователь', to=User, on_delete=models.CASCADE)
    image = models.CharField(verbose_name='Картинка уведомления', max_length=1000, null=True, default=None)
    text = models.CharField(verbose_name='Текст уведомления', max_length=1000)
    link = models.CharField(verbose_name='Ссылка', max_length=1000, null=True, default=None)
    type = models.IntegerField(verbose_name='Тип списка', choices=types)
    action = models.IntegerField(verbose_name='Действие', choices=action_choices, default=1)
    status = models.BooleanField(verbose_name='Прочитано?', default=False)
    date = models.DateTimeField(verbose_name='Дата', auto_now_add=True)

    class Meta:
        db_table = 'user_notifications'
        verbose_name = 'Увед'
        verbose_name_plural = 'Уведы'
