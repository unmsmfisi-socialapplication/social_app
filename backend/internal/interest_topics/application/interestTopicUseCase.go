package application

import (
	"database/sql"
)

type InterestTopics struct {
	db *sql.DB
}

// Create the instance initializated with a connection to the database
func SetInterestTopics(db *sql.DB) *InterestTopics {
	return &InterestTopics{db}
}

//Implementation