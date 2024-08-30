package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api"
	"goblog-backend/internal/service/web"
	"strconv"
)

func ListArticle(serv web.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := serv.Lists(c)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, res)
	}
}

func GetArticle(serv web.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			api.InternetServErr(c, errors.New("非法路径参数"))
			return
		}

		res, err := serv.Get(c, aid)
		if err != nil {
			api.InternetServErr(c, err)
			return
		}
		api.Success(c, res)
	}
}
