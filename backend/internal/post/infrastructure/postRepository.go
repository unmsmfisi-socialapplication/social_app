package infrastructure

import (
	"database/sql"
	"log"
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

func (p *PostsDBRepository) CreateHomeTimelineItem(user_id int64, post_id int64) error {
	query := `INSERT INTO SOC_APP_HOME_TIMELINE (user_id, post_id)
				VALUES ($1, $2)
				RETURNING id, created_at;`

	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction of TIMELINE type")
		return err
	}

	_, err = tx.Exec(query, user_id, post_id)
	if err != nil {
		return err
	} else {
		tx.Commit()
	}

	return err
}

func (p *PostsDBRepository) FanoutHomeTimeline(posts_id int64, following_id int64) error {
	query := `INSERT INTO SOC_APP_HOME_TIMELINE (user_id, post_id)
				SELECT 
					SOC_APP_USER_PROFILE_FOLLOW.follower_profile_id, 
					$1
				FROM SOC_APP_USER_PROFILE_FOLLOW
				WHERE SOC_APP_USER_PROFILE_FOLLOW.following_profile_id = $2
					ON CONFLICT (user_id, post_id) DO NOTHING
				RETURNING id, created_at;`

	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction of FanOutTimeline type")
		return err
	}

	_, err = tx.Exec(query, posts_id, following_id)
	if err != nil {
		return err
	} else {
		tx.Commit()
	}

	return err
}

func (p *PostsDBRepository) HomeTimeline(user_id int64) (*[]domain.TimelineRes, error) {
	query := `SELECT 
				SOC_APP_POSTS.post_id,
				SOC_APP_POSTS.user_id,
				SOC_APP_POSTS.title,
				SOC_APP_POSTS.description,
				SOC_APP_POSTS.multimedia,
				SOC_APP_POSTS.insertion_date,
				SOC_APP_USERS.user_name
			FROM SOC_APP_HOME_TIMELINE
				INNER JOIN SOC_APP_POSTS ON SOC_APP_HOME_TIMELINE.post_id = SOC_APP_POSTS.post_id
				INNER JOIN SOC_APP_USERS ON SOC_APP_POSTS.user_id = SOC_APP_USERS.user_id
			WHERE SOC_APP_HOME_TIMELINE.user_id = $1
			ORDER BY SOC_APP_HOME_TIMELINE.id DESC;`

	rows, err := p.db.Query(query, user_id)
	if err != nil {
		return nil, err
	}

	var timeline []domain.TimelineRes

	for rows.Next() {
		var t domain.TimelineRes
		err := rows.Scan(
			&t.UserId,
			&t.PostId,
			&t.Title,
			&t.Description,
			&t.Multimedia,
			&t.InsertionDate,
			&t.Username,
		)
		if err != nil {
			return nil, err
		}
		timeline = append(timeline, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &timeline, nil
}
