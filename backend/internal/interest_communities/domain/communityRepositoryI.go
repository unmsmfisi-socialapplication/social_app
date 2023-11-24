package domain

type CommunityRepository interface {
	CheckUserInterestTopics (user_id string) (bool,error)
	GetCommunities() ([]Community, error)
	GetCommunitiesByUserId(userId string) ([]Community, error)
}
