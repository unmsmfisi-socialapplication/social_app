package domain

import "time"

type Comment struct {
    CommentID       int64     `json:"commentID"`  
    UserID          int64     `json:"userID"`
    PostID          int64     `json:"postID"`
    Comment         string    `json:"comment"`
    InsertionDate   time.Time `json:"insertionDate"`
    UpdateDate      time.Time `json:"updateDate"`
    ParentCommentID *int64     `json:"parentCommentID"`  
}