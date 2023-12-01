package infrastructure

import (
	"context"
	"net/http"
	"strings"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type TokenMiddleware struct {
	useCase application.AuthUseCaseInterface
}

func NewTokenMiddleware(useCase application.AuthUseCaseInterface) *TokenMiddleware {
	return &TokenMiddleware{useCase: useCase}
}

func (m *TokenMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "No token provided")
			return
		}

		user, err := m.useCase.ValidateToken(tokenString)
		if err != nil {
			switch err {
			case application.ErrInvalidToken:
				utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Invalid token")
				return
			case application.ErrTokenExpired:
				utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Token has expired")
				return
			case application.ErrUnauthorized:
				utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Unauthorized access")
				return
			default:
				utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during token validation")
				return
			}
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
