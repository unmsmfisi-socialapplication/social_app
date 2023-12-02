package application

import (
	"errors"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	 // User-related errors
	 ErrUserNotFound = errors.New("user not found")
	 ErrUserUnauthorized = errors.New("user unauthorized")
	 ErrUserInactive = errors.New("user inactive")
 
	 // Post-related errors
	 ErrPostNotFound = errors.New("post not found")
	 ErrPostInvalid = errors.New("invalid post")
	 ErrPostAccessDenied = errors.New("access to post denied")
 
	 // External API-related errors
	 ErrAPIConnectionFailure = errors.New("failed to connect to external API")
	 ErrAPIResponseError = errors.New("error in external API response")
	 ErrAPITimeout = errors.New("external API request timed out")
 
	 // Database-related errors
	 ErrDatabaseConnection = errors.New("database connection error")
	 ErrDatabaseQueryFailed = errors.New("database query failed")
	 ErrRecordNotFound = errors.New("record not found in database") 
)

type PostUseCaseInterface interface {
	CreatePost(post domain.PostCreate) (*domain.PostResponse, error)
	GetPost(id int) (*domain.Post, error)
	GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error)
  MultipostPixelfeed(post domain.PostCreate) error
  MultipostMastodon(post domain.PostCreate) error
	RetrieveTimelinePosts(user_id, page_size, page_num int64) (*domain.QueryResult, error)

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
	HomeTimeline(user_id, page_size, page_num int64) (*domain.QueryResult, error)
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
    //For multipost
    pixelfeedAPI PixelfeedAPI
    mastodonAPI  MastodonAPI
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

func (puc *PostUseCase) RetrieveTimelinePosts(user_id, page_size, page_num int64) (*domain.QueryResult, error) {
	timeline, err := puc.repo.HomeTimeline(user_id, page_size, page_num)
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


func NewPostUseCaseWithApis(r PostRepository, pixelfeedAPI PixelfeedAPI, mastodonAPI MastodonAPI) *PostUseCase {
    return &PostUseCase{repo: r, pixelfeedAPI: pixelfeedAPI, mastodonAPI: mastodonAPI}
}


func (l *PostUseCase) MultipostPixelfeed(post domain.PostCreate) error {
    /// Logic for posting to Mastodon using pixelfeedAPI
    return nil
}

func (l *PostUseCase) MultipostMastodon(post domain.PostCreate) error {
    // Logic for posting to Mastodon using mastodonAPI
    return nil
}

func (p *PostUseCase) PostToMultiplePlatforms(post domain.PostCreate) error {
	return nil
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
