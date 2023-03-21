package router

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controller"
)

/*InternalRouter
* @Description: 私有路径和服务，故开源项目里，这里为空
* @param r
 */
func InternalRouter(r *gin.Engine) {
	BaseUrl := config.Config.Runtime.BaseUrl
	r.POST(BaseUrl+"test", controller.Test)

	//r.POST(BaseUrl+"createUser", BBSUserManage.CreateUser)
	InternalRouter(r)
}
