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
