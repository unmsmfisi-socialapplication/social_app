package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

var (
	ExistingUserInterestCommunity = errors.New("existing record")
)

type SetInterestCommunitiesUseCaseI interface {
	SetInterestCommunities(userId string, community_id []string) error
}

type SetInterestCommunitiesUseCase struct {
	repository domain.UserInterestCommunityRepository
}

func NewSetInterestCommunitiesUseCase(repository domain.UserInterestCommunityRepository) *SetInterestCommunitiesUseCase {
	return &SetInterestCommunitiesUseCase{
		repository: repository,
	}
}

func (usecase *SetInterestCommunitiesUseCase) SetInterestCommunities(userId string, communityId []string) error {
	var userInterestCommunities []domain.UserInterestCommunity

	for i := 0; i < len(communityId); i++ {
		userInterestCommunity := domain.UserInterestCommunity{
			UserId:     userId,
			CommunityId: communityId[i],
		}
		userInterestCommunities = append(userInterestCommunities, userInterestCommunity)
	}
	err := usecase.repository.Create(userInterestCommunities)
	if err != nil {
		return err
	}
	return nil
}
