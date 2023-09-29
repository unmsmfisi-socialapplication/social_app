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

	query := `SELECT * FROM create_post(
		p_user_id := $1,
		p_title := $2,
		p_description := $3,
		p_has_multimedia := $4,
		p_public := $5,
		p_multimedia := $6
	);`

	row := p.db.QueryRow(query, post.UserId, post.Title, post.Description, post.HasMultimedia, post.Public, post.Multimedia)

	var dbPost domain.Post
	err := row.Scan(&dbPost.Id, &dbPost.UserId, &dbPost.Title, &dbPost.Description, &dbPost.HasMultimedia, &dbPost.Public, &dbPost.Multimedia, &dbPost.InsertionDate, &dbPost.UpdateDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &dbPost, nil
}
