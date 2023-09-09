package infraestructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain/entity"
)

type TopicCategoryRepository struct {
	db *sql.DB
}

func NewTopicCategoryRepository(db *sql.DB) *TopicCategoryRepository {
	return &TopicCategoryRepository{db}
}

func (tr *TopicCategoryRepository) FindAll() ([]*entity.Topic_Category, error) {
	//Implementation
	return nil, nil
}
