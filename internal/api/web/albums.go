package web

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/requests"
	"goblog-backend/internal/service/web"
)

func ListAlbums(serv web.AlbumsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		albums, err := serv.Lists(c)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, albums)
	}
}

func CreateAlbum(serv web.AlbumsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := requests.AlbumRequest{}
		if ok := requests.Validate(c, &request, requests.AlbumManager); !ok {
			return
		}

		aid, err := serv.Create(c, request)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		album, err := serv.Get(c, aid)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Created(c, album)
	}
}
