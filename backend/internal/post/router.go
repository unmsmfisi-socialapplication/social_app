package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/infrastructure"
)

func PostModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	postRepository := infrastructure.NewPostDBRepository(dbInstance)
	postUseCase := application.NewPostUseCase(postRepository)
	postHandler := infrastructure.NewPostHandler(postUseCase)

	r.Post("/create", postHandler.HandleCreatePost)
	r.Get("/getall", postHandler.HandleGetAllPost)
	r.Get("/{id}", postHandler.HandleGetPost)

	// DELETE POST //
	r.Delete("/delete/{id}", postHandler.HandleDeletePost)
	// UPDATE POST //
	r.Put("/update/{id}", postHandler.HandleUpdatePost)
	// REPORT POST //
	r.Post("/report", postHandler.HandleReportPost)

	return r
}
