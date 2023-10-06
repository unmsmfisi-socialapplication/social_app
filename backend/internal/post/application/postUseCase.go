package application

import (
	"errors"
	"net/http"
	"encoding/json"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
)

type PostUseCaseInterface interface {
	CreatePost(postData domain.CreatePost) (map[string]*domain.Post, error)
}

type PostRepository interface {
	CreatePost(postData interface{}) (*domain.Post, error)
}

type PostUseCase struct {
	repo PostRepository
}

func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}

func (uc *PostUseCase) CreatePost(postData domain.CreatePost) (map[string]*domain.Post, error) {
	createdPosts := make(map[string]*domain.Post)

	// Create and store the post for Mastodon if applicable
	if postData.Mastodon != nil {
		mastodonPost, err := uc.repo.CreatePost(postData.Mastodon)
		if err != nil {
			return nil, err
		}
		createdPosts["mastodon"] = mastodonPost
	}

	// Create and store the post for Pixelfed if applicable
	if postData.Pixelfed != nil {
		pixelfedPost, err := uc.repo.CreatePost(postData.Pixelfed)
		if err != nil {
			return nil, err
		}
		createdPosts["pixelfed"] = pixelfedPost
	}

	return createdPosts, nil
}

type PostHandler struct {
	useCase PostUseCaseInterface
}

func NewPostHandler(useCase PostUseCaseInterface) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (ph *PostHandler) HandleCreateMultiPost(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		MastodonData *domain.CreatePost `json:"mastodon,omitempty"`
		PixelfedData *domain.CreatePost `json:"pixelfed,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	
	var mastodonPost *domain.Post
	var pixelfedPost *domain.Post

	if requestData.MastodonData != nil {
		/* mastodonPost, err := ph.useCase.CreatePost(*requestData.MastodonData)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
			return
		}*/
	}

	if requestData.PixelfedData != nil {
		/* pixelfedPost, err := ph.useCase.CreatePost(*requestData.PixelfedData)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
			return
		} */
	}

	response := map[string]interface{}{
		"mastodon": mastodonPost,
		"pixelfed": pixelfedPost,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	responseString := string(jsonResponse)

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", responseString)
}
