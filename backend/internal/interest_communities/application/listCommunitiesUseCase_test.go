package application

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type mockCommunityRepository struct {
	checkUserInterestTopicsFunc func(userId string) (string, error)
	getCommunitiesByUserIdFunc  func(userId, pageSize, pageNumber string) ([]domain.Community, error)
}

func (m *mockCommunityRepository) CheckUserInterestTopics(userId string) (string, error) {
	return m.checkUserInterestTopicsFunc(userId)
}

func (m *mockCommunityRepository) GetCommunitiesByUserId(userId, pageSize, pageNumber string) ([]domain.Community, error) {
	if m.getCommunitiesByUserIdFunc != nil {
		return m.getCommunitiesByUserIdFunc(userId, pageSize, pageNumber)
	}
	return nil, nil
}

func TestGetCommunitiesList(t *testing.T) {

	mockRepo := &mockCommunityRepository{}
	useCase := NewListCommunitiesUseCase(mockRepo)

	pageNumber := "1"
	pageSize := "2"

	// Test Case: User has no interest topics
	t.Run("Repository - GetCommunities", func(t *testing.T) {

		expectedCommunities := []domain.Community{
			{CommunityId: "1", CommunityName: "Programmers", CommunityDescription: "Programmers community", InterestId: "1"},
			{CommunityId: "2", CommunityName: "Musicians", CommunityDescription: "Musicians community", InterestId: "2"},
			{CommunityId: "3", CommunityName: "Engineers", CommunityDescription: "Engineers community", InterestId: "3"},
		}
		mockRepo.checkUserInterestTopicsFunc = func(userId string) (string, error) {
			return "-1", nil
		}

		mockRepo.getCommunitiesByUserIdFunc = func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
			return expectedCommunities, nil
		}

		userId := "11231"

		userId, err := mockRepo.checkUserInterestTopicsFunc(userId)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		communities, err := useCase.GetCommunitiesList(userId, pageSize, pageNumber)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if len(communities) != len(expectedCommunities) {
			t.Errorf("Expected the number of obtained communities to be %d, but got %d", len(expectedCommunities), len(communities))
		}

		for i, expectedCommunity := range expectedCommunities {
			community := communities[i]

			if community.CommunityId != expectedCommunity.CommunityId {
				t.Errorf("Expected communityId to be %s, but got %s", expectedCommunity.CommunityId, community.CommunityId)
			}
			if community.CommunityName != expectedCommunity.CommunityName {
				t.Errorf("Expected communityName to be %s, but got %s", expectedCommunity.CommunityName, community.CommunityName)
			}
			if community.CommunityDescription != expectedCommunity.CommunityDescription {
				t.Errorf("Expected communityDescription to be %s, but got %s", expectedCommunity.CommunityDescription, community.CommunityDescription)
			}
			if community.InterestId != expectedCommunity.InterestId {
				t.Errorf("Expected interestId to be %s, but got %s", expectedCommunity.InterestId, community.InterestId)
			}
		}
	})

	// Test Case: User has interest topics
	t.Run("Repository - GetCommunities", func(t *testing.T) {

		expectedCommunities := []domain.Community{
			{CommunityId: "2", CommunityName: "Musicians", CommunityDescription: "Musicians community", InterestId: "2"},
			{CommunityId: "3", CommunityName: "Engineers", CommunityDescription: "Engineers community", InterestId: "3"},
		}
		mockRepo.checkUserInterestTopicsFunc = func(userId string) (string, error) {
			return userId, nil
		}

		mockRepo.getCommunitiesByUserIdFunc = func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
			return expectedCommunities, nil
		}

		userId := "1"
		userId, err := mockRepo.checkUserInterestTopicsFunc(userId)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		communities, err := useCase.GetCommunitiesList(userId, pageSize, pageNumber)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if len(communities) != len(expectedCommunities) {
			t.Errorf("Expected the number of obtained communities to be %d, but got %d", len(expectedCommunities), len(communities))
		}

		for i, expectedCommunity := range expectedCommunities {
			community := communities[i]

			if community.CommunityId != expectedCommunity.CommunityId {
				t.Errorf("Expected communityId to be %s, but got %s", expectedCommunity.CommunityId, community.CommunityId)
			}
			if community.CommunityName != expectedCommunity.CommunityName {
				t.Errorf("Expected communityName to be %s, but got %s", expectedCommunity.CommunityName, community.CommunityName)
			}
			if community.CommunityDescription != expectedCommunity.CommunityDescription {
				t.Errorf("Expected communityDescription to be %s, but got %s", expectedCommunity.CommunityDescription, community.CommunityDescription)
			}
			if community.InterestId != expectedCommunity.InterestId {
				t.Errorf("Expected interestId to be %s, but got %s", expectedCommunity.InterestId, community.InterestId)
			}
		}
	})

}
