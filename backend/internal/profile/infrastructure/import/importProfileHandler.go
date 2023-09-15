package infrastructure_import

import (
	"encoding/json"
	"log"
	"net/http"

	application_import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/import"
)

type ImportProfileHandler struct {
    importProfileUseCase *application_import.ImportProfileUseCase
}

func NewImportProfileHandler(profileRepository *application_import.ImportProfileUseCase) *ImportProfileHandler {
    return &ImportProfileHandler{profileRepository}
}

func (iph *ImportProfileHandler) ImportProfile(w http.ResponseWriter, r *http.Request) {
    var requestData ImportProfileRequest
    var responseData ImportProfileResponse

    type responseError struct {
        Error string `json:"error"`
    }

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(responseData.Error(err.Error()))
        log.Println(err)
        return
    }

    err = requestData.Validate()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(responseData.Error(err.Error()))
        log.Println(err)
        return
    }

    p, err := iph.importProfileUseCase.ImportProfile(requestData.ToProfile())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(responseData.Error(err.Error()))
        log.Println(err)
        return
    }

    responseData.FromProfile(p)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(responseData)
}
