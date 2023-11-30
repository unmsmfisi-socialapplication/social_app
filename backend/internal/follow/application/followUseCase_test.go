package application

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type MockFollowerRepository struct{}

func (m *MockFollowerRepository) InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error) {
	return &domain.Follower{}, nil
}

func (m *MockFollowerRepository) IsFollowing(newFollower *domain.Follower) (*bool, error) {
	return new(bool), nil
}

func (m *MockFollowerRepository) ViewCommonFollowers(p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {
	return &domain.FollowerDataList{}, nil
}

func (m *MockFollowerRepository) ProfileFollowers(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {
	return &domain.FollowerDataList{}, nil
}

func TestNewFollowerUseCase(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	newFollower, err := useCase.FollowProfile(1, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if newFollower == nil {
		t.Error("newFollower is null, but a non-null value was expected.")
	}
}

func TestFollowProfile(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	newFollower, err := useCase.FollowProfile(-1, -2)
	if err != nil {
		t.Logf("Expected error: %v", err)
	} else {
		t.Error("Expected an error, but got none.")
	}

	if newFollower != nil {
		t.Error("newFollower is not null, but a null value was expected.")
	}
}

func TestIsFollowing(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	isFollowing, err := useCase.IsFollowing(1, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *isFollowing != false {
		t.Error("User 1 was expected not to follow user 2, but IsFollowing is true.")
	}

	isFollowing, err = useCase.IsFollowing(2, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *isFollowing != false {
		t.Error("It was expected that user 2 would not follow user 1, but IsFollowing is true.")
	}

	_, err = useCase.IsFollowing(-2, 1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.IsFollowing(2, -1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.IsFollowing(-2, -1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}
}

func TestCommonFollowers(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	commonFollowers, err := useCase.CommonFollowers(1, 2, 10, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if commonFollowers == nil {
		t.Error("commonFollowers is null, but a non-null value was expected.")
	}

	_, err = useCase.CommonFollowers(1, 1, 10, 1)
	if err != ErroEqualProfile {
		t.Errorf("Expected error: %v, but got: %v", ErroEqualProfile, err)
	} else {
		t.Logf("Expected error: %v", ErroEqualProfile)
	}

	_, err = useCase.CommonFollowers(-1, 2, 10, 1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.CommonFollowers(1, -2, 10, 1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}
}

func TestProfileFollowers(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	followers, err := useCase.ProfileFollowers(1, 10, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if followers == nil {
		t.Error("followers is null, but a non-null value was expected.")
	}

	_, err = useCase.ProfileFollowers(-1, 10, 1)
	if err == nil {
		t.Error("An error was expected as the profile_id provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowers(1, -10, 1)
	if err == nil {
		t.Error("An error was expected as the page_size provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowers(1, 10, -1)
	if err == nil {
		t.Error("An error was expected as the page_number provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowers(0, 10, 1)
	if err == nil {
		t.Error("An error was expected as the profile_id provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowers(1, 0, 1)
	if err == nil {
		t.Error("An error was expected as the page_size provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowers(1, 10, 0)
	if err == nil {
		t.Error("An error was expected as the page_number provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}
}

func TestProfileFollowingProfiles(t *testing.T) {
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	followingProfiles, err := useCase.ProfileFollowingProfiles(1, 10, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if followingProfiles == nil {
		t.Error("followingProfiles is null, but a non-null value was expected.")
	}

	_, err = useCase.ProfileFollowingProfiles(-1, 10, 1)
	if err == nil {
		t.Error("An error was expected as the profile_id provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowingProfiles(1, -10, 1)
	if err == nil {
		t.Error("An error was expected as the page_size provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowingProfiles(1, 10, -1)
	if err == nil {
		t.Error("An error was expected as the page_number provided is negative.")
	} else {
		expectedErrMsg := "NEGATIVE_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowingProfiles(0, 10, 1)
	if err == nil {
		t.Error("An error was expected as the profile_id provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowingProfiles(1, 0, 1)
	if err == nil {
		t.Error("An error was expected as the page_size provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	_, err = useCase.ProfileFollowingProfiles(1, 10, 0)
	if err == nil {
		t.Error("An error was expected as the page_number provided is zero.")
	} else {
		expectedErrMsg := "NULL_PARAMETERS"
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}
}
