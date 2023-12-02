package domain

import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"

type QueryResult struct {
    Results []domain.Profile
}

func NewQueryResult(profiles []domain.Profile) *QueryResult {
    return &QueryResult{Results: profiles}
}
