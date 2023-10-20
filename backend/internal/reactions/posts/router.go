package posts

import (
	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/infraestructure"
	"gorm.io/gorm"
	"net/http"
)

func PostReactionModuleRouter(dbInstance *gorm.DB) *chi.Mux {

	r := chi.NewRouter()

	postReactionRepo := infraestructure.NewPostReactionRepository(dbInstance)
	postReactionUseCase := application.NewPostReactionUseCase(postReactionRepo)
	postReactionHandler := infraestructure.NewPostReactionHandler(postReactionUseCase)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"hello\": \"from post reactions section\"}"))
	})

	r.Get("/{postReactionID:[0-9]+}", postReactionHandler.GetPostReactionByIDHandler)
	r.Post("/", postReactionHandler.CreatePostReactionHandler)
	r.Delete("/{postReactionID:[0-9]+", postReactionHandler.DeletePostReactionHandler)

	return r

}
