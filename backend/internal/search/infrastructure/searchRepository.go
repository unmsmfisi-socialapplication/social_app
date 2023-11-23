package infrastructure

import (
	"database/sql"

	profiledomain "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/search/domain"
)

type SearchRepository struct {
	db *sql.DB
}

func NewSearchRepository(database *sql.DB) *SearchRepository {
	return &SearchRepository{db: database}
}

func (repository *SearchRepository) GetProfilesByName(keyword string) *domain.QueryResult {
    var profiles []profiledomain.Profile

    rows, err := repository.db.Query("SELECT * FROM soc_app_user_profile WHERE name LIKE '%" + keyword + "%'")
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var profile profiledomain.Profile
        err := rows.Scan(&profile.Id_profile, &profile.Name, &profile.Biography, &profile.ProfileImage)
        if err != nil {
            panic(err)
        }
        profiles = append(profiles, profile)
    }

    result := domain.NewQueryResult(profiles)

    return result
}

func (repository *SearchRepository) GetsuggestionsProfiles(keyword string) *domain.QueryResult {
    var profiles []profiledomain.Profile

    rows, err := repository.db.Query("SELECT * FROM soc_app_user_profile WHERE name LIKE '%" + keyword + "%' limit 5")
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var profile profiledomain.Profile
        err := rows.Scan(&profile.Id_profile, &profile.Name, &profile.Biography, &profile.ProfileImage)
        if err != nil {
            panic(err)
        }
        profiles = append(profiles, profile)
    }

    result := domain.NewQueryResult(profiles)

    return result
}
