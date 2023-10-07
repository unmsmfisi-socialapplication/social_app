package infrastructure

import (
    "encoding/json"
    "net/http"

    "github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
    "github.com/unmsmfisi-socialapplication/social_app/pkg/utils"
)

// PostHandler handles HTTP requests related to posts.
type PostHandler struct {
    useCase application.PostUseCaseInterface
}

// NewPostHandler creates a new instance of PostHandler.
func NewPostHandler(useCase application.PostUseCaseInterface) *PostHandler {
    return &PostHandler{useCase: useCase}
}

// HandleCreateMultiPost handles the creation of posts for Mastodon and Pixelfed.
func (ph *PostHandler) HandleCreateMultiPost(w http.ResponseWriter, r *http.Request) {
    // Define a struct to hold the JSON request data
    var requestData struct {
        MastodonData domain.CreatePost `json:"mastodon,omitempty"`
        PixelfedData domain.CreatePost `json:"pixelfed,omitempty"`
    }

    // Decode the JSON request into the requestData struct
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "ERROR", "Invalid request payload")
        return
    }

    var mastodonPost *domain.Post
    var pixelfedPost *domain.Post

    // Uncomment and add your Mastodon-specific code here
    // if requestData.MastodonData != nil {
    //     mastodonPost, err := ph.useCase.CreatePost(*requestData.MastodonData)
    //     if err != nil {
    //         utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
    //         return
    //     }
    // }

    // Uncomment and add your Pixelfed-specific code here
    // if requestData.PixelfedData != nil {
    //     pixelfedPost, err := ph.useCase.CreatePost(*requestData.PixelfedData)
    //     if err != nil {
    //         utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
    //         return
    //     }
    // }

    // Create a response map with Mastodon and Pixelfed posts
    response := map[string]interface{}{
        "mastodon": mastodonPost,
        "pixelfed": pixelfedPost,
    }

    // Marshal the response map into JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        utils.SendJSONResponse(w, http.StatusInternalServerError, "ERROR", err.Error())
        return
    }

    // Convert the JSON response to a string
    responseString := string(jsonResponse)

    // Send the JSON response as a string with a "SUCCESS" status
    utils.SendJSONResponse(w, http.StatusOK, "SUCCESS", responseString)
}

// HandleCreatePost is a placeholder for handling single post creation (optional).
func (ph *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
}