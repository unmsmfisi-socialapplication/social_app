package application

import "github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"

type IUserRepository interface {
	GetUserByUsername(username string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	InsertUser(newUser *domain.User) (*domain.User, error)
}
