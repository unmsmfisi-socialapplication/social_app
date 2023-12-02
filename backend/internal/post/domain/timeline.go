package domain

import (
	"time"
)

type Timeline struct {
	Id        int64
	UserId    int64
	PostId    int64
	CreatedAt time.Time
}

type TimelineRes struct {
	UserId        int64     `json:"userId"`
	PostId        int64     `json:"postId"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Multimedia    string    `json:"multimedia"`
	InsertionDate time.Time `json:"insertionDate"`
	Username      string    `json:"username"`
}

type PaginatedRes struct {
	Results  []TimelineRes `json:"results"`
	Page     int           `json:"page"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
}

type QueryResult struct {
	Results []TimelineRes
}

func NewQueryResult(timelinePosts []TimelineRes) *QueryResult {
	return &QueryResult{Results: timelinePosts}
}
