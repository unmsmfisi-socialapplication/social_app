package application

import (
	"errors"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
	RetrieveTimelinePosts(user_id int64) (*[]domain.TimelineRes, error)
}

type PostRepository interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
	CreateHomeTimelineItem(user_id int64, post_id int64) error
	FanoutHomeTimeline(posts_id int64, following_id int64) error
	HomeTimeline(user_id int64) (*[]domain.TimelineRes, error)
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

	err = l.repo.CreateHomeTimelineItem(dbPost.UserId, dbPost.Id)
	if err != nil {
		log.Fatal("Error while indexing post with timeline user")
		return dbPost, nil
	}

	go func() {
		err = l.repo.FanoutHomeTimeline(dbPost.Id, dbPost.UserId)
		if err != nil {
			log.Fatal("Error While Running FanOut Broadcast")
		}
	}()

	return dbPost, nil
}

func (puc *PostUseCase) RetrieveTimelinePosts(user_id int64) (*[]domain.TimelineRes, error) {
	timeline, err := puc.repo.HomeTimeline(user_id)
	if err != nil {
		return nil, err
	}

	return timeline, err
}
