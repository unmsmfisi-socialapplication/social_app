package register

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/infrastructure"
)

func RegisterModule(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	registerRepository := infrastructure.NewUserRepository(dbInstance)
	registerUseCase := application.NewRegistrationUseCase(registerRepository)
	registerHandler := infrastructure.NewRegisterUserHandler(registerUseCase)

	r.Post("/", registerHandler.RegisterUser)

	return r
}
