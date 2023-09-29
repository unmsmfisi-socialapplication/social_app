package infrastructureroutes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/applicationimport"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/infrastructureimport"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/infrastructurerepository"
)

func ProfileHandler(r chi.Router) {
    // TODO: Add database

	importProfileRepository := infrastructurerepository.NewProfileRepository()
	importProfileUseCase := applicationimport.NewImportProfileUseCase(importProfileRepository)
	importProfileHandler := infrastructureimport.NewImportProfileHandler(importProfileUseCase)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        // GET Profile
        w.Write([]byte("Profile"))
    })

	r.Put("/import", importProfileHandler.ImportProfile)
}
