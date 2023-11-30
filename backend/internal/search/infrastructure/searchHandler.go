package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	profiledomain "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/search/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type SearchHandler struct {
	SearchProfileUseCase *application.SearchProfileUseCase
}

func NewSearchHandler(searchProfileUseCase *application.SearchProfileUseCase) *SearchHandler {
	return &SearchHandler{SearchProfileUseCase: searchProfileUseCase}
}

func (handler *SearchHandler) SearchProfileByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	var page, limit int

	spage := r.URL.Query().Get("page")
	slimit := r.URL.Query().Get("limit")

	if spage == "" {
		spage = "1"
	}

	if slimit == "" {
		slimit = "10"
	}

	page, err := strconv.Atoi(spage)
	if err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "Error: page must be a number", nil)
		return
	}

	limit, err = strconv.Atoi(slimit)
	if err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "Error: page must be a number", nil)
		return
	}

	result := handler.SearchProfileUseCase.SearchProfileByName(query, page, limit)

	next := "/search/profiles?query=" + query + "&page=" + strconv.Itoa(page+1) + "&limit=" + strconv.Itoa(limit)

	var previous string
	if page > 1 {
		previous = "/search/profiles?query=" + query + "&page=" + strconv.Itoa(page-1) + "&limit=" + strconv.Itoa(limit)
	} else {
		previous = ""
	}

	response := struct {
		Results  []profiledomain.Profile `json:"results"`
		Page     int                     `json:"page"`
		Next     string                  `json:"next"`
		Previous string                  `json:"previous"`
	}{Results: result.Results, Page: page, Next: next, Previous: previous}
    
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (handler *SearchHandler) SuggestProfileByName(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")

	result := handler.SearchProfileUseCase.SuggestProfileByName(keyword)
	response := struct {
		Results []profiledomain.Profile `json:"results"`
	}{Results: result.Results}

    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
