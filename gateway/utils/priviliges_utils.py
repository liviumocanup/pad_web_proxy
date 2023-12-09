from fastapi import Depends, HTTPException, Header

from models.user_models import JWT
from services.track_service import TrackService
from services.user_service import UserService

tracks_instance = TrackService.get_instance()
users_instance = UserService.get_instance()


def get_jwt_token(authorization: str = Header(...)) -> str:
    token_prefix = "Bearer "
    if authorization.startswith(token_prefix):
        return authorization[len(token_prefix):]
    else:
        raise HTTPException(status_code=401, detail="Invalid authorization header")


def get_current_user_id(jwt_token: str = Depends(get_jwt_token)) -> str:
    try:
        user_data = users_instance.validate(JWT(token=jwt_token))
        return user_data["id"]
    except HTTPException as e:
        if e.status_code == 400:
            raise HTTPException(status_code=401, detail="Invalid JWT token")
        raise


def authorize_track(track_id: str, user_id: str = Depends(get_current_user_id)):
    track = tracks_instance.get_info_by_id(track_id)
    if track["userId"] != user_id:
        raise HTTPException(status_code=403, detail="Not authorized to edit or delete this track.")
    return track


def delete_user(user_id: str, current_user_id: str = Depends(get_current_user_id)):
    return users_instance.delete_by_id(user_id, current_user_id)
