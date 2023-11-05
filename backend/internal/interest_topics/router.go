package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/infraestructure"
)

func InterestTopicsModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userInterestRepo := infraestructure.NewUserInterestsDBRepository(dbInstance)
	interestTopicUseCase := application.NewInterestTopicsUseCase(userInterestRepo)
	selectTopicHandler := infraestructure.NewSelectTopicHandler(interestTopicUseCase)

	r.Post("/", selectTopicHandler.HandleSelectTopic)

	return r
}
