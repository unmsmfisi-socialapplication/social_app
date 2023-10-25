package domain

import (
	"time"
)

type Post struct {
	Id            int64
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
	InsertionDate time.Time
	UpdateDate    time.Time
}

type CreatePost struct {
	UserId        int64  `json:"userId"`
	Title         string `json:"title" db:"title" validate:"max=100"`
	Description   string `json:"description" db:"description" validate:"max=1000"`
	HasMultimedia bool   `json:"hasMultimedia"`
	Public        bool   `json:"public"`
	Multimedia    string `json:"multimedia" db:"multimedia" validate:"max=1000"`
}

func CreatePostToPost(p CreatePost) Post {
	post := Post{
		UserId:        p.UserId,
		Title:         p.Title,
		Description:   p.Description,
		HasMultimedia: p.HasMultimedia,
		Public:        p.Public,
		Multimedia:    p.Multimedia,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	return post
}
