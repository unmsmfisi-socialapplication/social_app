package domain

//Access category data
type TopicCategoryRepository interface {
	FindAll() ([]*Topic_Category, error)
}

//Access subcategory data
type TopicSubcategoryRepository interface {
	FindAll() ([]*Topic_Subcategory, error)
}

//Select interest topics
type InterestTopicsUserRepository interface {
	Create(interest *Interest_Topic_Users) error
}
