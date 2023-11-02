package dto

import (
	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ImportProfileResponse struct {
	Response string
	Data     *activitypub.Person
}

func NewImportProfileResponse(profile *domain.Profile) *ImportProfileResponse {

	person := &activitypub.Person{
		ID: activitypub.IRI(profile.Id_profile),
		Name: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.Username)},
		},
		Type: activitypub.PersonType,
		PreferredUsername: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.Username)},
		},
		Summary: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.Biography)},
		},
		Icon: activitypub.Image{
			URL: activitypub.IRI(profile.ProfileImage),
		},
	}

	return &ImportProfileResponse{
		Response: "Profile Imported Succesfully",
		Data:     person,
	}
}
