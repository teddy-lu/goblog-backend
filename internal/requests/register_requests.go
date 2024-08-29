package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goblog-backend/internal/requests/validators"
)

type RegisterRequest struct {
	Username       string `json:"username,omitempty" valid:"username"`
	Password       string `json:"password,omitempty" valid:"password"`
	RepeatPassword string `json:"repeat_password,omitempty" valid:"repeat_password"`
	Email          string `json:"email" valid:"email"`
	Nickname       string `json:"nickname" valid:"nickname"`
}

func Register(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username":        []string{"required", "between:3,20", "alpha_num", "is_exists:users,username"},
		"password":        []string{"required", "min:6"},
		"nickname":        []string{"regex:^[\u4E00-\u9FFF0-9a-zA-Z_]+$"},
		"email":           []string{"required", "email", "is_exists:users,email"},
		"repeat_password": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required: 用户名不能为空",
			"between: 长度为3-20个字符",
			"alpha_num: 用户名格式错误，只允许数字和英文",
			"is_exists: 用户已存在",
		},
		"password": []string{
			"required: 密码不能为空",
			"min: 密码不能小于6位",
		},
		"email": []string{
			"required: 邮箱不能为空",
			"email: Email格式不正确，请提供有效的邮箱地址",
			"is_exists: 邮箱已存在",
		},
		"repeat_password": []string{
			"required: 确认密码框为必填项",
			"min: 密码不能小于6位",
		},
		"nickname": []string{"regex: 非法的字符串输入，只接受中文、英文、数字、下划线"},
	}

	errs := validate(data, rules, messages)
	_data := data.(*RegisterRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.RepeatPassword, errs)

	return errs
}
