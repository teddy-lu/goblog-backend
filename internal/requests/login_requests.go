package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Username string `json:"username,omitempty" valid:"required"`
	Password string `json:"password,omitempty" valid:"required"`
}

func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
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
