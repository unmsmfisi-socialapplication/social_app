package domain

import "errors"

type Follower struct {
	Follower_profile_id  int
	Following_profile_id int
}

func NewFollower(p_follower_profile_id, p_following_profile_id int) (*Follower, error) {
	if p_follower_profile_id < 0 || p_following_profile_id < 0 {
		return nil, errors.New("Los valores de perfil de seguidor o perfil seguido no pueden ser negativos")
	}
	return &Follower{
		Follower_profile_id:  p_follower_profile_id,
		Following_profile_id: p_following_profile_id,
	}, nil
}
