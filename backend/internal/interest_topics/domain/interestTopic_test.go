package domain

import (
	"testing"
)

func TestNewInterestTopic(t *testing.T) {
	interestId:= "1"
	interestName:= "Deporte"
	interestSummary:= "Deportes"

	interestTopic, err := NewInterestTopic(interestId,interestName,interestSummary)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if interestTopic.InterestId != interestId {
		t.Errorf("Expected interestId to be %s, but got %s",interestId, interestTopic.InterestId)
	}
	if interestTopic.InterestName != interestName {
		t.Errorf("Expected interestId to be %s, but got %s",interestName, interestTopic.InterestName)
	}
	if interestTopic.InterestSummary != interestSummary{
		t.Errorf("Expected interestId to be %s, but got %s",interestSummary, interestTopic.InterestSummary)
	}

}