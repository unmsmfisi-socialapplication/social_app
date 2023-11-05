package database_test

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func InitDatabaseTest(t *testing.T) {
	t.Log("InitDatabaseTest")

	db := database.InitDatabase()
	if db == nil {
		t.Errorf("Error: %v", errors.New("Can't connect to the database") )
	}

	t.Log("InitDatabaseTest: Success")
}

