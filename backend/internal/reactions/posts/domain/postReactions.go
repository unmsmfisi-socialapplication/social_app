package domain

import "gorm.io/gorm"

type PostReaction struct {
	gorm.Model
	ReactionID int64 `json:"reactionID"`
	UserID     int64 `json:"userID"`
	PostID     int64 `json:"postID"`
}
