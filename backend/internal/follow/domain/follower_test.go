package domain

import "testing"

func TestNewFollower(t *testing.T) {
	t.Run("Test con éxito", func(t *testing.T) {
		p_follower_profile_id := 1
		p_following_profile_id := 2

		follower, err := NewFollower(p_follower_profile_id, p_following_profile_id)

		if err != nil {
			t.Errorf("A null error was expected, but the result was: %v", err)
		}

		if follower == nil {
			t.Error("One follower was expected, but nil was obtained.")
		}

		if follower.Follower_profile_id != p_follower_profile_id || follower.Following_profile_id != p_following_profile_id {
			t.Errorf("Returned follower does not match expectations: %v", follower)
		}
	})

	t.Run("Test with negative values", func(t *testing.T) {
		p_follower_profile_id := -1
		p_following_profile_id := -2

		follower, err := NewFollower(p_follower_profile_id, p_following_profile_id)

		if err == nil {
			t.Error("An error was expected, but a null error was obtained.")
		}

		if follower != nil {
			t.Error("A null follower was expected, but a follower was obtained.")
		}
	})
}

func TestNewFollowerDataList(t *testing.T) {
	// Test case 1
	data1 := FollowerData{
		ProfileID:      1,
		UserID:         101,
		Name:           "John",
		LastName:       "Doe",
		ProfilePicture: "profile1.jpg",
		UserName:       "john_doe",
	}

	// Test case 2
	data2 := FollowerData{
		ProfileID:      2,
		UserID:         102,
		Name:           "Jane",
		LastName:       "Smith",
		ProfilePicture: "profile2.jpg",
		UserName:       "jane_smith",
	}

	// Test case 3 (empty list)
	emptyList := FollowerDataList{}

	// NewFollowerDataList function call with multiple data
	resultList := NewFollowerDataList(data1, data2)

	// Check the length of the resulting list
	if len(resultList) != 2 {
		t.Errorf("Expected length 2, got %d", len(resultList))
	}

	// Verify if the items in the list match the input data
	if resultList[0] != data1 {
		t.Errorf("Expected %v, got %v", data1, resultList[0])
	}

	if resultList[1] != data2 {
		t.Errorf("Expected %v, got %v", data2, resultList[1])
	}

	// Call to NewFollowerDataList function with empty list
	emptyResultList := NewFollowerDataList(emptyList...)

	// Check that the resulting list is empty
	if len(emptyResultList) != 0 {
		t.Errorf("Expected an empty list, got %v", emptyResultList)
	}
}
