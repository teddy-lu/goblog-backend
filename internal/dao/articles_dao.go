package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type ArticlesDao struct {
	db *gorm.DB
}

func NewArticlesDao(db *gorm.DB) ArticlesStore {
	return &ArticlesDao{db: db}
}

type ArticlesStore interface {
	List(ctx context.Context, page, limit int) ([]models.Article, error)
	Detail(ctx context.Context, id int) (models.Article, error)
}

func (a *ArticlesDao) List(ctx context.Context, page, limit int) ([]models.Article, error) {
	var articles []models.Article

	offset := (page - 1) * limit
	err := a.db.WithContext(ctx).
		Preload("User").
		Preload("Comments").
		Preload("Tags").
		Limit(limit).Offset(offset).
		Order("created_at desc").
		Find(&articles).Error
	return articles, err
}

func (a *ArticlesDao) Detail(ctx context.Context, id int) (models.Article, error) {
	var article models.Article
	err := a.db.WithContext(ctx).
		Preload("User").
		Preload("Comments").
		Preload("Tags").
		First(&article, id).Error
	return article, err
}
