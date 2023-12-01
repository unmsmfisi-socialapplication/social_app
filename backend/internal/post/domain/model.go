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

type PostResponse struct {
	Context string
	Type    string
	Object  Post
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

func PostToPostResponse(p Post) PostResponse {
	return PostResponse{
		Context: "https://www.w3.org/ns/activitystreams",
		Type:    "create",
		Object:  p,
	}
}

type SocialMediaPost struct {
    ID      int64
    Message string
}
