from fastapi import APIRouter, Body, Depends, UploadFile, File, HTTPException
from pydantic import ValidationError
import json

from models.track_models import TrackMetadata, EditTrackRequest
from services.track_service import TrackService
from utils.priviliges_utils import get_current_user_id, authorize_track
from utils import s3_utils
from config.config import config
from services import cache

router = APIRouter()
tracks_instance = TrackService.get_instance()


@router.get("/find_all")
async def find_all():
    cached_tracks = cache.get("all_tracks")

    if cached_tracks:
        return json.loads(cached_tracks)

    tracks = tracks_instance.find_all()

    cache.set("all_tracks", json.dumps(tracks), config.CACHE_INVALIDATION_TIME)

    return tracks


@router.post("/upload")
async def upload_track(file: UploadFile = File(...), request: str = Body(...),
                       user_id: str = Depends(get_current_user_id)):
    if file.content_type != "audio/mpeg":
        raise HTTPException(status_code=400, detail="Invalid file type. Please upload an MP3 file.")

    try:
        track_metadata = json.loads(request)
        request = TrackMetadata(**track_metadata)
    except ValidationError as e:
        raise HTTPException(status_code=400, detail=str(e))
    except ValueError:
        raise HTTPException(status_code=400, detail="Invalid JSON format for track metadata")

    unique_filename = s3_utils.create_unique_filename(user_id, file.filename)
    s3_utils.upload_file(file, unique_filename)

    request.userId = user_id
    request.url = f"https://{config.BUCKET_NAME}.s3.amazonaws.com/{unique_filename}"

    return tracks_instance.upload_track(request)


@router.put("/edit")
async def edit_info(request: EditTrackRequest = Body(...), user_id: str = Depends(get_current_user_id)):
    track_info = authorize_track(request.trackId, user_id)
    file_name = track_info["url"].split("/")[-1]

    cache_key = f"track:{request.trackId}"
    cache.delete(cache_key)

    request.metadata.userId = user_id
    request.metadata.url = f"https://{config.BUCKET_NAME}.s3.amazonaws.com/{file_name}"
    return tracks_instance.edit_track_info(request)


@router.get("/status")
async def get_status():
    return tracks_instance.get_status()


@router.delete("/{track_id}")
async def delete_by_id(track_id: str, user_id: str = Depends(get_current_user_id)):
    track_info = authorize_track(track_id, user_id)
    file_name = track_info["url"].split("/")[-1]

    unique_filename = s3_utils.create_unique_filename(user_id, file_name)

    cache_key = f"track:{track_id}"
    cache.delete(cache_key)

    s3_utils.delete_file(unique_filename)
    return tracks_instance.delete_by_id(track_id)


@router.get("/{track_id}")
async def get_info_by_id(track_id: str):
    cache_key = f"track:{track_id}"
    cached_response = cache.get(cache_key)

    if cached_response:
        return json.loads(cached_response)

    response = tracks_instance.get_info_by_id(track_id)

    cache.set(cache_key, json.dumps(response), config.CACHE_INVALIDATION_TIME)

    return response
