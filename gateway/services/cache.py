import redis
from config.config import config

redis_client = redis.StrictRedis(host=config.REDIS_HOST, port=config.REDIS_PORT, db=0)


def set(key: str, value: str, expiration: int = None):
    if expiration:
        redis_client.setex(key, expiration, value)
    else:
        redis_client.set(key, value)


def get(key: str) -> str:
    return redis_client.get(key)


def delete(key: str):
    redis_client.delete(key)
