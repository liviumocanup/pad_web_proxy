from fastapi import HTTPException
from google.protobuf import empty_pb2

from models.user_service import UserRequest, JWT
from generated import user_pb2, user_pb2_grpc
import grpc

from services import cache


def get_user_stub():
    channel = grpc.insecure_channel("localhost:50051")
    return user_pb2_grpc.UserServiceStub(channel)


def handle_grpc_error(f):
    """
    Decorator to handle gRPC errors in a unified manner.
    """

    def wrapper(*args, **kwargs):
        try:
            return f(*args, **kwargs)
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    return wrapper


@handle_grpc_error
def register(request: UserRequest):
    stub = get_user_stub()
    grpc_request = user_pb2.UserRequest(username=request.username, password=request.password)
    try:
        stub.Register(grpc_request)
        return {"status": "success"}
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())


@handle_grpc_error
def login(request: UserRequest):
    stub = get_user_stub()
    grpc_request = user_pb2.UserRequest(username=request.username, password=request.password)
    try:
        response = stub.Login(grpc_request)
        print(response)
        return {"token": response.token}
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())


@handle_grpc_error
def validate(request: JWT):
    stub = get_user_stub()
    grpc_request = user_pb2.JWT(token=request.token)
    try:
        response = stub.Validate(grpc_request)
        return {"id": response.id, "username": response.username}
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())


@handle_grpc_error
def find_by_id(user_id: str):
    # Attempt to fetch from cache first
    cached_user = cache.get(user_id)
    if cached_user:
        return cached_user

    stub = get_user_stub()
    grpc_request = user_pb2.IdRequest(id=user_id)
    try:
        response = stub.FindById(grpc_request)

        # If the user is found, store it in the cache.
        # Note: The cache value is serialized as a string.
        user_data = {"id": response.id, "username": response.username}
        cache.set(user_id, str(user_data))

        return user_data
    except grpc.RpcError as e:
        # invalidate the cache for that user_id (if exists).
        cache.delete(user_id)
        raise HTTPException(status_code=400, detail=e.details())


@handle_grpc_error
def find_by_username(username: str):
    stub = get_user_stub()
    grpc_request = user_pb2.UsernameRequest(username=username)
    try:
        response = stub.FindByUsername(grpc_request)
        return {"id": response.id, "username": response.username}
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())


@handle_grpc_error
def find_all():
    stub = get_user_stub()
    request = empty_pb2.Empty()  # This is the protobuf Empty message.
    try:
        response = stub.FindAll(request)
        # Convert the gRPC user list response to a list of dictionaries.
        return [{"id": user.id, "username": user.username} for user in response.users]
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())
