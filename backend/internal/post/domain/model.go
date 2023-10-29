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
	BasePost      PostBase
	Id            int64
	InsertionDate time.Time
	UpdateDate    time.Time
}

type PostCreate struct {
	PostBase
}

func PostCreateToPost(p PostCreate) Post {
	post := Post{
		BasePost: PostBase{
			UserId:        p.UserId,
			Title:         p.Title,
			Description:   p.Description,
			HasMultimedia: p.HasMultimedia,
			Public:        p.Public,
			Multimedia:    p.Multimedia,
		},
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	return post
}
