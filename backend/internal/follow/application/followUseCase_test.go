package application

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

// Mock para el repositorio
type MockFollowerRepository struct{}

func (m *MockFollowerRepository) InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error) {
	return &domain.Follower{}, nil
}

func (m *MockFollowerRepository) IsFollowing(newFollower *domain.Follower) (*bool, error) {
	return new(bool), nil
}

func TestNewFollowerUseCase(t *testing.T) {
	// Crear un Mock del repositorio
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	// Ejecutar la función y asegurarse de que no se produzca un error
	newFollower, err := useCase.FollowProfile(1, 2)
	if err != nil {
		t.Errorf("Error inesperado: %v", err)
	}
	if newFollower == nil {
		t.Error("newFollower es nulo, pero se esperaba un valor no nulo")
	}
}

func TestFollowProfile(t *testing.T) {
	// Crear un Mock del repositorio
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	// Ejecutar la función y asegurarse de que no se produzca un error
	newFollower, err := useCase.FollowProfile(1, 2)
	if err != nil {
		t.Errorf("Error inesperado: %v", err)
	}
	if newFollower == nil {
		t.Error("newFollower es nulo, pero se esperaba un valor no nulo")
	}
}

func TestIsFollowing(t *testing.T) {
	// Crear un Mock del repositorio
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	// Ejecutar la función y asegurarse de que no se produzca un error
	isFollowing, err := useCase.IsFollowing(1, 2)
	if err != nil {
		t.Errorf("Error inesperado: %v", err)
	}
	if isFollowing == nil {
		t.Error("isFollowing es nulo, pero se esperaba un valor no nulo")
	}
}
