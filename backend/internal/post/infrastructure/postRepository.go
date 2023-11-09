package infrastructure

import (
	"database/sql"
	"time"
	"fmt"
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

func (p *PostsDBRepository) CreatePost(post domain.CreatePost) (*domain.Post, error) {
	countPost++

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

func (p *PostsDBRepository) UploadMultimedia(postId int64, multimedia []byte) error {
    return nil
}

func (p *PostsDBRepository) GetMultimedia(postId int64) ([]byte, error) {
    // Database connection
    db := map[int64][]byte{
        1: []byte("imagen1.jpg"),
        2: []byte("video1.mp4"),
    }

    // Simulates a database query
    multimedia, ok := db[postId]
    if !ok {
        return nil, fmt.Errorf("multimedia not found for postId %d", postId)
    }

    return multimedia, nil
}

