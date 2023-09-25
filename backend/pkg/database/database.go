package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	config "github.com/unmsmfisi-socialapplication/social_app"
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
