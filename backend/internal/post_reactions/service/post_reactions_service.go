package service

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/post_reactions/model"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post_reactions/repository"
)

type PostReactionService struct {
	Repository *repository.PostReactionRepository
}

func NewPostReactionService(postReactionRepository *repository.PostReactionRepository) *PostReactionService {
	return &PostReactionService{
		Repository: postReactionRepository,
	}
}

func (s *PostReactionService) CreatePostReaction(postID int, userID int) error {
	postReaction := &model.PostReaction{
		PostID: postID,
		UserID: userID,
	}

	return s.Repository.CreatePostReaction(postReaction)
}

func (s *PostReactionService) GetReactionsForPost(postID int) (*[]model.PostReaction, error) {
	return s.Repository.GetReactionsForPost(postID)
}
