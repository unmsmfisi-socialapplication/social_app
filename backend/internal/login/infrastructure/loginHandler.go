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

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	isAuthenticated, err := lh.useCase.Authenticate(requestData.Username, requestData.Password)
	if err != nil {
		switch err {
		case application.ErrUserNotFound:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
			return
		case application.ErrInvalidCredentials:
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid password"))
			return
		default:
			http.Error(w, "Error during authentication", http.StatusInternalServerError)
			return
		}
	}

	if isAuthenticated {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authentication successful"))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authentication failed"))
	}
}
