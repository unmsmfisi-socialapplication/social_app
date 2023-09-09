package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"

type profileRequest interface {}

type ImportProfileUseCase struct {
    profileRepository domain.ProfileRepositoryI
}

func NewImportProfileUseCase(profileRepository domain.ProfileRepositoryI) *ImportProfileUseCase {
    return &ImportProfileUseCase{profileRepository}
}

func (ipuc *ImportProfileUseCase) ImportProfile(p *domain.Profile) (*domain.Profile, error) {
    profile, err := ipuc.profileRepository.UpdateProfile(p)
    if err != nil {
        return nil, err
    }
    return profile, nil
}
