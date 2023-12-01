package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type ListCommunitiesHandler struct {
	useCase application.ListCommunitiesUseCaseI
}

type requestData struct {
	UserId string `json:"user_id"`
	PageNumber string `json:"page_number"`
	PageSize   string `json:"page_size"`
}

func NewListCommunitiesHandler(useCase application.ListCommunitiesUseCaseI) *ListCommunitiesHandler {
	return &ListCommunitiesHandler{useCase: useCase}
}

func checkInvalidPayload(raw json.RawMessage, requestData *requestData) bool {

	if err := json.Unmarshal(raw, requestData); err != nil {
		return true
	}

	//Empty parameters
	if requestData.UserId == "" || requestData.PageNumber == "" || requestData.PageSize == "" {
		return true
	}
	
	var extra map[string]interface{}
	
	err := json.Unmarshal(raw, &extra)
	if err != nil {
		return true
	}

	delete(extra, "user_id")
	if len(extra) == 0 {
		return true
	}
	delete(extra, "page_number")
	if len(extra) == 0 {
		return true
	}

	delete(extra, "page_size")

	//extra parameters in payload
	return len(extra) != 0
}

func (handler *ListCommunitiesHandler) HandleListCommunities(writer http.ResponseWriter, request *http.Request) {

	var requestData requestData

	var raw json.RawMessage
	er := json.NewDecoder(request.Body).Decode(&raw)
	
	if checkInvalidPayload(raw, &requestData) || er!=nil{
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	var communities []domain.Community

	userId := requestData.UserId
	pageSize := requestData.PageSize
	pageNumber := requestData.PageNumber
	communities, err := handler.useCase.GetCommunitiesList(userId,pageSize,pageNumber)
	
	if err != nil {
		utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error while fetching data")
		fmt.Print(err.Error())
		return
	}
	if len(communities) < 1 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "There are no communities")
		return
	}
	utils.SendJSONResponse(writer, http.StatusOK, "OK", communities)

}
