package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBConnectionString string
	AppPort            string
}

func loadEnvFromFile(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	env := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	for key, value := range env {
		os.Setenv(key, value)
	}
}

func checkEnvVariables() error {
	requiredVariables := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_DBNAME", "DB_SSLMODE"}

	for _, variable := range requiredVariables {
		fmt.Println(os.Getenv(variable))
		if os.Getenv(variable) == "" {
			return fmt.Errorf("Environment variable %s is not set", variable)
		}
	}

	return nil
}

func LoadConfig() *Config {

	loadEnvFromFile(".env")

	err := checkEnvVariables()
	if err != nil {
		fmt.Println("Environment variables are incorrectly set")
		return nil
	}

	return &Config{
		DBConnectionString: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		),
	}
}
