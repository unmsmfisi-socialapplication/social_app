package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT user_id, email, user_name, password,name,phone FROM soc_app_users WHERE email = $1`

	row := u.db.QueryRow(query, email)
	prueba, _ := u.db.Exec(query, email)
	fmt.Println(prueba.RowsAffected())
	var user domain.User
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Name, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) InsertUser(newUser *domain.User) (*domain.User, error) {
	query := `INSERT INTO soc_app_users (insertion_date, email, 
		user_name, password, name, phone, photo) 
	VALUES (NOW(), $1, $2, $3, $4, $5, $6)
	RETURNING user_id
	`

	err := u.db.QueryRow(
		query,
		newUser.Email,
		newUser.Username,
		newUser.Password,
		newUser.Name,
		newUser.Phone,
		newUser.Photo,
	).Scan(&newUser.Id)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
