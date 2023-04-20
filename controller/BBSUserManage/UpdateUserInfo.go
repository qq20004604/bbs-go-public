package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*UpdateUserInfo
* @Description: 管理员修改其他用户信息
* @param c
 */
func UpdateUserInfo(c *gin.Context) {
	var data userService.UpdateUserInfoRequest
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

	if err := userService.UpdateUserInfo(data); err != nil {
		log.Error("用户信息更新失败，用户ID：", data.ID, "，数据：", data)
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "更新成功", gin.H{})
	}
}
