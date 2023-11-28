package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

var (
	ErrFollowHimself        = errors.New("PROFILE_FOLLOW_HIMSELF")
	ErrForeignKeyConstraint = errors.New("VIOLATES_FOREIGN_KEY_CONSTRAINT")
)

type FollowerRepository interface {
	InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error)
	IsFollowing(newFollower *domain.Follower) (*bool, error)
}

type FollowerUseCase struct {
	repo FollowerRepository
}

func NewFollowerUseCase(r FollowerRepository) *FollowerUseCase {
	return &FollowerUseCase{repo: r}
}

func (r *FollowerUseCase) FollowProfile(p_follower_profile_id, p_following_profile_id int) (*domain.Follower, error) {
	if p_follower_profile_id == p_following_profile_id {
		return nil, ErrFollowHimself
	}

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

func (r *FollowerUseCase) IsFollowing(p_follower_profile_id, p_following_profile_id int) (*bool, error) {
	if p_follower_profile_id == p_following_profile_id {
		return nil, ErrFollowHimself
	}

	newFollower, err := domain.NewFollower(p_follower_profile_id, p_following_profile_id)
	if err != nil {
		return nil, err
	}

	isFollowing, err := r.repo.IsFollowing(newFollower)
	if err != nil {
		return nil, err
	}
	return isFollowing, nil
}
