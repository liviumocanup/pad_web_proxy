from fastapi import HTTPException
from google.protobuf import empty_pb2

from config.config import config
from models.user_models import UserRequest, JWT
from proto import user_pb2, user_pb2_grpc
import grpc

from utils.circuit_breaker_manager import CircuitBreakerManager
from utils.grpc_utils import get_grpc_channel, handle_grpc_error


def get_user_stub():
    channel = get_grpc_channel(config.USER_SERVICE_HOST, config.USER_SERVICE_PORT)
    return user_pb2_grpc.UserServiceStub(channel)


breaker = CircuitBreakerManager.get_breaker("UserService")


class UserService:
    @staticmethod
    def get_instance():
        return UserService()

    def __init__(self):
        self.stub = get_user_stub()

    @handle_grpc_error
    @breaker
    def register(self, request: UserRequest):
        stub = self.stub
        grpc_request = user_pb2.UserRequest(username=request.username, password=request.password)
        try:
            stub.Register.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    def login(self, request: UserRequest):
        stub = self.stub
        grpc_request = user_pb2.UserRequest(username=request.username, password=request.password)
        try:
            response, call_status = stub.Login.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"token": response.token}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def validate(self, request: JWT):
        stub = self.stub
        grpc_request = user_pb2.JWT(token=request.token)
        try:
            response, call_status = stub.Validate.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"id": response.id, "username": response.username}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def find_by_id(self, user_id: str):
        stub = self.stub
        grpc_request = user_pb2.UserIdRequest(id=user_id)
        try:
            response, call_status = stub.FindById.with_call(grpc_request, timeout=config.TIMEOUT)

            return {"id": response.id, "username": response.username}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def find_by_username(self, username: str):
        stub = self.stub
        grpc_request = user_pb2.UsernameRequest(username=username)
        try:
            response, call_status = stub.FindByUsername.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"id": response.id, "username": response.username}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def delete_by_id(self, user_id: str, current_user_id: str):
        if user_id != current_user_id:
            raise HTTPException(status_code=403, detail="Not authorized to delete this user.")

        stub = self.stub
        grpc_request = user_pb2.UserIdRequest(id=user_id)
        try:
            stub.DeleteById.with_call(grpc_request, timeout=config.TIMEOUT)
            return {"status": "success"}
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def find_all(self):
        stub = self.stub
        grpc_request = empty_pb2.Empty()
        try:
            response, call_status = stub.FindAll.with_call(grpc_request, timeout=config.TIMEOUT)
            return [{"id": user.id, "username": user.username} for user in response.users]
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())

    @handle_grpc_error
    @breaker
    def get_status(self):
        stub = self.stub
        try:
            response, call_status = stub.Status.with_call(empty_pb2.Empty(), timeout=config.TIMEOUT)
            return {
                "status": response.status
            }
        except grpc.RpcError as e:
            raise HTTPException(status_code=400, detail=e.details())
