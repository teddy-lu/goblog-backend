package web

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
)

type ArticlesService struct {
	ArticleStore dao.ArticlesStore
}

func NewArticleService(store dao.ArticlesStore) *ArticlesService {
	return &ArticlesService{ArticleStore: store}
}

func (as *ArticlesService) Lists(ctx context.Context) ([]models.Article, error) {
	return as.ArticleStore.List(ctx, 1, 10)
}

func (as *ArticlesService) Get(ctx context.Context, id int) (models.Article, error) {
	return as.ArticleStore.Detail(ctx, id)
}
