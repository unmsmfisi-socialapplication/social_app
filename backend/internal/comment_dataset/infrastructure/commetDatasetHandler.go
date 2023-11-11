package infrastructure

import (
	"log"
	"net/http"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment_dataset/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type commentDatasetHandler struct {
	useCase *application.CommentDatasetUseCase
}

func NewCommentDatasetHandler(useCase *application.CommentDatasetUseCase) *commentDatasetHandler {
	return &commentDatasetHandler{useCase: useCase}
}

func (cdh *commentDatasetHandler) HandleRetrieveScopedComments(w http.ResponseWriter, r *http.Request) {

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	sd, err_sd := time.Parse(time.RFC3339, start_date)
	ed, err_ed := time.Parse(time.RFC3339, end_date)
	if err_sd != nil || err_ed != nil {
		log.Fatal("Date Format should be like: 2016-02-02T15:04:05Z")
		return
	}
	if ed.Before(sd) {
		log.Fatal("Incorrect Date range")
		return
	}

	comments, err := cdh.useCase.RetrieveDateScopedComments(start_date, end_date)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", comments)
}
