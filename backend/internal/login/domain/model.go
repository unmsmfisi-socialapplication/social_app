package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	Role     string
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
