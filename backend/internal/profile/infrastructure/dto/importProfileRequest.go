package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type ObjectInput struct {
	Id                string         `json:"id"`
	Type              string         `json:"type"`
	Name              string         `json:"name"`
	PreferredUsername string         `json:"preferredUsername"`
	Summary           string         `json:"summary"`
	ProfileImage      string         `json:"profileImage"`
	CoverImage        string         `json:"coverImage"`
	Endpoints         EndpointsInput `json:"endpoints"`
	ImportedFrom      string         `json:"importedFrom"`
}

type EndpointsInput struct {
	SharedInbox string `json:"sharedInbox"`
}

type ImportProfileRequest struct {
	Context json.RawMessage `json:"@context"`
	Type    string          `json:"type"`
	Actor   string          `json:"actor"`
	Name    string          `json:"name"`
	Object  ObjectInput     `json:"object"`
	To      []string        `json:"to"`
	CC      string          `json:"cc"`
}

func NewImportProfileRequest(requestBody *io.ReadCloser) (*ImportProfileRequest, error) {
	var request ImportProfileRequest

	err := json.NewDecoder(*requestBody).Decode(&request)
	if err != nil {
		return nil, err
	}

	fmt.Println(request)

	err = request.Validate()
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *ImportProfileRequest) ToProfile() *domain.Profile {
	var name, lastname string

	splittedName := strings.Split(r.Object.Name, " ")
	if len(splittedName) > 1 {
		name = splittedName[0]
		lastname = splittedName[1]
	} else {
		name = splittedName[0]
		lastname = ""
	}

	return domain.NewProfile(
		r.Object.Id,
		name,
		lastname,
		r.Object.PreferredUsername,
		r.Object.ProfileImage,
		r.Object.Summary,
	)
}

func (r *ImportProfileRequest) Validate() error {
	requiredFields := []string{"@context", "type", "actor", "name", "object", "to"}

	for _, field := range requiredFields {
		if field == "object" {
			if r.Object.Id == "" || r.Object.Type == "" || r.Object.Name == "" ||
				r.Object.PreferredUsername == "" || r.Object.Summary == "" ||
				r.Object.ProfileImage == "" || r.Object.CoverImage == "" ||
				r.Object.Endpoints.SharedInbox == "" || r.Object.ImportedFrom == "" {

				return fmt.Errorf("invalid %s field", field)
			}
		} else if field == "to" {
			if len(r.To) == 0 {
				return fmt.Errorf("invalid %s field", field)
			}
		} else {
			if r.Context == nil || r.Type == "" || r.Actor == "" || r.Name == "" {
				return fmt.Errorf("invalid %s field", field)
			}
		}
	}

	if r.Type != "Profile" {
		return fmt.Errorf("invalid type field")
	}

	if r.Object.Type != "Person" {
		return fmt.Errorf("invalid object.type field")
	}

	return nil
}
