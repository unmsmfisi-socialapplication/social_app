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

func (p *PostsDBRepository) CreatePost(post domain.PostCreate) (*domain.Post, error) {

	query := `INSERT INTO sa.soc_app_posts
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
				WHERE public = true
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

	dbPosts.TotalCount = len(dbPosts.Posts)
	dbPosts.CurrentPage = params.Page

	return &dbPosts, nil
}

func (p *PostsDBRepository) GetById(id int) (*domain.Post, error) {
	query := `
        SELECT post_id, user_id, title, description, has_multimedia, public, multimedia, insertion_date, update_date 
        FROM sa.soc_app_posts
        WHERE post_id = $1
    `

	dbPost := domain.Post{}
	err := p.db.QueryRow(query, id).Scan(
		&dbPost.Id, &dbPost.UserId, &dbPost.Title, &dbPost.Description, &dbPost.HasMultimedia,
		&dbPost.Public, &dbPost.Multimedia, &dbPost.InsertionDate, &dbPost.UpdateDate,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &dbPost, nil
}
