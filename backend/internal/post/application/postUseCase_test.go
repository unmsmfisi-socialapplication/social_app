package application

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

type MockPostRepository struct {
	UserExistFunc                  func(post domain.PostCreate) bool
	CreatePostFunc                 func(post domain.PostCreate) (*domain.Post, error)
	GetByIdFunc                    func(id int) (*domain.Post, error)
	GetAllFunc                     func(params domain.PostPaginationParams) (*domain.PostPagination, error)
	DeletePostFunc                 func(id int64) error
	PostExistsFunc                 func(id int64) bool
	UpdatePostFunc                 func(id int64, update domain.PostUpdate) error
	CreateReportFunc               func(report domain.PostReport) error
	ReporterExistsFunc             func(username string) bool
	HasUserAlreadyReportedPostFunc func(postId int64, username string) bool
}

// CreateHomeTimelineItem implements PostRepository.
func (*MockPostRepository) CreateHomeTimelineItem(user_id int64, post_id int64) error {
	panic("unimplemented")
}

// FanoutHomeTimeline implements PostRepository.
func (*MockPostRepository) FanoutHomeTimeline(posts_id int64, following_id int64) error {
	panic("unimplemented")
}

// HomeTimeline implements PostRepository.
func (*MockPostRepository) HomeTimeline(user_id int64, page_size int64, page_num int64) (*domain.QueryResult, error) {
	panic("unimplemented")
}

func (m *MockPostRepository) UserExist(post domain.PostCreate) bool {
	return m.UserExistFunc(post)
}

func (m *MockPostRepository) CreatePost(post domain.PostCreate) (*domain.Post, error) {
	return m.CreatePostFunc(post)
}

func (m *MockPostRepository) GetById(id int) (*domain.Post, error) {
	return m.GetByIdFunc(id)
}

func (m *MockPostRepository) GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error) {
	return m.GetAllFunc(params)
}

func (m *MockPostRepository) DeletePost(id int64) error {
	return m.DeletePostFunc(id)
}

func (m *MockPostRepository) PostExists(id int64) bool {
	return m.PostExistsFunc(id)
}

func (m *MockPostRepository) UpdatePost(id int64, update domain.PostUpdate) error {
	return m.UpdatePostFunc(id, update)
}

func (m *MockPostRepository) CreateReport(report domain.PostReport) error {
	return m.CreateReportFunc(report)
}

func (m *MockPostRepository) ReporterExists(username string) bool {
	return m.ReporterExistsFunc(username)
}

func (m *MockPostRepository) HasUserAlreadyReportedPost(postId int64, username string) bool {
	return m.HasUserAlreadyReportedPostFunc(postId, username)
}

func TestPostUseCase_CreatePost_UserNotFound(t *testing.T) {
	mockRepo := &MockPostRepository{
		UserExistFunc: func(post domain.PostCreate) bool {
			return false
		},
	}
	useCase := NewPostUseCase(mockRepo)
	testPostCreate := domain.PostCreate{
		PostBase: domain.PostBase{
			UserId:        1,
			Title:         "Test Title",
			Description:   "Test Description",
			HasMultimedia: false,
			Public:        true,
			Multimedia:    "",
		},
	}

	response, err := useCase.CreatePost(testPostCreate)

	expectedError := errors.New("user not found")

	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}

	if response != nil {
		t.Errorf("Expected nil response, but got %+v", response)
	}
}
