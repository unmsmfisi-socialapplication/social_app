package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT user_id,email,user_name,phone FROM soc_app_users WHERE email = $1`

	row := u.db.QueryRow(query, email)
	_, err := u.db.Exec(query, email)

	if err != nil {
		return nil, err
	}

	var user domain.User

	err = row.Scan(&user.Id, &user.Email, &user.Username, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) InsertUser(newUser *domain.User) (*domain.User, error) {
	query := `INSERT INTO soc_app_users ( insertion_date, email, user_name, password) VALUES (NOW(), $1, $2, $3)
	RETURNING user_id`

	err := u.db.QueryRow(query,
		newUser.Email,
		newUser.Username,
		newUser.Password,
	).Scan(&newUser.Id)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
