from fastapi import APIRouter, Body, Depends
from models.user_models import UserRequest, JWT
from utils.priviliges_utils import get_current_user_id
from services.user_service import UserService
from config.config import config
import json
from services import cache

router = APIRouter()
users_instance = UserService.get_instance()


@router.get("/find_all")
async def find_all():
    cached_users = cache.get("all_users")

    if cached_users:
        return json.loads(cached_users)

    users = users_instance.find_all()

    cache.set("all_users", json.dumps(users), config.CACHE_INVALIDATION_TIME)

    return users


@router.post("/register")
async def register(request: UserRequest = Body(...)):
    response = users_instance.register(request)
    cache.delete("all_users")
    return response


@router.post("/login")
async def login(request: UserRequest = Body(...)):
    return users_instance.login(request)


@router.post("/validate")
async def validate(request: JWT = Body(...)):
    return users_instance.validate(request)


@router.get("/status")
async def get_status():
    return users_instance.get_status()


@router.delete("/{user_id}")
async def delete_by_id(user_id: str, current_user_id: str = Depends(get_current_user_id)):
    response = users_instance.delete_by_id(user_id, current_user_id)
    cache.delete(f"user:{user_id}")
    cache.delete("all_users")
    return response


@router.get("/{user_id}")
async def find_by_id(user_id: str):
    cache_key = f"user:{user_id}"
    cached_response = cache.get(cache_key)

    if cached_response:
        return json.loads(cached_response)

    response = users_instance.find_by_id(user_id)

    cache.set(cache_key, json.dumps(response), config.CACHE_INVALIDATION_TIME)

    return response


@router.get("/")
async def find_by_username(username: str):
    cache_key = f"user:{username}"
    cached_response = cache.get(cache_key)

    if cached_response:
        return json.loads(cached_response)

    response = users_instance.find_by_username(username)

    cache.set(cache_key, json.dumps(response), config.CACHE_INVALIDATION_TIME)

    return response
