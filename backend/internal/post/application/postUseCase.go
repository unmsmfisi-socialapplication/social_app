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

// PostUseCaseInterface defines the interface for post-related use cases.
type PostUseCaseInterface interface {
	CreatePost(postData domain.CreatePost) (map[string]*domain.Post, error)
}

// PostRepository defines the interface for post-related data storage operations.
type PostRepository interface {
	CreatePost(postData interface{}) (*domain.Post, error)
}

// PostUseCase is responsible for handling post-related business logic.
type PostUseCase struct {
	repo PostRepository
}

// NewPostUseCase creates a new instance of PostUseCase.
func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}

// CreatePost creates a post based on the provided data for Mastodon and Pixelfed.
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

// PostHandler handles HTTP requests related to posts.
type PostHandler struct {
	useCase PostUseCaseInterface
}

// NewPostHandler creates a new instance of PostHandler.
func NewPostHandler(useCase PostUseCaseInterface) *PostHandler {
	return &PostHandler{useCase: useCase}
}

// HandleCreateMultiPost handles the creation of posts for Mastodon and Pixelfed.
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
/*
	if requestData.MastodonData != nil {
	}

	if requestData.PixelfedData != nil {
	}*/

	response := map[string]interface{}{
		"mastodon": mastodonPost,
		"pixelfed": pixelfedPost,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	// Convert the JSON response to a string
	responseString := string(jsonResponse)

	// Send the JSON response as a string
	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", responseString)
}