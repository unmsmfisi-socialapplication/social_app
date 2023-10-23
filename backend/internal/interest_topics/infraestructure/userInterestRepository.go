package infraestructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
)

type UserInterestsDBRepository struct {
	db *sql.DB
}

func NewUserInterestsDBRepository(database *sql.DB) domain.UserInterestsRepository {
	return &UserInterestsDBRepository{db: database}
}

func (dbRepository UserInterestsDBRepository) FindUserInterestTopics(interests []domain.UserInterestTopic) error {
	query := `SELECT user_id, interest_id FROM soc_app_users_interest_topics`

	rows, err := dbRepository.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	interestMap := make(map[string]bool)
	for _, interest := range interests {
		// Create a unique key for each item
		key := interest.UserId + "_" + interest.InterestId
		interestMap[key] = true
	}

	for rows.Next() {
		var interestTopic domain.UserInterestTopic
		err := rows.Scan(&interestTopic.UserId, &interestTopic.InterestId)
		if err != nil {
			return err
		}

		// Check if the key exists in the map
		key := interestTopic.UserId + "_" + interestTopic.InterestId
		if interestMap[key] {
			return application.ExistingUserInterestTopic // Attempted insertion of an existing user interest topic
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	// Enable insertion
	return nil
}

func (dbRepository *UserInterestsDBRepository) Create(interests []domain.UserInterestTopic) error {

	//Handling the atomicity of transaction
	tx, err := dbRepository.db.Begin()
	if err != nil {
		return err
	}
	err = dbRepository.FindUserInterestTopics(interests)
	if err != nil {
		return err
	}
	for _, interest := range interests {

		query := "INSERT INTO soc_app_users_interest_topics (user_id, interest_id) VALUES ($1, $2)"
		_, err := tx.Exec(query, interest.UserId, interest.InterestId)
		if err != nil {
			//If an error occurs, a rollback of the previous insertions is performed
			tx.Rollback()
			return err
		}

	}

	tx.Commit()
	return nil
}
