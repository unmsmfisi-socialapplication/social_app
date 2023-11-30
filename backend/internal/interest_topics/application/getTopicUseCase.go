package application

import (
	

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type GetInterestTopicsUseCase struct {
	repository domain.UserInterestsRepository
}

func NewGetInterestTopicsUseCase(repository domain.UserInterestsRepository) *GetInterestTopicsUseCase {
	return &GetInterestTopicsUseCase{
		repository: repository,
	}
}

func (usecase *GetInterestTopicsUseCase) GetInterestTopics(userId string) ([]string, error) {
	interests, err := usecase.repository.GetInterestTopics(userId)
	if err != nil {
		return nil, err
	}
	return interests, nil

	
}