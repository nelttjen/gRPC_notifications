# Generated by Django 3.2 on 2023-07-13 13:46

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('notifications', '0006_auto_20230712_2312'),
    ]

    operations = [
        migrations.AddField(
            model_name='chaptermock',
            name='is_paid',
            field=models.BooleanField(default=False, verbose_name='Платный?'),
        ),
    ]
