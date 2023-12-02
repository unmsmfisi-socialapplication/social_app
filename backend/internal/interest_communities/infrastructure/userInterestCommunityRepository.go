package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type UserInterestCommunitiesDBRepository struct {
	db *sql.DB
}

func NewUserInterestCommunitiesDBRepository(database *sql.DB) domain.UserInterestCommunityRepository {
	return &UserInterestCommunitiesDBRepository{db: database}
}
func (dbRepository UserInterestCommunitiesDBRepository) FindUserInterestCommunities(communities []domain.UserInterestCommunity) error {
	query := `SELECT user_id, community_id FROM soc_app_users_interest_communities`
	rows, err := dbRepository.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	communityMap := make(map[string]bool)
	for _, community := range communities {
		// Create a unique key for each item
		key := community.UserId + "_" + community.CommunityId
		communityMap[key] = true
	}
	for rows.Next() {
		var userInterestCommunity domain.UserInterestCommunity
		err := rows.Scan(&userInterestCommunity.UserId, &userInterestCommunity.CommunityId)
		if err != nil {
			return err
		}
		// Check if the key exists in the map
		key := userInterestCommunity.UserId + "_" + userInterestCommunity.CommunityId
		if communityMap[key] {
			return application.ExistingUserInterestCommunity // Attempted insertion of an existing record
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	// Enable insertion
	return nil
}
func (dbRepository UserInterestCommunitiesDBRepository) Create(communities []domain.UserInterestCommunity) error {
	//Handling the atomicity of transaction
	tx, err := dbRepository.db.Begin()
	if err != nil {
		return err
	}
	err = dbRepository.FindUserInterestCommunities(communities)
	if err != nil {
		return err
	}
	for _, community := range communities {
		query := "INSERT INTO soc_app_users_interest_communities (user_id, community_id) VALUES ($1, $2)"
		_, err := tx.Exec(query, community.UserId, community.CommunityId)
		if err != nil {
			//If an error occurs, a rollback of the previous insertions is performed
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
