package infrastructure

import (
	"database/sql"
	"strconv"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type InterestTopicsDBRepository struct {
	db *sql.DB
}

func NewInterestTopicsDBRepository(database *sql.DB) domain.InterestTopicsRepository {
	return &InterestTopicsDBRepository{db: database}
}

func (dbRepository *InterestTopicsDBRepository) GetAll(pageSize, pageNumber string) ([]domain.InterestTopic, error) {

	var interestTopics []domain.InterestTopic
	pageSizeInt, err := strconv.Atoi(pageSize)
    if err != nil {
        return nil, err
    }

    pageNumberInt, err := strconv.Atoi(pageNumber)
    if err != nil {
        return nil, err
    }

    offset := (pageNumberInt - 1) * pageSizeInt

	query := `SELECT interest_id, interest_name, interest_summary FROM soc_app_m_users_interests LIMIT $1 OFFSET $2`
	rows, err := dbRepository.db.Query(query,pageSizeInt,offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var interestTopic domain.InterestTopic
		err := rows.Scan(&interestTopic.InterestId, &interestTopic.InterestName, &interestTopic.InterestSummary)
		if err != nil {
			return nil, err
		}
		interestTopics = append(interestTopics, interestTopic)
	}
	return interestTopics, nil
}