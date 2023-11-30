package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
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
	} else {
		responseData := struct {
			Response string `json:"response"`
			Follow   struct {
				Follower_profile_id  int `json:"follower_profile_id"`
				Following_profile_id int `json:"following_profile_id"`
			} `json:"follow"`
		}{
			Response: "Profile Followed Succesfully",
			Follow: struct {
				Follower_profile_id  int `json:"follower_profile_id"`
				Following_profile_id int `json:"following_profile_id"`
			}{
				Follower_profile_id:  data.Follower_profile_id,
				Following_profile_id: data.Following_profile_id,
			},
		}
		json.NewEncoder(w).Encode(responseData)
	}
}

func (rh *FollowerUserHandler) IsFollowing(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Follower_profile_id  int `json:"follower_profile_id"`
		Following_profile_id int `json:"following_profile_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isFollowing, er := rh.useCase.IsFollowing(data.Follower_profile_id, data.Following_profile_id)
	if er != nil {
		switch er {
		case application.ErrFollowHimself:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Profile Follow Himself")
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", er.Error())
		}
	}

	var isFollowingResponse struct {
		Follower_profile_id  int   `json:"follower_profile_id"`
		Following_profile_id int   `json:"following_profile_id"`
		IsFollowing          *bool `json:"is_following"`
	}

	isFollowingResponse = struct {
		Follower_profile_id  int   `json:"follower_profile_id"`
		Following_profile_id int   `json:"following_profile_id"`
		IsFollowing          *bool `json:"is_following"`
	}{
		data.Follower_profile_id,
		data.Following_profile_id,
		isFollowing,
	}

	json.NewEncoder(w).Encode(isFollowingResponse)
}

func (rh *FollowerUserHandler) ViewCommonFollowers(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Own_profile_id    int `json:"own_profile_id"`
		Viewed_profile_id int `json:"viewed_profile_id"`
		Page_size         int `json:"page_size"`
		Page_number       int `json:"page_number"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common_profiles, er := rh.useCase.CommonFollowers(data.Own_profile_id, data.Viewed_profile_id, data.Page_size, data.Page_number)
	if er != nil {
		switch er {
		case application.ErrFollowHimself:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "The same profiles have been entered")
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", er.Error())
		}
	} else {
		responseData := struct {
			Response       string                   `json:"response"`
			Own_profile_id int                      `json:"own_profile_id"`
			Page_size      int                      `json:"page_size"`
			Page_number    int                      `json:"page_number"`
			Data           *domain.FollowerDataList `json:"data"`
		}{
			Response:       "Common Profiles Succesfully",
			Own_profile_id: data.Own_profile_id,
			Page_size:      data.Page_size,
			Page_number:    data.Page_number,
			Data:           common_profiles,
		}
		json.NewEncoder(w).Encode(responseData)
	}
}

func (rh *FollowerUserHandler) ProfileFollowers(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Profile_id  int `json:"profile_id"`
		Page_size   int `json:"page_size"`
		Page_number int `json:"page_number"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile_followers, er := rh.useCase.ProfileFollowers(data.Profile_id, data.Page_size, data.Page_number)
	if er != nil {
		switch er {
		case application.ErrFollowHimself:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "The same profiles have been entered")
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", er.Error())
		}
	} else {
		responseData := struct {
			Response    string                   `json:"response"`
			Profile_id  int                      `json:"profile_id"`
			Page_size   int                      `json:"page_size"`
			Page_number int                      `json:"page_number"`
			Data        *domain.FollowerDataList `json:"data"`
		}{
			Response:    "Profile Following Succesfully",
			Profile_id:  data.Profile_id,
			Page_size:   data.Page_size,
			Page_number: data.Page_number,
			Data:        profile_followers,
		}
		json.NewEncoder(w).Encode(responseData)
	}
}

func (rh *FollowerUserHandler) ProfileFollowingProfiles(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Profile_id  int `json:"profile_id"`
		Page_size   int `json:"page_size"`
		Page_number int `json:"page_number"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile_followers, er := rh.useCase.ProfileFollowers(data.Profile_id, data.Page_size, data.Page_number)
	if er != nil {
		switch er {
		case application.ErrFollowHimself:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "The same profiles have been entered")
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", er.Error())
		}
	} else {
		responseData := struct {
			Response    string                   `json:"response"`
			Profile_id  int                      `json:"profile_id"`
			Page_size   int                      `json:"page_size"`
			Page_number int                      `json:"page_number"`
			Data        *domain.FollowerDataList `json:"data"`
		}{
			Response:    "Profile Following Profiles Succesfully",
			Profile_id:  data.Profile_id,
			Page_size:   data.Page_size,
			Page_number: data.Page_number,
			Data:        profile_followers,
		}
		json.NewEncoder(w).Encode(responseData)
	}
}
