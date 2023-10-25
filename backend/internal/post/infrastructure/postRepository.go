package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

type PostsDBRepository struct {
	db *sql.DB
}

func NewPostDBRepository(database *sql.DB) application.PostRepository {
	return &PostsDBRepository{db: database}
}

func (p *PostsDBRepository) CreatePost(post domain.CreatePost) (*domain.Post, error) {

	if !p.ValidateUserExist(post) {
		return nil, fmt.Errorf("user not exist %d", post.UserId)
	}

	query := `INSERT INTO soc_app_posts
	(user_id,title, description, has_multimedia, public, multimedia, insertion_date, update_date)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING post_id;
	`
	dbPost := domain.CreatePostToPost(post)

	err := p.db.QueryRow(
		query,
		dbPost.UserId,
		dbPost.Title,
		dbPost.Description,
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

func (p *PostsDBRepository) ValidateUserExist(post domain.CreatePost) bool {
	query := `SELECT user_id FROM sa.soc_app_users WHERE user_id =  $1`

	var dbUserId int64

	err := p.db.QueryRow(query, post.UserId).Scan(&dbUserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("error find user: %v", err)
		return false
	}

	return dbUserId != 0
}
