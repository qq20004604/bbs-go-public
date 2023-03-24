package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	"main/service/userService"
	"main/utils"
)

/*IsOnline
* @Description: 判断用户是否已登录成功
* @param c
 */
func IsOnline(c *gin.Context) {
	if isOnline := userService.IsOnline(c); isOnline {
		utils.SuccessJson(c, "已登录", gin.H{})
	} else {
		utils.ErrorJson(c, "未登录")
	}
}
