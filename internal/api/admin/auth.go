package admin

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/admin"
)

//type LoginRequest struct {
//	Username string `json:"username,omitempty" valid:"required"`
//	Password string `json:"password,omitempty" valid:"required"`
//}

func AdminLogin(service admin.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := requests.LoginRequest{}
		if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
			return
		}

		//if err := c.ShouldBind(c.Request); err != nil {
		//	api.Error(c, 400, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", nil)
		//	return
		//}
		//
		//rules := govalidator.MapData{
		//	"username": []string{"required"},
		//	"password": []string{"required"},
		//}
		//
		//messages := govalidator.MapData{
		//	"username": []string{"required:用户名称必须"},
		//	"password": []string{"required:密码必须"},
		//}
		//
		//opts := govalidator.Options{
		//	TagIdentifier: "valid",
		//	Rules:         rules,
		//	Data:          &LoginRequest{},
		//	Messages:      messages,
		//}
		//
		//v := govalidator.New(opts).ValidateStruct()
		//logger.Debug("validate res:", v)

		u := service.Login(c)
		api.Json(c, 200, "success", u)
	}
}
