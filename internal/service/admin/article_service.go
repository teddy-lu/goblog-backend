package admin

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
)

type ArticleService struct {
	ArticleStore dao.ArticlesStore
}

func NewArticleService(store dao.ArticlesStore) *ArticleService {
	return &ArticleService{ArticleStore: store}
}

func (ats *ArticleService) List(ctx context.Context) ([]models.Article, error) {
	var a []models.Article
	a, err := ats.ArticleStore.List(ctx, 1, 10)
	if err != nil {
		return nil, err
	}
	return a, nil
}
