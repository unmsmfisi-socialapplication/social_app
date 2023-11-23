package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/search/application"
)

type SearchHandler struct {
	SearchProfileUseCase *application.SearchProfileUseCase
}

func NewSearchHandler(searchProfileUseCase *application.SearchProfileUseCase) *SearchHandler {
	return &SearchHandler{SearchProfileUseCase: searchProfileUseCase}
}

func (handler *SearchHandler) SearchProfileByName(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    result := handler.SearchProfileUseCase.SearchProfileByName(query)

    json.NewEncoder(w).Encode(result)
}

func (handler *SearchHandler) SuggestProfileByName(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")

    result := handler.SearchProfileUseCase.SuggestProfileByName(keyword)
    json.NewEncoder(w).Encode(result)
}
