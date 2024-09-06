package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ImageRequest struct {
	UserID      int    `json:"user_id,omitempty" valid:"user_id"`
	AlbumID     int    `json:"album_id,omitempty" valid:"album_id"`
	Title       string `json:"title,omitempty" valid:"title"`
	Path        string `json:"path,omitempty" valid:"path"`
	Alt         string `json:"alt,omitempty" valid:"alt"`
	Description string `json:"description,omitempty" valid:"description"`
	MediaType   string `json:"media_type,omitempty" valid:"media_type"`
}

func ImageManager(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:2", "max_cn:30"},
		"path":        []string{"required"},
		"alt":         []string{"required", "min_cn:2", "max_cn:50"},
		"description": []string{"required", "min_cn:2", "max_cn:100"},
		"media_type":  []string{"required", "in:jpg,png,jpeg"},
		"user_id":     []string{"required"},
		"album_id":    []string{"required", "not_exists:albums,album_id"},
	}

	message := govalidator.MapData{
		"title":       []string{"required:标题为必填项", "min_cn:标题长度需大于2", "max_cn:标题长度需小于30"},
		"path":        []string{"required:图片路径为必填项"},
		"alt":         []string{"required:图片别名为必填项", "min_cn:图片别名长度需大于2", "max_cn:图片别名长度需小于50"},
		"description": []string{"required:图片描述为必填项", "min_cn:图片描述长度需大于2", "max_cn:图片描述长度需小于100"},
		"media_type":  []string{"required:图片类型为必填项", "in:图片类型只能是jpg,png,jpeg"},
		"user_id":     []string{"required:用户ID为必填项"},
		"album_id":    []string{"required:相册ID为必填项", "not_exists:相册不存在"},
	}

	return validate(data, rules, message)
}
