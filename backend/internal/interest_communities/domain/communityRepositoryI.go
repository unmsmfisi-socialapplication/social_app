package domain

type CommunityRepository interface {
	CheckUserInterestTopics (user_id string) (bool,error)
	GetCommunities(pageSize, pageNumber string) ([]Community, error)
	GetCommunitiesByUserId(userId, pageSize, pageNumber string) ([]Community, error)
}
