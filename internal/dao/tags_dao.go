package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type TagsDao struct {
	db *gorm.DB
}

func NewTagsDao(db *gorm.DB) TagsStore {
	return &TagsDao{db: db}
}

type TagsStore interface {
	List(ctx context.Context) ([]models.Tag, error)
}

func (t *TagsDao) List(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag
	err := t.db.Find(&tags).Error
	return tags, err
}
