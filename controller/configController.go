package controller

import (
	"github.com/gin-gonic/gin"
	"lan/gin-web/config"
	"lan/gin-web/utils"
	"time"
)

/**
获取配置
 */
func GetConfig(c *gin.Context)  {
	c.JSON(200,gin.H{"lan":config.DatabaseSetting.Type})
}

//登录
func Login(c *gin.Context)  {
	id := utils.CreateSnowFlake()
	jwt,_ := utils.JwtSign(id,"lan",300*time.Second)
	c.JSON(200,gin.H{"lan": jwt})
}

//uuid
func GetUUID(c *gin.Context)  {
	//c.JSON(200,gin.H{"lan":utils.CreateUUID()})
	//c.JSON(200,gin.H{"lan":utils.CreateSnowFlake()})
	id := utils.CreateSnowFlake()
	jwt,_ := utils.JwtSign(id,"lan",300*time.Second)
	c.JSON(200,gin.H{"lan": jwt})
	user,_ := utils.JwtParse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjI3MTEwNDczOTY0MjIzMjAxMjksIlVzZXJOYW1lIjoibGFuIiwiZXhwIjoxNjE1OTEwNDcwLCJpYXQiOjE2MTU5MTAxNzAsIm5iZiI6MTYxNTkxMDE3MH0.2lvRgLhl1Z29xNYEgdI5NzXX2jFstJph0XGro55hNOg")
	c.JSON(200,gin.H{"lan": user})
}
