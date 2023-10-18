import re

import grpc
from grpc import StatusCode
from fastapi import HTTPException

GRPC_STATUS_TO_HTTP = {
    StatusCode.INVALID_ARGUMENT: 400,
    StatusCode.UNAUTHENTICATED: 401,
    StatusCode.PERMISSION_DENIED: 403,
    StatusCode.NOT_FOUND: 404,
}


def get_grpc_channel(host: str, port: int) -> grpc.Channel:
    return grpc.insecure_channel(f"{host}:{port}")


def handle_grpc_error(f):
    """
    Decorator to handle gRPC errors in a unified manner.
    """

    def wrapper(*args, **kwargs):
        try:
            return f(*args, **kwargs)
        except grpc.RpcError as e:
            http_code = GRPC_STATUS_TO_HTTP.get(e.code(), 400)
            raise HTTPException(status_code=http_code, detail=e.details())
        except HTTPException as e:
            print(e.detail)
            if "DNS resolution failed" in e.detail:
                pattern = r'(\S+:\d+)'
                service_name = re.search(pattern, e.detail).group(1)
                raise HTTPException(status_code=503, detail=service_name + " unavailable.")
            if "failed to connect to all addresses" in e.detail:
                raise HTTPException(status_code=503, detail="Service unavailable.")
            raise e
        except Exception as e:
            if "Circuit is OPEN" in str(e):
                raise HTTPException(status_code=503, detail="Circuit Breaker OPEN: Service is currently unavailable.")
            raise e

    return wrapper
