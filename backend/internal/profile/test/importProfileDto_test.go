package test

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure"
)

var importProfileRequest infrastructure.ImportProfileRequest
var importProfileResponse infrastructure.ImportProfileResponse

func ToProfileTest(t *testing.T) {
    t.Log("ToProfileTest")

    importProfileRequest = infrastructure.ImportProfileRequest{
        Username:       "test",
        ProfilePicture: "test",
        Biography:      "test",
    }

    profile := importProfileRequest.ToProfile()

    if profile.Username != importProfileResponse.Username {
        t.Errorf("Error: %v", "Username")
    }

    if profile.ProfilePicture != importProfileResponse.ProfilePicture {
        t.Errorf("Error: %v", "ProfilePicture")
    }

    if profile.Biography != importProfileResponse.Biography {
        t.Errorf("Error: %v", "Biography")
    }

    t.Log("ToProfileTest: Success")
}

func FromProfileTest(t *testing.T) {
    t.Log("FromProfileTest")

    p := &domain.Profile{
        Id_profile:     1,
        Username:       "test",
        ProfilePicture: "test",
        Biography:      "test",
    }

    importProfileResponse.FromProfile(p)

    if importProfileResponse.Username != p.Username {
        t.Errorf("Error: %v", "Username")
    }

    if importProfileResponse.ProfilePicture != p.ProfilePicture {
        t.Errorf("Error: %v", "ProfilePicture")
    }

    if importProfileResponse.Biography != p.Biography {
        t.Errorf("Error: %v", "Biography")
    }

    t.Log("FromProfileTest: Success")
}
