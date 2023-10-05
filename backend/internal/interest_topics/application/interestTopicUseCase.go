package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

var (
	ErrInvalidInsertion = errors.New("invalid insertion")
)

type InterestTopicsUseCaseI interface {
	SetInterestTopics(user_id, interest_id string) (bool, error)
}

type InterestTopicsUseCase struct {
	repo domain.UserInterestsRepository
}

func NewInterestTopicsUseCase(repo domain.UserInterestsRepository) *InterestTopicsUseCase {
	return &InterestTopicsUseCase{
		repo: repo,
	}
}

func (itus *InterestTopicsUseCase) SetInterestTopics(user_id, interest_id string) (bool, error) {
	userInterest := &domain.UserInterestTopics{
		User_id:     user_id,
		Interest_id: interest_id,
	}

	err := itus.repo.Create(userInterest)
	if err != nil {
		return false, err
	}
	return true, nil
}
