package router

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controller"
)

func LoadRoute(r *gin.Engine) {
	BaseUrl := config.Config.Runtime.BaseUrl
	r.POST(BaseUrl+"test", controller.Test)

	//r.POST(BaseUrl+"createUser", BBSUserManage.CreateUser)

}
