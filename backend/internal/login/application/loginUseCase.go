package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/domain"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type LoginUsecaseInterface interface {
	Authenticate(username, password string) (bool, error)
}

type UserRepository interface {
	GetUserByUsername(username string) (*domain.User, error)
}

type LoginUseCase struct {
	repo UserRepository
}

func NewLoginUseCase(r UserRepository) *LoginUseCase {
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

	if !user.ComparePassword(password) {
		return false, ErrInvalidCredentials
	}

	return true, nil
}
