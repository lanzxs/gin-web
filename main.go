package main

import (
	"github.com/gin-gonic/gin"
	"lan/gin-web/config"
	"lan/gin-web/router"
	"lan/gin-web/utils"
)
func main() {
	//配置文件初始化
	config.Setup()
	//日志初始化
	utils.InitLogger()

	r := gin.Default()
	//路由配置初始化
	router.RouterInit(r)
	r.Run()
}
