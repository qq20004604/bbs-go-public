package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/service/userService"
	"main/utils"
)

/*BBSUserLogout
* @Description: 登出
 */
func BBSUserLogout(c *gin.Context) {
	// 1. 先拿到cookie
	token, errCookie := c.Cookie(config.Config.Common.HeaderTokenName)
	// 如果报错了，一般是没找到这个 cookie，但无论如何，把cookie清除掉
	if errCookie != nil {
		utils.ErrorJson(c, "你尚未登录，无需登出")
		return
	}
	userService.Logout(c, token)

	// 登出成功
	utils.SuccessJson(c, "登出成功", gin.H{})
}
