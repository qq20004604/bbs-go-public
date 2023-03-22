package userService

import "github.com/gin-gonic/gin"

/*IsTooManyLogin
* @Description: 检查当前用户是否登录太多次了
* @param c
* @return bool	true 表示登录太多了，false表示可以正常登录
 */
func IsTooManyLogin(c *gin.Context) bool {
	// todo 待完善
	return false
}
