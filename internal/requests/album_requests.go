package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AlbumRequest struct {
	Title   string `json:"title,omitempty" valid:"title"`
	Slug    string `json:"slug,omitempty" valid:"slug"`
	Private int    `json:"private,omitempty" valid:"private"`
}

func AlbumManager(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":   []string{"required", "between:0,30"},
		"slug":    []string{"required", "between:0,100"},
		"private": []string{"required", "in:0,1"},
	}

	messages := govalidator.MapData{
		"slug": []string{
			"required: 内容不能为空",
			"between: 内容长度需在 0~100 之间",
		},
		"title": []string{
			"required: 标题不能为空",
			"between: 标题长度需在 0~30 之间",
		},
		"private": []string{
			"required: 权限不能为空",
			"in: 权限只能为 0 或 1",
		},
	}

	return validate(data, rules, messages)
}
