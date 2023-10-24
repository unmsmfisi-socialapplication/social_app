package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
    Email    string
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

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
