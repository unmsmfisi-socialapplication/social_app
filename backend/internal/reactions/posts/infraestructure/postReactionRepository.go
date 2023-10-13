package infraestructure

import "gorm.io/gorm"

type PostReactionRepository struct {
	Database *gorm.DB
}

func NewPostReactionRepository(database *gorm.DB) *PostReactionRepository {
	return &PostReactionRepository{
		Database: database,
	}
}
