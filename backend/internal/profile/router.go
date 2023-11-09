package profile

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure"
)

func ProfileModuleRouter(dbInstance *sql.DB) *chi.Mux {
    r := chi.NewRouter()

	importProfileRepository := infrastructure.NewProfileRepository(dbInstance)
	importProfileUseCase := application.NewImportProfileUseCase(importProfileRepository)
	importProfileHandler := infrastructure.NewImportProfileHandler(importProfileUseCase)

	createProfileRepository := infrastructure.NewProfileRepository(dbInstance)
	createProfileUseCase := application.NewCreateProfileUseCase(createProfileRepository)
	createProfileHandler := infrastructure.NewCreateProfileHandler(createProfileUseCase)
	

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        // GET Profile
        w.Write([]byte("Profile"))
    })

	r.Put("/import", importProfileHandler.ImportProfile)
	r.Post("/create", createProfileHandler.CreateProfile)

    return r
}
