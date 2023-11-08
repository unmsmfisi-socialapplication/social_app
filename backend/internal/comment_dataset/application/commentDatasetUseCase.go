package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment_dataset/domain"
)

var (
	ErrType = errors.New("")
)

type CommentDatasetRepository interface {
	GetCommentsByTimestamp(start_date string, end_date string) (*[]domain.CommentDataset, error)
}

type CommentDatasetUseCase struct {
	repo CommentDatasetRepository
}

func NewCommentDatasetUseCase(r CommentDatasetRepository) *CommentDatasetUseCase {
	return &CommentDatasetUseCase{repo: r}
}

func (r *CommentDatasetUseCase) RetrieveDateScopedComments(start_date string, end_date string) (*[]domain.CommentDataset, error) {
	comments, err := r.repo.GetCommentsByTimestamp(start_date, end_date)
	if err != nil {
		return nil, err
	}

	return comments, err
}
