package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"lan/gin-web/globar"
	"time"
)

func RateLimiter() gin.HandlerFunc {
	// 创建速率配置
	rate := limiter.Rate{
		Period: time.Minute,
		Limit:  1,
	}
	// 将数据存入内存
	store := memory.NewStore()

	// 创建速率实例, 必须是真实的请求
	instance := limiter.New(store, rate, limiter.WithTrustForwardHeader(true))

	globar.Logger.Infof("初始化速率限制中间件完成")

	// 生成gin中间件
	return mgin.NewMiddleware(instance)
}
