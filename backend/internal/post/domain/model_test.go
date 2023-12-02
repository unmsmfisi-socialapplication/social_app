package domain

import (
	"testing"
	"time"
)

func TestPostCreateToPost(t *testing.T) {
	postCreate := PostCreate{
		PostBase: PostBase{
			UserId:        1,
			Title:         "Test Title",
			Description:   "Test Description",
			HasMultimedia: false,
			Public:        true,
			Multimedia:    "",
		},
	}

	resultPost := PostCreateToPost(postCreate)

	expectedPost := Post{
		PostBase:      postCreate.PostBase,
		InsertionDate: resultPost.InsertionDate,
		UpdateDate:    resultPost.UpdateDate,
	}

	if expectedPost.PostBase != resultPost.PostBase {
		t.Errorf("Expected PostBase %+v, but got %+v", expectedPost.PostBase, resultPost.PostBase)
	}

	if resultPost.InsertionDate.IsZero() {
		t.Errorf("Expected non-zero InsertionDate, but got zero")
	}

	if resultPost.UpdateDate.IsZero() {
		t.Errorf("Expected non-zero UpdateDate, but got zero")
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
			HasMultimedia: false,
			Public:        true,
			Multimedia:    "",
		},
	}

	resultPostResponse := PostToPostResponse(post)

	expectedPostResponse := PostResponse{
		Context: "https://www.w3.org/ns/activitystreams",
		Type:    "create",
		Object:  post,
	}

	if expectedPostResponse != resultPostResponse {
		t.Errorf("Expected PostResponse %+v, but got %+v", expectedPostResponse, resultPostResponse)
	}
}
