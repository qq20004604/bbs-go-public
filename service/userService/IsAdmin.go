package userService

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/db/model"
)

/*IsAdmin
* @Description: 判断用户当前是否是管理员
* @param c
* @return bool	是否是管理员
* @return int	如果是管理员，则返回管理员权限级别
 */
func IsAdmin(c *gin.Context) (bool, int) {
	// 1. 在线检查
	isOnline, userID := IsOnline(c)
	if !isOnline {
		return false, 0
	}

	// 2. 查询用户数据
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return false, 0
	}

	// 3. 检查用户状态，如果不正常，则返回其状态。如果禁止某个状态用户登录，则修改这个状态
	// 但理论上，能登录，就意味着这个账号能正常获取自己的个人数据
	if user.IsUserStatusNormal() == false {
		return false, 0
	}

	// 4. 此时已经获取到用户信息了
	if user.IsAdmin >= 10 {
		return true, user.IsAdmin
	} else {
		return false, 0
	}
}
