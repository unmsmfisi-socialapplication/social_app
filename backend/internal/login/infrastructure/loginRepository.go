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
	var user domain.User
	query := `SELECT user_name, password FROM soc_app_users WHERE user_name = $1`
	err := u.db.QueryRow(query, username).Scan(&user.Username, &user.Password)
	if err != nil {
		return &domain.User{
			Username: "myuser12",
			Password: "$2y$10$Fii1qaKs9l77Omi1S0EZsOhwLS6eIWIDKX23JW36PpP4OhNl88/8u", // this is gonna be in https://bcrypt.online/ for: "Social@12" (without quotes)
		}, err
	}
	return &user, nil
}
