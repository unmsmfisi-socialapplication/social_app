package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment_dataset/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment_dataset/infrastructure"
)

func CommentDatasetModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	cdRepository := infrastructure.NewCommentDatasetRepository(dbInstance)
	cdUseCase := application.NewCommentDatasetUseCase(cdRepository)
	cdHandler := infrastructure.NewCommentDatasetHandler(cdUseCase)

	r.Post("/comments", cdHandler.HandleRetrieveScopedComments)

	return r
}
