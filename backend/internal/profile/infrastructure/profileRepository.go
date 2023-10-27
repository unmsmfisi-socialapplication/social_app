package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ProfileRepository struct {
    db *sql.DB
}

func NewProfileRepository(database *sql.DB) *ProfileRepository {
    return &ProfileRepository{db: database}
}

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) (error) {
    query := `UPDATE sa.SOC_APP_USER_PROFILE SET name = $1, profile_picture = $2, about_me = $3, update_date = NOW() WHERE name = $4`

    tx, err := pr.db.Begin()
    if err != nil {
        return fmt.Errorf("could not begin transaction: %v", err)
    }

    _, err = tx.Exec(query, profile.Username, profile.ProfileImage, profile.Biography, profile.Username)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("could not update profile: %v", err)
    }

    tx.Commit()

	return nil
}
