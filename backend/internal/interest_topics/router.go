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
	getInterestTopicUseCase := application.NewGetInterestTopicsUseCase(userInterestRepo)
	getTopicHandler := infraestructure.NewGetTopicHandler(*getInterestTopicUseCase)

	r.Post("/", selectTopicHandler.HandleSelectTopic)
	r.Get("/", getTopicHandler.GetTopic)
	return r
}
