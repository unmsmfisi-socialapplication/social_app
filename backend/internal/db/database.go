package infrastructure

import (
    "database/sql"
	config "github.com/unmsmfisi-socialapplication/social_app"
)
func InitDatabase() *sql.DB {
	//Return a connection to the database
	db, err := sql.Open("postgres", config.LoadConfig().DBConnectionString)
    if err != nil {
        panic(err)
    }
    return db
}