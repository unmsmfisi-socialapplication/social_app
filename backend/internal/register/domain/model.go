package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint
	Phone    string
	Name     string
	Email    string
	Username string
	Password string
	Photo    string
}

type UserResponse struct {
	Id                uint
	Type              string
	Name              string
	Username          string
	Email             string
	Photo             string
	Phone             string
	Role              uint
	PreferredUsername string
	Summary           string
}

type UserRequest struct {
	Email    string
	Username string
	Photo    string
	Name     string
	Password string
	Phone    string
}

func NewUser(userReq UserRequest) (*User, error) {
	hashedPassword, err := HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		Phone:    userReq.Phone,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Username: userReq.Username,
		Password: hashedPassword,
		Photo:    userReq.Photo,
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
	return UserResponse{
		Id:                u.Id,
		Username:          u.Username,
		Name:              u.Name,
		Type:              "user",
		Role:              1,
		Email:             u.Email,
		Phone:             u.Phone,
		PreferredUsername: u.Username,
		Photo:             u.Photo,
		Summary:           "I am a developer"}
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

	if u.Phone == "" {
		return errors.New("Phone is required")
	}

	if u.Name == "" {
		return errors.New("Name is required")
	}
	

	return nil
}
