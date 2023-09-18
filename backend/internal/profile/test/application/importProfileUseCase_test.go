package test

import (
	"testing"

	application_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/import"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	infrastructure_repository "github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/repository"
)

func TestImportProfileUseCase(t *testing.T) {
	t.Log("TestImportProfileUseCase")

	p := &domain.Profile{
		Id_profile:   "test",
		Username:     "test",
		ProfileImage: "test",
		Biography:    "test",
	}

	var profile *domain.Profile

	profileRepository := infrastructure_repository.NewProfileRepository()
	importProfileUseCase := application_import.NewImportProfileUseCase(profileRepository)

	profile, err := importProfileUseCase.ImportProfile(p)

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
