package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

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

func (ch *CommentHandler) HandleGetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := ch.useCase.GetAll()
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error getting comments")
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", comments)
}

func (ch *CommentHandler) HandleGetCommentByID(w http.ResponseWriter, r *http.Request) {
	commentIDStr := chi.URLParam(r, "commentID")
	if commentIDStr == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid commentID")
		return
	}

	commentID, _ := strconv.ParseInt(commentIDStr, 10, 64)
	comment, err := ch.useCase.GetByID(commentID)
	if err != nil {
		if commentID > 0{
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Comment not found")
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error getting comment")
		}
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", comment)
}

func (ch *CommentHandler) HandleGetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")
	if postIDStr == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid postID")
		return
	}

	postID, _ := strconv.ParseInt(postIDStr, 10, 64)
	comments, err := ch.useCase.GetByPostID(postID)
	if err != nil {
		if postID > 0{
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Post not found")
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error getting comments")
		}
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", comments)
}

func (ch *CommentHandler) HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	var commentData struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID *int64     `json:"parentCommentID"`  
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
	commentIDStr := chi.URLParam(r, "commentID")
	if commentIDStr == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid commentID")
		return
	}
	commentID, _ := strconv.ParseInt(commentIDStr, 10, 64)

	var commentData struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID *int64     `json:"parentCommentID"`  
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
	err := ch.useCase.Update(commentID, comment)
	if err != nil {
		if commentID > 0{
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Comment not found")
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error updating comment")
		}
		fmt.Println(err.Error())
		return
	} 

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Comment updated successfully")
}

func (ch *CommentHandler) HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	commentIDStr := chi.URLParam(r, "commentID")
	
	if commentIDStr == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid commentID")
		return
	}

	commentID, _ := strconv.ParseInt(commentIDStr, 10, 64)
	err := ch.useCase.Delete(commentID)
	
	if err != nil {
		if commentID > 0{
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Comment not found")
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error deleting comment")
		}
		fmt.Println(err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Comment deleted successfully")
}