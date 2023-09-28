package test

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

func TestNewUser(t *testing.T) {

	phone := "123456789"
	email := "testtest@gmail.com"
	username := "testtest"
	password := "testtest"

	user, err := domain.NewUser(phone, email, username, password)
	if err != nil {
		t.Fatal(err)
	}

	if user.Phone != phone {
		t.Errorf("expected %s, got %s", phone, user.Phone)
	}
	if user.Email != email {
		t.Errorf("expected %s, got %s", email, user.Email)
	}
	if user.User_name != username {
		t.Errorf("expected %s, got %s", username, user.User_name)
	}
	if user.Password == password {
		t.Errorf("no se encripto la contraseña")
	}

}