package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/reactions/posts/infraestructure"
	"gorm.io/gorm"
	"log"
)

type PostReactionUseCaseInterface interface {
	GetByID(postReactionID int64) (*domain.PostReaction, error)
	Create(postReaction *domain.PostReaction) error
	// Update(postReactionID int64, postReaction *domain.PostReaction) error
	Delete(postReactionID int64) error
}

type PostReactionUseCase struct {
	repository *infraestructure.PostReactionRepository
}

func NewPostReactionUseCase(r *infraestructure.PostReactionRepository) *PostReactionUseCase {
	return &PostReactionUseCase{
		repository: r,
	}
}

func (p *PostReactionUseCase) GetByID(postReactionID int64) (*domain.PostReaction, error) {
	var postReaction domain.PostReaction
	if err := p.repository.Database.First(&postReaction, postReactionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Reaction not found")
		}
		return nil, err
	}

	return &postReaction, nil
}

func (p *PostReactionUseCase) Create(postReaction domain.PostReaction) error {
	err := p.repository.Database.Create(&postReaction).Error

	if err != nil {
		log.Println("Cannot insert that reaction")
		return err
	}

	return nil

}

func (p *PostReactionUseCase) Delete(PostReactionId int64) error {
	var PostReaction domain.PostReaction

	if err := p.repository.Database.First(&PostReaction, PostReactionId).Error; err != nil {
		log.Println("Reaction not found")
		return err
	}

	if err := p.repository.Database.Delete(&PostReaction).Error; err != nil {
		log.Println("Cannot delete that PostReaction")
		return err
	}

	return nil
}
