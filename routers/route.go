package routers

import (
	"github.com/gin-gonic/gin"
	"goblog-backend/internal/api/admin"
	"goblog-backend/internal/api/index"
	"goblog-backend/internal/api/web"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/middlewares"
	admServ "goblog-backend/internal/service/admin"
	idx "goblog-backend/internal/service/index"
	webServ "goblog-backend/internal/service/web"
	"goblog-backend/pkg/logger"
	"net/http"
)

type MyServer struct {
	demoService    *idx.DemoService
	admAuthService *admServ.AuthService
	webAuthService *webServ.AuthService
}

func NewServer(
	demoStore dao.DemoStore,
	userStore dao.UsersStore,
) *MyServer {
	var demoService = idx.NewDemoService(demoStore)
	var adminAuthService = admServ.NewAuthService(userStore)
	var webAuthService = webServ.NewAuthService(userStore)

	return &MyServer{
		demoService:    demoService,
		admAuthService: adminAuthService,
		webAuthService: webAuthService,
	}
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

	// 前端页面的路由
	webG := g.Group("/api/v1")
	webG.POST("/login", web.WebLogin(*srv.webAuthService))

	// 后台页面的路由
	adminG := g.Group("/admin")
	adminG.POST("/login", admin.AdminLogin(*srv.admAuthService))
	adminG.GET("/articles", middlewares.AdminAuth(), admin.AdminLogin(*srv.admAuthService))

	return g
}
