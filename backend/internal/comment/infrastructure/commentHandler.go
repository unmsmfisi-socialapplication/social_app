package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type CommentHandler struct {
	useCase application.CommentUseCaseInterface
}

func NewCommentHandler(useCase application.CommentUseCaseInterface) *CommentHandler {
	return &CommentHandler{useCase: useCase}
}

func (ch *CommentHandler) HandleGetCommentByID(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		CommentID int64 `json:"commentID"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid commentID")
		return
	}

	comment, err := ch.useCase.GetByID(requestData.CommentID)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error getting comment")
		fmt.Println(err.Error())
		return
	}

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error marshaling comment to JSON")
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", string(commentJSON))
}

func (ch *CommentHandler) HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	var commentData struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`  
	}

	errStructure := json.NewDecoder(r.Body).Decode(&commentData)

	if errStructure != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid comment")
		return
	}

	comment := &domain.Comment{
        UserID:          commentData.UserID,
        PostID:          commentData.PostID,
        Comment:         commentData.Comment,
        InsertionDate:   commentData.InsertionDate,
        UpdateDate:      commentData.UpdateDate,
        ParentCommentID: commentData.ParentCommentID,
    }

	err := ch.useCase.Create(comment)
	
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error creating comment")
		fmt.Println(err.Error())
		return
	} 

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Comment created successfully")
}

func (ch *CommentHandler) HandleUpdateComment(w http.ResponseWriter, r *http.Request) {
	var commentData struct {
		CommentID       int64     `json:"commentID"`
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`  
	}

	errStructure := json.NewDecoder(r.Body).Decode(&commentData)

	if errStructure != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid comment")
		return
	}

	comment := &domain.Comment{
		CommentID:       commentData.CommentID,
		UserID:          commentData.UserID,
		PostID:          commentData.PostID,
		Comment:         commentData.Comment,
		InsertionDate:   commentData.InsertionDate,
		UpdateDate:      commentData.UpdateDate,
		ParentCommentID: commentData.ParentCommentID,
	}

	err := ch.useCase.Update(comment)
	
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error updating comment")
		fmt.Println(err.Error())
		return
	} 

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Comment updated successfully")
}

func (ch *CommentHandler) HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		CommentID int64 `json:"commentID"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid commentID")
		return
	}

	err = ch.useCase.Delete(requestData.CommentID)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error deleting comment")
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Comment deleted successfully")
}