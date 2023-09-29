package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post_reactions/model"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post_reactions/service"
)

type PostReactionHandler struct {
	Service *service.PostReactionService
}

func NewPostReactionHandler(postReactionService *service.PostReactionService) *PostReactionHandler {
	return &PostReactionHandler{
		Service: postReactionService,
	}
}

func (s *PostReactionHandler) CreatePostReactionHandler(w http.ResponseWriter, r *http.Request) {

	var newPostReaction model.PostReaction

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newPostReaction); err != nil {
		http.Error(w, "There was an error on JSON decoding", http.StatusBadRequest)
		return
	}

	fmt.Println(newPostReaction.PostID)
	fmt.Println(newPostReaction.UserID)

	err := s.Service.CreatePostReaction(newPostReaction.PostID, newPostReaction.UserID)

	if err != nil {
		http.Error(w, "There was an error on creating the new reaction for this post", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Successful post reaction creation",
	}
	w.Header().Set("Content-Type", "service/json")
	json.NewEncoder(w).Encode(response)
}

func (s *PostReactionHandler) GetReactionsForPost(w http.ResponseWriter, r *http.Request) {

	num, err := strconv.Atoi(chi.URLParam(r, "post_id"))

	if err != nil {
		http.Error(w, "Pfff", http.StatusInternalServerError)
		return
	}

	fmt.Println(num)

	postReactions, err := s.Service.GetReactionsForPost(num)

	if err != nil {
		http.Error(w, "There was an error on retrieve reactions for this post", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(postReactions)
}
