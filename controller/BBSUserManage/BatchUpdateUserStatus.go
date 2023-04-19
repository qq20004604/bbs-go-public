package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*BatchUpdateUserStatus
* @Description: 批量更新用户状态
* @param c
 */
func BatchUpdateUserStatus(c *gin.Context) {
	var data userService.BatchUpdateUserStatusRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}
	// 1. 当前用户是否是管理员
	isAdmin, userInfo := userService.IsAdmin(c)
	if !isAdmin {
		msg := "只有管理员才能更新用户状态"
		utils.ErrorJson(c, msg)
		return
	}

	// 2. 判重，被更改的用户不能包括当前用户自己
	for _, item := range data.List {
		if item == userInfo.ID {
			msg := "管理员不能更新自己的账号状态"
			utils.ErrorJson(c, msg)
			return
		}
	}

	// 3. 更新数据，如果有报错说明更新失败
	if err := userService.BatchUpdateUserStatus(data); err != nil {
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "更新成功", gin.H{})
	}
}
