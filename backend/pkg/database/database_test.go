package database_test

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func InitDatabaseTest(t *testing.T) {
	t.Log("InitDatabaseTest")

	err := database.InitDatabase()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Log("InitDatabaseTest: Success")
}

func GetDBTest(t *testing.T) {
	t.Log("GetDBTest")

	err := database.InitDatabase()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	db := database.GetDB()
	if db == nil {
		t.Errorf("Error: %v", "DB is nil")
	}

	t.Log("GetDBTest: Success")
}
