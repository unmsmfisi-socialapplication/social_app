package application

import (
	"errors"
	"regexp"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

var (
	ErrEmailInUse    = errors.New("EMAIL_IN_USE")
	ErrUsernameInUse = errors.New("USERNAME_IN_USE")
	ErrFormat        = errors.New("INVALID_PASSWORD")
	ErrPhone         = errors.New("INVALID_PHONE")
)

type RegistrationUseCase struct {
	repo IUserRepository
}

func NewRegistrationUseCase(r IUserRepository) *RegistrationUseCase {
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

func (r *RegistrationUseCase) RegisterUser(email, username, password string) (*domain.User, error) {
	existingUser, err := r.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrEmailInUse
	}

	existingUser, err = r.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUsernameInUse
	}

	if !isValidPassword(password) {
		return nil, ErrFormat
	}

	newUser, err := domain.NewUser(email, username, password)
	if err != nil {
		return nil, err
	}

	newUser, err = r.repo.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
