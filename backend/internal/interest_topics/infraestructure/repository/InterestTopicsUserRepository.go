package infraestructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

type InterestTopicsUserRepository struct {
	db *sql.DB
}

func NewInterestTopicsUserRepository(db *sql.DB) *InterestTopicsUserRepository {
	return &InterestTopicsUserRepository{db}
}

func (itr *InterestTopicsUserRepository) Create(interest *entity.Interest_Topic_Users) error {
	//Implementation
	return nil
}
