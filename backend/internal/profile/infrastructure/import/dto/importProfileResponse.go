package dto

import "github.com/go-ap/activitypub"

type ImportProfileResponse struct {
    Response string
    Data *activitypub.Person
}

func NewImportProfileResponse(person *activitypub.Person) *ImportProfileResponse {
    return &ImportProfileResponse{
        Response: "Profile Imported Succesfully",
        Data: person,
    }
}
