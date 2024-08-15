package web

import (
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/web"

	"github.com/gin-gonic/gin"
)

func WebLogin(serv web.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		serv.Login()
		api.Json(c, 200, "success", nil)
	}
}
