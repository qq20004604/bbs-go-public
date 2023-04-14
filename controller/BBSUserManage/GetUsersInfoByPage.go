package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*GetUsersInfoByPage
* @Description: 批量获取所有用户信息
*				1. 只有管理员权限的人才能调用
* @param c
 */
func GetUsersInfoByPage(c *gin.Context) {
	var data userService.GetAllUsersInfoRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 当前用户是否是管理员
	isAdmin, _ := userService.IsAdmin(c)
	if !isAdmin {
		msg := "只有管理员才能获取所有用户信息"
		utils.ErrorJson(c, msg)
		return
	}

	// 2. 获取用户信息
	res := userService.GetUsersInfoByPage(data.Page)

	utils.SuccessJson(c, "查询成功", res)
}
