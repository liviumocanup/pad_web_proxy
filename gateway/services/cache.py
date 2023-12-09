import redis
from hashring import HashRing

from config.config import config

# Define Redis instances
servers = {
    'redis1': {'host': config.REDIS_HOST1, 'port': config.REDIS_PORT},
    'redis2': {'host': config.REDIS_HOST2, 'port': config.REDIS_PORT},
    'redis3': {'host': config.REDIS_HOST3, 'port': config.REDIS_PORT}
}

# Create HashRing
ring = HashRing(servers.keys())

# Initialize Redis clients
redis_clients = {
    name: redis.StrictRedis(host=details['host'], port=details['port'], db=0)
    for name, details in servers.items()
}


def get_redis_client(key: str):
    server_name = ring.get_node(key)
    return redis_clients[server_name]


def set(key: str, value: str, expiration: int = None):
    client = get_redis_client(key)
    if expiration:
        client.setex(key, expiration, value)
    else:
        client.set(key, value)


def get(key: str) -> str:
    client = get_redis_client(key)
    return client.get(key)


def delete(key: str):
    client = get_redis_client(key)
    client.delete(key)
