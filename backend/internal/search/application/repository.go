package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/search/domain"

type SearchRepository interface {
    GetProfilesByName(keyword string) *domain.QueryResult
    GetsuggestionsProfiles(keyword string) *domain.QueryResult
}
