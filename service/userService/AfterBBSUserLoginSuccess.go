package userService

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/db/model"
)

/*AfterBBSUserLoginSuccess
* @Description: 	更新用户信息里的最后登录IP和登录地址
* @param c
* @param bbsUserID	当前用户的ID
 */
func AfterBBSUserLoginSuccess(c *gin.Context, bbsUserID uint) {
	ip := GetClientIP(c)
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("id = ?", bbsUserID).First(&user)
	user.UpdateAfterLogin(ip)
	db.DbItem.Save(&user)
}
