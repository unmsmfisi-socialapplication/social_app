package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type RegisterUserHandler struct {
	useCase *application.RegistrationUseCase
}

func NewRegisterUserHandler(uc *application.RegistrationUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{useCase: uc}
}

func (rh *RegisterUserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user domain.UserCreate

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbuser, err := rh.useCase.RegisterUser(user)

	if err != nil {
		switch err {
		case application.ErrEmailInUse:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Email already in use")
			return
		case application.ErrFormat:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Not secure password")
			return
		case application.ErrPhone:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid phone format")
			return
		default:
			utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Unexpected error")
			return
		}
	}

	userResponse := domain.UserToUserReponse(*dbuser)

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", userResponse)
}
