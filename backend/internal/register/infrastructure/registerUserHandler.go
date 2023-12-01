package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"

	login_app "github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	login_infra "github.com/unmsmfisi-socialapplication/social_app/internal/login/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)
type RegisterUserHandler struct {
	useCase *application.RegistrationUseCase
}

func NewRegisterUserHandler(uc *application.RegistrationUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{useCase: uc}
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
	dbInstance := database.GetDB()
	loginRepo := login_infra.NewUserDBRepository(dbInstance)

	loginUseCase := login_app.NewLoginUseCase(loginRepo)
	tokenge, err := loginUseCase.GenerateToken(data.Username, "role")
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Error generating token")
		return
	}

	_, er := rh.useCase.RegisterUser(data.Email, data.Username, data.Password)
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
	loginRepo.InsertSession(data.Username)
	jti,err:= loginUseCase.ExtractJTI(tokenge)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Error extracting jti")
		return
	}
	loginUseCase.StoreJTIForSession(data.Username, jti)
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
