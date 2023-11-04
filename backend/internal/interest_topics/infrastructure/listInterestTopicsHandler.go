package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type ListInterestTopicsHandler struct {
	useCase application.ListInterestTopicsUseCaseI
}

func NewListInterestTopicsHandler(useCase application.ListInterestTopicsUseCaseI) *ListInterestTopicsHandler {
	return &ListInterestTopicsHandler{useCase: useCase}
}

func (handler *ListInterestTopicsHandler) HandleListTopics(writer http.ResponseWriter, r *http.Request) {

	var interestTopics []domain.InterestTopics
	interestTopics, err := handler.useCase.GetInteresTopics()
	if err != nil {
		utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error while fetching BFD data")
		fmt.Print(err.Error())
		return
	}
	if len(interestTopics) < 1 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "There are no interest topics")
		return
	}
	utils.SendJSONResponse(writer, http.StatusOK, "OK", interestTopics)

}
