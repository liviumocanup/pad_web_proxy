import os
from urllib.parse import unquote
from fastapi import APIRouter, Body, Depends, HTTPException
from models.playback_models import (CreatePlaylistRequest,
                                    AddTracksToPlaylistRequest, RemoveTracksFromPlaylistRequest, RemovePlaylistRequest)
from services.auth import get_current_user_id
from services.playback_service import PlaybackService
from utils.s3_utils import download_file
from services import cache
from config.config import config
import json

router = APIRouter()
playback_instance = PlaybackService.get_instance()


@router.post("/create")
async def create_playlist(request: CreatePlaylistRequest = Body(...), user_id: str = Depends(get_current_user_id)):
    request.userId = user_id
    result = playback_instance.create_playlist(request)
    cache.delete(f"playlist:{result.get('playlistId')}")
    return result


@router.delete("/remove")
async def remove_playlist(request: RemovePlaylistRequest = Body(...), user_id: str = Depends(get_current_user_id)):
    request.userId = user_id
    cache.delete(f"playlist:{request.playlistId}")
    return playback_instance.remove_playlist(request)


@router.post("/add_tracks")
async def add_tracks_to_playlist(request: AddTracksToPlaylistRequest = Body(...),
                                 user_id: str = Depends(get_current_user_id)):
    request.userId = user_id
    cache.delete(f"playlist:{request.playlistId}")
    return playback_instance.add_tracks_to_playlist(request)


@router.delete("/remove_tracks")
async def remove_tracks_from_playlist(request: RemoveTracksFromPlaylistRequest = Body(...),
                                      user_id: str = Depends(get_current_user_id)):
    request.userId = user_id
    cache.delete(f"playlist:{request.playlistId}")
    return playback_instance.remove_tracks_from_playlist(request)


@router.get("/status")
async def get_status():
    return playback_instance.get_status()


@router.get("/play/{playlist_id}")
async def play_playlist(playlist_id: str, user_id: str = Depends(get_current_user_id)):
    response = playback_instance.play_playlist(playlist_id)

    if not response:
        raise HTTPException(status_code=404, detail="Playlist not found")

    # Cache handling for playlist play
    cache_key = f"play_playlist:{playlist_id}"
    cached_response = cache.get(cache_key)
    if cached_response:
        return json.loads(cached_response)

    # Create the required folders
    playlist_name = response.get("playlistName")
    destination_folder = os.path.join("downloads", f"{user_id}_{playlist_name}")
    os.makedirs(destination_folder, exist_ok=True)

    # Download each track to the specified folder
    for track in response.get("tracks", []):
        s3_url = track.get("url")
        # Extracting the original filename from the URL after decoding
        file_name = unquote(s3_url.split("/")[-1])
        destination_path = os.path.join(destination_folder, file_name)

        download_file(s3_url, destination_path)

    cache.set(cache_key, json.dumps(response), config.CACHE_INVALIDATION_TIME)

    return response


@router.get("/{playlist_id}")
async def get_playlist_by_id(playlist_id: str):
    cache_key = f"playlist:{playlist_id}"
    cached_response = cache.get(cache_key)

    if cached_response:
        return json.loads(cached_response)

    response = playback_instance.get_playlist_by_id(playlist_id)
    cache.set(cache_key, json.dumps(response), config.CACHE_INVALIDATION_TIME)

    return response
