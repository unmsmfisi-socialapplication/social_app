package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email string
	Password string
	Name string
	LastName string
	Birthday string

}

func NewUser(username, email, password, name, lastName, birthday string) (*User ,error){
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Email: email,
		Password: hashedPassword,
		Name: name,
		LastName: lastName,
		Birthday: birthday,
	},nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!= nil {
		return "", err
	}
	return string(hashedPassword), nil
}