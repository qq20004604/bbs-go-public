package userService

import (
	"github.com/gin-gonic/gin"
)

/*Logout
* @Description: 登出
* @param c
 */
func Logout(c *gin.Context, token string) {
	// 1. 先根据 token 去删除 redis 里2个键值对
	ClearTokenByRedis(c, token)

	// 2. 再清除 cookie
	ClearLoginByCookie(c)
}
