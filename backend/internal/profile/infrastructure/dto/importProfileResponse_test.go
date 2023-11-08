package dto

import (
	"testing"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

func TestNewImportProfileResponse(t *testing.T) {

	expectedPerson := &activitypub.Actor{
		ID:                "https://appsocial.com/sofia/id1234",
		Type:              "Person",
		Name:              activitypub.DefaultNaturalLanguageValue("Sofia Rodriguez"),
		PreferredUsername: activitypub.DefaultNaturalLanguageValue("SofiR"),
		Summary:           activitypub.DefaultNaturalLanguageValue("Amante de la naturaleza y entusiasta de la tecnología."),
		Icon:              activitypub.IRI("https://kenzoishii.example.com/image/165987aklre4"),
	}
	
	profile := domain.NewProfileToImport(
		"Sofia",
		"https://appsocial.com/sofia/id1234",
		"Rodriguez",
		"SofiR",
		"https://appsocial.com/sofia/profileImage.jpg",
		"Amante de la naturaleza y entusiasta de la tecnología.",
	)

	response := NewImportProfileResponse(profile)

	messageExcepted := "Profile Imported Succesfully"
	if response.Response != "Profile Imported Succesfully" {
		t.Errorf("Expected message response '%s', but got '%s'", messageExcepted, response.Response)
	}

	if response.Data.ID != expectedPerson.ID {
		t.Errorf("Expected ID '%s', but got '%s'", expectedPerson.ID, response.Data.ID)
	}
}
