package domain

type UserInterestTopic struct {
	UserInterestId string `json:"user_interest_id"`
	UserId          string `json:"user_id"`
	InterestId      string `json:"interest_id"`
}

func NewUserInterestTopic(userInterestId, userId, interestId string) (*UserInterestTopic , error) {
	return &UserInterestTopic {
		UserInterestId:userInterestId,
		UserId: userId,
		InterestId: interestId,
	}, nil
}