from pydantic import BaseModel


class UserRequest(BaseModel):
    username: str
    password: str


class UserIdRequest(BaseModel):
    id: str


class UsernameRequest(BaseModel):
    username: str


class JWT(BaseModel):
    token: str
