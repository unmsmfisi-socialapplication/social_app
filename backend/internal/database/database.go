package database
import (

	"database/sql"
	
	_ "github.com/lib/pq"
	

)

var db *sql.DB

func InitDatabase() error {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:root@localhost:5432/social_app?sslmode=disable")

	return err
}

func GetDB() *sql.DB {
	return db
}

