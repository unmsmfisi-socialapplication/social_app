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
		ID: activitypub.IRI(profile.Name),
		Name: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.Name)},
		},
		Type: activitypub.PersonType,
		PreferredUsername: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.Name)},
		},
		Summary: activitypub.NaturalLanguageValues{
			{Ref: activitypub.NilLangRef, Value: activitypub.Content(profile.AboutMe)},
		},
		Icon: activitypub.Image{
			URL: activitypub.IRI(profile.ProfilePicture),
		},
	}

	return &ImportProfileResponse{
		Response: "Profile Imported Succesfully",
		Data:     person,
	}
}
