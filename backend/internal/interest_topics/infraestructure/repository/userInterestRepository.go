package infraestructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

type UserInterestsDBRepository struct {
	db *sql.DB
}

func NewUserInterestsDBRepository(db *sql.DB) domain.UserInterestsRepository {
	return &UserInterestsDBRepository{db: db}
}

func (ui *UserInterestsDBRepository) Create(interest *entity.UserInterestTopics) error {
	query := `INSERT INTO soc_app_users_interest_topics (user_id, interest_id) VALUES ($1,$2)`

	_, err := ui.db.Exec(query, interest.User_id, interest.Interest_id)
	if err != nil {
		return err
	}

	return nil
}
