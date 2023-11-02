package infrastructure

import (
    "database/sql"
    "time"
    "errors"

    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var countPost int64 = 0

var localPost = []domain.Post{}

// PostsDBRepository is a repository for managing posts in a database.
type PostsDBRepository struct {
    db *sql.DB
}

// NewPostDBRepository creates a new instance of PostsDBRepository.
func NewPostDBRepository(database *sql.DB) *PostsDBRepository {
    return &PostsDBRepository{db: database}
}

// CreatePost creates a new post in the database based on the provided data.
func (p *PostsDBRepository) CreatePost(postData interface{}) (*domain.Post, error) {
    countPost++

    var post domain.CreatePost
    if typed, ok := postData.(domain.CreatePost); ok {
        post = typed
    } else {
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
