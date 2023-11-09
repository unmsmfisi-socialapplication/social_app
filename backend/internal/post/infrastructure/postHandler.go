package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"

	//"github.com/unmsmfisi-socialapplication/social_app/internal/post/infrastructure/postRepository"
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

func (ph *PostHandler) HandleGetMultimedia(w http.ResponseWriter, r *http.Request) {
    // Get the post ID from the request
    postIdStr := r.URL.Query().Get("postId")

    // Validate the post ID
    if postIdStr == "" {
        utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Post ID must not be empty")
        return
    }

    // Convert postIdStr to int64
    postId, err := strconv.ParseInt(postIdStr, 10, 64)
    if err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid Post ID format")
        return
    }

    // Get the multimedia data from the repository
    multimedia, err := ph.useCase.GetMultimedia(postId)

    // Handle errors
    if err != nil {
        if err == application.ErrPostNotFound {
            utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Post not found")
        } else {
            utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
        }
        return
    }

    // Send a successful response with the multimedia data
    utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", multimedia)
}
