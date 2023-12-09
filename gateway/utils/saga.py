from fastapi import HTTPException

from models.track_models import TrackMetadata


class SagaCoordinator:
    def __init__(self, track_service, user_service):
        self.track_service = track_service
        self.user_service = user_service

    async def execute_saga(self, track_id, user_id):
        state = self.track_service.save_state(track_id)

        try:
            # Step 1: Delete track
            delete_track_response = self.track_service.delete_by_id(track_id)
            print("delete_track_response", delete_track_response)
            if delete_track_response["status"] != "success":
                raise HTTPException(status_code=400, detail="Failed to delete track")

            # Step 2: Delete user
            delete_user_response = self.user_service.delete_by_id(user_id, user_id)
            print("delete_user_response", delete_user_response)
            if delete_user_response["status"] != "success":
                raise HTTPException(status_code=400, detail="Failed to delete user")

            return {"status": "success"}

        except HTTPException as e:
            if e.detail == "Not authorized to delete this user." or e.detail == "Failed to delete user":
                self.compensate_delete_track(state)
            return {"status": "failure", "reason": e.detail.title()}

    def compensate_delete_track(self, state):
        self.track_service.restore_state(TrackMetadata(**state))
