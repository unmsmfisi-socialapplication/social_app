package infrastructure

import (
	"database/sql"
	"strconv"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type CommunityDBRepository struct {
	db *sql.DB
}

func NewCommunityDBRepository(database *sql.DB) domain.CommunityRepository {
	return &CommunityDBRepository{db: database}
}

func (dbRepository CommunityDBRepository) CheckUserInterestTopics(user_id string) (bool, error) {
	query := `SELECT count(user_id) FROM soc_app_users_interest_topics WHERE user_id=$1`
	var count int

	err := dbRepository.db.QueryRow(query, user_id).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func parsePageSizeAndNumber(pageSize, pageNumber string) (int, int, error) {
    pageSizeInt, err := strconv.Atoi(pageSize)
    if err != nil {
        return 0, 0, err
    }

    pageNumberInt, err := strconv.Atoi(pageNumber)
    if err != nil {
        return 0, 0, err
    }

    return pageSizeInt, pageNumberInt, nil
}

func processRows(rows *sql.Rows) ([]domain.Community, error) {
    var communities []domain.Community

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

func (dbRepository CommunityDBRepository) GetCommunities(pageSize, pageNumber string) ([]domain.Community, error) {

	pageSizeInt, pageNumberInt, err := parsePageSizeAndNumber(pageSize, pageNumber)
    if err != nil {
        return nil, err
    }
	offset := (pageNumberInt - 1) * pageSizeInt

	query := `SELECT community_id, community_name, community_description,interest_id FROM soc_app_communities LIMIT $1 OFFSET $2`
	rows, err := dbRepository.db.Query(query,pageSizeInt,offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return processRows(rows)
}

func (dbRepository CommunityDBRepository) GetCommunitiesByUserId(userId, pageSize, pageNumber string) ([]domain.Community, error) {

	pageSizeInt, pageNumberInt, err := parsePageSizeAndNumber(pageSize, pageNumber)
    if err != nil {
        return nil, err
    }
	offset := (pageNumberInt - 1) * pageSizeInt
	
	query := `SELECT com.community_id, com.community_name, com.community_description, com.interest_id FROM soc_app_communities com
				INNER JOIN soc_app_users_interest_topics userit ON com.interest_id=userit.interest_id
				WHERE userit.user_id=$1
				LIMIT $2 OFFSET $3`

	rows, err := dbRepository.db.Query(query, userId,pageSizeInt,offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return processRows(rows)
}
