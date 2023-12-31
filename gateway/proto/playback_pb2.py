# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: playback.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from proto import common_pb2 as common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0eplayback.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x0c\x63ommon.proto\"5\n\x15\x43reatePlaylistRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06userId\x18\x02 \x01(\t\"\'\n\x11PlaylistIdRequest\x12\x12\n\nplaylistId\x18\x01 \x01(\t\";\n\x15RemovePlaylistRequest\x12\x12\n\nplaylistId\x18\x01 \x01(\t\x12\x0e\n\x06userId\x18\x02 \x01(\t\"R\n\x1a\x41\x64\x64TracksToPlaylistRequest\x12\x12\n\nplaylistId\x18\x01 \x01(\t\x12\x10\n\x08trackIds\x18\x02 \x03(\t\x12\x0e\n\x06userId\x18\x03 \x01(\t\"W\n\x1fRemoveTracksFromPlaylistRequest\x12\x12\n\nplaylistId\x18\x01 \x01(\t\x12\x10\n\x08trackIds\x18\x02 \x03(\t\x12\x0e\n\x06userId\x18\x03 \x01(\t\"[\n\x10PlaylistResponse\x12\x12\n\nplaylistId\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x06tracks\x18\x03 \x03(\x0b\x32\x15.common.TrackMetadata\"@\n\x11TrackPlayMetadata\x12\x0f\n\x07trackId\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x0b\n\x03url\x18\x03 \x01(\t\"P\n\x14PlayPlaylistResponse\x12\x14\n\x0cplaylistName\x18\x01 \x01(\t\x12\"\n\x06tracks\x18\x02 \x03(\x0b\x32\x12.TrackPlayMetadata2\xe1\x03\n\x0fPlaybackService\x12;\n\x0e\x43reatePlaylist\x12\x16.CreatePlaylistRequest\x1a\x11.PlaylistResponse\x12@\n\x0eRemovePlaylist\x12\x16.RemovePlaylistRequest\x1a\x16.google.protobuf.Empty\x12J\n\x13\x41\x64\x64TracksToPlaylist\x12\x1b.AddTracksToPlaylistRequest\x1a\x16.google.protobuf.Empty\x12T\n\x18RemoveTracksFromPlaylist\x12 .RemoveTracksFromPlaylistRequest\x1a\x16.google.protobuf.Empty\x12\x38\n\x0fGetPlaylistById\x12\x12.PlaylistIdRequest\x1a\x11.PlaylistResponse\x12\x39\n\x0cPlayPlaylist\x12\x12.PlaylistIdRequest\x1a\x15.PlayPlaylistResponse\x12\x38\n\x06Status\x12\x16.google.protobuf.Empty\x1a\x16.common.StatusResponseB\x18Z\x16playback_service/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'playback_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\026playback_service/proto'
  _globals['_CREATEPLAYLISTREQUEST']._serialized_start=61
  _globals['_CREATEPLAYLISTREQUEST']._serialized_end=114
  _globals['_PLAYLISTIDREQUEST']._serialized_start=116
  _globals['_PLAYLISTIDREQUEST']._serialized_end=155
  _globals['_REMOVEPLAYLISTREQUEST']._serialized_start=157
  _globals['_REMOVEPLAYLISTREQUEST']._serialized_end=216
  _globals['_ADDTRACKSTOPLAYLISTREQUEST']._serialized_start=218
  _globals['_ADDTRACKSTOPLAYLISTREQUEST']._serialized_end=300
  _globals['_REMOVETRACKSFROMPLAYLISTREQUEST']._serialized_start=302
  _globals['_REMOVETRACKSFROMPLAYLISTREQUEST']._serialized_end=389
  _globals['_PLAYLISTRESPONSE']._serialized_start=391
  _globals['_PLAYLISTRESPONSE']._serialized_end=482
  _globals['_TRACKPLAYMETADATA']._serialized_start=484
  _globals['_TRACKPLAYMETADATA']._serialized_end=548
  _globals['_PLAYPLAYLISTRESPONSE']._serialized_start=550
  _globals['_PLAYPLAYLISTRESPONSE']._serialized_end=630
  _globals['_PLAYBACKSERVICE']._serialized_start=633
  _globals['_PLAYBACKSERVICE']._serialized_end=1114
# @@protoc_insertion_point(module_scope)
