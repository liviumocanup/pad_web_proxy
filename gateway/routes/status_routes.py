from fastapi import APIRouter, HTTPException
from services.playback_service import PlaybackService
from services.track_service import TrackService
from services.user_service import UserService

router = APIRouter()

SERVICES = {
    "user": UserService.get_instance(),
    "track": TrackService.get_instance(),
    "playback": PlaybackService.get_instance()
}


@router.get("/")
def read_status():
    return {"status": "ok"}


@router.get("/discovery")
def discover_services():
    discovered = {}
    for name, instance in SERVICES.items():
        try:
            response = instance.get_status()
            if response['status'] == 'ok':
                discovered[name] = "online"
            else:
                discovered[name] = "unavailable"
        except HTTPException as e:
            discovered[name] = "unavailable"
    return discovered
