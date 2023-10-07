package comment

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/infrastructure"
)

func CommentModuleRouter(dbInstance *sql.DB) *chi.Mux{

	r := chi.NewRouter()

	commentRepo := infrastructure.NewCommentRepository(dbInstance)
	commentUseCase := application.NewCommentUseCase(commentRepo)
	commentHandler := infrastructure.NewCommentHandler(commentUseCase)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"hello\": \"from comments section\"}"))
	})
	r.Get("/{commentID}", commentHandler.HandleGetCommentByID)
	r.Post("/", commentHandler.HandleCreateComment)
	r.Put("/", commentHandler.HandleUpdateComment) //refactor this route
	r.Delete("/{commentID}", commentHandler.HandleDeleteComment)

	return r
}