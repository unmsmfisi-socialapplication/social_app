package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type ListCommunitiesUseCaseI interface {
	GetCommunitiesList(userId, pageSize, pageNumber string) ([]domain.Community, error)
}

type ListCommunitiesUseCase struct {
	repository domain.CommunityRepository
}

func NewListCommunitiesUseCase(repository domain.CommunityRepository) *ListCommunitiesUseCase {
	return &ListCommunitiesUseCase{repository: repository}
}


func (usecase *ListCommunitiesUseCase) GetCommunitiesList(userId, pageSize, pageNumber string) ([]domain.Community, error) {

	var communities []domain.Community

	userid,err:=usecase.repository.CheckUserInterestTopics(userId)
	if err != nil {
		return nil, err
	}
	communities, err = usecase.repository.GetCommunitiesByUserId(userid, pageSize, pageNumber)
	if err != nil {
		return nil, err
	}
	return communities, nil

}
