package infrastructure_import

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ImportProfileRequest struct {
	Context []interface{} `json:"@context"`
	Type    string        `json:"type"`
	Actor   string        `json:"actor"`
	Name    string        `json:"name"`
	Object  struct {
		Id                string `json:"id"`
		Type              string `json:"type"`
		Name              string `json:"name"`
		PreferredUsername string `json:"preferredUsername"`
		Summary           string `json:"summary"`
		ProfileImage      string `json:"profileImage"`
		CoverImage        string `json:"coverImage"`
		Endpoints         struct {
			SharedInbox string `json:"sharedInbox"`
		} `json:"endpoints"`
		ImportedFrom string `json:"importedFrom"`
	} `json:"object"`
	To []string `json:"to"`
	Cc string   `json:"cc"`
}

func (r *ImportProfileRequest) Validate() error {
    requiredFields := []struct {
        field string
        value string
    } {
        {"type", r.Type},
        {"actor", r.Actor},
        {"name", r.Name},
    }

    for _, requiredField := range requiredFields {
        if requiredField.value == "" {
            return &domain.ErrorRequiredField{Message: requiredField.field}
        }
    }

    return nil
}

func (r *ImportProfileRequest) ToProfile() *domain.Profile {
	return &domain.Profile{
		Id_profile:   r.Object.Id,
		Username:     r.Object.Name,
		Biography:    r.Object.Summary,
		ProfileImage: r.Object.ProfileImage,
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
	w.ProfilePicture = p.ProfileImage
}

func (w *ImportProfileResponse) Error(responseError string) struct {Error string `json:"error"`} {
    return struct {
        Error string `json:"error"`
    }{
        Error: responseError,
    }
}
