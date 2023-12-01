package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

var (
	ExistingUserInterestTopic = errors.New("existing user interest topic")
)

type SelectTopicUseCaseI interface {
	SetInterestTopics(userId string, interest_id []string) error
}

type SelectTopicUseCase struct {
	repository domain.UserInterestsRepository
}

func NewSelectTopicUseCase(repository domain.UserInterestsRepository) *SelectTopicUseCase {
	return &SelectTopicUseCase{
		repository: repository,
	}
}

func (usecase *SelectTopicUseCase) SetInterestTopics(userId string, interestId []string) error {
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
