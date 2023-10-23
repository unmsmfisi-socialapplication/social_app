package domain

//Interfaces must be implemented at the infraestructure layer

// Select interest topics
type UserInterestsRepository interface {
	FindUserInterestTopics(interests []UserInterestTopic)  error
	Create(interests[] UserInterestTopic) error
}
