package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
}

func LoginByPassword(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:用户名称必须"},
		"password": []string{"required:密码必须"},
	}

	return validate(data, rules, messages)
}
