package search

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/search/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/search/infrastructure"
)

func SuggestModuleRouter(searchHandler *infrastructure.SearchHandler) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/profile", searchHandler.SuggestProfileByName)

    return r
}

func SearchModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

    searchRepository := infrastructure.NewSearchRepository(dbInstance)
    searchUseCase := application.NewSearchProfileUseCase(searchRepository)
    searchHandler := infrastructure.NewSearchHandler(searchUseCase)

    r.Get("/profile", searchHandler.SearchProfileByName)

    r.Mount("/suggest", SuggestModuleRouter(searchHandler))

	return r
}
