from pydantic.v1 import BaseSettings
import os


class Config(BaseSettings):
    AWS_ACCESS_KEY: str = os.environ.get("AWS_ACCESS_KEY", "default_value")
    AWS_SECRET_KEY: str = os.environ.get("AWS_SECRET_KEY", "default_value")
    REDIS_HOST: str = "redis-service"
    REDIS_PORT: int = 6379
    USER_SERVICE_HOST: str = "user-service"
    USER_SERVICE_PORT: int = 50051
    TRACK_SERVICE_HOST: str = "track-service"
    TRACK_SERVICE_PORT: int = 50052
    PLAYBACK_SERVICE_HOST: str = "playback-service"
    PLAYBACK_SERVICE_PORT: int = 50053
    BUCKET_NAME: str = "tracks-bucket-pad"
    TIMEOUT: int = 5
    CACHE_INVALIDATION_TIME: int = 10


config = Config()
