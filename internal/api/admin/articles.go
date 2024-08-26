package admin

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/admin"
	"goblog-backend/pkg/logger"
)

func ArticlesList(service admin.ArticleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.MustGet("uid").(int)
		logger.Info("uid: %v", uid)

		articles, err := service.List(c)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, articles)
	}
}
