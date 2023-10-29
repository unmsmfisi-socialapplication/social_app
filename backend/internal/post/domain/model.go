package domain

import (
	"time"
)

type BasePost struct {
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
}

type Post struct {
	BasePost      BasePost
	Id            int64
	InsertionDate time.Time
	UpdateDate    time.Time
}

type CreatePost struct {
	BasePost
}

func CreatePostToPost(p CreatePost) Post {
	post := Post{
		BasePost: BasePost{
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
