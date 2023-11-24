package communities_interest

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/infrastructure"
)

func InterestTopicsModuleRouter(dbInstance *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	CommunityRepo := infrastructure.NewCommunityDBRepository(dbInstance)
	ListInterestCommunitiesUseCase := application.NewListCommunitiesUseCase(CommunityRepo)
	ListInterestCommunitiesHandler := infrastructure.NewListCommunitiesHandler(ListInterestCommunitiesUseCase)

	r.Post("/list", ListInterestCommunitiesHandler.HandleListCommunities)

	return r
}
