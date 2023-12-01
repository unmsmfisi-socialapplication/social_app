package domain

type Community struct {
	CommunityId          string `json:"community_id"`
	CommunityName        string `json:"community_name"`
	CommunityDescription string `json:"community_description"`
	InterestId           string `json:"interest_id"`
}

func NewCommunity(communityId, communityName, communityDescription, interestId string) (*Community, error) {
	return &Community{
		CommunityId:          communityId,
		CommunityName:        communityName,
		CommunityDescription: communityDescription,
		InterestId:           interestId,
	}, nil
}
