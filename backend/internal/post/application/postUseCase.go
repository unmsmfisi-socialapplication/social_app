package application

import (
	"errors"

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
}

func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}

func (l *PostUseCase) CreatePost(post domain.CreatePost) (*domain.Post, error) {
	dbPost, err := l.repo.CreatePost(post)

	if dbPost == nil {
		return dbPost, err
	}

	return dbPost, nil
}
