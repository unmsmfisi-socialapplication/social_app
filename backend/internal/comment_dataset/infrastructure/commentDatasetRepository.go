package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment_dataset/domain"
)

type commentDatasetRepository struct {
	db *sql.DB
}

func NewCommentDatasetRepository(database *sql.DB) *commentDatasetRepository {
	return &commentDatasetRepository{db: database}
}

func (u *commentDatasetRepository) GetCommentsByTimestamp(start_date string, end_date string) (*[]domain.CommentDataset, error) {
	query := `SELECT * FROM FN_SOC_APP_GET_COMMENTS($1, $2)`
	rows, err := u.db.Query(query, start_date, end_date)
	if err != nil {
		return nil, err
	}

	var comments []domain.CommentDataset

	for rows.Next() {
		var com domain.CommentDataset
		err := rows.Scan(
			&com.InsertionDate,
			&com.CommentID,
			&com.PostID,
			&com.Username,
			&com.Comment,
			&com.Name,
			&com.Lastname,
			&com.Genre,
			&com.Country,
			&com.City,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}
