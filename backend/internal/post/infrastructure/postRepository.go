package infrastructure

import (
	"database/sql"
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
	query := `SELECT user_id FROM soc_app_users WHERE user_id =  $1`

	var dbUserId int64

	p.db.QueryRow(query, post.UserId).Scan(&dbUserId)

	return dbUserId != 0
}

func (p *PostsDBRepository) GetAll(params domain.PostPaginationParams) (*domain.PostPagination, error) {
	query := `SELECT post_id ,user_id, title, description, has_multimedia, 
				public, multimedia, insertion_date, update_date 
				FROM soc_app_posts
				WHERE public = true  AND is_deleted = FALSE
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
        FROM soc_app_posts
        WHERE post_id = $1 AND is_deleted = FALSE
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

// DELETE POST //

func (p *PostsDBRepository) DeletePost(id int64) error {
	query := `UPDATE soc_app_posts SET is_deleted = TRUE WHERE post_id = $1;`
	_, err := p.db.Exec(query, id)
	return err
}

func (p *PostsDBRepository) PostExists(id int64) bool {
	query := `SELECT EXISTS(SELECT 1 FROM soc_app_posts WHERE post_id = $1 AND is_deleted = FALSE)  ;`
	var exists bool
	err := p.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// UPDATE POST //
func (p *PostsDBRepository) UpdatePost(id int64, update domain.PostUpdate) error {
	query := `UPDATE soc_app_posts SET title=$2, description=$3, has_multimedia=$4, public=$5, multimedia=$6, update_date=NOW() WHERE post_id=$1 ;`
	_, err := p.db.Exec(query, id, update.Title, update.Description, update.HasMultimedia, update.Public, update.Multimedia)
	if err != nil {
		return err
	}
	return nil
}

// REPORT POST //

func (p *PostsDBRepository) CreateReport(report domain.PostReport) error {
	query := `INSERT INTO soc_app_post_reports (post_id, reported_by, reason, report_date) VALUES ($1, $2, $3, NOW());`
	_, err := p.db.Exec(query, report.PostId, report.ReportedBy, report.Reason)
	return err
}

func (p *PostsDBRepository) ReporterExists(username string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM soc_app_users WHERE user_name = $1);`
	var exists bool
	err := p.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
func (p *PostsDBRepository) HasUserAlreadyReportedPost(postId int64, username string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM soc_app_post_reports WHERE post_id = $1 AND reported_by = $2);`
	var exists bool
	err := p.db.QueryRow(query, postId, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
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
