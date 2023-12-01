package infrastructure

import (
	"database/sql"
	"errors"
	"log"

	"github.com/unmsmfisi-socialapplication/social_app/internal/follow/domain"
)

type FollowerRepository struct {
	db *sql.DB
}

func NewFollowerRepository(database *sql.DB) *FollowerRepository {
	return &FollowerRepository{db: database}
}

func (u *FollowerRepository) InsertNewFollower(newFollower *domain.Follower) (*domain.Follower, error) {
	query := `INSERT INTO SOC_APP_USER_PROFILE_FOLLOW (follow_date, follower_profile_id , following_profile_id ) VALUES (NOW(), $1, $2)`
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	_, err = tx.Exec(query, newFollower.Follower_profile_id, newFollower.Following_profile_id)
	//fmt.Print(err)
	if err != nil {
		return nil, err
	} else {
		tx.Commit()
		return newFollower, nil
	}

}

func (u *FollowerRepository) IsFollowing(newFollower *domain.Follower) (*bool, error) {
	query := `
        SELECT COUNT(1) is_following
        FROM soc_app_user_profile_follow
        WHERE follower_profile_id = $1 AND following_profile_id = $2
    `
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	rows, err := tx.Query(query, newFollower.Follower_profile_id, newFollower.Following_profile_id)
	if err != nil {
		log.Println("Error while executing the query")
		return nil, err
	}
	defer rows.Close()

	var isFollowing bool

	if rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			log.Println("Error while scanning the rows")
			return nil, err
		}
		isFollowing = (count == 1)
	} else {
		log.Println("No rows returned from the query")
		return nil, errors.New("no rows returned")
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error while retrieving the rows")
		return nil, err
	}

	return &isFollowing, nil
}

func (u *FollowerRepository) ViewCommonFollowers(p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {
	query := `
        SELECT
			profile_id,
			user_id,
			name,
			last_name,
			profile_picture,
			user_name
        FROM FN_SOC_APP_GET_COMMON_FOLLOWERS($1,$2,$3,$4)
    `
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	rows, err := tx.Query(query, p_own_profile_id, p_viewed_profile_id, p_page_size, p_page_number)
	if err != nil {
		log.Println("Error while executing the query")
		return nil, err
	}
	defer rows.Close()

	var followers domain.FollowerDataList

	for rows.Next() {
		var follower domain.FollowerData
		err := rows.Scan(
			&follower.ProfileID,
			&follower.UserID,
			&follower.Name,
			&follower.LastName,
			&follower.ProfilePicture,
			&follower.UserName,
		)
		if err != nil {
			log.Println("Error while scanning the rows")
			return nil, err
		}
		followers = append(followers, follower)
	}

	if len(followers) == 0 {
		log.Println("No rows returned from the query")
		return nil, errors.New("no rows returned")
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error while retrieving the rows")
		return nil, err
	}

	return &followers, nil
}

func (u *FollowerRepository) ProfileFollowers(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {
	query := `
        SELECT
			profile_id,
			user_id,
			name,
			last_name,
			profile_picture,
			user_name
        FROM FN_SOC_APP_GET_PROFILE_FOLLOWERS($1,$2,$3)
    `
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	rows, err := tx.Query(query, p_profile_id, p_page_size, p_page_number)
	if err != nil {
		log.Println("Error while executing the query")
		return nil, err
	}
	defer rows.Close()

	var followers domain.FollowerDataList

	for rows.Next() {
		var follower domain.FollowerData
		err := rows.Scan(
			&follower.ProfileID,
			&follower.UserID,
			&follower.Name,
			&follower.LastName,
			&follower.ProfilePicture,
			&follower.UserName,
		)
		if err != nil {
			log.Println("Error while scanning the rows")
			return nil, err
		}
		followers = append(followers, follower)
	}

	if len(followers) == 0 {
		log.Println("No rows returned from the query")
		return nil, errors.New("no rows returned")
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error while retrieving the rows")
		return nil, err
	}

	return &followers, nil
}

func (u *FollowerRepository) ProfileFollowingProfiles(p_profile_id, p_page_size, p_page_number int) (*domain.FollowerDataList, error) {
	query := `
        SELECT
			profile_id,
			user_id,
			name,
			last_name,
			profile_picture,
			user_name
        FROM FN_SOC_APP_GET_PROFILE_FOLLOWING($1,$2,$3)
    `
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Error while starting the transaction")
		return nil, err
	}

	rows, err := tx.Query(query, p_profile_id, p_page_size, p_page_number)
	if err != nil {
		log.Println("Error while executing the query")
		return nil, err
	}
	defer rows.Close()

	var followers domain.FollowerDataList

	for rows.Next() {
		var follower domain.FollowerData
		err := rows.Scan(
			&follower.ProfileID,
			&follower.UserID,
			&follower.Name,
			&follower.LastName,
			&follower.ProfilePicture,
			&follower.UserName,
		)
		if err != nil {
			log.Println("Error while scanning the rows")
			return nil, err
		}
		followers = append(followers, follower)
	}

	if len(followers) == 0 {
		log.Println("No rows returned from the query")
		return nil, errors.New("no rows returned")
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error while retrieving the rows")
		return nil, err
	}

	return &followers, nil
}
