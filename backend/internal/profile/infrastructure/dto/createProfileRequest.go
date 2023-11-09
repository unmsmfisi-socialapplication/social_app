package dto

import (
	"encoding/json"
	"io"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

//creame un createprofilerequest
// ProfileID      int64
// 	UserID         int64
//     BirthDate      time.Time
//     Name           string
//     LastName       string
//     AboutMe        string
//     Genre          string
//     Address        string
//     Country        string
//     City           string
//     InsertionDate  time.Time
//     UpdateDate     time.Time
//     ProfilePicture string

type CreateProfileRequest struct {
	UserID         int64     `json:"user_id"`
	BirthDate      time.Time `json:"birth_date"`
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	AboutMe        string    `json:"about_me"`
	Genre          string    `json:"genre"`
	Address        string    `json:"address"`
	Country        string    `json:"country"`
	City           string    `json:"city"`
	InsertionDate  time.Time `json:"insertion_date"`
	UpdateDate     time.Time `json:"update_date"`
	ProfilePicture string    `json:"profile_picture"`
}

func NewCreateProfileRequest(body *io.ReadCloser) (*CreateProfileRequest, error) {
	var request CreateProfileRequest
	err := json.NewDecoder(*body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (cpr *CreateProfileRequest) ToProfile() *domain.Profile {
	return domain.NewProfile(
		
		cpr.UserID,
		cpr.BirthDate,
		cpr.Name,
		cpr.LastName,
		cpr.AboutMe,
		cpr.Genre,
		cpr.Address,
		cpr.Country,
		cpr.City,
		cpr.InsertionDate,
		cpr.UpdateDate,
		cpr.ProfilePicture,
	)
}

