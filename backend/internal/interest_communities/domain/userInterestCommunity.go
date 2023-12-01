package domain

type UserInterestCommunity struct {
	UserInterestCommunityId string `json:"user_interest_community_id"`
	UserId                  string `json:"user_id"`
	CommunityId             string `json:"community_id"`
}

func NewUserInterestCommunity(userInterestCommunityId, userId, CommunityId string) (*UserInterestCommunity, error) {
	return &UserInterestCommunity{
		UserInterestCommunityId: userInterestCommunityId,
		UserId:                  userId,
		CommunityId:             CommunityId,
	}, nil
}
