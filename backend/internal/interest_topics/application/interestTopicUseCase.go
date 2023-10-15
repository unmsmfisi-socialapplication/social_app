package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

var (
	ExistingUserInterestTopic = errors.New("existing user interest topic")
)

type InterestTopicsUseCaseI interface {
	SetInterestTopics(user_id string, interest_id[] string)  error
}

type InterestTopicsUseCase struct {
	repository domain.UserInterestsRepository
}

func NewInterestTopicsUseCase(repository domain.UserInterestsRepository) *InterestTopicsUseCase {
	return &InterestTopicsUseCase{
		repository: repository,
	}
}

func (usecase *InterestTopicsUseCase) SetInterestTopics(user_id string, interest_id[] string) error {
	var userInterests []*domain.UserInterestTopics
		
	for i := 0; i < len(interest_id); i++ {
		userInterest := &domain.UserInterestTopics{
			User_id:     user_id,
			Interest_id: interest_id[i],
		}
		userInterests = append(userInterests, userInterest)
	}
	err := usecase.repository.Create(userInterests)
	if err != nil {
		return err
	}
	return nil
}
