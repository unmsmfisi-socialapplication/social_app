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

func NewListCommunitiesHandler(useCase application.ListCommunitiesUseCaseI) *ListCommunitiesHandler {
	return &ListCommunitiesHandler{useCase: useCase}
}

func (handler *ListCommunitiesHandler) HandleListCommunities(writer http.ResponseWriter, request *http.Request) {

	var requestData struct {
		UserId string `json:"user_id"`
	}

	er := json.NewDecoder(request.Body).Decode(&requestData)
	if er != nil {
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
	}

	var communities []domain.Community

	userId := requestData.UserId
	communities, err := handler.useCase.GetCommunitiesList(userId)
	
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
