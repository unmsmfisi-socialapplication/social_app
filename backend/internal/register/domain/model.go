package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int
	Email         string
	Phone         string
	Username      string
	Password      string
	InsertionDate time.Time
}

type UserCreate struct {
	Email    string
	Username string
	Password string
}

type UserReponse struct {
	Id       int
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UserToUserReponse(user User) UserReponse {
	return UserReponse{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}
}

func NewUser(user UserCreate) (*User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:         user.Email,
		Username:      user.Username,
		Password:      hashedPassword,
		InsertionDate: time.Now(),
	}, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
