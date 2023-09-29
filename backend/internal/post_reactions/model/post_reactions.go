package model

import "time"

type PostReaction struct {
	// PostReactionID uint
	PostReactionID int
	PostID         int
	UserID         int
	InsertionDate  time.Time
	// UpdateDate     time.Time
}
