package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type CommentsDao struct {
	db *gorm.DB
}

func NewCommentsDao(db *gorm.DB) CommentsStore {
	return &CommentsDao{db: db}
}

type CommentsStore interface {
	List(ctx context.Context) ([]models.Comment, error)
}

func (dao CommentsDao) List(ctx context.Context) ([]models.Comment, error) {
	var comments []models.Comment
	err := dao.db.Find(&comments).Error
	return comments, err
}
