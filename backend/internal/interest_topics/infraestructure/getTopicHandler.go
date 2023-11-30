package infraestructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type GetTopicHandler struct {
	useCase application.GetInterestTopicsUseCase
}

type getRequestData struct {
	ProfileId     string   `json:"profile_id"`
	
}

func NewGetTopicHandler(useCase application.GetInterestTopicsUseCase) *GetTopicHandler {

	return &GetTopicHandler{useCase: useCase}
}

func (gt *GetTopicHandler) GetTopic(writer http.ResponseWriter, request *http.Request) {

	var requestData getRequestData
	
	
	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	

	interests,err := gt.useCase.GetInterestTopics(requestData.ProfileId)
	if err != nil {
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", err.Error())
		return
	}
	if interests == nil {
		interests = []string{}
	}
	response:=struct {
		Interests []string `json:"interests"`
	}{
		Interests: interests,
	}

	utils.SendJSONResponse(writer, http.StatusOK, "OK", response)
}

