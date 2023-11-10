package domain

import (
	"time"
)

type PostBase struct {
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
}

type Post struct {
	Id            int64
	InsertionDate time.Time
	UpdateDate    time.Time
	PostBase
}

type PostCreate struct {
	PostBase
}

type PostPaginationParams struct {
	Page  int
	Limit int
}

type PostPagination struct {
	Posts       []Post
	TotalCount  int
	CurrentPage int
}

func PostCreateToPost(p PostCreate) Post {
	post := Post{
		PostBase:      p.PostBase,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}
	return post
}
