package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/web"
	"goblog-backend/pkg/logger"
)

func Login(serv web.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 表单验证
		request := requests.LoginRequest{}
		if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
			return
		}

		// 登录校验
		u, err := serv.Login(c, request.Username, request.Password)
		if err != nil {
			logger.Error(err.Error())
			api.InternetServErr(c, err)
			return
		}

		// 获取登陆信息
		res, err := serv.Auth(&u)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Json(c, 200, "success", res)
	}
}

func Register(serv web.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := requests.RegisterRequest{}
		fmt.Println(request)
		if ok := requests.Validate(c, &request, requests.Register); !ok {
			return
		}

		// 注册用户
		u, err := serv.RegisterWithPwd(c, request)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Json(c, 200, "success", u)
	}

}
