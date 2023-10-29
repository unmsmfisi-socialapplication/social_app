package application

import (
	"errors"
	"regexp"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

var (
	ErrEmailInUse         = errors.New("email is already in use")
	ErrFormat             = errors.New("invalid passowrd")
	ErrPhone              = errors.New("invalid phone")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	InsertUser(newUser *domain.User) (*domain.User, error)
}

type RegistrationUseCase struct {
	repo UserRepository
}

func NewRegistrationUseCase(r UserRepository) *RegistrationUseCase {
	return &RegistrationUseCase{repo: r}
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[!@#$%^&*()_+{}\[\]:;<>,.?~\\-]`).MatchString(password) {
		return false
	}

	return true
}

func (r *RegistrationUseCase) RegisterUser(user domain.UserCreate) (*domain.User, error) {

	existingUser, err := r.repo.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrEmailInUse
	}

	if !isValidPassword(user.Password) {
		return nil, ErrFormat
	}

	newUser, err := domain.NewUser(user)

	if err != nil {
		return nil, err
	}

	newUser, err = r.repo.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
