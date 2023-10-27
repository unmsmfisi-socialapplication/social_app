package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	config "github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

var db *sql.DB

func InitDatabase() error {
	db_config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	database, err := sql.Open("postgres", db_config.DBConnectionString)
	if err != nil {
		return err
	}

	if err = database.Ping(); err != nil {
		return err
	}

	log.Print("successful database connection")

	db = database
	return nil
}

func GetDB() *sql.DB {
	return db
}

type UserDBRepository struct {
	db *sql.DB
}
