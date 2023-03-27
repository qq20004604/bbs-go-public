package router

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controller"
	"main/middleware"
)

/*InternalRouter
* @Description: 私有路径和服务，不对外开放的，都放在这里
* @param r
 */
func InternalRouter(r *gin.Engine) {
	BaseUrl := config.Config.Runtime.BaseUrl
	r.POST(BaseUrl+"test", middleware.SetRateLimiter(0.5), controller.Test)
}
