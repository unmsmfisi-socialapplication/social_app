package application

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type mockUserInterestCommunityRepository struct {
	findFunc func(communities []domain.UserInterestCommunity) error
	createFunc func(communities []domain.UserInterestCommunity) error	
}


func (m *mockUserInterestCommunityRepository) FindUserInterestCommunities(communities []domain.UserInterestCommunity) error {
	if m.findFunc != nil {
		return m.findFunc(communities)
	}
	return nil
}

func (m *mockUserInterestCommunityRepository) Create(communities []domain.UserInterestCommunity) error {
	if m.createFunc != nil {
		return m.createFunc(communities)
	}
	return nil
}

func TestSetInterestCommunitiesUseCase_SetInterestCommunities(t *testing.T) {

	mockRepo := &mockUserInterestCommunityRepository{}
	useCase := NewSetInterestCommunitiesUseCase(mockRepo)

	// Test Case: Successful Insertion
	t.Run("Successful Insertion", func(t *testing.T) {

		mockRepo.createFunc = func(communities []domain.UserInterestCommunity) error {
			return nil
		}

		userId := "1"
		communityId := []string{"1", "2","3"}

		err := useCase.SetInterestCommunities(userId, communityId)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})

	// Test Case: Failed Insertion
	t.Run("Failed Insertion", func(t *testing.T) {
		mockError:= errors.New("Insertion Error")

		mockRepo.createFunc = func(communities []domain.UserInterestCommunity) error {
			return mockError
		}

		userId := "1"
		communityId := []string{"1", "2","1"}

		err := useCase.SetInterestCommunities(userId, communityId)

		if err == nil {
			t.Error("Expected error, but received none")
		} else if err != mockError {
			t.Errorf("expected error %v, got %v", mockError, err)
		}
	})
}
