package infrastructure

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
)

func TestCreateProfileHandler_CreateProfile(t *testing.T){
	profileRepository := newMockProfileRepository()
	createProfileUseCase := application.NewCreateProfileUseCase(profileRepository)
	handler:= NewCreateProfileHandler(createProfileUseCase)

	data := map[string]string{
		"user_id":         "123456789",
		"birth_date":      "02-01-2006",
		"name":            "testuser",
		"last_name":       "testLastname",
		"about_me":        "testAboutme",
		"genre":           "testGenre",
		"address":         "testAddress",
		"country":         "testCountry",
		"city":            "testCity",
		"profile_picture": "testProfilePicture",
	}
	requestData, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/profile/create",  bytes.NewReader(requestData))
	rr := httptest.NewRecorder()
	handler.CreateProfile(rr, req)

	if rr.Code != 200 {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}