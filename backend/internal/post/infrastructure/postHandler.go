package infrastructure

import (
    "encoding/json"
    "net/http"

    "github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
    "github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

type PostHandler struct {
    useCase application.PostUseCaseInterface
}

func NewPostHandler(useCase application.PostUseCaseInterface) *PostHandler {
    return &PostHandler{useCase: useCase}
}

func (ph *PostHandler) HandleCreateMultiPost(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
        MastodonData domain.CreatePost `json:"mastodon,omitempty"`
        PixelfedData domain.CreatePost `json:"pixelfed,omitempty"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
        return
    }

    var mastodonPost *domain.Post
    var pixelfedPost *domain.Post

    // if requestData.MastodonData != nil {
        // mastodonPost, err := ph.useCase.CreatePost(*requestData.MastodonData)
        // if err != nil {
        // 	utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
        // 	return
        // }
    // }

    // if requestData.PixelfedData != nil {
        // pixelfedPost, err := ph.useCase.CreatePost(*requestData.PixelfedData)
        // if err != nil {
        // 	utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
        // 	return
        // }
    //  }

    response := map[string]interface{}{
        "mastodon": mastodonPost,
        "pixelfed": pixelfedPost,
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
        return
    }

    // Convert the JSON response to a string
    responseString := string(jsonResponse)

    // Send the JSON response as a string
    utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", responseString)
}