package domain

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

//Interfaces must be implemented at the infraestructure layer

//Access category data
type TopicCategoryRepository interface {
	FindAll() ([]*entity.Topic_Category, error)
}

//Access subcategory data
type TopicSubcategoryRepository interface {
	FindAll() ([]*entity.Topic_Subcategory, error)
}

//Select interest topics
type InterestTopicsUserRepository interface {
	Create(interest *entity.Interest_Topic_Users) error
}
