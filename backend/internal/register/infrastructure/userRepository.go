package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT email, user_name, password FROM soc_app_users WHERE email = $1`

	row := u.db.QueryRow(query, email)
	prueba, _ := u.db.Exec(query, email)
	fmt.Println(prueba.RowsAffected())
	var user domain.User
	err := row.Scan(&user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) InsertUser(newUser *domain.User) (*domain.User, error) {
	query := `INSERT INTO soc_app_users ( insertion_date, email, user_name, password) VALUES (NOW(), $1, $2, $3)`

	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	_, err = tx.Exec(query, newUser.Email, newUser.Username, newUser.Password)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return newUser, nil
}
