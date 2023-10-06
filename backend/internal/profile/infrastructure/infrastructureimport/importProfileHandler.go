package infrastructureimport

import (
	"fmt"
	"net/http"

	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application/applicationimport"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/infrastructureimport/dto"
)

type importProfileHandler struct {
	importProfileUseCase *applicationimport.ImportProfileUseCase
}

func NewImportProfileHandler(ipUseCase *applicationimport.ImportProfileUseCase) *importProfileHandler {
	return &importProfileHandler{ipUseCase}
}

func (iph *importProfileHandler) ImportProfile(w http.ResponseWriter, r *http.Request) {
    request, err := dto.NewImportProfileRequest(&r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
        return
    }

    profile := request.ToProfile()

    err = iph.importProfileUseCase.ImportProfile(profile)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
        return
    }

    response := dto.NewImportProfileResponse(request.Person)

    resp, err := activitypub.MarshalJSON(response.Data)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf(`{"response": "%s", "data": %s}`, response.Response, string(resp))))
}
