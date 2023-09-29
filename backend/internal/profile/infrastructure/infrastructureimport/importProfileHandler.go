package infrastructureimport

import (
	"fmt"
	"net/http"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/applicationimport"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/infrastructureimport/dto"
)

type ImportProfileHandler struct {
	importProfileUseCase *applicationimport.ImportProfileUseCase
}

func NewImportProfileHandler(ipUseCase *applicationimport.ImportProfileUseCase) *ImportProfileHandler {
	return &ImportProfileHandler{ipUseCase}
}

func (iph *ImportProfileHandler) ImportProfile(w http.ResponseWriter, r *http.Request) {
    request, err := dto.NewImportProfileRequest(&r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
    }

    profile := request.ToProfile()

    err = iph.importProfileUseCase.ImportProfile(profile)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
    }

    response := dto.NewImportProfileResponse(request.Person)

    resp, err := activitypub.MarshalJSON(response.Data)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(fmt.Sprintf(`{"response": "%s", "data": %s}`, response.Response, string(resp))))
}
