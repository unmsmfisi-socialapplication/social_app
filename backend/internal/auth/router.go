package infrastructure

import (
	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/infrastructure"
)

func AuthModuleRouter(useCase application.AuthUseCaseInterface) *chi.Mux {
	router := chi.NewRouter()

	tokenMiddleware := infrastructure.NewTokenMiddleware(useCase)
	router.Use(tokenMiddleware.Middleware)

	authHandler := infrastructure.NewAuthHandler(useCase)

	router.Get("/logout", authHandler.LogoutUser)

	return router
}
