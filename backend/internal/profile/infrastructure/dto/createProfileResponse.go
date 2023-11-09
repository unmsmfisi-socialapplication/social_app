package dto

import (
	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type CreateProfileResponse struct {
	Response string
	Data     *activitypub.Person
}

func NewCreateProfileResponse(profile *domain.Profile) *CreateProfileResponse {
	
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

	return &CreateProfileResponse{
		Response: "Profile Created Succesfully",
		Data:     person,
	}
}