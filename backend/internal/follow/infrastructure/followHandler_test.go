package infrastructure

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/application"
)

func TestNewFollowerHandler(t *testing.T) {
	useCase := new(application.FollowerUseCase)
	handler := NewFollowerHandler(useCase)

	if handler == nil {
		t.Errorf("NewFollowerHandler was incorrect, got: nil")
	}
}
