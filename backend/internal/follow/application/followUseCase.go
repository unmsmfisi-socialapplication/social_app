package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

var (
	ErrFollowHimself        = errors.New("PROFILE_FOLLOW_HIMSELF")
	ErrForeignKeyConstraint = errors.New("VIOLATES_FOREIGN_KEY_CONSTRAINT")
	ErroEqualProfile        = errors.New("EQUAL_PROFILES")
	ErrNullParameters       = errors.New("NULL_PARAMETERS")
	ErrNegativeParameters   = errors.New("NEGATIVE_PARAMETERS")
)

type FollowerRepository interface {
	InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error)
	IsFollowing(newFollower *domain.Follower) (*bool, error)
	ViewCommonFollowers(p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error)
	ProfileFollowers(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error)
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

func (r *FollowerUseCase) CommonFollowers(p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {

	if p_own_profile_id < 0 || p_viewed_profile_id < 0 || p_page_size < 0 || p_page_number < 0 {
		return nil, ErrNegativeParameters
	}

	if p_own_profile_id == 0 || p_viewed_profile_id == 0 || p_page_size == 0 || p_page_number == 0 {
		return nil, ErrNullParameters
	}

	if p_own_profile_id == p_viewed_profile_id {
		return nil, ErroEqualProfile
	}

	newFollowerDataList, err := r.repo.ViewCommonFollowers(p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number)
	if err != nil {
		return nil, err
	}
	return newFollowerDataList, nil
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

func (r *FollowerUseCase) ProfileFollowers(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {

	if p_profile_id < 0 || p_page_size < 0 || p_page_number < 0 {
		return nil, ErrNegativeParameters
	}

	if p_profile_id == 0 || p_page_size == 0 || p_page_number == 0 {
		return nil, ErrNullParameters
	}

	newFollowerDataList, err := r.repo.ProfileFollowers(p_profile_id, p_page_size, p_page_number)
	if err != nil {
		return nil, err
	}
	return newFollowerDataList, nil
}

func (r *FollowerUseCase) ProfileFollowingProfiles(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {

	if p_profile_id < 0 || p_page_size < 0 || p_page_number < 0 {
		return nil, ErrNegativeParameters
	}

	if p_profile_id == 0 || p_page_size == 0 || p_page_number == 0 {
		return nil, ErrNullParameters
	}

	newFollowerDataList, err := r.repo.ProfileFollowers(p_profile_id, p_page_size, p_page_number)
	if err != nil {
		return nil, err
	}
	return newFollowerDataList, nil
}
