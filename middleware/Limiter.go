package middleware

import (
	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
)

const limitMessage = `{"code": 0, "msg": "接口请求次数过多，请稍后再试"}`
const defaultRate = 1.0 // 每秒1次

/*
SetRateLimiter
  - @Description: 对单个接口进行限流设置，配置示例：
    r.POST(BaseUrl+"login", middleware.Limiter(limiter), BBSUserManage.UserLogin)
  - @param customRate		意思是每秒可以调用的次数，默认是1次
  - @return gin.HandlerFunc
*/
func SetRateLimiter(customRate ...float64) gin.HandlerFunc {
	rate := defaultRate
	if len(customRate) > 0 {
		rate = customRate[0]
	}

	limiter := tollbooth.NewLimiter(rate, nil)
	limiter.SetMessage(limitMessage)
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, limiter.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
