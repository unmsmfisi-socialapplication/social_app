package profile

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure"
)

func ProfileModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	importProfileRepository := infrastructure.NewProfileRepository(dbInstance)
	importProfileUseCase := application.NewImportProfileUseCase(importProfileRepository)
	importProfileHandler := infrastructure.NewImportProfileHandler(importProfileUseCase)

	r.Put("/import", importProfileHandler.ImportProfile)

	return r
}
