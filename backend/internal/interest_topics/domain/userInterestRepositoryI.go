package domain

//Interfaces must be implemented at the infraestructure layer

// Select interest topics
type UserInterestsRepository interface {
	Create(interest *UserInterestTopics) error
}
