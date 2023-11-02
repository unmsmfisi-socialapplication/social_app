package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(database *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: database}
}

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) error {

	// Update according to the name and lastname,
	// TODO: change to user_id
	query := fmt.Sprintf(`UPDATE %s.SOC_APP_USER_PROFILE SET name = $1, last_name = $2, about_me = $3, profile_picture = $4, update_date = NOW()
			WHERE name = $5 and last_name = $6`, os.Getenv("DB_SCHEMA"))

	tx, err := pr.db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = tx.Exec(query, profile.Name, profile.Lastname, profile.Biography, profile.ProfileImage, profile.Name, profile.Lastname)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update profile: %v", err)
	}

	tx.Commit()

	return nil
}
