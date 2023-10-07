package infraestructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type SelectTopicHandler struct {
	useCase application.InterestTopicsUseCaseI
}

func NewSelectTopicHandler(useCase application.InterestTopicsUseCaseI) *SelectTopicHandler {

	return &SelectTopicHandler{useCase: useCase}
}

func (slh *SelectTopicHandler) HandleSelectTopic(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		User_id     string   `json:"user_id"`
		Interest_id []string `json:"interest_id"`
	}

	var raw json.RawMessage
	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	err = json.Unmarshal(raw, &requestData)
	if err != nil || requestData.User_id == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	var extra map[string]interface{}
	err = json.Unmarshal(raw, &extra)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	//user_id not defined
	delete(extra, "user_id")
	if len(extra) == 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}
	delete(extra, "interest_id")

	//extra parameters in payload
	if len(extra) != 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	//Avoid duplicate interest_id
	seen := make(map[string]bool)
	for _, value := range requestData.Interest_id  {
		if seen[value] {
			utils.SendJSONResponse(w, http.StatusConflict, "ERROR", "Error during insertion")
			return
		}
		seen[value] = true
	}

	if len(requestData.Interest_id) == 0 {
		utils.SendJSONResponse(w, http.StatusOK, "OK", "Skipped setting interest topics")
		return
	}

	for i := 0; i < len(requestData.Interest_id); i++ {
		err := slh.useCase.SetInterestTopics(requestData.User_id, requestData.Interest_id[i])
		if err != nil {
			switch err {
			case application.ErrInvalidInsertion:
				utils.SendJSONResponse(w, http.StatusConflict, "ERROR", "Attempted insertion of an existing user interest topic")
				fmt.Println(err.Error())
				return
			default:
				utils.SendJSONResponse(w, http.StatusNotFound, "NOTFOUND", "Error during insertion")
				fmt.Println(err.Error())
				return
			}

		}
	}

	utils.SendJSONResponse(w, http.StatusOK, "OK", "Insertion successful")

}
