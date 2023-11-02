package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type FollowerUserHandler struct {
	useCase *application.FollowerUseCase
}

func NewFollowerHandler(uc *application.FollowerUseCase) *FollowerUserHandler {
	return &FollowerUserHandler{useCase: uc}
}

func (rh *FollowerUserHandler) FollowProfile(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Follower_profile_id  int `json:"follower_profile_id"`
		Following_profile_id int `json:"following_profile_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, er := rh.useCase.FollowProfile(data.Follower_profile_id, data.Following_profile_id)
	if er != nil {
		switch er {
		case application.ErrFollowHimself:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Profile Follow Himself")
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", er.Error())
		}
	}
	json.NewEncoder(w).Encode(data)
}
