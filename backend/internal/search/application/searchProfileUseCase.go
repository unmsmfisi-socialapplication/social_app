package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/search/domain"

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
