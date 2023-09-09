package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() (err error) {
    db, err = sql.Open("postgres", "postgres://postgres:luiggi_123@localhost:5432/social_app?sslmode=disable")
	return err
}

func GetDB() *sql.DB {
	return db
}
