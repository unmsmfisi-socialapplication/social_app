package test

import (
	"testing"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/import/dto"
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

	// Crear una nueva respuesta ImportProfileResponse
	response := dto.NewImportProfileResponse(expectedPerson)

	// Verificar que los campos de la respuesta sean los esperados
	if response.Response != "Profile Imported Succesfully" {
		t.Errorf("Se esperaba el valor 'Profile Imported Succesfully', pero se obtuvo '%s'", response.Response)
	}

    if expectedResponse.Data != response.Data {
        t.Errorf("Se esperaba la respuesta '%s', pero se obtuvo '%s'", expectedResponse, response)
    }
}
