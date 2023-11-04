package domain

import (
	"testing"
)

func TestNewUserInterestTopic(t *testing.T) {
	userInterestId := "1"
	userId := "2"
	interestId := "3"

	userInterestTopic, err := NewUserInterestTopic(userInterestId, userId, interestId)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if userInterestTopic.UserInterestId != userInterestId {
		t.Errorf("Expected userInterestId to be %s, but got %s", userInterestId, userInterestTopic.UserInterestId)
	}

	if userInterestTopic.UserId != userId {
		t.Errorf("Expected userId to be %s, but got %s",userId, userInterestTopic.UserId)
	}

	if userInterestTopic.InterestId != interestId {
		t.Errorf("Expected interestId to be %s, but got %s",interestId, userInterestTopic.InterestId)
	}

}