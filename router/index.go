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
	// 登录检测
	r.POST(BaseUrl+"isOnline", BBSUserManage.IsOnline)

	InternalRouter(r)
}
