package domain

type CommunityRepository interface {
	CheckUserInterestTopics (userId string) (string,error)
	GetCommunitiesByUserId(userId, pageSize, pageNumber string) ([]Community, error)
}
