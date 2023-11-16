from redis.sentinel import Sentinel
from config.config import config

# Initialize Sentinel instance
sentinel = Sentinel([(config.REDIS_SENTINEL_HOST, config.REDIS_SENTINEL_PORT)], socket_timeout=0.1)

# Connect to master
redis_client = sentinel.master_for('mymaster', socket_timeout=0.1, db=0)


def set(key: str, value: str, expiration: int = None):
    if expiration:
        redis_client.setex(key, expiration, value)
    else:
        redis_client.set(key, value)


def get(key: str) -> str:
    return redis_client.get(key)


def delete(key: str):
    redis_client.delete(key)
