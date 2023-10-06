package dto

import (
	"fmt"
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

    err = activity.UnmarshalJSON(body)
    if err != nil {
        return nil, fmt.Errorf("Invalid JSON")
    }

    obj := activity.Object

    p, ok := obj.(*activitypub.Person)
    if !ok {
        return nil, fmt.Errorf("Invalid object type")
    }

    if p.Name == nil || p.Icon == nil || p.Summary == nil {
        return nil, fmt.Errorf("Missing required fields")
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
