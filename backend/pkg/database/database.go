package database

import (
	"database/sql"
	
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func InitDatabase() *Database {
	
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/social_app?sslmode=disable")
	log.Println("Conectando a la base de datos...")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	
	
	
	return &Database{db}
}


func(db *Database) Close() error {
	return db.DB.Close()
}

