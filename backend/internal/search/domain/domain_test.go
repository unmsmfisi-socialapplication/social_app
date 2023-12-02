package domain

import (
	"reflect"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

func TestNewQueryResult(t *testing.T)  {
    profiles := []domain.Profile{}
    result := NewQueryResult(profiles)

    if !reflect.DeepEqual(result.Results, profiles) {
        t.Errorf("Expected results: %v, results obtained: %v", profiles, result.Results)
    }
}
