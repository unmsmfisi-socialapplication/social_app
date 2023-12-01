package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Phone    string
	Email    string
	Username string
	Password string
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"user_name"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func NewUser(email, username, password string) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func UserToUserResponse(u User) UserResponse {
	return UserResponse{Username: u.Username, Email: u.Email}
}

func ValidateUserRequest(u UserRequest) error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Username == "" {
		return errors.New("Username is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}
