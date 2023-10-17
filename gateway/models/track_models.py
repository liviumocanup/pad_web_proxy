from typing import Optional

from pydantic import BaseModel


class TrackMetadata(BaseModel):
    title: str
    artist: str
    album: str
    genre: str
    userId: Optional[str] = None
    url: Optional[str] = None


class TrackIdRequest(BaseModel):
    id: str


class TrackResponse(BaseModel):
    trackId: str
    url: str
    userId: str


class TrackInfoResponse(BaseModel):
    trackId: str
    title: str
    artist: str
    album: str
    genre: str
    url: str
    userId: str


class EditTrackRequest(BaseModel):
    trackId: str
    metadata: TrackMetadata
