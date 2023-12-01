package application

import (
	"testing"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type mockInterestTopicsRepository struct {
	GetAllFunc func(pageSize, pageNumber string) ([]domain.InterestTopic, error)
}

func (m *mockInterestTopicsRepository) GetAll(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
	return m.GetAllFunc(pageSize, pageNumber)
}

func TestGetInterestTopics(t *testing.T) {

	mockRepo := &mockInterestTopicsRepository{}
	usecase := NewListInterestTopicsUseCase(mockRepo)

	expectedTopics := []domain.InterestTopic{
		{InterestId: "1", InterestName: "Topic 1", InterestSummary: "Summary 1"},
		{InterestId: "2", InterestName: "Topic 2", InterestSummary: "Summary 2"},
	}
	mockRepo.GetAllFunc = func(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
		return expectedTopics, nil
	}
	pageSize:="1"
	pageNumber:="2"
	topics, err := usecase.GetInteresTopics(pageSize, pageNumber)

	// Verify the function's result
	if err != nil {
		t.Errorf("Expected error nil, but got error: %v", err)
	}

	if len(topics) != len(expectedTopics) {
		t.Errorf("Expected the number of obtained interest topics to be %d, but got %d", len(expectedTopics), len(topics))
	}

	for i, expectedTopic := range expectedTopics {
        topic := topics[i]

        if topic.InterestId != expectedTopic.InterestId {
            t.Errorf("Expected interestId to be %s, but got %s", expectedTopic.InterestId, topic.InterestId)
        }

        if topic.InterestName != expectedTopic.InterestName {
            t.Errorf("Expected interestName to be %s, but got %s", expectedTopic.InterestName, topic.InterestName)
        }

        if topic.InterestSummary != expectedTopic.InterestSummary {
            t.Errorf("Expected interestSummary to be %s, but got %s", expectedTopic.InterestSummary, topic.InterestSummary)
        }
    }
}