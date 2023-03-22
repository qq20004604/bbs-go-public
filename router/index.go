package router

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controller/BBSUserManage"
)

func LoadRoute(r *gin.Engine) {
	BaseUrl := config.Config.Runtime.BaseUrl
	// 登录
	r.POST(BaseUrl+"login", BBSUserManage.UserLogin)

	InternalRouter(r)
}
