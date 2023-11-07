package application

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type MockRepo struct {
	insertedFollower *domain.Follower
	insertError      error
}

func (m *MockRepo) InsertNewFollower(follower *domain.Follower) (*domain.Follower, error) {
	return m.insertedFollower, m.insertError
}

func TestFollowProfile(t *testing.T) {
	mockRepo := &MockRepo{}

	useCase := NewFollowerUseCase(mockRepo)

	t.Run("Test con éxito", func(t *testing.T) {
		mockRepo.insertedFollower = &domain.Follower{Follower_profile_id: 1, Following_profile_id: 2}

		follower, err := useCase.FollowProfile(1, 2)

		if err != nil {
			t.Errorf("Se esperaba un error nulo, pero se obtuvo: %v", err)
		}

		if follower == nil {
			t.Error("Se esperaba un seguidor, pero se obtuvo nulo")
		}

		if follower.Follower_profile_id != 1 || follower.Following_profile_id != 2 {
			t.Errorf("El seguidor devuelto no coincide con las expectativas: %v", follower)
		}
	})

	t.Run("Test con error al insertar", func(t *testing.T) {
		mockRepo.insertError = errors.New("Error al insertar")

		follower, err := useCase.FollowProfile(1, 2)

		if err == nil {
			t.Error("Se esperaba un error, pero se obtuvo un error nulo")
		}

		if follower != nil {
			t.Error("Se esperaba un seguidor nulo, pero se obtuvo un seguidor")
		}
	})
}
