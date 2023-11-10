package domain


type InterestTopicsRepository interface {
	FindAll() ([]InterestTopic, error)
}