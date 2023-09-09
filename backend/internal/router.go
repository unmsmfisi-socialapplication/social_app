package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func Router() http.Handler {
	r := chi.NewRouter()
    
    err := database.InitDatabase()
    if err != nil {
        panic(err)
    }

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

    db := database.GetDB()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Simulates some hard work.
		//
		// We want this handler to complete successfully during a shutdown signal,
		// so consider the work here as some background routine to fetch a long running
		// search query to find as many results as possible, but, instead we cut it short
		// and respond with what we have so far. How a shutdown is handled is entirely
		// up to the developer, as some code blocks are preemptible, and others are not.
		time.Sleep(5 * time.Second)

		w.Write([]byte(fmt.Sprintf("{\"response\": \"all done slow\"}")))
	})

    importProfileRepository := infrastructure.NewProfileRepository(db)
    importProfileUseCase := application.NewImportProfileUseCase(importProfileRepository)
    importProfileHandler := infrastructure.NewImportProfileHandler(importProfileUseCase)

    r.Put("/import-profile", importProfileHandler.ImportProfile)

	return r
}
