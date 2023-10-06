package infraestructure

import (
	"database/sql"
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
)

type UserInterestsDBRepository struct {
	db *sql.DB
}

func NewUserInterestsDBRepository(db *sql.DB) domain.UserInterestsRepository {
	return &UserInterestsDBRepository{db: db}
}

func (ui *UserInterestsDBRepository) Create(interest *domain.UserInterestTopics) error {

	userIDValues := []string{"1", "2"}
	interestsIDValues := []string{"1", "2", "3", "4", "5", "6", "7"}
	userInterestTopics := [][]string{{"1", "2"}, {"1", "3"}, {"1", "4"}}

	if !contains(userIDValues, interest.User_id) || !contains(interestsIDValues, interest.Interest_id) {
		return errors.New("invalid insertion")
	}
	if containsSlice(userInterestTopics, []string{interest.User_id, interest.Interest_id}) {
		return application.ErrInvalidInsertion
	}
	return nil
}

// Temporary function for searching in a one dimensional array
func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

// Temporary function for searching in a bidimensional array
func containsSlice(arr2D [][]string, arrayToFind []string) bool {
	for _, slice := range arr2D {
		if len(slice) != len(arrayToFind) {
			continue
		}

		match := true
		for i := range slice {
			if slice[i] != arrayToFind[i] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}
	return false
}
