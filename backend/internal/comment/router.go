package comment

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/infrastructure"
)

func CommentModuleRouter(dbInstance *sql.DB) *chi.Mux{

	r := chi.NewRouter()

	commentRepo := infrastructure.NewCommentRepository(dbInstance)
	commentUseCase := application.NewCommentUseCase(commentRepo)
	commentHandler := infrastructure.NewCommentHandler(commentUseCase)

	r.Get("/", commentHandler.HandleGetAllComments)
	r.Get("/{commentID:[0-9]+}", commentHandler.HandleGetCommentByID)
	r.Get("/post/{postID:[0-9]+}", commentHandler.HandleGetCommentsByPostId)
	r.Post("/", commentHandler.HandleCreateComment)
	r.Put("/{commentID:[0-9]+}", commentHandler.HandleUpdateComment)
	r.Delete("/{commentID:[0-9]+}", commentHandler.HandleDeleteComment)

	return r
}