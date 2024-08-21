package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/admin"
	"goblog-backend/pkg/jwt"
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

		// 获取jwt token
		authJWT, err := jwt.NewJWT()
		if err != nil {
			logger.Error(err.Error())
			api.InternetServErr(c, err)
			return
		}
		token, expiredAt := authJWT.IssueToken(u.ID, u.Username)
		if token == "" {
			api.InternetServErr(c, errors.New("token生成失败"))
			return
		}

		// 返回响应
		res := UserInfo{
			User:      u,
			Token:     token,
			ExpiredAt: expiredAt,
		}
		api.Json(c, 200, "success", res)
	}
}
