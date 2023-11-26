package infrastructure

import (
	"encoding/json"
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

func (handler *ListInterestTopicsHandler) HandleListTopics(writer http.ResponseWriter, request *http.Request) {
	var requestData struct {
		PageNumber string `json:"page_number"`
		PageSize   string `json:"page_size"`
	}

	er := json.NewDecoder(request.Body).Decode(&requestData)
	if er != nil {
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
	}

	var interestTopics []domain.InterestTopic
	pageSize := requestData.PageSize
	pageNumber := requestData.PageNumber

	interestTopics, err := handler.useCase.GetInteresTopics(pageSize, pageNumber)
	if err != nil {
		utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error while fetching data")
		fmt.Print(err.Error())
		return
	}
	if len(interestTopics) < 1 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "There are no interest topics")
		return
	}
	utils.SendJSONResponse(writer, http.StatusOK, "OK", interestTopics)
}
