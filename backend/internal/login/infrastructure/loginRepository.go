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
		Username: "my_user",
		Password: "$2y$10$JFvY.pJMUgQqPYYGvPWIBOdSvrfLKpvDTbYKCU4VCt1SRA3wdY7Iq", // this is gonna be in https://bcrypt.online/ for: "hopefully this is not a plain text password" (without quotes)
	}, nil
}
