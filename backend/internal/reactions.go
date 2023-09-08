package internal

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type reactionsResource struct{
	reaction_id         int       	// Unique id for reaction
    user_id  			int       	// User's id who reacted
    post_id 			int    		// Id of post which this reaction belongs
    reaction_type_id	int    		// reaction type (like, dislike, weird, ...)
    createdAt			time.Time 	
}

// Routes creates a REST router for the reactions resource
func (rs reactionsResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Post("/", rs.Create)

	return r
}



func (rs reactionsResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reaction created"))
}

