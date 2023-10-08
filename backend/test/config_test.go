package test

import (
	"os"
	"testing"

	config "github.com/unmsmfisi-socialapplication/social_app"
)

func TestLoadEnvFromFile(t *testing.T) {
	tempFile := "test.env"
	content := []byte("DB_HOST=localhost\nDB_USER=testuser\nDB_PASSWORD=testpass\nDB_DBNAME=testdb\nDB_SSLMODE=disable")
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
		"DB_DBNAME":   "testdb",
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
	os.Setenv("DB_DBNAME", "testdb")
	os.Setenv("DB_SSLMODE", "disable")

	err := config.CheckEnvVariables()
	if err != nil {
		t.Errorf("Se esperaba que todas las variables de entorno estuvieran configuradas, pero se obtuvo un error: %v", err)
	}

	os.Unsetenv("DB_HOST")

	err = config.CheckEnvVariables()
	if err == nil {
		t.Error("Se esperaba un error ya que falta una variable de entorno requerida, pero no se obtuvo ningún error")
	}
}
