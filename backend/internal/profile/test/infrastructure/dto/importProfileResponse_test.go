package test

import (
	"testing"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/infrastructureimport/dto"
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

    expectedResponse := dto.ImportProfileResponse{
        Response: "Profile Imported Succesfully",
        Data: expectedPerson,
    }

	response := dto.NewImportProfileResponse(expectedPerson)

    messageExcepted := "Profile Imported Succesfully"
	if response.Response != "Profile Imported Succesfully" {
        t.Errorf("Expected response '%s', but got '%s'", messageExcepted, response.Response)
	}

    if expectedResponse.Data != response.Data {
        t.Errorf("Expected response '%s', but got '%s'", expectedResponse, response)
    }
}
