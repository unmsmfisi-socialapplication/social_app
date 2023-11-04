package application

import (

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type ListInterestTopicsUseCaseI interface {
	GetInteresTopics() ([]domain.InterestTopics, error)
}

type ListInterestTopicsUseCase struct {
	repository domain.InterestTopicsRepository
}

func NewInterestTopicsUseCase(repository domain.InterestTopicsRepository) *ListInterestTopicsUseCase {
	return &ListInterestTopicsUseCase{repository: repository}
}

func (usecase *ListInterestTopicsUseCase) GetInteresTopics() ([]domain.InterestTopics, error) {
	var interestTopics []domain.InterestTopics
	interestTopics, err := usecase.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return interestTopics, nil

}
