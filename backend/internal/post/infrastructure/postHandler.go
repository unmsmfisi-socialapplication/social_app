package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type PostHandler struct {
	useCase application.PostUseCaseInterface
}

func NewPostHandler(useCase application.PostUseCaseInterface) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (ph *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var requestData domain.CreatePost

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	postCreate, err := ph.useCase.CreatePost(requestData)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", postCreate)
}
