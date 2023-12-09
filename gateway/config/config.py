import os


def extract_port(uri: str) -> int:
    return int(uri.split(":")[-1])


class Config:
    AWS_ACCESS_KEY = os.environ.get("AWS_ACCESS_KEY", "default_value")
    AWS_SECRET_KEY = os.environ.get("AWS_SECRET_KEY", "default_value")
    REDIS_HOST1 = os.environ.get("REDIS_HOST1", "redis-service-1")
    REDIS_HOST2 = os.environ.get("REDIS_HOST2", "redis-service-2")
    REDIS_HOST3 = os.environ.get("REDIS_HOST3", "redis-service-3")
    REDIS_PORT = extract_port(os.environ.get("REDIS_PORT", "6379"))
    USER_SERVICE_HOST = os.environ.get("USER_SERVICE_HOST", "user-service")
    USER_SERVICE_PORT = extract_port(os.environ.get("USER_SERVICE_PORT", "50051"))
    TRACK_SERVICE_HOST = os.environ.get("TRACK_SERVICE_HOST", "track-service")
    TRACK_SERVICE_PORT = extract_port(os.environ.get("TRACK_SERVICE_PORT", "50052"))
    PLAYBACK_SERVICE_HOST = os.environ.get("PLAYBACK_SERVICE_HOST", "playback-service")
    PLAYBACK_SERVICE_PORT = extract_port(os.environ.get("PLAYBACK_SERVICE_PORT", "50053"))
    BUCKET_NAME = os.environ.get("BUCKET_NAME", "tracks-bucket-pad")
    TIMEOUT = extract_port(os.environ.get("TIMEOUT", "5"))
    CACHE_INVALIDATION_TIME = extract_port(os.environ.get("CACHE_INVALIDATION_TIME", "10"))
    CIRCUIT_FAIL_MAX = extract_port(os.environ.get("CIRCUIT_FAIL_MAX", "3"))


config = Config()
