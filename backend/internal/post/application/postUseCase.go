package application

import (
	"errors"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.PostCreate) (*domain.Post, error)
	GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error)
	RetrieveTimelinePosts(user_id int64) (*[]domain.TimelineRes, error)
}

type PostRepository interface {
	CreatePost(post domain.PostCreate) (*domain.Post, error)
	UserExist(post domain.PostCreate) bool
	GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error)
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

func (l *PostUseCase) CreatePost(post domain.PostCreate) (*domain.Post, error) {

	if !l.repo.UserExist(post) {
		return nil, ErrUserNotFound
	}

	dbPost, err := l.repo.CreatePost(post)

	if err != nil {
		return nil, err
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

func (l *PostUseCase) GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error) {

	dbPosts, err := l.repo.GetAll(params)

	if err != nil {
		return nil, err
	}

	return dbPosts, nil
}

func (puc *PostUseCase) RetrieveTimelinePosts(user_id int64) (*[]domain.TimelineRes, error) {
	timeline, err := puc.repo.HomeTimeline(user_id)
	if err != nil {
		return nil, err
	}
	return timeline, err
}
