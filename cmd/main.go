package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"goblog-backend/config"
	"goblog-backend/pkg/logger"
	rsapkg "goblog-backend/pkg/rsa"
	"log"
)

var conf = pflag.StringP("conf", "c", "", "config filepath")

func main() {
	pflag.Parse()

	// 读取配置
	c, cErr := config.Run(*conf)
	if cErr != nil {
		panic(cErr)
	}

	// 读取环境配置mode模式
	mode := viper.GetString("mode")
	// 初始化日志
	logger.InitLogger()
	// 初始化db类
	dbm := createDb(c)

	// 生成rsa密钥对
	rsapkg.GenerateRsaKey()

	// 实例化server参数，并启动gin
	fmt.Println("server start...")
	g := createGinServer(dbm, mode)

	// Listen and Server in 0.0.0.0:8080
	//if err := g.Run(fmt.Sprintf(":%d", c.Addr)); err != nil {
	//	panic(err)
	//	return
	//}

	appSrv := createServEngine(c, g)
	log.Printf("Server: http://127.0.0.1:%d", c.Addr)
	for _, r := range g.Routes() {
		fmt.Printf("path: %s, method: %s\n", r.Path, r.Method)
	}
	//log.Fatalln()是一个日志函数，用于记录一条错误日志。如果ListenAndServe()方法返回错误，它会调用log.Fatalln()来记录这条错误日志，并导致程序立即退出
	log.Fatalln(appSrv.server.ListenAndServe())
}
