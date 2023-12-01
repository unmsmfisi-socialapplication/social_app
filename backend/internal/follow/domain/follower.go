package domain

import "errors"

type Follower struct {
	Follower_profile_id  int
	Following_profile_id int
}

func NewFollower(p_follower_profile_id, p_following_profile_id int) (*Follower, error) {
	if p_follower_profile_id < 0 || p_following_profile_id < 0 {
		return nil, errors.New("Follower profile or followed profile values cannot be negative.")
	}
	return &Follower{
		Follower_profile_id:  p_follower_profile_id,
		Following_profile_id: p_following_profile_id,
	}, nil
}

type FollowerData struct {
	ProfileID      int
	UserID         int
	Name           string
	LastName       string
	ProfilePicture string
	UserName       string
}

type FollowerDataList []FollowerData

func NewFollowerDataList(dataList ...FollowerData) FollowerDataList {
	return FollowerDataList(dataList)
}
