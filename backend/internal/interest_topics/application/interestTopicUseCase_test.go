package application

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type mockUserInterestsRepository struct {
	findFunc func(interests []domain.UserInterestTopic) error
	createFunc func(interests []domain.UserInterestTopic) error	
}


func (m *mockUserInterestsRepository) FindUserInterestTopics(interests []domain.UserInterestTopic) error {
	if m.findFunc != nil {
		return m.findFunc(interests)
	}
	return nil
}

func (m *mockUserInterestsRepository) Create(interests []domain.UserInterestTopic) error {
	if m.createFunc != nil {
		return m.createFunc(interests)
	}
	return nil
}

func TestInterestTopicsUseCase_SetInterestTopics(t *testing.T) {

	mockRepo := &mockUserInterestsRepository{}
	useCase := NewInterestTopicsUseCase(mockRepo)

	// Test Case: Successful Insertion
	t.Run("Successful Insertion", func(t *testing.T) {

		mockRepo.createFunc = func(interests []domain.UserInterestTopic) error {
			return nil
		}

		userId := "1"
		interestId := []string{"1", "2","3"}

		err := useCase.SetInterestTopics(userId, interestId)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})

	// Test Case: Failed Insertion
	t.Run("Failed Insertion", func(t *testing.T) {
		mockError:= errors.New("Insertion Error")

		mockRepo.createFunc = func(interests []domain.UserInterestTopic) error {
			return mockError
		}

		userId := "1"
		interestId := []string{"1", "2","1"}

		err := useCase.SetInterestTopics(userId, interestId)

		if err == nil {
			t.Error("Expected error, but received none")
		} else if err != mockError {
			t.Errorf("expected error %v, got %v", mockError, err)
		}
	})
}
