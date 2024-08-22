package middlewares

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/pkg/jwt"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, _ := jwt.NewJWT()
		claims, err := t.ParseToken(c)
		if err != nil {
			api.Unauthorized(c, err)
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.Username)

		c.Next()
	}
}
