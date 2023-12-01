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

	var data domain.UserRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = domain.ValidateUserRequest(data)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	dbUser, err := rh.useCase.RegisterUser(data.Email, data.Username, data.Password)
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
		}
	}

	reponse := domain.UserToUserResponse(*dbUser)

	utils.SendJSONResponse(w, http.StatusOK, "OK", reponse)
}
