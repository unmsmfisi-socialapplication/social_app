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

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) error {

	// Update according to the name and lastname,
	// TODO: change to user_id
	query := `UPDATE SOC_APP_USER_PROFILE SET name = $1, last_name = $2, about_me = $3, profile_picture = $4, update_date = NOW()
			WHERE name = $5 and last_name = $6`

	tx, err := pr.db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = tx.Exec(query, profile.Name, profile.LastName, profile.AboutMe, profile.ProfilePicture, profile.Name, profile.LastName)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update profile: %v", err)
	}

	tx.Commit()

	return nil
}

func (pr *ProfileRepository) CreateProfile(profile *domain.Profile) error {

	query := `INSERT INTO SOC_APP_USER_PROFILE (user_id, birth_date, name, last_name, about_me, genre, address, country, city, insertion_date, update_date, profile_picture)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW(), $10)`

	tx, err := pr.db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = tx.Exec(query, profile.UserID, profile.BirthDate, profile.Name, profile.LastName, profile.AboutMe, profile.Genre, profile.Address, profile.Country, profile.City, profile.ProfilePicture)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not create profile: %v", err)
	}

	tx.Commit()

	return nil

}
