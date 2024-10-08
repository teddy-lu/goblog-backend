package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog-backend/config"
	"goblog-backend/db"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/pkg/logger"
	"goblog-backend/routers"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type app struct {
	config *config.Config
	server *http.Server
}

func newApp(config *config.Config, server *http.Server) *app {
	return &app{config: config, server: server}
}

func createServEngine(cfg *config.Config, g *gin.Engine) *app {
	mainApp := newApp(cfg, &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Addr),
		Handler:      g,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})
	return mainApp
}

func createDb(cfg *config.Config) *gorm.DB {
	// 链接mysql
	dbm, m := db.GetMysqlPool().InitPool(cfg)
	if !m {
		logger.Error("init database pool failure...")
		panic("mysql init failed")
	}
	mErr := migratorDb(dbm)
	if mErr != nil {
		logger.Error("migrator database failure...")
		panic("mysql database migrator failed")
	}
	// 链接redis
	db.InitRedis(cfg)

	return dbm
}

func migratorDb(dbm *gorm.DB) error {
	log.Println("migrator database...")
	return dbm.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Comment{},
		&models.Tag{},
		&models.LifeLogs{},
		&models.Image{},
		&models.Album{},
		&models.MessageBoard{},
		&models.ArticleLikes{},
		//&models.Category{},
	)
}

func createGinServer(dbm *gorm.DB, mode string) *gin.Engine {
	demoDao := dao.NewDemoDao(dbm)
	userDao := dao.NewUsersDao(dbm)
	articleDao := dao.NewArticlesDao(dbm)
	lifeLogsDao := dao.NewLifeLogsDao(dbm)
	albumDao := dao.NewAlbumsDao(dbm)
	imageDao := dao.NewImagesDao(dbm)
	//commentDao := dao.NewCommentsDao(dbm)
	//tagDao := dao.NewTagsDao(dbm)
	//commentDao.SetTagDao(tagDao)
	//articleDao.SetCommentDao(commentDao)

	serv := routers.NewServer(demoDao, userDao, articleDao, lifeLogsDao, albumDao, imageDao)
	if mode == "debug" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	serv.SetRouter(g)
	return g
}
