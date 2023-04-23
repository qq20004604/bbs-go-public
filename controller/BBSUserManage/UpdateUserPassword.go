package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*UpdateUserPassword
* @Description: 管理员修改用户密码
* @param c
 */
func UpdateUserPassword(c *gin.Context) {
	var data userService.UpdateUserPasswordRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 当前用户是否是管理员
	isAdmin, _ := userService.IsAdmin(c)
	if !isAdmin {
		msg := "只有管理员才能执行本操作"
		utils.ErrorJson(c, msg)
		return
	}

	// 2. 然后执行密码更新
	if err := userService.UpdateUserPassword(data.ID, data.Password); err != nil {
		log.Error("用户密码更新失败，用户ID：", data.ID, "，数据：", data)
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "修改密码成功，下次登录时生效", gin.H{})
	}
}
