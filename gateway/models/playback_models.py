from typing import List, Optional
from pydantic import BaseModel
from models.track_models import TrackMetadata


class CreatePlaylistRequest(BaseModel):
    name: str
    userId: Optional[str] = None


class PlaylistIdRequest(BaseModel):
    playlistId: str


class RemovePlaylistRequest(BaseModel):
    playlistId: str
    userId: Optional[str] = None


class AddTracksToPlaylistRequest(BaseModel):
    playlistId: str
    trackIds: List[str]
    userId: Optional[str] = None


class RemoveTracksFromPlaylistRequest(BaseModel):
    playlistId: str
    trackIds: List[str]
    userId: Optional[str] = None


class PlaylistResponse(BaseModel):
    playlistId: str
    name: str
    tracks: List[TrackMetadata]


class TrackPlayMetadata(BaseModel):
    trackId: str
    title: str
    url: str


class PlayPlaylistResponse(BaseModel):
    playlistName: str
    tracks: List[TrackPlayMetadata]
