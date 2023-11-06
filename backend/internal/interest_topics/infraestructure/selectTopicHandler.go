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
	UserId     string   `json:"user_id"`
	InterestId []string `json:"interest_id"`
}

func NewSelectTopicHandler(useCase application.InterestTopicsUseCaseI) *SelectTopicHandler {

	return &SelectTopicHandler{useCase: useCase}
}

func checkInvalidPayload(raw json.RawMessage, requestData *requestData) bool {

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

	//interest_id not defined
	delete(extra, "user_id")
	if len(extra) == 0 {
		return true
	}
	delete(extra, "interest_id")

	//extra parameters in payload
	return len(extra) != 0
}

func extraValidations(writer http.ResponseWriter, requestData requestData) bool {

	if len(requestData.InterestId) < 3 && len(requestData.InterestId) > 0 {
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "At least 3 interest topics must be selected")
		return true
	}

	//Avoid duplicate interest_id
	seen := make(map[string]bool)
	for _, value := range requestData.InterestId {
		if seen[value] {
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Duplicate interest topic")
			return true
		}
		seen[value] = true
	}

	//Interest topics not selected
	if len(requestData.InterestId) == 0 {
		utils.SendJSONResponse(writer, http.StatusOK, "OK", "Skipped setting interest topics")
		return true
	}

	return false

}
func (handler *SelectTopicHandler) HandleSelectTopic(writer http.ResponseWriter, request *http.Request) {

	var requestData requestData

	var raw json.RawMessage
	err := json.NewDecoder(request.Body).Decode(&raw)

	if checkInvalidPayload(raw, &requestData) || err!=nil{
		utils.SendJSONResponse(writer, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	if extraValidations(writer, requestData) {
		return
	}

	err = handler.useCase.SetInterestTopics(requestData.UserId, requestData.InterestId)
	if err != nil {
		if err == application.ExistingUserInterestTopic {
			utils.SendJSONResponse(writer, http.StatusConflict, "ERROR", "Attempted insertion of an existing user interest topic")
			return
		} else {
			utils.SendJSONResponse(writer, http.StatusInternalServerError, "ERROR", "Error during insertion")
			return
		}

	}

	utils.SendJSONResponse(writer, http.StatusOK, "OK", "Insertion successful")

}
