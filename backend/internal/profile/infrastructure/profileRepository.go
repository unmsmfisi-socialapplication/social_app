package infrastructure

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ProfileRepository struct {
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (pr *ProfileRepository) UpdateProfile(profile *domain.Profile) (error) {
	return nil
}
