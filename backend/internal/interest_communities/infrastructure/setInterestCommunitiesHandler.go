package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type SetInterestCommunitiesHandler struct {
	useCase application.SetInterestCommunitiesUseCaseI
}

type requestData_intCommunities struct {
	UserId     string   `json:"user_id"`
	CommunityId []string `json:"community_id"`
}

func NewSetInterestCommunitiesHandler(useCase application.SetInterestCommunitiesUseCaseI) *SetInterestCommunitiesHandler {

	return &SetInterestCommunitiesHandler{useCase: useCase}
}

func checkInvalidPayload_intCommunities(raw json.RawMessage, requestData *requestData_intCommunities) bool {

	//user_id not defined
	err := json.Unmarshal(raw, &requestData)
	if err != nil || requestData.UserId == "" {
		return true
	}
	var extra map[string]interface{}
	err = json.Unmarshal(raw, &extra)
	if err != nil {
		return true
	}

	//community_id not defined
	delete(extra, "user_id")
	if len(extra) == 0 {
		return true
	}
	delete(extra, "community_id")

	//extra parameters in payload
	return len(extra) != 0
}

func extraValidations(writer http.ResponseWriter, requestData requestData_intCommunities) bool {

	//Avoid duplicate community_id
	seen := make(map[string]bool)
	for _, value := range requestData.CommunityId {
		if seen[value] {
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Duplicate interest community")
			return true
		}
		seen[value] = true
	}

	//Interest communities not selected
	if len(requestData.CommunityId) == 0 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "Skipped setting interest communities")
		return true
	}

	return false

}
func (handler *SetInterestCommunitiesHandler) HandleSetInterestCommunities(writer http.ResponseWriter, request *http.Request) {

	var requestData requestData_intCommunities

	var raw json.RawMessage
	err := json.NewDecoder(request.Body).Decode(&raw)

	if checkInvalidPayload_intCommunities(raw, &requestData) || err!=nil{
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	if extraValidations(writer, requestData) {
		return
	}

	err = handler.useCase.SetInterestCommunities(requestData.UserId, requestData.CommunityId)
	if err != nil {
		if err == application.ExistingUserInterestCommunity {
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Attempted insertion of an existing record")
			return
		} else {
			utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error during insertion")
			return
		}

	}

	utils.SendJSONResponse(writer, http.StatusOK, "OK", "Insertion successful")

}
