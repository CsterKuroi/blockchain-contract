# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: response-data.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='response-data.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x13response-data.proto\x12\x06protos\"5\n\x0cResponseData\x12\n\n\x02ok\x18\x01 \x01(\x08\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\tB)\n\x14\x63om.uniledger.protosB\x11ProtoResponseDatab\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_RESPONSEDATA = _descriptor.Descriptor(
  name='ResponseData',
  full_name='protos.ResponseData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ok', full_name='protos.ResponseData.ok', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='msg', full_name='protos.ResponseData.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='data', full_name='protos.ResponseData.data', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=84,
)

DESCRIPTOR.message_types_by_name['ResponseData'] = _RESPONSEDATA

ResponseData = _reflection.GeneratedProtocolMessageType('ResponseData', (_message.Message,), dict(
  DESCRIPTOR = _RESPONSEDATA,
  __module__ = 'response_data_pb2'
  # @@protoc_insertion_point(class_scope:protos.ResponseData)
  ))
_sym_db.RegisterMessage(ResponseData)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\024com.uniledger.protosB\021ProtoResponseData'))
# @@protoc_insertion_point(module_scope)