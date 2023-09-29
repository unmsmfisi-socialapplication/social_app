package repository

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post_reactions/model"
)

type PostReactionRepository struct {
	DB *sql.DB
}

func NewPostReactionRepository(db *sql.DB) *PostReactionRepository {
	return &PostReactionRepository{
		DB: db,
	}
}

func (r *PostReactionRepository) CreatePostReaction(postReaction *model.PostReaction) error {
	query := "INSERT INTO post_reactions (post_id, user_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, postReaction.PostID, postReaction.UserID)

	return err
}

func (r *PostReactionRepository) GetReactionsForPost(postID int) (*[]model.PostReaction, error) {
	var postReactions []model.PostReaction

	query := "SELECT * FROM post_reactions WHERE post_id=$1"
	rows, err := r.DB.Query(query, postID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// postReactions := []*model.PostReaction{}

	for rows.Next() {
		var postReaction model.PostReaction
		// err := rows.Scan(&postReaction.PostReactionID, &postReaction.PostID, &postReaction.UserID, &postReaction.InsertionDate)
		err := rows.Scan(&postReaction.PostReactionID, &postReaction.PostID, &postReaction.UserID, &postReaction.InsertionDate)

		if err != nil {
			return nil, err
		}

		postReactions = append(postReactions, postReaction)
	}

	return &postReactions, nil
}
