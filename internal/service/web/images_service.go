package web

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
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
