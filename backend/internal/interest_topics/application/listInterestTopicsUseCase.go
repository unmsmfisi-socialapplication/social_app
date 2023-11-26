package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type ListInterestTopicsUseCaseI interface {
	GetInteresTopics(pageSize, pageNumber string) ([]domain.InterestTopic, error)
}

type ListInterestTopicsUseCase struct {
	repository domain.InterestTopicsRepository
}

func NewListInterestTopicsUseCase(repository domain.InterestTopicsRepository) *ListInterestTopicsUseCase {
	return &ListInterestTopicsUseCase{repository: repository}
}

func (usecase *ListInterestTopicsUseCase) GetInteresTopics(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
	var interestTopics []domain.InterestTopic
	interestTopics, err := usecase.repository.GetAll(pageSize, pageNumber)
	if err != nil {
		return nil, err
	}
	return interestTopics, nil
}
