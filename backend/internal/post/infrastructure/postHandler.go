package infrastructure

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/helpers"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type PostHandler struct {
	useCase application.PostUseCaseInterface
}

func NewPostHandler(useCase application.PostUseCaseInterface) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (ph *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var requestData domain.PostCreate

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	if requestData.PostBase.UserId == 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request User")
		return
	}

	if requestData.PostBase.Title == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request Title")
		return
	}

	postCreate, err := ph.useCase.CreatePost(requestData)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", postCreate)
}

func (ph *PostHandler) HandleGetAllPost(w http.ResponseWriter, r *http.Request) {

	params, err := helpers.ParsePaginationParams(r.URL.Query())
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	posts, err := ph.useCase.GetPosts(params)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", posts)
}

func (ph *PostHandler) HandleGetPost(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "param error.")
		return
	}

	post, err := ph.useCase.GetPost(id)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	if post == nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "SUCCESS", "Post not found")
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", post)
}

func (ph *PostHandler) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid post ID")
		return
	}

	err = ph.useCase.DeletePost(id)
	if err != nil {
		if err.Error() == "post not found" {
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Post not found")
			return
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", "Post deleted successfully")
}

func (ph *PostHandler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid post ID")
		return
	}

	var updateData domain.PostUpdate
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	err = ph.useCase.UpdatePost(id, updateData)
	if err != nil {
		if err.Error() == "post not found" {
			utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", "Post not found")
			return
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", "Post edited successfully")
}

func (ph *PostHandler) HandleReportPost(w http.ResponseWriter, r *http.Request) {
	var reportData domain.PostReport

	if err := json.NewDecoder(r.Body).Decode(&reportData); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	err := ph.useCase.ReportPost(reportData)

	switch {
	case errors.Is(err, domain.ErrReporterUserDoesNotExist):
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", err.Error())
	case errors.Is(err, domain.ErrPostNotFound):
		utils.SendJSONResponse(w, http.StatusNotFound, "ERROR", err.Error())
	case errors.Is(err, domain.ErrUserHasAlreadyReportedPost):
		utils.SendJSONResponse(w, http.StatusConflict, "ERROR", err.Error())
	case err != nil:
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Internal Server Error")
	default:
		utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", "Post reported successfully")
	}
}
