package infrastructure

import (
	"encoding/json"
	"net/http"

	login_app "github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type RegisterUserHandler struct {
	registerUseCase *application.RegistrationUseCase
	authUseCase     *login_app.LoginUseCase
}

func NewRegisterUserHandler(ru *application.RegistrationUseCase, au *login_app.LoginUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{registerUseCase: ru, authUseCase: au}
}

func (rh *RegisterUserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data struct {
		Email    string `json:"email"`
		Username string `json:"user_name"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, er := rh.registerUseCase.RegisterUser(data.Email, data.Username, data.Password)
	if er != nil {
		switch er {
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

	tokenge, err := rh.authUseCase.GenerateToken(data.Username, "role")
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Error generating token")
		return
	}
	jti, err := rh.authUseCase.ExtractJTI(tokenge)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Error extracting jti")
		return
	}
	err = rh.authUseCase.StoreJTIForSession(data.Username, jti)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Error storing jti")
		return
	}

	var outputData struct {
		Email    string `json:"email"`
		Username string `json:"user_name"`
		Token    string `json:"token_result"`
	}
	{
		outputData.Email = data.Email
		outputData.Username = data.Username
		outputData.Token = tokenge
	}
	json.NewEncoder(w).Encode(outputData)
}
