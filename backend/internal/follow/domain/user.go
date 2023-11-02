package domain

type Follow struct {
	Profile_id  int
	Follower_id int
}

func NewFollower(p_profile_id, p_follower_id int) (*Follow, error) {

	return &Follow{
		Profile_id:  p_profile_id,
		Follower_id: p_follower_id,
	}, nil
}
