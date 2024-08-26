package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Err     any         `json:"err"`
}

func Json(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
		Err:     nil,
	})
}

func Success(c *gin.Context, data interface{}) {
	code := http.StatusOK
	Json(c, code, "Success", data)
}

func Created(c *gin.Context, data interface{}) {
	code := http.StatusCreated
	Json(c, code, "Created Success", data)
}

func Deleted(c *gin.Context) {
	c.String(http.StatusNoContent, "Deleted Success")
}

func Error(c *gin.Context, code int, message string, err error) {
	c.AbortWithStatusJSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
		Err:     err.Error(),
	})
}

func ValidateErr(c *gin.Context, errs map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Response{
		Code:    http.StatusUnprocessableEntity,
		Message: "请求验证不通过，具体查看errors",
		Data:    nil,
		Err:     errs,
	})
}

func InternetServErr(c *gin.Context, err error) {
	Error(c, http.StatusInternalServerError, "服务器内部错误", err)
}

func Unauthorized(c *gin.Context, err error) {
	Error(c, http.StatusUnauthorized, "未授权", err)
}
