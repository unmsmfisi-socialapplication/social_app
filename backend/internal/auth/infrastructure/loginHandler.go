package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type LoginHandler struct {
	useCase application.LoginUsecaseInterface
}

func NewLoginHandler(useCase application.LoginUsecaseInterface) *LoginHandler {
	return &LoginHandler{useCase: useCase}
}

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var raw json.RawMessage
	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(raw, &requestData)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	var extra map[string]interface{}
	err = json.Unmarshal(raw, &extra)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	delete(extra, "username")
	delete(extra, "password")
	if len(extra) != 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	isAuthenticated, err := lh.useCase.Authenticate(requestData.Username, requestData.Password)
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
			fmt.Println(err.Error())
			return
		}
	}

	if isAuthenticated {
		utils.SendJSONResponse(w, http.StatusOK, "OK", "Authentication successful")
	} else {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Authentication failed")
	}
}
