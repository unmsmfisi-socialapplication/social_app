package infrastructure

import (
	"database/sql"
	"time"

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
	query := `SELECT user_name, password,role FROM soc_app_users WHERE user_name = $1`
	err := u.db.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) StoreJTIForSession(username string, jti string) error {
	currentTime := time.Now()
	query := `UPDATE SOC_APP_auth_sessions SET timestamp = $2, jti = $3 WHERE user_name = $1 and logged = true`
	_, err := u.db.Exec(query, username, currentTime, jti)
	return err
}

func (u *UserDBRepository) InsertSession(username string) error {
	query := `INSERT INTO soc_app_auth_sessions (user_name, logged) VALUES ($1, true)`
	_, err := u.db.Exec(query, username)
	return err
}
func (u *UserDBRepository) CheckExistingSession(username string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM soc_app_auth_sessions WHERE user_name = $1)`
	var exists bool
	err := u.db.QueryRow(query, username).Scan(&exists)
	return exists, err
}

func (u *UserDBRepository) UpdateSession(username string) error {
	query := `UPDATE soc_app_auth_sessions SET logged = true, timestamp = CURRENT_TIMESTAMP WHERE user_name = $1`
	_, err := u.db.Exec(query, username)
	return err
}
