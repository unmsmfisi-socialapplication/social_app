package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type FollowerRepository struct {
	db *sql.DB
}

func NewFollowerRepository(database *sql.DB) *FollowerRepository {
	return &FollowerRepository{db: database}
}

func (u *FollowerRepository) InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error) {
	schema := os.Getenv("DB_SCHEMA")
	query := fmt.Sprintf(`INSERT INTO %s.soc_app_profile_followers (follow_date, follower_profile_id , following_profile_id ) VALUES (NOW(), $1, $2)`, schema)

	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	_, err = tx.Exec(query, newFollower.Follower_profile_id, newFollower.Following_profile_id)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return newFollower, nil
}
