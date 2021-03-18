package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lan/gin-web/globar"
	"net/http"
	"runtime/debug"
)

// 全局异常处理中间件
func RecoverException(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 将异常写入日志
			globar.Logger.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			// 服务器异常

			// 以json方式写入响应
			c.JSON(http.StatusOK, gin.H{"errmsg": err})
			c.Abort()
			return
		}
	}()
	c.Next()
}
