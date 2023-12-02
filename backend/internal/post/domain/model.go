package domain

import (
	"errors"
	"time"
)

var (
	ErrReporterUserDoesNotExist   = errors.New("Reporter user does not exist")
	ErrPostNotFound               = errors.New("Post not found")
	ErrUserHasAlreadyReportedPost = errors.New("User has already reported this post")
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

// POST UPDATE //
type PostUpdate struct {
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
}

// // // // // //

// POST REPORT //

type PostReport struct {
	ReportId   int64
	PostId     int64
	ReportedBy string
	Reason     string
	ReportDate time.Time
}

// // // // // //

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
