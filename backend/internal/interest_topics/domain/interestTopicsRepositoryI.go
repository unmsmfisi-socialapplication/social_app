package domain


type InterestTopicsRepository interface {
	FindAll() ([]InterestTopics, error)
}