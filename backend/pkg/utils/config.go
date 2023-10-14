package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadEnvFromFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}

	defer file.Close()

	env := make(map[string]interface{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			if strings.Contains(value, "[") {

				value = strings.ReplaceAll(value, "[", "")
				value = strings.ReplaceAll(value, "]", "")
				env[key] = strings.Split(value, ",")

			} else {
				env[key] = value
			}
		}
	}

	for key, value := range env {
		switch val := value.(type) {
		case string:
			os.Setenv(key, val)
		case []string:

		}

	}
}

func CheckEnvVariables() error {
	requiredVariables := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE", "DB_SCHEMA", "ALLOW_ORIGINS"}

	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			return fmt.Errorf("Environment variable %s is not set", variable)
		}
	}

	return nil
}

func LoadConfig() (*Config, error) {

	LoadEnvFromFile(".env")

	err := CheckEnvVariables()
	if err != nil {
		log.Fatal("Environment variables are incorrectly set")
		return nil, err
	}

	return &Config{
		DBConnectionString: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
			os.Getenv("DB_SCHEMA"),
		),
	}, nil
}
