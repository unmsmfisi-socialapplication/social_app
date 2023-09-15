package infrastructure_import

import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"

type ImportProfileRequest struct {
	Username       string `json:"username"`
	Biography      string `json:"biography"`
	ProfilePicture string `json:"profile_picture"`
}

func (r *ImportProfileRequest) ToProfile() *domain.Profile {
	return &domain.Profile{
        Id_profile:     0,
		Username:       r.Username,
		Biography:      r.Biography,
		ProfilePicture: r.ProfilePicture,
	}
}

type ImportProfileResponse struct {
	Username       string `json:"username"`
	Biography      string `json:"biography"`
	ProfilePicture string `json:"profile_picture"`
}

func (w *ImportProfileResponse) FromProfile(p *domain.Profile) {
	w.Username = p.Username
	w.Biography = p.Biography
	w.ProfilePicture = p.ProfilePicture
}
