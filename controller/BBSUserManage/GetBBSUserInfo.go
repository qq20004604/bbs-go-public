package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	"main/service/userService"
	"main/utils"
)

/*GetBBSUserInfo
* @Description: 获取用户信息
* @param c
 */
func GetBBSUserInfo(c *gin.Context) {
	// 1. 去获取该用户信息
	errGet, resUserData := userService.GetAdvanceBBSUserInfoBySelf(c)
	if errGet != nil {
		utils.ErrorJson(c, errGet.Error())
		return
	}

	// 2. 返回数据给用户
	utils.SuccessJson(c, "查询成功", resUserData)
}
