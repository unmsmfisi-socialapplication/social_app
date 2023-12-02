package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestPostCreateToPost(t *testing.T) {
	postCreate := PostCreate{
		PostBase: PostBase{
			UserId:        1,
			Title:         "Test Title",
			Description:   "Test Description",
			HasMultimedia: true,
			Public:        true,
			Multimedia:    "test.jpg",
		},
	}

	expectedPost := Post{
		PostBase:      postCreate.PostBase,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	createdPost := PostCreateToPost(postCreate)

	if !reflect.DeepEqual(createdPost, expectedPost) {
		t.Error("Converting PostCreate to Post did not produce the expected result")
	}

	if !createdPost.InsertionDate.After(time.Now().Add(-time.Second)) || !createdPost.UpdateDate.After(time.Now().Add(-time.Second)) {
		t.Error("Invalid insertion and update dates")
	}
}

func TestPostToPostResponse(t *testing.T) {

	post := Post{
		Id:            1,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
		PostBase: PostBase{
			UserId:        1,
			Title:         "Test Title",
			Description:   "Test Description",
			HasMultimedia: true,
			Public:        true,
			Multimedia:    "test.jpg",
		},
	}

	expectedPostResponse := PostResponse{
		Context: "https://www.w3.org/ns/activitystreams",
		Type:    "create",
		Object:  post,
	}

	createdPostResponse := PostToPostResponse(post)

	if !reflect.DeepEqual(createdPostResponse, expectedPostResponse) {
		t.Error("Conversion from Post to PostResponse did not produce the expected result")
	}
}
