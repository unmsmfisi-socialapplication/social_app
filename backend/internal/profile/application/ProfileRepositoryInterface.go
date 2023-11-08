package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"

type IProfileRepository interface {
    UpdateProfile(profile *domain.Profile) (error)
}
