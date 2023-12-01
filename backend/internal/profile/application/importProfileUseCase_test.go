package application

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type MockProfileRepository struct{}

func newMockProfileRepository() *MockProfileRepository {
	return &MockProfileRepository{}
}

func (mr *MockProfileRepository) UpdateProfile(p *domain.Profile) error {
	return nil
}

func (mr *MockProfileRepository) CreateProfile(p *domain.Profile) error {
	return nil
}

func TestImportProfileUseCase(t *testing.T) {
	t.Log("TestImportProfileUseCase")

	p := domain.NewProfileToImport("id_profile", "test", "test", "test", "test", "test")
	
	profile := domain.NewProfileToImport("id_profile", p.Name, p.LastName, "test", p.ProfilePicture, p.AboutMe)

    profileRepository := newMockProfileRepository()

	importProfileUseCase := NewImportProfileUseCase(profileRepository)

	err := importProfileUseCase.ImportProfile(p)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if profile.Name != p.Name {
		t.Errorf("Error: %v", err)
	}

	if profile.ProfilePicture != p.ProfilePicture {
		t.Errorf("Error: %v", err)
	}

	if profile.AboutMe != p.AboutMe {
		t.Errorf("Error: %v", err)
	}

	t.Log("TestImportProfileUseCase: Success")
}
