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

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
		return
	}

	for i := 0; i < len(requestData.Interest_id); i++ {
		insertion, err := slh.useCase.SetInterestTopics(requestData.User_id, requestData.Interest_id[i])
		if err != nil {
			switch err {
			case application.ErrInvalidInsertion:
				utils.SendJSONResponse(w, http.StatusNotFound, "NOTFOUND", "User not found")
				return
			default:
				utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Error during insertion")
				fmt.Println(err.Error())
				return
			}

		}
		if insertion {
			utils.SendJSONResponse(w, http.StatusOK, "OK", "Insertion successful")
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", "Insertion failed")
		}
	}

}
