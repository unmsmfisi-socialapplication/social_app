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

    // Create and publish posts on selected social networks
    var mastodonPost *domain.Post
    var pixelfedPost *domain.Post

    if requestData.MastodonData != (domain.CreatePost{}) {
        mastodonPost, err := ph.useCase.CreatePost(requestData.MastodonData)
        if err != nil {
            utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
            return
        }
    }

    if requestData.PixelfedData != (domain.CreatePost{}) {
        pixelfedPost, err := ph.useCase.CreatePost(requestData.PixelfedData)
        if err != nil {
            utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
            return
        }
    }


    // Send a response with the created posts (you can customize this part)
    response := map[string]interface{}{
        "mastodon": mastodonPost,
        "pixelfed": pixelfedPost,
    }

    utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", response)
}
