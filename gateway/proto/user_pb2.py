# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x1a\x1bgoogle/protobuf/empty.proto\" \n\x0eStatusResponse\x12\x0e\n\x06status\x18\x01 \x01(\t\"1\n\x0bUserRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x1b\n\rUserIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\"#\n\x0fUsernameRequest\x12\x10\n\x08username\x18\x01 \x01(\t\",\n\x0cUserResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\"\x14\n\x03JWT\x12\r\n\x05token\x18\x01 \x01(\t\"0\n\x10UserListResponse\x12\x1c\n\x05users\x18\x01 \x03(\x0b\x32\r.UserResponse2\xfa\x02\n\x0bUserService\x12\x30\n\x08Register\x12\x0c.UserRequest\x1a\x16.google.protobuf.Empty\x12\x1b\n\x05Login\x12\x0c.UserRequest\x1a\x04.JWT\x12\x1f\n\x08Validate\x12\x04.JWT\x1a\r.UserResponse\x12)\n\x08\x46indById\x12\x0e.UserIdRequest\x1a\r.UserResponse\x12\x31\n\x0e\x46indByUsername\x12\x10.UsernameRequest\x1a\r.UserResponse\x12\x34\n\x07\x46indAll\x12\x16.google.protobuf.Empty\x1a\x11.UserListResponse\x12\x34\n\nDeleteById\x12\x0e.UserIdRequest\x1a\x16.google.protobuf.Empty\x12\x31\n\x06Status\x12\x16.google.protobuf.Empty\x1a\x0f.StatusResponseB\x14Z\x12user_service/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\022user_service/proto'
  _globals['_STATUSRESPONSE']._serialized_start=43
  _globals['_STATUSRESPONSE']._serialized_end=75
  _globals['_USERREQUEST']._serialized_start=77
  _globals['_USERREQUEST']._serialized_end=126
  _globals['_USERIDREQUEST']._serialized_start=128
  _globals['_USERIDREQUEST']._serialized_end=155
  _globals['_USERNAMEREQUEST']._serialized_start=157
  _globals['_USERNAMEREQUEST']._serialized_end=192
  _globals['_USERRESPONSE']._serialized_start=194
  _globals['_USERRESPONSE']._serialized_end=238
  _globals['_JWT']._serialized_start=240
  _globals['_JWT']._serialized_end=260
  _globals['_USERLISTRESPONSE']._serialized_start=262
  _globals['_USERLISTRESPONSE']._serialized_end=310
  _globals['_USERSERVICE']._serialized_start=313
  _globals['_USERSERVICE']._serialized_end=691
# @@protoc_insertion_point(module_scope)
