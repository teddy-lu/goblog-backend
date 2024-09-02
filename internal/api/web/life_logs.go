package web

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/web"
)

func ListLogs(serv web.LifeLogsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := serv.Lists(c)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, res)
	}
}
