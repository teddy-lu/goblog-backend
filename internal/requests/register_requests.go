package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goblog-backend/internal/requests/validators"
)

type RegisterRequest struct {
	Username       string `json:"username,omitempty" valid:"username"`
	Password       string `json:"password,omitempty" valid:"password"`
	Email          string `json:"email" valid:"email"`
	RepeatPassword string `json:"repeat_password,omitempty" valid:"repeat_password"`
}

func Register(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username":        []string{"required", "between:3,10", "is_exists:users,username"},
		"password":        []string{"required", "min:6"},
		"email":           []string{"required", "email", "is_exists:users,email"},
		"repeat_password": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名不能为空",
			"between:长度为3-10个字符",
			"is_exists:用户已存在",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码不能小于6位",
		},
		"email": []string{
			"required:邮箱不能为空",
			"email:Email格式不正确，请提供有效的邮箱地址",
			"is_exists:邮箱已存在",
		},
		"repeat_password": []string{
			"required:确认密码框为必填项",
			"min:密码不能小于6位",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*RegisterRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.RepeatPassword, errs)

	return errs
}
