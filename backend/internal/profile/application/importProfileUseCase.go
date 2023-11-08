package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ImportProfileUseCase struct {
    profileRepository IProfileRepository
}

func NewImportProfileUseCase(profileRepository IProfileRepository) *ImportProfileUseCase {
    return &ImportProfileUseCase{profileRepository}
}

func (ipuc *ImportProfileUseCase) ImportProfile(p *domain.Profile) (error) {
    err := ipuc.profileRepository.UpdateProfile(p)
    if err != nil {
        return err
    }
    return nil
}
