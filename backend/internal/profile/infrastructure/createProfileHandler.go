package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

type createProfileHandler struct {
	createProfileUseCase *application.CreateProfileUseCase
}

func NewCreateProfileHandler(cpUseCase *application.CreateProfileUseCase) *createProfileHandler {
	return &createProfileHandler{cpUseCase}
}

func (cph *createProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	
	var data struct {
		UserID         int64  `json:"user_id"`
		BirthDate      string `json:"birth_date"`
		Name           string `json:"name"`
		LastName       string `json:"last_name"`
		AboutMe        string `json:"about_me"`
		Genre          string `json:"genre"`
		Address        string `json:"address"`
		Country        string `json:"country"`
		City           string `json:"city"`
		ProfilePicture string `json:"profile_picture"`

	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.BirthDate == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, "BirthDate is required")))
		return
	}
	
	birthDate, err := time.Parse("02-01-2006", data.BirthDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, "Invalid BirthDate format")))
		return
	}
	profile:= domain.NewProfile(data.UserID, birthDate, data.Name, data.LastName, data.AboutMe, data.Genre, data.Address, data.Country, data.City, data.ProfilePicture)

	err = cph.createProfileUseCase.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	var outputData struct {
		UserID int64 `json:"user_id"`
		BirthDate      time.Time `json:"birth_date"`
		Name           string `json:"name"`
		LastName       string `json:"last_name"`
		AboutMe        string `json:"about_me"`
		Genre          string `json:"genre"`
		Country        string `json:"country"`




	}
	{
		outputData.UserID = data.UserID
		outputData.BirthDate = birthDate
		outputData.Name = data.Name
		outputData.LastName = data.LastName
		outputData.AboutMe = data.AboutMe
		outputData.Genre = data.Genre
		outputData.Country = data.Country
		
		
	}
	
	
	
	json.NewEncoder(w).Encode(outputData) 
	
}