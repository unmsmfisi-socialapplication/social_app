package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/helpers"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type PostHandler struct {
	useCase application.PostUseCaseInterface
}

func NewPostHandler(useCase application.PostUseCaseInterface) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (ph *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var requestData domain.PostCreate

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	if requestData.PostBase.UserId == 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request User")
		return
	}

	if requestData.PostBase.Title == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request Title")
		return
	}

	postCreate, err := ph.useCase.CreatePost(requestData)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", postCreate)
}

func (ph *PostHandler) HandleGetAllPost(w http.ResponseWriter, r *http.Request) {

	params, err := helpers.ParsePaginationParams(r.URL.Query())
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	posts, err := ph.useCase.GetPosts(params)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", posts)
}

func (ph *PostHandler) HandleTimeline(w http.ResponseWriter, r *http.Request) {

	var data struct {
		UserId int `json:"user_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeline, err := ph.useCase.RetrieveTimelinePosts(int64(data.UserId))
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", timeline)
}
