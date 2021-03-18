package middleware

import (
	"github.com/gin-gonic/gin"
	"lan/gin-web/globar"
	"lan/gin-web/utils"
)

//简单的jwt验证
func JwtCheckMiddleware()  gin.HandlerFunc  {
	return func(c *gin.Context) {
		//token := c.GetHeader("token")
		token := c.Query("token")
		userinfo,err := utils.JwtParse(token)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "验证失败",
			})
			c.Abort()
			return
		}
		//fmt.Println("lllll:" , userinfo)
		globar.Logger.Warn("阿拉啦啦啦","ni hao a ",userinfo)
	}
}