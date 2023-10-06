package application

import (
	"errors"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.CreatePost) (map[string]*domain.Post, error)
}

type PostRepository interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
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

	// Add logic to create and store posts for other social networks as needed

	return createdPosts, nil
}
