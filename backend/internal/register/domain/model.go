package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Phone    string
	Email    string
	Username string
	Password string
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
