package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	config "github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

var db *sql.DB

func InitDatabase() error {
	var err error
	db_config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	db, err = sql.Open("postgres", db_config.DBConnectionString)

	return err
}

func GetDB() *sql.DB {
	return db
}

type UserDBRepository struct {
	db *sql.DB
}
