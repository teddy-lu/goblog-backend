package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/web"
	"strconv"
)

func ListImages(serv web.ImagesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			api.InternetServErr(c, errors.New("非法路径参数"))
			return
		}

		images, err := serv.List(c, aid)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, images)
	}
}

func CreateImage(serv web.ImagesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := requests.ImageRequest{}
		if ok := requests.Validate(c, &request, requests.ImageManager); !ok {
			return
		}

		imgID, err := serv.Create(c, request)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		image, err := serv.Get(c, imgID)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, image)
	}
}
