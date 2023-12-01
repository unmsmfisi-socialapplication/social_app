package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

type AuthSessionDBRepository struct {
	db *sql.DB
}

func NewAuthSessionDBRepository(database *sql.DB) domain.AuthSessionRepository {
	return &AuthSessionDBRepository{db: database}
}

func (a *AuthSessionDBRepository) CheckValidSession(username string) (bool, error) {
	query := `SELECT logged FROM SOC_APP_auth_sessions WHERE user_name = $1 ORDER BY timestamp DESC LIMIT 1`
	var logged bool
	err := a.db.QueryRow(query, username).Scan(&logged)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("No valid sessions for this user")
		}
		return false, err
	}
	return logged, nil
}

func (a *AuthSessionDBRepository) GetJTIByUsername(username string) (string, error) {
	query := `SELECT jti FROM SOC_APP_auth_sessions WHERE user_name = $1 and logged = true ORDER BY timestamp DESC LIMIT 1`
	var jti string
	err := a.db.QueryRow(query, username).Scan(&jti)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", application.ErrUnauthorized
		}
		return "", err
	}
	return jti, nil
}

func (a *AuthSessionDBRepository) InvalidateSession(username string) error {
	query := `UPDATE SOC_APP_auth_sessions SET logged = false, jti = NULL WHERE user_name = $1`
	_, err := a.db.Exec(query, username)
	return err
}
