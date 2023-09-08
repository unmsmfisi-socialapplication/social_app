package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
	config "github.com/unmsmfisi-socialapplication/social_app"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/domain"
)

var db *sql.DB

func InitDatabase() error {
	var err error
	db, err = sql.Open("postgres", config.LoadConfig().DBConnectionString)

	return err
}

func GetDB() *sql.DB {
	return db
}

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(database *sql.DB) application.UserRepository {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserByUsername(username string) (*domain.User, error) {
	query := `SELECT username, password FROM users WHERE username = $1`

	row := u.db.QueryRow(query, username)

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
