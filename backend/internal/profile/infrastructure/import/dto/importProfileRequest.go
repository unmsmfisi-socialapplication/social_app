package dto

import (
	"errors"
	"io"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ImportProfileRequest struct {
    Person *activitypub.Person
}

func NewImportProfileRequest(requestBody *io.ReadCloser) (*ImportProfileRequest, error) {
    var activity activitypub.Activity

    body, err := io.ReadAll(*requestBody)
    if err != nil {
        return nil, err
    }

    activity.UnmarshalJSON(body)

    obj := activity.Object

    p := obj.(*activitypub.Person)
    if p.Name == nil || p.Icon == nil || p.Summary == nil {
        return nil, errors.New("invalid actor")
    }

    return &ImportProfileRequest{Person: p}, nil
}

func (r *ImportProfileRequest) ToProfile() *domain.Profile {
	return &domain.Profile{
		Id_profile:   r.Person.GetID().String(),
		Username:     r.Person.Name.String(),
		Biography:    r.Person.Summary.First().String(),
		ProfileImage: r.Person.Icon.GetLink().String(),
	}
}
