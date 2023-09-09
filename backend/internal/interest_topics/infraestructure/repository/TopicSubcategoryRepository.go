package infraestructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

type TopicSubcategoryRepository struct {
	db *sql.DB
}

func NewTopicSubcategoryRepository(db *sql.DB) *TopicSubcategoryRepository {
    return &TopicSubcategoryRepository{db}
}

func (tsr *TopicSubcategoryRepository) FindAll() ([]*entity.Topic_Subcategory, error) {
	//Implementation
	return nil, nil
}
