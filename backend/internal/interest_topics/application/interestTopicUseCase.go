package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

var (
	ExistingUserInterestTopic = errors.New("existing user interest topic")
)

type InterestTopicsUseCaseI interface {
	SetInterestTopics(userId string, interest_id []string) error
}

type InterestTopicsUseCase struct {
	repository domain.UserInterestsRepository
}

func NewInterestTopicsUseCase(repository domain.UserInterestsRepository) *InterestTopicsUseCase {
	return &InterestTopicsUseCase{
		repository: repository,
	}
}

func (usecase *InterestTopicsUseCase) SetInterestTopics(userId string, interestId []string) error {
	var userInterests []domain.UserInterestTopic

	for i := 0; i < len(interestId); i++ {
		userInterest := domain.UserInterestTopic{
			UserId:     userId,
			InterestId: interestId[i],
		}
		userInterests = append(userInterests, userInterest)
	}
	err := usecase.repository.Create(userInterests)
	if err != nil {
		return err
	}
	return nil
}
