package application

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

var jwtKey string

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token has expired")
	ErrUnauthorized = errors.New("unauthorized access")
)

type AuthUseCaseInterface interface {
	ValidateToken(tokenString string) (*domain.User, error)
	CheckValidSession(username string) (bool, error)
	GetJTIByUsername(username string) (string, error)
	Logout(username string) error
}

type TokenGenerator interface {
	GenerateToken(username string) (string, error)
}

type AuthUseCase struct {
	tokenGen    TokenGenerator
	sessionRepo domain.AuthSessionRepository
	jwtKey      string
}

func NewAuthUseCase(sr domain.AuthSessionRepository) *AuthUseCase {
	jwtKey = os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		panic("JWT_SECRET_KEY is not set in environment variables")
	}
	return &AuthUseCase{sessionRepo: sr}
}

func (a *AuthUseCase) ValidateToken(tokenString string) (*domain.User, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, ErrInvalidToken
	}

	storedJTI, err := a.GetJTIByUsername(claims.Subject)
	if err != nil {
		return nil, err
	}

	if claims.Id != storedJTI {
		return nil, ErrInvalidToken
	}

	tokenExpiryTime := time.Unix(claims.ExpiresAt, 0)

	now := time.Now()

	if now.After(tokenExpiryTime) {
		return nil, ErrTokenExpired
	}

	user := &domain.User{Username: claims.Subject}
	return user, nil
}

func (a *AuthUseCase) CheckValidSession(username string) (bool, error) {
	return a.sessionRepo.CheckValidSession(username)
}
func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (a *AuthUseCase) GetJTIByUsername(username string) (string, error) {
	return a.sessionRepo.GetJTIByUsername(username)
}

func (a *AuthUseCase) Logout(username string) error {
	return a.sessionRepo.InvalidateSession(username)
}
