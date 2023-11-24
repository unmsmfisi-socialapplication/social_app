package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type CommunityDBRepository struct {
	db *sql.DB
}

func NewCommunityDBRepository(database *sql.DB) domain.CommunityRepository {
	return &CommunityDBRepository{db: database}
}

func (dbRepository CommunityDBRepository) CheckUserInterestTopics(user_id string) (bool, error) {
	query := `SELECT user_id FROM soc_app_users_interest_topics WHERE user_id=$1`
	rows, err := dbRepository.db.Query(query, user_id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (dbRepository CommunityDBRepository) GetCommunities() ([]domain.Community, error) {
	var communities []domain.Community
	query := `SELECT community_id, community_name, community_description FROM soc_app_interest_communities`
	rows, err := dbRepository.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var community domain.Community
		err := rows.Scan(&community.CommunityId, &community.CommunityName, &community.CommunityDescription, &community.InterestId)
		if err != nil {
			return nil, err
		}
		communities = append(communities, community)
	}
	return communities, nil
}

func (dbRepository CommunityDBRepository) GetCommunitiesByUserId(userId string) ([]domain.Community, error) {
	var communities []domain.Community
	query := `SELECT community_id, community_name, community_description FROM soc_app_interest_communities com
				INNER JOIN soc_app_users_interest_topics userit ON com.interest_id=userit.interest_id
				WHERE userit.user_id=$1`
	rows, err := dbRepository.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var community domain.Community
		err := rows.Scan(&community.CommunityId, &community.CommunityName, &community.CommunityDescription)
		if err != nil {
			return nil, err
		}
		communities = append(communities, community)
	}

	return communities, nil
}
