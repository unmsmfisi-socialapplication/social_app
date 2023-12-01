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
	CreatePost(post domain.PostCreate) (*domain.PostResponse, error)
	GetPost(id int) (*domain.Post, error)
	GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error)
	RetrieveTimelinePosts(user_id int64) (*[]domain.TimelineRes, error)
	// DELETE POST //
	DeletePost(id int64) error

	//UPDATE POST //
	UpdatePost(id int64, update domain.PostUpdate) error

	//REPORT POST //
	ReportPost(report domain.PostReport) error
}

type PostRepository interface {
	CreatePost(post domain.PostCreate) (*domain.Post, error)
	UserExist(post domain.PostCreate) bool
	GetById(id int) (*domain.Post, error)
	GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error)
	CreateHomeTimelineItem(user_id int64, post_id int64) error
	FanoutHomeTimeline(posts_id int64, following_id int64) error
	HomeTimeline(user_id int64) (*[]domain.TimelineRes, error)
	// DELETE POST //
	DeletePost(id int64) error
	PostExists(id int64) bool
	// UPDATE POST //
	UpdatePost(id int64, update domain.PostUpdate) error
	CreateReport(report domain.PostReport) error
	ReporterExists(username string) bool
	HasUserAlreadyReportedPost(postId int64, username string) bool
}

type PostUseCase struct {
	repo PostRepository
}

func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}

func (l *PostUseCase) CreatePost(post domain.PostCreate) (*domain.PostResponse, error) {

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
		response := domain.PostToPostResponse(*dbPost)
		return &response, nil
	}

	go func() {
		err = l.repo.FanoutHomeTimeline(dbPost.Id, dbPost.UserId)
		if err != nil {
			log.Fatal("Error While Running FanOut Broadcast")
		}
	}()

	response := domain.PostToPostResponse(*dbPost)

	return &response, nil
}

func (l *PostUseCase) GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error) {

	if params.Page == 0 && params.Limit == 0 {
		params.Limit = 10
		params.Page = 1
	}

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

func (l *PostUseCase) GetPost(id int) (*domain.Post, error) {
	dbPost, err := l.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return dbPost, nil
}

// DELETE POST //

func (l *PostUseCase) DeletePost(id int64) error {
	if !l.repo.PostExists(id) {
		return errors.New("post not found")
	}

	return l.repo.DeletePost(id)
}

// UPDATE POST //
func (l *PostUseCase) UpdatePost(id int64, update domain.PostUpdate) error {

	if !l.repo.PostExists(id) {
		return errors.New("post not found")
	}

	return l.repo.UpdatePost(id, update)
}

// REPORT POST //
func (l *PostUseCase) ReportPost(report domain.PostReport) error {
	if !l.repo.ReporterExists(report.ReportedBy) {
		return domain.ErrReporterUserDoesNotExist
	}
	if !l.repo.PostExists(report.PostId) {
		return domain.ErrPostNotFound
	}

	if l.repo.HasUserAlreadyReportedPost(report.PostId, report.ReportedBy) {
		return domain.ErrUserHasAlreadyReportedPost
	}
	return l.repo.CreateReport(report)
}
