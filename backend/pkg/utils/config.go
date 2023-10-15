package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	env := make(map[string]string)

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
			}

			env[key] = value
		}
	}

	for key, value := range env {
		os.Setenv(key, value)
	}
}

func GetCorsOrigins() []string {
	value := os.Getenv("CORS_ORIGINS")
	return strings.Split(value, ",")
}

func GetCorsMaxAge() int {
	maxAge, err := strconv.Atoi(os.Getenv("CORS_MAXAGE"))

	if err != nil {
		return 300
	}

	return maxAge
}

func CheckEnvVariables() error {
	requiredVariables := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE", "DB_SCHEMA", "CORS_ORIGINS", "CORS_MAXAGE"}

	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			return fmt.Errorf("environment variable %s is not set", variable)
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

func ConfigCors(router *chi.Mux) {
	origins := GetCorsOrigins()
	maxAge := GetCorsMaxAge()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	})
	router.Use(corsMiddleware.Handler)
}
