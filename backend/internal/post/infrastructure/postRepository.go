package infrastructure

import (
	"database/sql"
	"time"
	"encoding/json"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

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

func (p *PostsDBRepository) CreatePost(post domain.PostCreate) (*domain.Post, error) {

	query := `INSERT INTO soc_app_posts
	(user_id,title, description, has_multimedia, public, multimedia, insertion_date, update_date)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING post_id;	
	`

	dbPost := domain.PostCreateToPost(post)

	err := p.db.QueryRow(
		query,
		dbPost.UserId,
		dbPost.Description,
		dbPost.Title,
		dbPost.HasMultimedia,
		dbPost.Public,
		dbPost.Multimedia,
		dbPost.InsertionDate,
		dbPost.UpdateDate,
	).Scan(&dbPost.Id)

	if err != nil {
		return nil, err
	}

	return &dbPost, nil
}

func (p *PostsDBRepository) UserExist(post domain.PostCreate) bool {
	query := `SELECT user_id FROM sa.soc_app_users WHERE user_id =  $1`

	var dbUserId int64

	p.db.QueryRow(query, post.UserId).Scan(&dbUserId)

	return dbUserId != 0
}

func (p *PostsDBRepository) GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error) {
	query := `SELECT post_id ,user_id, title, description, has_multimedia, 
				public, multimedia, insertion_date, update_date 
				FROM sa.soc_app_posts 
				ORDER BY insertion_date DESC 
              	LIMIT $1 OFFSET $2`

	offset := (params.Page - 1) * params.Limit

	rows, err := p.db.Query(query, params.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dbPosts domain.PostPagination
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Description, &post.HasMultimedia, &post.Public, &post.Multimedia, &post.InsertionDate, &post.UpdateDate)
		if err != nil {
			return nil, err
		}
		dbPosts.Posts = append(dbPosts.Posts, post)
	}

	totalRowsQuery := "SELECT COUNT(*) FROM sa.soc_app_posts "
	err = p.db.QueryRow(totalRowsQuery).Scan(&dbPosts.TotalCount)
	if err != nil {
		return nil, err
	}

	dbPosts.CurrentPage = params.Page

	return &dbPosts, nil
}

func (p *PostsDBRepository) SavePost(post *domain.Post) error {
    // Update the post in the repository
    p.localPost[post.Id] = *post

    return nil
}

//Code to update a post in the database
func (p *PostsDBRepository) UpdatePost(postId int64, post domain.PostCreate) (*domain.Post, error) {
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
    err := p.SavePost(&dbPost)
    if err != nil {
        return nil, err
    }

    return &dbPost, nil
}


func (p *PostsDBRepository) DeletePost(postId int64) error {
    // Get the post from the repository
    dbPost, ok := p.localPost[postId]
    if !ok {
        return fmt.Errorf("post not found for postId %d", postId)
    }
    fmt.Println(dbPost)
    // Delete the post from the repository
    delete(p.localPost, postId)

    // Return a success message
    return nil
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


