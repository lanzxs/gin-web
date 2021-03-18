package router

import (
	"github.com/gin-gonic/gin"
	"lan/gin-web/middleware"
)
import "lan/gin-web/controller"

func RouterInit(engine *gin.Engine)  {

	//获取配置
	engine.GET("/config",controller.GetConfig)
	//登录
	engine.GET("/login",controller.Login)
	//uuid
	engine.GET("/uuid",middleware.JwtCheckMiddleware(),controller.GetUUID)
}
