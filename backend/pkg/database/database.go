package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

var db *sql.DB

func InitDatabase() error {
	var err error
	db, err = sql.Open("postgres", utils.LoadConfig().DBConnectionString)

	return err
}

func GetDB() *sql.DB {
	return db
}

type UserDBRepository struct {
	db *sql.DB
}
