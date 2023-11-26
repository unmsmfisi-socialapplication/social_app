package domain


type InterestTopicsRepository interface {
	GetAll(pageSize, pageNumber string) ([]InterestTopic, error)
}