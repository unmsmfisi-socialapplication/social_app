package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfig() *Config {
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
