package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type LoginHandler struct {
	useCase application.LoginUsecaseInterface
}

func NewLoginHandler(useCase application.LoginUsecaseInterface) *LoginHandler {
	return &LoginHandler{useCase: useCase}
}

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	if len(requestData) != 2 || requestData["username"] == nil || requestData["password"] == nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	username, ok1 := requestData["username"].(string)
	password, ok2 := requestData["password"].(string)
	if !ok1 || !ok2 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	token, err := lh.useCase.Authenticate(username, password)
	if err != nil {
		switch err {
		case application.ErrUserNotFound:
			utils.SendJSONResponse(w, http.StatusNotFound, "NOTFOUND", "User not found")
			return
		case application.ErrInvalidCredentials:
			utils.SendJSONResponse(w, http.StatusUnauthorized, "BADCREDENTIALS", "Invalid credentials")
			return
		default:
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during authentication")
			return
		}
	}
	jti, err := lh.useCase.ExtractJTI(token)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error extracting JTI from token")
		return
	}

	err = lh.useCase.StoreJTIForSession(username, jti)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error saving JTI in database")
		return
	}
	responseData := map[string]string{
		"token": token,
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", responseData)
}
