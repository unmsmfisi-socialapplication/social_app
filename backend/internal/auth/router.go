package auth

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
)

func AuthModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userRepository := infrastructure.NewUserRepository(dbInstance)

    loginUseCase := application.NewLoginUseCase(userRepository)
    loginHandler := infrastructure.NewLoginHandler(loginUseCase)
	r.Post("/login", loginHandler.HandleLogin)

    registerUseCase := application.NewRegistrationUseCase(userRepository)
    registerHandler := infrastructure.NewRegisterUserHandler(registerUseCase)
    r.Post("/register", registerHandler.RegisterUser)

	return r
}
