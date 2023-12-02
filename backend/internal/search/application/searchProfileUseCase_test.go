package application

import (
	"reflect"
	"testing"

	domainprofile "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/search/domain"
)

type mockSearchRepository struct{}

func (m *mockSearchRepository) GetProfilesByName(keyword string, psize, pnumber int) *domain.QueryResult {
	result := domain.NewQueryResult([]domainprofile.Profile{})
	return result
}

func (m *mockSearchRepository) GetsuggestionsProfiles(keyword string) *domain.QueryResult {
	result := domain.NewQueryResult([]domainprofile.Profile{})
	return result
}

func TestSearchProfileByName(t *testing.T) {
    useCase := NewSearchProfileUseCase(&mockSearchRepository{})

    result := useCase.SearchProfileByName("test", 10, 1)

    expectedResults := []domainprofile.Profile{}
    if !reflect.DeepEqual(result.Results, expectedResults) {
        t.Errorf("Expected results: %v, results obtained: %v", expectedResults, result.Results)
    }
}

func TestSuggestProfileByName(t *testing.T) {
    useCase := NewSearchProfileUseCase(&mockSearchRepository{})

    result := useCase.SuggestProfileByName("keyword")

    expectedResults := []domainprofile.Profile{}
    if !reflect.DeepEqual(result.Results, expectedResults) {
        t.Errorf("Resultados esperados: %v, Resultados obtenidos: %v", expectedResults, result.Results)
    }
}
