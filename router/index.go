package router

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controller/BBSUserManage"
	"main/middleware"
)

func LoadRoute(r *gin.Engine) {
	BaseUrl := config.Config.Runtime.BaseUrl
	// 参数2意思是每秒可以调用几次。使用中间件但不添加参数是默认 1次/秒，完全不添加这个中间件的意思是无限制

	// 登录
	r.POST(BaseUrl+"login", middleware.SetRateLimiter(0.2), BBSUserManage.UserLogin)
	// 在线检测
	r.POST(BaseUrl+"isOnline", middleware.SetRateLimiter(4), BBSUserManage.IsOnline)
	// 登出
	r.POST(BaseUrl+"logout", middleware.SetRateLimiter(), BBSUserManage.BBSUserLogout)
	// 注册账号。每秒0.1次，意思是10秒只能调用一次
	r.POST(BaseUrl+"register", middleware.SetRateLimiter(0.1), BBSUserManage.RegisterBBSUser)
	// 获取用户信息
	r.POST(BaseUrl+"getUserInfo", middleware.SetRateLimiter(), BBSUserManage.GetBBSUserInfo)
	// 管理员分页获取所有用户信息
	r.POST(BaseUrl+"getUsersInfoByPage", middleware.SetRateLimiter(), BBSUserManage.GetUsersInfoByPage)
	// 批量更新用户状态
	r.POST(BaseUrl+"updateUserStatus", middleware.SetRateLimiter(), BBSUserManage.BatchUpdateUserStatus)
	// 更新本人信息
	r.POST(BaseUrl+"updateSelfInfo", middleware.SetRateLimiter(), BBSUserManage.UpdateSelfInfo)
	// 管理员修改其他用户信息
	r.POST(BaseUrl+"updateUserInfo", middleware.SetRateLimiter(), BBSUserManage.UpdateUserInfo)

	InternalRouter(r)
}
