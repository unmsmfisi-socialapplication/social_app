package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/infrastructure"
)

func PostModuleRouter(dbInstance *sql.DB, eventManager *events.EventManager) *chi.Mux {
	r := chi.NewRouter()

	postRepository := infrastructure.NewPostDBRepository(dbInstance)
	postUseCase := application.NewPostUseCase(postRepository, eventManager)
	postHandler := infrastructure.NewPostHandler(postUseCase)

	r.Post("/create", postHandler.HandleCreatePost)

	return r
}
