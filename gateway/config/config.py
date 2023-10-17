from pydantic.v1 import BaseSettings


class Config(BaseSettings):
    REDIS_HOST: str = "localhost"
    REDIS_PORT: int = 6379
    USER_SERVICE_HOST: str = "localhost"
    USER_SERVICE_PORT: int = 50051
    TRACK_SERVICE_HOST: str = "localhost"
    TRACK_SERVICE_PORT: int = 50052
    PLAYBACK_SERVICE_HOST: str = "localhost"
    PLAYBACK_SERVICE_PORT: int = 50053
    AWS_ACCESS_KEY: str = "AKIAS7M73W4WKXNEBD7C"
    AWS_SECRET_KEY: str = "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"
    BUCKET_NAME: str = "tracks-bucket-pad"
    TIMEOUT: int = 5
    CACHE_INVALIDATION_TIME: int = 10


config = Config()
