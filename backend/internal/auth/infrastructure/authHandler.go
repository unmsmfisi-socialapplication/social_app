package infrastructure

import (
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type AuthHandler struct {
	useCase application.AuthUseCaseInterface
}

func NewAuthHandler(useCase application.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{useCase: useCase}
}

func (ah *AuthHandler) ValidateUserToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "No token provided")
		return
	}

	user, err := ah.useCase.ValidateToken(tokenString)
	if err != nil {
		switch err {
		case application.ErrInvalidToken:
			utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Invalid token")
			return
		case application.ErrUnauthorized:
			utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Unauthorized access")
			return
		default:
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during token validation")
			return
		}
	}

	responseData := struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}{
		Username: user.Username,
		Role:     user.Role,
	}
	utils.SendJSONResponse(w, http.StatusOK, "OK", responseData)
}

func (ah *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "No token provided")
		return
	}

	user, err := ah.useCase.ValidateToken(tokenString)
	if err != nil {
		switch err {
		case application.ErrInvalidToken:
			utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Invalid token")
			return
		case application.ErrUnauthorized:
			utils.SendJSONResponse(w, http.StatusUnauthorized, "ERROR", "Unauthorized access")
			return
		default:
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during token validation")
			return
		}
	}

	err = ah.useCase.Logout(user.Username)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during logout process")
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", "User logged out successfully")
}
