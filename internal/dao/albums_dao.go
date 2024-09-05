package dao

import (
	"context"
	"goblog-backend/internal/models"
	"gorm.io/gorm"
)

type AlbumsDao struct {
	db *gorm.DB
}

func NewAlbumsDao(db *gorm.DB) AlbumsStore {
	return &AlbumsDao{db: db}
}

type AlbumsStore interface {
	List(ctx context.Context, page, limit int) ([]models.Album, error)
	Detail(ctx context.Context, id int64) (models.Album, error)
	CreateAlbum(ctx context.Context, album *models.Album) (id int64, err error)
	UpdateAlbum(ctx context.Context, id int, album *models.Album) error
	DeleteAlbum(ctx context.Context, id int) error
}

func (a *AlbumsDao) List(ctx context.Context, page, limit int) ([]models.Album, error) {
	var albums []models.Album

	offset := (page - 1) * limit
	err := a.db.WithContext(ctx).
		Limit(limit).Offset(offset).
		Order("created_at desc").
		Find(&albums).Error
	return albums, err
}

func (a *AlbumsDao) Detail(ctx context.Context, id int64) (models.Album, error) {
	var album models.Album
	err := a.db.WithContext(ctx).
		Where(&models.Album{BaseModel: models.BaseModel{ID: id}}).
		First(&album).Error
	return album, err
}

func (a *AlbumsDao) CreateAlbum(ctx context.Context, album *models.Album) (id int64, err error) {
	err = a.db.WithContext(ctx).Create(album).Error
	id = album.ID
	return id, err
}

func (a *AlbumsDao) UpdateAlbum(ctx context.Context, id int, album *models.Album) error {
	err := a.db.WithContext(ctx).
		Model(&models.Album{}).
		Where("id =?", id).
		Save(album).Error
	return err
}

func (a *AlbumsDao) DeleteAlbum(ctx context.Context, id int) error {
	err := a.db.WithContext(ctx).Delete(&models.Album{}, id).Error
	return err
}
