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
	// 在线检测
	r.POST(BaseUrl+"isOnline", BBSUserManage.IsOnline)
	// 注册账号
	r.POST(BaseUrl+"register", BBSUserManage.RegisterBBSUser)

	InternalRouter(r)
}
