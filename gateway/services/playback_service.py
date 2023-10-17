import functools
import grpc
from fastapi import HTTPException
from google.protobuf import empty_pb2

from proto import playback_pb2, playback_pb2_grpc
from models.playback_models import (CreatePlaylistRequest,
                                    AddTracksToPlaylistRequest, RemoveTracksFromPlaylistRequest, RemovePlaylistRequest)
from config.config import config
from utils.grpc_utils import get_grpc_channel, handle_grpc_error


def get_playback_stub():
    channel = get_grpc_channel(config.PLAYBACK_SERVICE_HOST, config.PLAYBACK_SERVICE_PORT)
    return playback_pb2_grpc.PlaybackServiceStub(channel)


class PlaybackService:
    @staticmethod
    @functools.cache
    def get_instance():
        return PlaybackService()

    def __init__(self):
        self.stub = get_playback_stub()

    @handle_grpc_error
    def create_playlist(self, request: CreatePlaylistRequest):
        stub = self.stub
        grpc_request = playback_pb2.CreatePlaylistRequest(name=request.name, userId=request.userId)
        try:
            response, call_status = stub.CreatePlaylist.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"playlistId": response.playlistId}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def remove_playlist(self, request: RemovePlaylistRequest):
        stub = self.stub
        grpc_request = playback_pb2.RemovePlaylistRequest(playlistId=request.playlistId, userId=request.userId)
        try:
            stub.RemovePlaylist.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def add_tracks_to_playlist(self, request: AddTracksToPlaylistRequest):
        stub = self.stub
        grpc_request = playback_pb2.AddTracksToPlaylistRequest(playlistId=request.playlistId, trackIds=request.trackIds,
                                                               userId=request.userId)
        try:
            stub.AddTracksToPlaylist.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def remove_tracks_from_playlist(self, request: RemoveTracksFromPlaylistRequest):
        stub = self.stub
        grpc_request = playback_pb2.RemoveTracksFromPlaylistRequest(playlistId=request.playlistId,
                                                                    trackIds=request.trackIds, userId=request.userId)
        try:
            stub.RemoveTracksFromPlaylist.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def get_playlist_by_id(self, playlist_id: str):
        stub = self.stub
        grpc_request = playback_pb2.PlaylistIdRequest(playlistId=playlist_id)
        try:
            response, call_status = stub.GetPlaylistById.with_call(grpc_request, timeout=config.TIMEOUT)
            return {
                "playlistId": response.playlistId,
                "name": response.name,
                "tracks": [
                    {
                        "title": track.title,
                        "artist": track.artist,
                        "album": track.album,
                        "genre": track.genre,
                        "url": track.url,
                        "userId": track.userId
                    } for track in response.tracks
                ]
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def play_playlist(self, playlist_id: str):
        stub = self.stub
        grpc_request = playback_pb2.PlaylistIdRequest(playlistId=playlist_id)
        try:
            response, call_status = stub.PlayPlaylist.with_call(grpc_request, timeout=config.TIMEOUT)
            return {
                "playlistName": response.playlistName,
                "tracks": [
                    {
                        "trackId": track.trackId,
                        "title": track.title,
                        "url": track.url
                    } for track in response.tracks
                ]
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def get_status(self):
        stub = self.stub
        try:
            response, call_status = stub.Status.with_call(empty_pb2.Empty(), timeout=config.TIMEOUT)
            return {
                "status": response.status
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())
