package domain

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

//Interfaces must be implemented at the infraestructure layer

// Select interest topics
type UserInterestsRepository interface {
	Create(interest *entity.UserInterestTopics) error
}
