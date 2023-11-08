package domain

import "time"

type CommentDataset struct {
	InsertionDate time.Time `json:"insertionDate"`
	CommentID     int64     `json:"commentID"`
	PostID        int64     `json:"postID"`
	Username      string    `json:"username"`
	Comment       string    `json:"comment"`
	Name          string    `json:"name"`
	Lastname      string    `json:"lastname"`
	Genre         string    `json:"genre"`
	Country       string    `json:"country"`
	City          string    `json:"city"`
}
