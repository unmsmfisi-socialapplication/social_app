package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
)

type LoginHandler struct {
	useCase *application.LoginUseCase
}

func NewLoginHandler(useCase *application.LoginUseCase) *LoginHandler {
	return &LoginHandler{useCase: useCase}
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, status string, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse := map[string]string{
		"status":   status,
		"response": response,
	}
	json.NewEncoder(w).Encode(jsonResponse)
}

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, "error", "Invalid request payload")
		return
	}

	isAuthenticated, err := lh.useCase.Authenticate(requestData.Username, requestData.Password)
	if err != nil {
		switch err {
		case application.ErrUserNotFound:
			sendJSONResponse(w, http.StatusNotFound, "NOTFOUND", "User not found")
			return
		case application.ErrInvalidCredentials:
			sendJSONResponse(w, http.StatusUnauthorized, "NOPASSWORD", "Invalid password")
			return
		default:
			sendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during authentication")
			return
		}
	}

	if isAuthenticated {
		sendJSONResponse(w, http.StatusOK, "OK", "Authentication successful")
	} else {
		sendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Authentication failed")
	}
}
