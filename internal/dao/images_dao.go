package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type ImagesDao struct {
	db *gorm.DB
}

func NewImagesDao(db *gorm.DB) ImagesStore {
	return &ImagesDao{db: db}
}

type ImagesStore interface {
	List(ctx context.Context, albumID, page, limit int) ([]models.Image, error)
	//Create(ctx context.Context, image *models.Image) (id int64, err error)
	//Update(ctx context.Context, id int, image *models.Image) error
	//Delete(ctx context.Context, id int) error
	//Detail(ctx context.Context, id int64) (models.Image, error)
}

func (i *ImagesDao) List(ctx context.Context, albumID, page, limit int) ([]models.Image, error) {
	var images []models.Image
	offset := (page - 1) * limit
	err := i.db.WithContext(ctx).
		Where(&models.Image{AlbumID: uint(albumID)}).
		Limit(limit).Offset(offset).
		Order("created_at desc").
		Find(&images).Error
	return images, err
}
