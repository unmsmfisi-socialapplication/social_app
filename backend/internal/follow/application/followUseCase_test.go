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
		t.Errorf("Unexpected error: %v", err)
	}
	if newFollower == nil {
		t.Error("newFollower is null, but a non-null value was expected.")
	}
}

func TestFollowProfile(t *testing.T) {
	// Crear un Mock del repositorio
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	// Ejecutar la función y asegurarse de que se produzca un error
	newFollower, err := useCase.FollowProfile(-1, -2)
	if err != nil {
		t.Logf("Expected error: %v", err)
	} else {
		t.Error("Expected an error, but got none.")
	}

	// Asegurarse de que newFollower es nulo
	if newFollower != nil {
		t.Error("newFollower is not null, but a null value was expected.")
	}
}

func TestIsFollowing(t *testing.T) {
	// Crear un Mock del repositorio
	repo := &MockFollowerRepository{}
	useCase := NewFollowerUseCase(repo)

	// Caso 1: Verificar si el usuario 1 sigue al usuario 2 (esperando falso)
	isFollowing, err := useCase.IsFollowing(1, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *isFollowing != false {
		t.Error("User 1 was expected not to follow user 2, but IsFollowing is true.")
	}

	// Caso 2: Verificar si el usuario 2 sigue al usuario 1 (esperando falso)
	isFollowing, err = useCase.IsFollowing(2, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *isFollowing != false {
		t.Error("It was expected that user 2 would not follow user 1, but IsFollowing is true.")
	}

	// Caso 3: Verificar si el usuario -2 sigue al usuario 1 (esperando error)
	_, err = useCase.IsFollowing(-2, 1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	// Caso 4: Verificar si el usuario 2 sigue al usuario -1 (esperando error)
	_, err = useCase.IsFollowing(2, -1)
	if err == nil {
		t.Error("An error was expected as the user provided is negative.")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}

	// Caso 5: Verificar si el usuario -2 sigue al usuario -1 (esperando error)
	_, err = useCase.IsFollowing(-2, -1)
	if err == nil {
		t.Error("Se esperaba un error ya que el usuario proporcionado es negativo")
	} else {
		expectedErrMsg := "Follower profile or followed profile values cannot be negative."
		if err.Error() != expectedErrMsg {
			t.Errorf("The error message '%s' was expected, but '%s' was obtained.", expectedErrMsg, err.Error())
		}
	}
}
