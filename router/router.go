package router

import (
	"github.com/gin-gonic/gin"
	"lan/gin-web/middleware"
)
import "lan/gin-web/controller"

func RouterInit(engine *gin.Engine) {

	//全局异常处理
	engine.Use(middleware.RecoverException)
	//访问日志
	engine.Use(middleware.AccessLog)
	//初始化速率限制
	engine.Use(middleware.RateLimiter())
	//获取配置
	engine.GET("/config", controller.GetConfig)
	//登录
	engine.GET("/login", controller.Login)
	//uuid
	engine.GET("/uuid", middleware.JwtCheckMiddleware(), controller.GetUUID)
}
