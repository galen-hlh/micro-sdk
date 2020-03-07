# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: idProduce/idProduce.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='idProduce/idProduce.proto',
  package='idProduce',
  syntax='proto3',
  serialized_options=_b('\312\002\024MicroProto\\IdProduce'),
  serialized_pb=_b('\n\x19idProduce/idProduce.proto\x12\tidProduce\"\x1f\n\x10IdProduceRequest\x12\x0b\n\x03len\x18\x01 \x01(\r\".\n\x11IdProduceResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\r\x12\x0b\n\x03ids\x18\x02 \x03(\x04\x32Z\n\x06Helper\x12P\n\x0fGetDistributeId\x12\x1b.idProduce.IdProduceRequest\x1a\x1c.idProduce.IdProduceResponse(\x01\x30\x01\x42\x17\xca\x02\x14MicroProto\\IdProduceb\x06proto3')
)




_IDPRODUCEREQUEST = _descriptor.Descriptor(
  name='IdProduceRequest',
  full_name='idProduce.IdProduceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='len', full_name='idProduce.IdProduceRequest.len', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=71,
)


_IDPRODUCERESPONSE = _descriptor.Descriptor(
  name='IdProduceResponse',
  full_name='idProduce.IdProduceResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='idProduce.IdProduceResponse.code', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ids', full_name='idProduce.IdProduceResponse.ids', index=1,
      number=2, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=73,
  serialized_end=119,
)

DESCRIPTOR.message_types_by_name['IdProduceRequest'] = _IDPRODUCEREQUEST
DESCRIPTOR.message_types_by_name['IdProduceResponse'] = _IDPRODUCERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IdProduceRequest = _reflection.GeneratedProtocolMessageType('IdProduceRequest', (_message.Message,), {
  'DESCRIPTOR' : _IDPRODUCEREQUEST,
  '__module__' : 'idProduce.idProduce_pb2'
  # @@protoc_insertion_point(class_scope:idProduce.IdProduceRequest)
  })
_sym_db.RegisterMessage(IdProduceRequest)

IdProduceResponse = _reflection.GeneratedProtocolMessageType('IdProduceResponse', (_message.Message,), {
  'DESCRIPTOR' : _IDPRODUCERESPONSE,
  '__module__' : 'idProduce.idProduce_pb2'
  # @@protoc_insertion_point(class_scope:idProduce.IdProduceResponse)
  })
_sym_db.RegisterMessage(IdProduceResponse)


DESCRIPTOR._options = None

_HELPER = _descriptor.ServiceDescriptor(
  name='Helper',
  full_name='idProduce.Helper',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=121,
  serialized_end=211,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetDistributeId',
    full_name='idProduce.Helper.GetDistributeId',
    index=0,
    containing_service=None,
    input_type=_IDPRODUCEREQUEST,
    output_type=_IDPRODUCERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_HELPER)

DESCRIPTOR.services_by_name['Helper'] = _HELPER

# @@protoc_insertion_point(module_scope)