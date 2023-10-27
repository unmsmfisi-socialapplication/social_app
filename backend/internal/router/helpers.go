package router

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/cors"
)

func getCorsOrigins() []string {
	value := os.Getenv("CORS_ORIGINS")
	return strings.Split(value, ",")
}

func getCorsMaxAge() int {
	maxAge, err := strconv.Atoi(os.Getenv("CORS_MAXAGE"))

	if err != nil {
		return 300
	}

	return maxAge
}

func configCorsMiddleware() func(http.Handler) http.Handler {
	origins := getCorsOrigins()
	maxAge := getCorsMaxAge()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	})
	return corsMiddleware.Handler
}
