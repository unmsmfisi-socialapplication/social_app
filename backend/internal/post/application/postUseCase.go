package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.PostCreate) (*domain.PostResponse, error)
	GetPost(id int) (*domain.Post, error)
	GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error)
    
    //For multipost
    MultipostPixelfeed(post domain.PostCreate) error
    MultipostMastodon(post domain.PostCreate) error
}


type PostRepository interface {
	CreatePost(post domain.PostCreate) (*domain.Post, error)
	UserExist(post domain.PostCreate) bool
	GetById(id int) (*domain.Post, error)
	GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error)
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