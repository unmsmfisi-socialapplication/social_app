package domain

type Follower struct {
	Follower_profile_id  int
	Following_profile_id int
}

func NewFollower(p_follower_profile_id, p_following_profile_id int) (*Follower, error) {
	return &Follower{
		Follower_profile_id:  p_follower_profile_id,
		Following_profile_id: p_following_profile_id,
	}, nil
}
