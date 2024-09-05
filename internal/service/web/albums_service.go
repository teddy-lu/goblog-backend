package web

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
)

type AlbumsService struct {
	AlbumStore dao.AlbumsStore
}

func NewAlbumsService(store dao.AlbumsStore) *AlbumsService {
	return &AlbumsService{AlbumStore: store}
}

func (as *AlbumsService) Lists(ctx context.Context) ([]models.Album, error) {
	return as.AlbumStore.List(ctx, 1, 10)
}

func (as *AlbumsService) Get(ctx context.Context, id int64) (models.Album, error) {
	return as.AlbumStore.Detail(ctx, id)
}

func (as *AlbumsService) Create(ctx context.Context, album requests.AlbumRequest) (int64, error) {
	data := models.Album{
		UserID:  1,
		Title:   album.Title,
		Slug:    album.Slug,
		Private: album.Private,
	}
	return as.AlbumStore.CreateAlbum(ctx, &data)
}
