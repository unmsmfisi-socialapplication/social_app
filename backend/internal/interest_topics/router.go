package post

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/infrastructure"
)

func InterestTopicsModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userInterestRepo := infrastructure.NewUserInterestsDBRepository(dbInstance)
	interestTopicUseCase := application.NewInterestTopicsUseCase(userInterestRepo)
	selectTopicHandler := infrastructure.NewSelectTopicHandler(interestTopicUseCase)

	r.Post("/", selectTopicHandler.HandleSelectTopic)

	return r
}
