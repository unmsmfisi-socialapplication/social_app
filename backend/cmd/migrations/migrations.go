package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func main() {

	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Re-create the database and apply migrations
	fmt.Println("Creating a new tables and applying migrations...")
	err = m.Up()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Tables created and migrations applied successfully.")

	version, _, err := m.Version()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Current migration version: %d\n", version)
}
