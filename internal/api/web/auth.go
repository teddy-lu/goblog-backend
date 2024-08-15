package web

import (
	"goblog-backend/internal/api"

	"github.com/gin-gonic/gin"
)

func WebLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		api.Json(c, 200, "success", nil)
	}
}
