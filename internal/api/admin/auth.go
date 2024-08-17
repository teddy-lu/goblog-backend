package admin

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/admin"
	"net/http"
)

type UserInfo struct {
	User      models.User `json:"user_info"`
	Token     string      `json:"token"`
	ExpiredAt int64       `json:"expired_at"`
}

func AdminLogin(service admin.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 表单验证
		request := requests.LoginRequest{}
		if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
			return
		}

		// 登录校验
		u, err := service.Login(c, request.Username, request.Password)
		if err != nil {
			api.Error(c, http.StatusInternalServerError, "服务器内部错误", err)
			return
		}

		// 获取jwt token

		// 返回响应
		res := UserInfo{
			User:      u,
			Token:     "123456",
			ExpiredAt: 10000,
		}
		api.Json(c, 200, "success", res)
	}
}
