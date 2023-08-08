import pydantic
import datetime

from typing import Union, Optional


class NotificationText(pydantic.BaseModel):
    text: str = pydantic.Field(alias='Text')


class Title(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    name: str = pydantic.Field(alias='Name')


class Chapter(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    title: 'Title' = pydantic.Field(alias='Title')
    is_paid: bool = pydantic.Field(alias='IsPaid')
    index: int = pydantic.Field(alias='Index')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'


class Comment(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    reply_to: Optional['Comment'] = pydantic.Field(alias='ReplyTo', default=None)
    text: str = pydantic.Field(alias='Text')
    date: datetime.datetime = pydantic.Field(alias='Date')


class Billing(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    sum: float = pydantic.Field(alias='Sum')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'


class SpecialOffer(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    need_sum: float = pydantic.Field(alias='NeedSum')
    reward: float = pydantic.Field(alias='Reward')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'


class Badge(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    name: str = pydantic.Field(alias='Name')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'


class UserNotificationResponse(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    target: Optional[dict] = pydantic.Field(alias='Target')
    target_type: int = pydantic.Field(alias='TargetType')

    text: Optional[str] = pydantic.Field(alias='Text')
    notification_text: Optional['NotificationText'] = pydantic.Field(alias='NotificationText')

    link: Optional[str] = pydantic.Field(alias='Link')
    image: Optional[str] = pydantic.Field(alias='Image')
    action: Optional[str] = pydantic.Field(alias='Action')

    read: bool = pydantic.Field(alias='Read')
    important: bool = pydantic.Field(alias='Important')
    confirmation: bool = pydantic.Field(alias='Confirmation')

    date: datetime.datetime = pydantic.Field(alias='Date')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'

    def target_validate(self):
        models = {
            1: Title,
            2: Chapter,
            3: Comment,
            4: Billing,
            5: SpecialOffer,
            6: Badge,
        }
        if self.target_type not in models.keys():
            return

        pd_model = models.get(self.target_type)
        validated = pd_model.model_validate(self.target)
        self.target = validated.model_dump()


class MassNotificationResponse(pydantic.BaseModel):
    target: Optional[dict] = pydantic.Field(alias='Target')
    target_type: int = pydantic.Field(alias='TargetType')
    text: str = pydantic.Field(alias='Text')

    type: int = pydantic.Field(alias='Type')
    important: bool = pydantic.Field(alias='Important')

    link: Optional[str] = pydantic.Field(alias='Link')
    image: Optional[str] = pydantic.Field(alias='Image')
    action: Optional[str] = pydantic.Field(alias='Action')

    date: datetime.datetime = pydantic.Field(alias='Date')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'


class UserMassNotificationResponse(pydantic.BaseModel):
    id: int = pydantic.Field(alias='ID')
    read: bool = pydantic.Field(alias='Read')

    notification: 'MassNotificationResponse' = pydantic.Field(alias='Notification')

    class Config(pydantic.ConfigDict):
        extra = 'ignore'

    def target_validate(self):
        models = {
            1: Title,
            2: Chapter,
            3: Comment,
            4: Billing,
            5: SpecialOffer,
            6: Badge,
        }
        if self.notification.target_type not in models.keys():
            return

        pd_model = models.get(self.notification.target_type)
        validated = pd_model.model_validate(self.notification.target)
        self.notification.target = validated.model_dump()