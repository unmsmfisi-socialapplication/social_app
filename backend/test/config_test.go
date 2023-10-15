package test

import (
	"os"
	"testing"

	config "github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

func TestLoadEnvFromFile(t *testing.T) {
	tempFile := "test.env"
	content := []byte("DB_HOST=localhost\nDB_USER=testuser\nDB_PASSWORD=testpass\nDB_NAME=testdb\nDB_SSLMODE=disable")
	err := os.WriteFile(tempFile, content, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile)

	config.LoadEnvFromFile(tempFile)

	expectedEnv := map[string]string{
		"DB_HOST":     "localhost",
		"DB_USER":     "testuser",
		"DB_PASSWORD": "testpass",
		"DB_NAME":     "testdb",
		"DB_SSLMODE":  "disable",
	}

	for key, expectedValue := range expectedEnv {
		actualValue := os.Getenv(key)
		if actualValue != expectedValue {
			t.Errorf("Variable de entorno %s: se esperaba %s pero se obtuvo %s", key, expectedValue, actualValue)
		}
	}
}

func TestCheckEnvVariables(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("DB_SCHEMA", "sa")
	os.Setenv("CORS_ORIGINS", "[http://localhost:3000,https://social-app-frontend-service-7bg5siys2q-uc.a.run.app/]")
	os.Setenv("CORS_MAXAGE", "300")

	err := config.CheckEnvVariables()
	if err != nil {
		t.Errorf("Se esperaba que todas las variables de entorno estuvieran configuradas, pero se obtuvo un error: %v", err)
	}
}
