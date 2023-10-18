from fastapi import FastAPI
import uvicorn
from routes import user_routes, track_routes, playback_routes, status_routes

app = FastAPI()
user_router = user_routes.router
track_router = track_routes.router
playback_router = playback_routes.router
status_router = status_routes.router

app.include_router(user_router, prefix="/user", tags=["Users"])
app.include_router(track_router, prefix="/track", tags=["Tracks"])
app.include_router(playback_router, prefix="/playback", tags=["Playback"])
app.include_router(status_router, prefix="/status", tags=["Status"])

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
