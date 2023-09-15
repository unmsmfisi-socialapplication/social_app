package infrastructure

import (
	"database/sql"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) (*domain.Profile, error) {

    tx, _ := pr.db.Begin()
    
    _, err := tx.Exec("UPDATE Profiles SET profile_picture = $1, biography = $2 WHERE username = $3", profile.ProfileImage, profile.Biography, profile.Username)
	if err != nil {
        log.Println(err)
		return nil, err
	}

    err = tx.Commit()
	if err != nil {
        log.Println(err)
		return nil, err
	}

	return profile, nil
}
