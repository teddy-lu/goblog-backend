package admin

import (
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/admin"

	"github.com/gin-gonic/gin"
)

func AdminLogin(service admin.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Login()
		api.Json(c, 200, "success", nil)
	}
}
