package infrastructure_repository

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ProfileRepository struct {
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return profile, nil
}
