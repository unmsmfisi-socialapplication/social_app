package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	application_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/import"
	infrastructure_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/import"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func ProfileHandler(r chi.Router) {
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	db := database.GetDB()
	importProfileRepository := NewProfileRepository(db)

	importProfileUseCase := application_import.NewImportProfileUseCase(importProfileRepository)
	importProfileHandler := infrastructure_import.NewImportProfileHandler(importProfileUseCase)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        // GET Profile
        w.Write([]byte("Profile"))
    })

	r.Put("/import", importProfileHandler.ImportProfile)
}
