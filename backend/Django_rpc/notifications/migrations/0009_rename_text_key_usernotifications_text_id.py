# Generated by Django 3.2 on 2023-07-14 15:05

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('notifications', '0008_auto_20230713_1716'),
    ]

    operations = [
        migrations.RenameField(
            model_name='usernotifications',
            old_name='text_key',
            new_name='text_id',
        ),
    ]
