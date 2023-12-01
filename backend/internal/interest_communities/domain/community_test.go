package domain

import (
	"testing"
)

func TestNewUserInterestTopic(t *testing.T) {
	communityId := "1"
	communityName := "Grupo de Programadores"
	communityDescription := "Grupo de Programadores"
	interestId := "1"

	community, err := NewCommunity(communityId,communityName,communityDescription,interestId)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if community.CommunityId != communityId {
		t.Errorf("Expected communityId to be %s, but got %s", communityId, community.CommunityId)
	}

	if community.CommunityName != communityName{
		t.Errorf("Expected communityName to be %s, but got %s",communityName, community.CommunityName)
	}

	if community.CommunityDescription!= communityDescription {
		t.Errorf("Expected communityDescription to be %s, but got %s",communityDescription, community.CommunityDescription)
	}

	if community.InterestId != interestId {
		t.Errorf("Expected interestId to be %s, but got %s",interestId, community.InterestId)
	}

}