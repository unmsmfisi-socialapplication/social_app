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

func TestImportProfileUseCase(t *testing.T) {
	t.Log("TestImportProfileUseCase")

	p := &domain.Profile{
		Id_profile:   "test",
		Username:     "test",
		ProfileImage: "test",
		Biography:    "test",
	}

    profile := &domain.Profile{
        Id_profile:   p.Id_profile,
        Username:     p.Username,
        ProfileImage: p.ProfileImage,
        Biography:    p.Biography,
    }

	profileRepository := newMockProfileRepository()
	importProfileUseCase := NewImportProfileUseCase(profileRepository)

	err := importProfileUseCase.ImportProfile(p)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if profile.Username != p.Username {
		t.Errorf("Error: %v", err)
	}

	if profile.ProfileImage != p.ProfileImage {
		t.Errorf("Error: %v", err)
	}

	if profile.Biography != p.Biography {
		t.Errorf("Error: %v", err)
	}

	t.Log("TestImportProfileUseCase: Success")
}
