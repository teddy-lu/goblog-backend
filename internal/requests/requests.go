package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goblog-backend/internal/api"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := c.ShouldBind(obj); err != nil {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求解析失败，请使用 application/json 请求格式",
		// 	"error":   err.Error(),
		// })
		api.Error(c, 400, "请求解析失败，请使用 application/json 请求格式", err)
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		api.ValidateErr(c, errs)
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的struct标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
