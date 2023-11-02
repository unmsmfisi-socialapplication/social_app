package application

import (
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type FollowerRepository interface {
	InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error)
}

type FollowerUseCase struct {
	repo FollowerRepository
}

func NewFollowerUseCase(r FollowerRepository) *FollowerUseCase {
	return &FollowerUseCase{repo: r}
}

func (r *FollowerUseCase) FollowProfile(p_follower_profile_id, p_following_profile_id int) (*domain.Follower, error) {
	if p_follower_profile_id != p_following_profile_id {
		newFollower, err := domain.NewFollower(p_follower_profile_id, p_following_profile_id)
		if err != nil {
			return nil, err
		}
		newFollower, err = r.repo.InsertNewFollower(newFollower)
		if err != nil {
			return nil, err
		}
		return newFollower, nil
	}
	return nil, fmt.Errorf("the user cannot follow himself")
}
