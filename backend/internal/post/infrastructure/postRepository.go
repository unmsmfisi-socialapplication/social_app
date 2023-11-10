package infrastructure

import (
	"database/sql"
	"time"
	"encoding/json"
	"fmt"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var countPost int64 = 0
var localPost = []domain.Post{}

type PostsDBRepository struct {
	db *sql.DB
	localPost map[int64]domain.Post
}

func NewPostsDBRepository() application.PostRepository {
	return &PostsDBRepository{localPost: make(map[int64]domain.Post)}
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

//Code to update a post in the database
func (p *PostsDBRepository) UpdatePost(postId int64, post domain.CreatePost) (*domain.Post, error) {
    // Get the post from the repository
    dbPost, ok := p.localPost[postId]
    if !ok {
        return nil, fmt.Errorf("post not found for postId %d", postId)
    }

    // Update the post fields
    dbPost.Title = post.Title
    dbPost.Description = post.Description
    dbPost.HasMultimedia = post.HasMultimedia
    dbPost.Public = post.Public
    dbPost.Multimedia = post.Multimedia
    dbPost.UpdateDate = time.Now()

    // Save the updated post to the repository
    err := p.SavePost(dbPost)
    if err != nil {
        return nil, err
    }

    return &dbPost, nil
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


func (p *PostsDBRepository) GetPostsUser(userId int64) ([]byte, error) {
    // Database connection
    db := map[int64]map[string]interface{}{
        1: {
            "Id": 1,
            "UserId": 456,
            "Title": "Mi primera publicación",
            "Description": "Esta es mi primera publicación en la plataforma.",
            "HasMultimedia": false,
            "Public": true,
            "InsertionDate": "2023-11-07T10:00:00.000000000Z",
        },
        2: {
            "Id": 2,
            "UserId": 456,
            "Title": "Otra publicación",
            "Description": "Una segunda publicación para compartir contenido interesante.",
            "HasMultimedia": true,
            "Public": true,
            "Multimedia": []interface{}{"imagen1.jpg", "imagen2.jpg"},
            "InsertionDate": "2023-11-08T09:30:00.000000000Z",
        },
    }

    // Simulates a database query
    post, ok := db[userId]
    if !ok {
        return nil, fmt.Errorf("post not found for userId %d", userId)
    }

    // Convert the post to JSON
    jsonData, err := json.Marshal(post)
    if err != nil {
        return nil, err
    }

    return jsonData, nil
}

func (p *PostsDBRepository) GetPost(postId int64) (*domain.Post, error) {
    // Get the post from the repository
    dbPost, ok := p.localPost[postId]
    if !ok {
        return nil, fmt.Errorf("post not found for postId %d", postId)
    }

    return &dbPost, nil
}


