package infrastructure

import (
	"database/sql"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type FollowerRepository struct {
	db *sql.DB
}

func NewFollowerRepository(database *sql.DB) *FollowerRepository {
	return &FollowerRepository{db: database}
}

func (u *FollowerRepository) InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error) {
	query := `INSERT INTO SOC_APP_USER_PROFILE_FOLLOW (follow_date, follower_profile_id , following_profile_id ) VALUES (NOW(), $1, $2)`
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	_, err = tx.Exec(query, newFollower.Follower_profile_id, newFollower.Following_profile_id)
	//fmt.Print(err)
	if err != nil {
		return nil, err
	} else {
		tx.Commit()
		return newFollower, nil
	}

}

func (u *FollowerRepository) IsFollowing(newFollower *domain.Follower) (*bool, error) {
	query := `
		SELECT COUNT(1) is_following
		FROM soc_app_user_profile_follow
		WHERE follower_profile_id = $1 AND following_profile_id = $2
	`
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	rows, err := tx.Query(query, newFollower.Follower_profile_id, newFollower.Following_profile_id)
	if err != nil {
		log.Println("Error while executing the query")
		return nil, err
	}
	defer rows.Close()

	var isFollowing bool
	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			log.Println("Error while scanning the rows")
			return nil, err
		}
		isFollowing = (count == 1)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error while retrieving the rows")
		return nil, err
	}

	return &isFollowing, nil
}
