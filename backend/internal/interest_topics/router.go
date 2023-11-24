package interest_topics

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/infrastructure"
)

func InterestTopicsModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userInterestRepo := infrastructure.NewUserInterestsDBRepository(dbInstance)
	interestTopicUseCase := application.NewSelectTopicUseCase(userInterestRepo)
	selectTopicHandler := infrastructure.NewSelectTopicHandler(interestTopicUseCase)

	InterestTopicsRepo := infrastructure.NewInterestTopicsDBRepository(dbInstance)
	ListinterestTopicUseCase := application.NewListInterestTopicsUseCase(InterestTopicsRepo)
	listInterestTopicHandler := infrastructure.NewListInterestTopicsHandler(ListinterestTopicUseCase)

	r.Get("/list", listInterestTopicHandler.HandleListTopics)

	r.Post("/select", selectTopicHandler.HandleSelectTopic)

	return r
}
