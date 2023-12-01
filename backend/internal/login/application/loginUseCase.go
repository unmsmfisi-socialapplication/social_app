package application

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/domain"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type LoginUsecaseInterface interface {
	Authenticate(username, password string) (string, error)
	StoreJTIForSession(username string, jti string) error
	ExtractJTI(tokenString string) (string, error)
}

type UserRepository interface {
	GetUserByUsername(username string) (*domain.User, error)
	StoreJTIForSession(username string, jti string) error
	InsertSession(username string) error
	UpdateSession(username string) error
	CheckExistingSession(username string) (bool, error)
}

type LoginUseCase struct {
	repo   UserRepository
	jwtKey string
}

func NewLoginUseCase(r UserRepository) *LoginUseCase {
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		panic("JWT_SECRET_KEY is not set in environment variables")
	}
	return &LoginUseCase{repo: r, jwtKey: jwtKey}
}

func (l *LoginUseCase) Authenticate(username, password string) (string, error) {
	user, err := l.repo.GetUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUserNotFound
		}
		return "", err
	}

	if !user.ComparePassword(password) {
		return "", ErrInvalidCredentials
	}

	if err != nil {
		return "", err
	}

	token, err := l.GenerateToken(user.Username, user.Role)

	exists, err := l.repo.CheckExistingSession(username)
	if err != nil {
		return "", err
	}

	if exists {
		err = l.repo.UpdateSession(username)
		if err != nil {
			return "", err
		}
	} else {
		err = l.repo.InsertSession(username)
		if err != nil {
			return "", err
		}
	}
	if err != nil {
		return "", err
	}

	return token, nil
}

func (l *LoginUseCase) GenerateToken(username, role string) (string, error) {
	jti := uuid.New().String()
	claims := &jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"jti": jti,
		"rol": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.jwtKey))
}
func (l *LoginUseCase) StoreJTIForSession(username string, jti string) error {
	return l.repo.StoreJTIForSession(username, jti)
}
func (l *LoginUseCase) ExtractJTI(tokenString string) (string, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.jwtKey), nil
	})

	if err != nil {
		return "", err
	}
	return claims.Id, nil
}
