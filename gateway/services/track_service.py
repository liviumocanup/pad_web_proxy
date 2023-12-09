import functools
import grpc
from fastapi import HTTPException
from google.protobuf import empty_pb2

from proto import track_pb2, track_pb2_grpc, common_pb2
from models.track_models import TrackMetadata, EditTrackRequest
from config.config import config
from utils.grpc_utils import get_grpc_channel, handle_grpc_error


def get_track_stub():
    channel = get_grpc_channel(config.TRACK_SERVICE_HOST, config.TRACK_SERVICE_PORT)
    return track_pb2_grpc.TrackServiceStub(channel)


class TrackService:
    @staticmethod
    @functools.cache
    def get_instance():
        return TrackService()

    def __init__(self):
        self.stub = get_track_stub()

    @handle_grpc_error
    def upload_track(self, request: TrackMetadata):
        stub = self.stub

        grpc_request = common_pb2.TrackMetadata(
            title=request.title,
            artist=request.artist,
            album=request.album,
            genre=request.genre,
            userId=request.userId,
            url=request.url
        )

        try:
            response, call_status = stub.Upload.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"trackId": response.trackId, "url": response.url}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def restore_state(self, request: TrackMetadata):
        stub = self.stub

        grpc_request = common_pb2.TrackMetadata(
            title=request.title,
            artist=request.artist,
            album=request.album,
            genre=request.genre,
            userId=request.userId,
            url=request.url
        )

        try:
            response, call_status = stub.Upload.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"trackId": response.trackId, "url": response.url}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def get_info_by_id(self, track_id: str):
        stub = self.stub
        grpc_request = track_pb2.TrackIdRequest(id=track_id)
        try:
            response, call_status = stub.GetInfoById.with_call(grpc_request, timeout=config.TIMEOUT)
            return {
                "trackId": response.trackId,
                "title": response.title,
                "artist": response.artist,
                "album": response.album,
                "genre": response.genre,
                "url": response.url,
                "userId": response.userId
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def save_state(self, track_id: str):
        stub = self.stub
        grpc_request = track_pb2.TrackIdRequest(id=track_id)
        try:
            response, call_status = stub.GetInfoById.with_call(grpc_request, timeout=config.TIMEOUT)
            return {
                "trackId": response.trackId,
                "title": response.title,
                "artist": response.artist,
                "album": response.album,
                "genre": response.genre,
                "url": response.url,
                "userId": response.userId
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def edit_track_info(self, request: EditTrackRequest):
        stub = self.stub
        metadata = request.metadata
        grpc_metadata = common_pb2.TrackMetadata(
            title=metadata.title,
            artist=metadata.artist,
            album=metadata.album,
            genre=metadata.genre,
            userId=metadata.userId,
            url=metadata.url,
        )
        grpc_request = track_pb2.EditTrackRequest(trackId=request.trackId, metadata=grpc_metadata)
        try:
            stub.EditInfo.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def delete_by_id(self, track_id: str):
        stub = self.stub
        grpc_request = track_pb2.TrackIdRequest(id=track_id)
        try:
            stub.DeleteById.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def find_all(self):
        stub = self.stub
        grpc_request = empty_pb2.Empty()
        try:
            response, call_status = stub.FindAll.with_call(grpc_request, timeout=config.TIMEOUT)
            return [
                {
                    "trackId": track.trackId,
                    "title": track.title,
                    "artist": track.artist,
                    "album": track.album,
                    "genre": track.genre,
                    "url": track.url,
                    "userId": track.userId
                } for track in response.tracks
            ]
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
