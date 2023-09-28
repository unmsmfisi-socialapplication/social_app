package domain

import (


	"golang.org/x/crypto/bcrypt"
)

type User struct {
	
	Phone string
	Email string
	User_name string
	Password string

}

func NewUser( phone, email, username,password string) (*User ,error){
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	
	return &User{
		
		Phone: phone,
		Email: email,
		User_name: username,
		Password: hashedPassword,
	
	},nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!= nil {
		return "", err
	}
	return string(hashedPassword), nil
}