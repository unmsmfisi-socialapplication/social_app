package infrastructure

import (
	"database/sql"

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

	query := `INSERT INTO soc_app_posts
	(user_id,title, description, has_multimedia, public, multimedia, insertion_date, update_date)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING post_id;	
	`

	dbPost := domain.CreatePostToPost(post)

	err := p.db.QueryRow(
		query,
		dbPost.BasePost.UserId,
		dbPost.BasePost.Description,
		dbPost.BasePost.Title,
		dbPost.BasePost.HasMultimedia,
		dbPost.BasePost.Public,
		dbPost.BasePost.Multimedia,
		dbPost.InsertionDate,
		dbPost.UpdateDate,
	).Scan(&dbPost.Id)

	if err != nil {
		return nil, err
	}

	return &dbPost, nil
}

func (p *PostsDBRepository) UserExist(post domain.CreatePost) bool {
	query := `SELECT user_id FROM sa.soc_app_users WHERE user_id =  $1`

	var dbUserId int64

	err := p.db.QueryRow(query, post.UserId).Scan(&dbUserId)

	if err != nil {		
		return false
	}

	return dbUserId != 0
}
