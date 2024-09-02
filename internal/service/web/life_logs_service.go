package web

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
)

type LifeLogsService struct {
	LifeLogsStore dao.LifeLogsStore
}

func NewLifeLogsService(store dao.LifeLogsStore) *LifeLogsService {
	return &LifeLogsService{
		LifeLogsStore: store,
	}
}

func (lls *LifeLogsService) Lists(ctx context.Context) ([]models.LifeLogs, error) {
	return lls.LifeLogsStore.List(ctx, 1, 10)
}
