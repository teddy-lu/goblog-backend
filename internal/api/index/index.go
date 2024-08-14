package index

import (
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/index"
	"goblog-backend/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("test info.....")
		logger.Debug("test debug.....")
		logger.Error("test error.....")

		api.Json(c, 200, "success", nil)
	}
}

func Demo(service index.DemoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := service.List(c)
		if err != nil {
			api.Error(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		api.Json(c, 200, "success", data)
	}
}
