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

func (repository *SearchRepository) GetProfilesByName(keyword string, psize, pnumber int) *domain.QueryResult {
    var profiles []profiledomain.Profile

    query := "SELECT profile_id, name, about_me, profile_picture FROM soc_app_user_profile WHERE name LIKE $1 LIMIT $2 OFFSET $3"
    rows, err := repository.db.Query(query, "%" + keyword + "%", psize, pnumber)
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

    query := "SELECT profile_id, name, about_me, profile_picture FROM soc_app_user_profile WHERE name LIKE $1 LIMIT $2 OFFSET $3"
    rows, err := repository.db.Query(query, "%" + keyword + "%", 5, 1)
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
