package database

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

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

func NewGormInstance() (*gorm.DB, error) {

	dbConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dbConfig.DBConnectionString), &gorm.Config{})

	if err != nil {
		panic("Cant connect to prod database: " + err.Error())
	}

	log.Println("Succesful conection to prod database")

	return db, nil
}
