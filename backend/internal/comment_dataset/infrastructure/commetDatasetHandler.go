package infrastructure

import (
	"net/http"

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

	start_date := r.URL.Query().Get("fromDate")
	end_date := r.URL.Query().Get("toDate")

	comments, err := cdh.useCase.RetrieveDateScopedComments(start_date, end_date)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", comments)
}
