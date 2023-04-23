package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*UpdateSelfPassword
* @Description: 用户本人修改密码
* @param c
 */
func UpdateSelfPassword(c *gin.Context) {
	var data userService.UpdateSelfPasswordRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 先判断是否已登录
	isOnline, userID := userService.IsOnline(c)
	if !isOnline {
		utils.ErrorJson(c, "请登录后再进行操作")
		return
	}

	// 2. 然后执行密码更新
	if err := userService.UpdateUserPassword(userID, data.Password); err != nil {
		log.Error("用户密码更新失败，用户ID：", userID, "，数据：", data)
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "更新密码成功，下次登录时生效", gin.H{})
	}
}
