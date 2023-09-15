package test

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	infrastructure_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/import"
)

var importProfileRequest infrastructure_import.ImportProfileRequest
var importProfileResponse infrastructure_import.ImportProfileResponse

func ValidateTest(t *testing.T) {
	t.Log("ValidateTest")

	request := infrastructure_import.ImportProfileRequest{
		Context: []interface{}{
			"https://www.w3.org/ns/activitystreams",
			map[string]interface{}{"@language": "es"},
		},
		Type:  "Profile",
		Actor: "https://appsocial.com/sofia/",
		Name:  "Perfil de Sofia",
		Object: struct {
			Id                string `json:"id"`
			Type              string `json:"type"`
			Name              string `json:"name"`
			PreferredUsername string `json:"preferredUsername"`
			Summary           string `json:"summary"`
			ProfileImage      string `json:"profileImage"`
			CoverImage        string `json:"coverImage"`
			Endpoints         struct {
				SharedInbox string `json:"sharedInbox"`
			} `json:"endpoints"`
			ImportedFrom string `json:"importedFrom"`
		}{
			Id:                "https://appsocial.com/sofia/id1234",
			Type:              "Person",
			Name:              "Sofia Rodriguez",
			PreferredUsername: "SofiR",
			Summary:           "Amante de la naturaleza y entusiasta de la tecnología.",
			ProfileImage:      "https://appsocial.com/sofia/profileImage.jpg",
			CoverImage:        "https://appsocial.com/sofia/coverImage.jpg",
			Endpoints: struct {
				SharedInbox string `json:"sharedInbox"`
			}{
				SharedInbox: "https://appsocial.com/inbox",
			},
			ImportedFrom: "https://mastodon.ejemplo.com/@sofiR",
		},
		To: []string{"https://appsocial.com/usuarios/"},
		Cc: "https://appsocial.com/seguidores/sofia",
	}

	err := request.Validate()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Log("ValidateTest: Success")
}

func ToProfileTest(t *testing.T) {
	t.Log("ToProfileTest")

	importProfileRequest = infrastructure_import.ImportProfileRequest{}

	profile := importProfileRequest.ToProfile()

	if profile.Username != importProfileResponse.Username {
		t.Errorf("Error: %v", "Username")
	}

	if profile.ProfileImage != importProfileResponse.ProfilePicture {
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
		Id_profile:   "test",
		Username:     "test",
		ProfileImage: "test",
		Biography:    "test",
	}

	importProfileResponse.FromProfile(p)

	if importProfileResponse.Username != p.Username {
		t.Errorf("Error: %v", "Username")
	}

	if importProfileResponse.ProfilePicture != p.ProfileImage {
		t.Errorf("Error: %v", "ProfilePicture")
	}

	if importProfileResponse.Biography != p.Biography {
		t.Errorf("Error: %v", "Biography")
	}

	t.Log("FromProfileTest: Success")
}
