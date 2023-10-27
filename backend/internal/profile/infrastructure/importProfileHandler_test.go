package infrastructure

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type MockProfileRepository struct{}

func newMockProfileRepository() *MockProfileRepository {
	return &MockProfileRepository{}
}

func (mr *MockProfileRepository) UpdateProfile(p *domain.Profile) error {
	return nil
}

func TestImportProfileHandler_ImportProfile(t *testing.T) {
	profileRepository := newMockProfileRepository()
	importProfileUseCase := application.NewImportProfileUseCase(profileRepository)
	handler := NewImportProfileHandler(importProfileUseCase)

	requestBody := []byte(`
    {
      "@context": [
        "https://www.w3.org/ns/activitystreams",
        {
          "@language": "es"
        }
      ],
      "type": "Profile",
      "actor": "https://appsocial.com/sofia/",
      "name": "Perfil de Sofia",
      "object": {
        "id": "https://appsocial.com/sofia/id1234",
        "type": "Person",
        "name": "Sofia Rodriguez",
        "preferredUsername": "SofiR",
        "summary": "Amante de la naturaleza y entusiasta de la tecnología.",
        "icon": "https://kenzoishii.example.com/image/165987aklre4",
        "endpoints": {
          "sharedInbox": "https://appsocial.com/inbox"
        },
        "importedFrom": "https://mastodon.ejemplo.com/@sofiR"
      },
      "to": [
        "https://appsocial.com/usuarios/"
      ],
      "cc": "https://appsocial.com/seguidores/sofia",
      "icon": "https://appsocial.com/image/sofia"
    }`)

	req := httptest.NewRequest("POST", "/import-profile", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	handler.ImportProfile(res, req)

	expectedStatus := http.StatusOK

	if res.Code != expectedStatus {
		t.Errorf("Expected status code %d, but got %d", expectedStatus, res.Code)
	}
}
