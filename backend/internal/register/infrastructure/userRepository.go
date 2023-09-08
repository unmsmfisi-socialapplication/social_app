package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(database *sql.DB) (*UserDBRepository) {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT username, password FROM users WHERE email = $1`

	row := u.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) RegisterUser(username, email, password, name, lastName, birthday string) (*domain.User, error) {

	existingUser, err := u.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("el correo electrónico ya está en uso")
	}


	newUser, err := domain.NewUser(username, email, password, name, lastName, birthday) // Utilizamos el correo electrónico como identificador
	if err != nil {
		return nil, err
	}


	query := `INSERT INTO users (username, email, password, name, lastName, birthday) VALUES (?, ? ,? ,? ,? ,?)`
	
	tx, _:= u.db.Begin()
	res, err := tx.Exec(query, newUser.Username, newUser.Email, newUser.Password, newUser.Name, newUser.LastName, newUser.Birthday) // Usamos el correo como username
	fmt.Println(res)
	_=tx.Commit()

	if err != nil {
		return nil, err
	}

	return newUser, nil
}