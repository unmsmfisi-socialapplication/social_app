package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/search/domain"

type SearchRepository interface {
    GetProfilesByName(keyword string) *domain.QueryResult
    GetsuggestionsProfiles(keyword string) *domain.QueryResult
}

type SearchProfileUseCase struct {
    repository SearchRepository
}

func NewSearchProfileUseCase(repository SearchRepository) *SearchProfileUseCase {
    return &SearchProfileUseCase{repository: repository}
}

func (useCase *SearchProfileUseCase) SearchProfileByName(keyword string) *domain.QueryResult {
    return useCase.repository.GetProfilesByName(keyword)
}

func (useCase *SearchProfileUseCase) SuggestProfileByName(keyword string) *domain.QueryResult {
    return useCase.repository.GetsuggestionsProfiles(keyword)
}
