package infraestructure

import (
	"encoding/json"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type SelectTopicHandler struct {
	useCase application.InterestTopicsUseCaseI
}

type requestData struct {
	User_id     string   `json:"user_id"`
	Interest_id []string `json:"interest_id"`
}

func NewSelectTopicHandler(useCase application.InterestTopicsUseCaseI) *SelectTopicHandler {

	return &SelectTopicHandler{useCase: useCase}
}

func checkInvalidPayload(request *http.Request, requestData *requestData) bool{

	var raw json.RawMessage
	err := json.NewDecoder(request.Body).Decode(&raw)
	if err != nil {
		return true
	}
	err = json.Unmarshal(raw, &requestData)
	if err != nil || requestData.User_id == "" {
		return true
	}
	var extra map[string]interface{}
	err = json.Unmarshal(raw, &extra)
	if err != nil {
		return true
	}

	//user_id not defined
	delete(extra, "user_id")
	if len(extra) == 0 {
		return true
	}
	delete(extra, "interest_id")

	//extra parameters in payload
	return len(extra) != 0
}

func (handler *SelectTopicHandler) HandleSelectTopic(writer http.ResponseWriter, request *http.Request) {

	var requestData requestData
	if checkInvalidPayload(request,&requestData){
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	//Avoid duplicate interest_id
	seen := make(map[string]bool)
	for _, value := range requestData.Interest_id  {
		if seen[value] {
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Duplicate interest topic")
			return
		}
		seen[value] = true
	}

	err := handler.useCase.SetInterestTopics(requestData.User_id, requestData.Interest_id)
	if err != nil {
		if err ==application.ExistingUserInterestTopic{
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Attempted insertion of an existing user interest topic")
			return
		}else
		{
			utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error during insertion")
			return
		}

	}
	if len(requestData.Interest_id) == 0 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "Skipped setting interest topics")
		return
	}
	utils.SendJSONResponse(writer, http.StatusOK, "OK", "Insertion successful")

}
