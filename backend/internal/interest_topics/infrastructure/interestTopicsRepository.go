package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
	
)

type InterestTopicsDBRepository struct {
	db *sql.DB
}

func NewInterestTopicsDBRepository(database *sql.DB) domain.InterestTopicsRepository{
	return &InterestTopicsDBRepository{db:	database}
}
func (dbRepository InterestTopicsDBRepository) FindAll () ([]domain.InterestTopics,error){
	var interestTopics []domain.InterestTopics
	query := `SELECT interest_id, interest_name, interest_summary FROM sa.soc_app_m_users_interests`
	rows,err := dbRepository.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var interestTopic domain.InterestTopics
		err:=rows.Scan(&interestTopic.InterestId,&interestTopic.InterestName,&interestTopic.InterestSummary)
		if err != nil {
			return nil, err
		}
		interestTopics = append(interestTopics, interestTopic)
	}
	return interestTopics, nil
}