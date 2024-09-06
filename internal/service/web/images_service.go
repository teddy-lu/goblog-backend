package web

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
)

type ImagesService struct {
	ImageStore dao.ImagesStore
}

func NewImagesService(imageStore dao.ImagesStore) *ImagesService {
	return &ImagesService{ImageStore: imageStore}
}

func (s *ImagesService) List(ctx context.Context, id int) ([]models.Image, error) {
	return s.ImageStore.List(ctx, id, 1, 10)
}

func (s *ImagesService) Create(ctx context.Context, image requests.ImageRequest) (int64, error) {
	data := models.Image{
		AlbumID: uint(image.AlbumID),
		UserID:  uint(image.UserID),
		Title:   image.Title,
		Path:    image.Path,
		//Url:     image.Url,
		Alt: image.Alt,
		//IsPrivate: image.IsPrivate,
		Description: image.Description,
		MediaType:   image.MediaType,
	}
	return s.ImageStore.Create(ctx, &data)
}

func (s *ImagesService) Get(ctx context.Context, id int64) (models.Image, error) {
	return s.ImageStore.Detail(ctx, id)
}
