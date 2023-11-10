package domain

type InterestTopic struct {
	InterestId      string `json:"interest_id"`
	InterestName    string `json:"interest_name"`
	InterestSummary string `json:"interest_summary"`
}

func NewInterestTopic(interestId, interestName, interestSummary string) (*InterestTopic, error) {
	return &InterestTopic{
		InterestId:      interestId,
		InterestName:    interestName,
		InterestSummary: interestSummary,
	}, nil
}
