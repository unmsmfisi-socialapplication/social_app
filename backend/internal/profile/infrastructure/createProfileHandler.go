package infrastructure

import (
	"fmt"
	"net/http"
	"github.com/go-ap/activitypub"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/infrastructure/dto"
)

type createProfileHandler struct {
	createProfileUseCase *application.CreateProfileUseCase
}

func NewCreateProfileHandler(cpUseCase *application.CreateProfileUseCase) *createProfileHandler {
	return &createProfileHandler{cpUseCase}
}

func (cph *createProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	request, err := dto.NewCreateProfileRequest(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	profile := request.ToProfile()

	err = cph.createProfileUseCase.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response := dto.NewCreateProfileResponse(request.ToProfile())

	resp, err := activitypub.MarshalJSON(response.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"response": "%s", "data": %s}`, response.Response, string(resp))))
}