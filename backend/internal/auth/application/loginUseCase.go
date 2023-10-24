package application

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type LoginUsecaseInterface interface {
	Authenticate(username, password string) (bool, error)
}

type LoginUseCase struct {
	repo IUserRepository
}

func NewLoginUseCase(r IUserRepository) *LoginUseCase {
	return &LoginUseCase{repo: r}
}

func (l *LoginUseCase) Authenticate(username, password string) (bool, error) {
	user, err := l.repo.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

    if user == nil {
        return false, ErrUserNotFound
    }

	if username != user.Username {
		return false, ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		return false, ErrInvalidCredentials
	}

	return true, nil
}
