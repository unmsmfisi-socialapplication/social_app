package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
)

type ImportProfileHandler struct {
    importProfileUseCase *application.ImportProfileUseCase
}

func NewImportProfileHandler(profileRepository *application.ImportProfileUseCase) *ImportProfileHandler {
    return &ImportProfileHandler{profileRepository}
}

func (iph *ImportProfileHandler) ImportProfile(w http.ResponseWriter, r *http.Request) {
    var requestData ImportProfileRequest

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Println(err)
        return
    }

    p, err := iph.importProfileUseCase.ImportProfile(requestData.ToProfile())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println(err)
        return
    }

    var responseData ImportProfileResponse
    responseData.FromProfile(p)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(responseData)
}
