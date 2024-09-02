package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type LifeLogsDao struct {
	db *gorm.DB
}

func NewLifeLogsDao(db *gorm.DB) LifeLogsStore {
	return &LifeLogsDao{db: db}
}

type LifeLogsStore interface {
	Create(ctx context.Context, lifeLog *models.LifeLogs) error
	Update(ctx context.Context, lifeLog *models.LifeLogs) error
	Delete(ctx context.Context, id int) error
	Detail(ctx context.Context, id int) (models.LifeLogs, error)
	List(ctx context.Context, page, limit int) ([]models.LifeLogs, error)
}

func (lld *LifeLogsDao) Create(ctx context.Context, lifeLog *models.LifeLogs) error {
	return lld.db.WithContext(ctx).Create(lifeLog).Error
}

func (lld *LifeLogsDao) Update(ctx context.Context, lifeLog *models.LifeLogs) error {
	return lld.db.WithContext(ctx).Model(lifeLog).Updates(lifeLog).Error
}

func (lld *LifeLogsDao) Delete(ctx context.Context, id int) error {
	return lld.db.WithContext(ctx).Delete(&models.LifeLogs{}, id).Error
}

func (lld *LifeLogsDao) Detail(ctx context.Context, id int) (models.LifeLogs, error) {
	var lifeLog models.LifeLogs
	err := lld.db.WithContext(ctx).First(&lifeLog, id).Error
	return lifeLog, err
}

func (lld *LifeLogsDao) List(ctx context.Context, page, limit int) ([]models.LifeLogs, error) {
	var lifeLogs []models.LifeLogs
	offset := (page - 1) * limit
	err := lld.db.WithContext(ctx).
		Limit(limit).Offset(offset).
		Order("created_at desc").
		Find(&lifeLogs).Error
	return lifeLogs, err
}
