package infraestructure

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/domain"
	"log"
	"net/http"
	"strconv"
)

type PostReactionHandler struct {
	useCase *application.PostReactionUseCase
}

func NewPostReactionHandler(useCase *application.PostReactionUseCase) *PostReactionHandler {
	return &PostReactionHandler{
		useCase: useCase,
	}
}

func (h *PostReactionHandler) GetPostReactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Println("Cannot convert string productId to int")
		return
	}

	product, err := h.useCase.GetByID(int64(productId))

	if err != nil {
		w.Write([]byte("Cannot retrieve post reaction by given id"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(product)

	if err != nil {
		log.Println("Cannot encode the post reaction to json")
	}

}

func (h *PostReactionHandler) CreatePostReactionHandler(w http.ResponseWriter, r *http.Request) {
	var newPostReaction domain.PostReaction

	if err := json.NewDecoder(r.Body).Decode(&newPostReaction); err != nil {
		http.Error(w, "There was an error on decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.useCase.Create(newPostReaction)

	if err != nil {
		w.Write([]byte("Cannot insert that post reaction"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post reaction succesful registered to the database"))

}

func (h *PostReactionHandler) DeletePostReactionHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Cannot read post reaction id sent: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.Delete(int64(productId))

	if err != nil {
		http.Error(w, "Cannot delete that post reaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted post reaction"))

}
