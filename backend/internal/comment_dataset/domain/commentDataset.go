package domain

import "time"

type CommentDataset struct {
	InsertionDate time.Time `json:"insertion_date"`
	CommentID     int64     `json:"comment_id"`
	PostID        int64     `json:"post_id"`
	Username      string    `json:"username"`
	Comment       string    `json:"comment"`
	Name          string    `json:"name"`
	Lastname      string    `json:"lastname"`
	Genre         string    `json:"genre"`
	Country       string    `json:"country"`
	City          string    `json:"city"`
}
