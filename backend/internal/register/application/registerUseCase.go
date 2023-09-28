package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

var (
	ErrUserNotFound       = errors.New("usuario no encontrado")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
)

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
}

type RegistrationUseCase struct {
	repo UserRepository
}

func NewRegistrationUseCase(r UserRepository) *RegistrationUseCase {
	return &RegistrationUseCase{repo: r}
}

func (r *RegistrationUseCase) Register(phone, email, username,password string) (*domain.User, error) {
	
	existingUser, err := r.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	
	newUser, err := domain.NewUser(phone, email, username,password)
	if err != nil {
		return nil, err
	}

	
	return newUser, nil
}
