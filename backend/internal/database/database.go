package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/unmsmfisi-socialapplication/social_app/configs"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	if err := configs.LoadEnv("configs/.env"); err != nil {
		log.Fatal(err)
	}

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Connection string
    dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    // Open connection
    db, err := sql.Open("postgres", dbInfo)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Access to database
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    fmt.Println("Successfully connected to database!")
    return db, nil
}
