package routers

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api/index"
	"goblog-backend/internal/dao"
	idx "goblog-backend/internal/service/index"
	"goblog-backend/pkg/logger"
	"net/http"
)

type MyServer struct {
	demoService *idx.DemoService
}

func NewServer(demoStore dao.DemoStore) *MyServer {
	var demoService = idx.NewDemoService(demoStore)
	return &MyServer{demoService: demoService}
}

func (srv *MyServer) SetRouter(g *gin.Engine) *gin.Engine {
	g.Use(logger.MyLogger())
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	})

	// 设置路由
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	g.GET("/index", index.Index())
	g.GET("/demo", index.Demo(*srv.demoService))

	return g
}
