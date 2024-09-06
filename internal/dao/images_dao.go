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
	Create(ctx context.Context, image *models.Image) (id int64, err error)
	Detail(ctx context.Context, id int64) (models.Image, error)
	//Update(ctx context.Context, id int, image *models.Image) error
	//Delete(ctx context.Context, id int) error
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

func (i *ImagesDao) Create(ctx context.Context, image *models.Image) (id int64, err error) {
	err = i.db.WithContext(ctx).Create(&image).Error
	if err != nil {
		return 0, err
	}
	return image.ID, nil
}

func (i *ImagesDao) Detail(ctx context.Context, id int64) (models.Image, error) {
	var image models.Image
	err := i.db.WithContext(ctx).
		Model(models.Image{}).
		Where(&models.Image{BaseModel: models.BaseModel{ID: id}}).
		First(&image).Error
	return image, err
}
