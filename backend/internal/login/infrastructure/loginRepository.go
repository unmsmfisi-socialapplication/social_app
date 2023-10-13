package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/domain"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(database *sql.DB) application.UserRepository {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserByUsername(username string) (*domain.User, error) {

	return &domain.User{
		Username: "myuser12",
		Password: "$2y$10$Fii1qaKs9l77Omi1S0EZsOhwLS6eIWIDKX23JW36PpP4OhNl88/8u", // this is gonna be in https://bcrypt.online/ for: "Social@12" (without quotes)
	}, nil
}
