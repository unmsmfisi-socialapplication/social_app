package application

import (
	"testing"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type mockInterestTopicsRepository struct {
	FindAllFunc func() ([]domain.InterestTopic, error)
}

func (m *mockInterestTopicsRepository) FindAll() ([]domain.InterestTopic, error) {
	return m.FindAllFunc()
}

func TestGetInterestTopics(t *testing.T) {

	mockRepo := &mockInterestTopicsRepository{}
	usecase := NewListInterestTopicsUseCase(mockRepo)

	expectedTopics := []domain.InterestTopic{
		{InterestId: "1", InterestName: "Topic 1", InterestSummary: "Summary 1"},
		{InterestId: "2", InterestName: "Topic 2", InterestSummary: "Summary 2"},
	}
	mockRepo.FindAllFunc = func() ([]domain.InterestTopic, error) {
		return expectedTopics, nil
	}
	topics, err := usecase.GetInteresTopics()

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