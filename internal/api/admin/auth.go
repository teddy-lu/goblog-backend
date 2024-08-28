package admin

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/admin"
	"goblog-backend/pkg/logger"
	"time"
)

type UserInfo struct {
	User      models.User `json:"user_info"`
	Token     string      `json:"token"`
	ExpiredAt time.Time   `json:"expired_at"`
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
			logger.Error(err.Error())
			api.InternetServErr(c, err)
			return
		}

		// 获取登陆信息
		res, err := service.Auth(&u)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Json(c, 200, "success", res)
	}
}
