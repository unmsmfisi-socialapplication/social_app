package domain

//Interfaces must be implemented at the infraestructure layer

type UserInterestCommunityRepository interface {
	FindUserInterestCommunities(communities []UserInterestCommunity) error
	Create(communities []UserInterestCommunity) error
}
