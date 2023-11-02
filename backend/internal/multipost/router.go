package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/infrastructure"
)

func PostModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	clientID := "ClienteID" // Reemplaza "ClienteID" con el valor real
	clientSecret := "ClienteSecret" // Reemplaza "ClienteSecret" con el valor real
	instanceURL := "URLDeInstancia" // Reemplaza "URLDeInstancia" con la URL real

	postRepository := infrastructure.NewPostDBRepository(dbInstance)
	postUseCase := application.NewPostUseCase(postRepository)
	postHandler := infrastructure.NewPostHandler(postUseCase, clientID, clientSecret, instanceURL)

	r.Post("/create", postHandler.HandleCreatePost)

	return r
}