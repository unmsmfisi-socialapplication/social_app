package test

import (
	"testing"

	application_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/import"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure"
)

func TestImportProfileUseCase(t *testing.T) {
	t.Log("TestImportProfileUseCase")

	p := &domain.Profile{
		Id_profile:     1,
		Username:       "test",
		ProfilePicture: "test",
		Biography:      "test",
	}

	var profile *domain.Profile

	profileRepositoryTest := infrastructure.NewProfileRepositoryTest()
	importProfileUseCase := application_import.NewImportProfileUseCase(profileRepositoryTest)

	profile, err := importProfileUseCase.ImportProfile(p)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if profile.Username != p.Username {
		t.Errorf("Error: %v", err)
	}

	if profile.ProfilePicture != p.ProfilePicture {
		t.Errorf("Error: %v", err)
	}

	if profile.Biography != p.Biography {
		t.Errorf("Error: %v", err)
	}

	t.Log("TestImportProfileUseCase: Success")
}
