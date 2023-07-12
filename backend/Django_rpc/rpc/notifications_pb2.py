# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: notifications.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import rpc.struct_pb2 as struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13notifications.proto\x12\x03\x61pi\x1a\x0cstruct.proto\"\xf2\x01\n\x19NotificationCreateRequest\x12\x0e\n\x06\x61\x63tion\x18\x01 \x01(\t\x12\x16\n\ttarget_id\x18\x02 \x01(\x03H\x00\x88\x01\x01\x12\x18\n\x0btarget_type\x18\x03 \x01(\x11H\x01\x88\x01\x01\x12\x11\n\timportant\x18\x04 \x01(\x08\x12\x0c\n\x04type\x18\x05 \x01(\x05\x12\x11\n\x04link\x18\x07 \x01(\tH\x02\x88\x01\x01\x12\x12\n\x05image\x18\x08 \x01(\tH\x03\x88\x01\x01\x12\x11\n\x04text\x18\t \x01(\tH\x04\x88\x01\x01\x42\x0c\n\n_target_idB\x0e\n\x0c_target_typeB\x07\n\x05_linkB\x08\n\x06_imageB\x07\n\x05_text\"\xec\x01\n\x1fNotificationCreateManualRequest\x12\x10\n\x08user_ids\x18\x01 \x03(\x05\x12\x16\n\ttarget_id\x18\x02 \x01(\x03H\x00\x88\x01\x01\x12\x18\n\x0btarget_type\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x11\n\x04text\x18\x04 \x01(\tH\x02\x88\x01\x01\x12\x11\n\timportant\x18\x05 \x01(\x08\x12\x11\n\x04link\x18\x06 \x01(\tH\x03\x88\x01\x01\x12\x12\n\x05image\x18\x07 \x01(\tH\x04\x88\x01\x01\x42\x0c\n\n_target_idB\x0e\n\x0c_target_typeB\x07\n\x05_textB\x07\n\x05_linkB\x08\n\x06_image\"0\n\x1aNotificationCreateResponse\x12\x12\n\nis_created\x18\x01 \x01(\x08\"\xc8\x01\n\x0cNotification\x12\n\n\x02id\x18\x01 \x01(\x05\x12\'\n\x06target\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0btarget_type\x18\x03 \x01(\t\x12\r\n\x05image\x18\x04 \x01(\t\x12\x0c\n\x04link\x18\x05 \x01(\t\x12\x0c\n\x04text\x18\x06 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x07 \x01(\t\x12\x11\n\timportant\x18\x08 \x01(\x08\x12\x14\n\x0c\x63onfirmation\x18\t \x01(\x08\x12\x0c\n\x04read\x18\n \x01(\x08\"\xa2\x01\n\x18UserNotificationsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\x05\x12\x0c\n\x04type\x18\x02 \x01(\x05\x12\x1b\n\x0eonly_important\x18\x03 \x01(\x08H\x00\x88\x01\x01\x12\x11\n\x04read\x18\x04 \x01(\x08H\x01\x88\x01\x01\x12\x0c\n\x04page\x18\x05 \x01(\x05\x12\r\n\x05\x63ount\x18\x06 \x01(\x05\x42\x11\n\x0f_only_importantB\x07\n\x05_read\"E\n\x19UserNotificationsResponse\x12(\n\rnotifications\x18\x01 \x03(\x0b\x32\x11.api.Notification\"S\n\x14UserMassNotification\x12-\n\x0cnotification\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0c\n\x04read\x18\x02 \x01(\x08\"\xc5\x01\n\x1bUserMassNotificationRequest\x12\x0c\n\x04type\x18\x01 \x01(\x05\x12\x0f\n\x07user_id\x18\x02 \x01(\x05\x12\x13\n\x06\x61\x63tion\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x1b\n\x0eonly_important\x18\x04 \x01(\x08H\x01\x88\x01\x01\x12\x11\n\x04read\x18\x05 \x01(\x08H\x02\x88\x01\x01\x12\x0c\n\x04page\x18\x06 \x01(\x05\x12\r\n\x05\x63ount\x18\x07 \x01(\x05\x42\t\n\x07_actionB\x11\n\x0f_only_importantB\x07\n\x05_read\"P\n\x1cUserMassNotificationResponse\x12\x30\n\rnotifications\x18\x01 \x03(\x0b\x32\x19.api.UserMassNotification\"F\n\x19NotificationManageRequest\x12\x18\n\x10notification_ids\x18\x01 \x03(\x05\x12\x0f\n\x07user_id\x18\x02 \x01(\x05\"-\n\x1aNotificationManageResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32\xa7\x04\n\x13\x43reateNotifications\x12\\\n\x19\x43reateNotificationsAction\x12\x1e.api.NotificationCreateRequest\x1a\x1f.api.NotificationCreateResponse\x12O\n\x06\x43reate\x12$.api.NotificationCreateManualRequest\x1a\x1f.api.NotificationCreateResponse\x12Q\n\x10GetNotifications\x12\x1d.api.UserNotificationsRequest\x1a\x1e.api.UserNotificationsResponse\x12Z\n\x13GetMassNotification\x12 .api.UserMassNotificationRequest\x1a!.api.UserMassNotificationResponse\x12V\n\x13\x44\x65leteNotifications\x12\x1e.api.NotificationManageRequest\x1a\x1f.api.NotificationManageResponse\x12Z\n\x17MarkAsReadNotifications\x12\x1e.api.NotificationManageRequest\x1a\x1f.api.NotificationManageResponseB\x07Z\x05.;apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'notifications_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\005.;api'
  _NOTIFICATIONCREATEREQUEST._serialized_start=43
  _NOTIFICATIONCREATEREQUEST._serialized_end=285
  _NOTIFICATIONCREATEMANUALREQUEST._serialized_start=288
  _NOTIFICATIONCREATEMANUALREQUEST._serialized_end=524
  _NOTIFICATIONCREATERESPONSE._serialized_start=526
  _NOTIFICATIONCREATERESPONSE._serialized_end=574
  _NOTIFICATION._serialized_start=577
  _NOTIFICATION._serialized_end=777
  _USERNOTIFICATIONSREQUEST._serialized_start=780
  _USERNOTIFICATIONSREQUEST._serialized_end=942
  _USERNOTIFICATIONSRESPONSE._serialized_start=944
  _USERNOTIFICATIONSRESPONSE._serialized_end=1013
  _USERMASSNOTIFICATION._serialized_start=1015
  _USERMASSNOTIFICATION._serialized_end=1098
  _USERMASSNOTIFICATIONREQUEST._serialized_start=1101
  _USERMASSNOTIFICATIONREQUEST._serialized_end=1298
  _USERMASSNOTIFICATIONRESPONSE._serialized_start=1300
  _USERMASSNOTIFICATIONRESPONSE._serialized_end=1380
  _NOTIFICATIONMANAGEREQUEST._serialized_start=1382
  _NOTIFICATIONMANAGEREQUEST._serialized_end=1452
  _NOTIFICATIONMANAGERESPONSE._serialized_start=1454
  _NOTIFICATIONMANAGERESPONSE._serialized_end=1499
  _CREATENOTIFICATIONS._serialized_start=1502
  _CREATENOTIFICATIONS._serialized_end=2053
# @@protoc_insertion_point(module_scope)
