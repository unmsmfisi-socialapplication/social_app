package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
}

type PostRepository interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
}

type PostUseCase struct {
	repo PostRepository
    eventManager *events.EventManager
}

func NewPostUseCase(r PostRepository, eventManager *events.EventManager) *PostUseCase {
	return &PostUseCase{repo: r, eventManager: eventManager}
}

func (l *PostUseCase) CreatePost(post domain.CreatePost) (*domain.Post, error) {
	dbPost, err := l.repo.CreatePost(post)

	if dbPost == nil {
		return dbPost, err
	}

    l.eventManager.TriggerEvent("postCreated", dbPost)

	return dbPost, nil
}
