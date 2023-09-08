package application

import (
    "database/sql"
)

type InterestService struct {
    db *sql.DB
}

//Create the instance initializated with a connection to the database
func SetInterestService(db *sql.DB) *InterestService {
    return &InterestService{db}
}
