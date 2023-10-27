package application

import (
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/domain"
)

type CommentUseCaseInterface interface {
	GetByID(commentID int64) (*domain.Comment, error)
	Create(comment *domain.Comment) error
	Update(commentID int64, comment *domain.Comment) error
	Delete(commentID int64) error
}

type CommentRepository interface {
	GetCommentByID(commentID int64) (*domain.Comment, error)
	CreateComment(comment *domain.Comment) error
	UpdateComment(commentID int64, comment *domain.Comment) error
	DeleteComment(commentID int64) error
}

type CommentUseCase struct {
	repo CommentRepository
}

func NewCommentUseCase(r CommentRepository) *CommentUseCase {
	return &CommentUseCase{repo: r}
}

func (c *CommentUseCase) GetByID(commentID int64) (*domain.Comment, error) {
	comment, err := c.repo.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}

func (c *CommentUseCase) Create(comment *domain.Comment) error {
	return c.repo.CreateComment(comment)
}

func (c *CommentUseCase) Update(commentID int64, comment *domain.Comment) error {
	return c.repo.UpdateComment(commentID, comment)
}

func (c *CommentUseCase) Delete(commentID int64) error {
	return c.repo.DeleteComment(commentID)
}