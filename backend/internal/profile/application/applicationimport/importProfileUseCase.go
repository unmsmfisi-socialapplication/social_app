package applicationimport

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ImportProfileUseCase struct {
    profileRepository domain.ProfileRepositoryI
}

func NewImportProfileUseCase(profileRepository domain.ProfileRepositoryI) *ImportProfileUseCase {
    return &ImportProfileUseCase{profileRepository}
}

func (ipuc *ImportProfileUseCase) ImportProfile(p *domain.Profile) (error) {
    err := ipuc.profileRepository.UpdateProfile(p)
    if err != nil {
        return err
    }
    return nil
}