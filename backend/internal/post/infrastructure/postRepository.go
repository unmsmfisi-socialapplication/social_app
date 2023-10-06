package infrastructure

import (
	"database/sql"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var countPost int64 = 0
var localPost = []domain.Post{}

type PostsDBRepository struct {
	db *sql.DB
}

func NewPostDBRepository(database *sql.DB) application.PostRepository {
	return &PostsDBRepository{db: database}
}

func (p *PostsDBRepository) CreatePost(postData interface{}) (*domain.Post, error) {
	countPost++

	post, ok := postData.(domain.CreatePost)
	if !ok {
		return nil, errors.New("invalid postData type")
	}

	dbPost := &domain.Post{
		Id:            countPost,
		UserId:        post.UserId,
		Title:         post.Title,
		Description:   post.Description,
		HasMultimedia: post.HasMultimedia,
		Public:        post.Public,
		Multimedia:    post.Multimedia,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	localPost = append(localPost, *dbPost)

	return dbPost, nil
}