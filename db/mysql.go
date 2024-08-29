package db

import (
	"errors"
	"fmt"
	"goblog-backend/config"
	zapLog "goblog-backend/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type MysqlPool struct {
}

var instance *MysqlPool
var once sync.Once

var Db *gorm.DB
var err error

func GetMysqlPool() *MysqlPool {
	once.Do(func() {
		// 初始化连接池
		instance = &MysqlPool{}
	})
	return instance
}

func (pool MysqlPool) InitPool(cfg *config.Config) (orm *gorm.DB, isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
		cfg.Db.Charset,
	)
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	zapLogger := zapLog.New()
	Db, err = gorm.Open(dbConfig, &gorm.Config{
		Logger:                                   zapLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(errors.New("init mysql pool failed"))
		return nil, false
	}

	dbase, err := Db.DB()
	if err != nil {
		panic(errors.New("get mysql DB failed"))
		return nil, false
	}

	// 设置最大空闲连接数
	dbase.SetMaxIdleConns(cfg.Db.MaxIdleCons)
	// 设置最大连接数
	dbase.SetMaxOpenConns(cfg.Db.MaxOpenCons)
	// 设置每个链接的过期时间
	//dbase.SetConnMaxLifetime(time.Duration(viper.GetInt("db.conn_max_lifetime")) * time.Second)

	return Db, true
}
