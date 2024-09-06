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
	demoService        *idx.DemoService
	admAuthService     *admServ.AuthService
	webAuthService     *webServ.AuthService
	admArticleService  *admServ.ArticleService
	webArticleService  *webServ.ArticlesService
	webLifeLogsService *webServ.LifeLogsService
	webAlbumsService   *webServ.AlbumsService
	webImagesService   *webServ.ImagesService
}

func NewServer(
	demoStore dao.DemoStore,
	userStore dao.UsersStore,
	ArticleStore dao.ArticlesStore,
	LifeLogsStore dao.LifeLogsStore,
	AlbumsStore dao.AlbumsStore,
	ImagesStore dao.ImagesStore,
	// CommentStore dao.CommentsStore,
	// TagStore dao.TagsStore,
) *MyServer {
	var demoService = idx.NewDemoService(demoStore)
	var adminAuthService = admServ.NewAuthService(userStore)
	var webAuthService = webServ.NewAuthService(userStore)
	var adminArticleService = admServ.NewArticleService(ArticleStore)
	var webArticleService = webServ.NewArticleService(ArticleStore)
	var webLifeLogsService = webServ.NewLifeLogsService(LifeLogsStore)
	var webAlbumsService = webServ.NewAlbumsService(AlbumsStore)
	var webImagesService = webServ.NewImagesService(ImagesStore)

	return &MyServer{
		demoService:        demoService,
		admAuthService:     adminAuthService,
		webAuthService:     webAuthService,
		admArticleService:  adminArticleService,
		webArticleService:  webArticleService,
		webLifeLogsService: webLifeLogsService,
		webAlbumsService:   webAlbumsService,
		webImagesService:   webImagesService,
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
	webG := g.Group("/front-api/v1")
	webG.POST("/login", web.Login(*srv.webAuthService))
	webG.POST("/register", web.Register(*srv.webAuthService))
	webG.GET("/articles", web.ListArticle(*srv.webArticleService))
	webG.GET("/article/:id", web.GetArticle(*srv.webArticleService))
	webG.GET("/life-logs", web.ListLogs(*srv.webLifeLogsService))
	webG.GET("/albums", web.ListAlbums(*srv.webAlbumsService))
	webG.POST("/albums", web.CreateAlbum(*srv.webAlbumsService))
	webG.GET("/albums/:id/images", web.ListImages(*srv.webImagesService))
	webG.POST("/albums/:id/images", web.CreateImage(*srv.webImagesService))

	// 后台页面的路由
	adminG := g.Group("/admin-api")
	adminG.POST("/login", admin.Login(*srv.admAuthService))
	adminG.GET("/articles", middlewares.AdminAuth(), admin.ArticlesList(*srv.admArticleService))

	return g
}
