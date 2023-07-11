# Generated by Django 3.2 on 2023-07-09 15:46

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion
import notifications.models


class Migration(migrations.Migration):

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ('notifications', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='MassNotifications',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('type', models.IntegerField(choices=[(0, 'Обновление'), (1, 'Социальное'), (2, 'Важное')], verbose_name='Тип списка')),
                ('action', models.IntegerField(choices=[(0, 'Действие 1'), (1, 'Действие 2'), (2, 'Действие 3')], default=1, verbose_name='Действие')),
            ],
            bases=(models.Model, notifications.models.NotificationBase),
        ),
        migrations.CreateModel(
            name='TitleMock',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=128, verbose_name='Название')),
            ],
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='action',
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='date',
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='image',
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='link',
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='text',
        ),
        migrations.RemoveField(
            model_name='usernotifications',
            name='type',
        ),
        migrations.AddField(
            model_name='usernotifications',
            name='confirmation',
            field=models.BooleanField(default=False, verbose_name='Нужно подтвердить?'),
        ),
        migrations.CreateModel(
            name='UserMassNotifications',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('status', models.BooleanField(default=False, verbose_name='Прочитано?')),
                ('date', models.DateTimeField(auto_now_add=True, verbose_name='Дата')),
                ('notification', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='notifications.massnotifications', verbose_name='Уведомление')),
                ('user', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='Пользователь')),
            ],
        ),
        migrations.CreateModel(
            name='CommentMock',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('text', models.CharField(max_length=1000, verbose_name='Текст комментария')),
                ('date', models.DateTimeField(auto_now_add=True, verbose_name='Дата')),
                ('user', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='Пользователь')),
            ],
        ),
    ]
