package domain

import (
	"testing"
)

func TestNewUserInterestCommunity(t *testing.T) {
	userInterestCommunityId := "1"
	userId := "2"
	communityId := "3"

	userInterestCommunity, err := NewUserInterestCommunity(userInterestCommunityId, userId, communityId)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if userInterestCommunity.UserInterestCommunityId != userInterestCommunityId {
		t.Errorf("Expected userInterestCommunityId to be %s, but got %s", userInterestCommunityId, userInterestCommunity.UserInterestCommunityId)
	}

	if userInterestCommunity.UserId != userId {
		t.Errorf("Expected userId to be %s, but got %s",userId, userInterestCommunity.UserId)
	}

	if userInterestCommunity.CommunityId != communityId {
		t.Errorf("Expected communityId to be %s, but got %s",communityId, userInterestCommunity.CommunityId)
	}

}