from fastapi import FastAPI, Body
from models import user_service
from models.user_service import UserRequest, JWT
from services import user_service

app = FastAPI()


@app.post("/user/register/")
async def register_endpoint(request: UserRequest = Body(...)):
    return user_service.register(request)


@app.post("/user/login/")
async def login_endpoint(request: UserRequest = Body(...)):
    return user_service.login(request)


@app.post("/user/validate/")
async def validate_endpoint(request: JWT = Body(...)):
    return user_service.validate(request)


@app.get("/user/find_by_id/")
async def find_by_id_endpoint(user_id: str):
    return user_service.find_by_id(user_id)


@app.get("/user/find_by_username/")
async def find_by_username_endpoint(username: str):
    return user_service.find_by_username(username)


@app.get("/user/find_all/")
async def find_all_endpoint():
    return user_service.find_all()


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=8000)
