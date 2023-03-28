package userService

import (
	"github.com/gin-gonic/gin"
	"main/config"
)

/*SetLoginByCookie
* @Description: 将token写入cookie
* @param c
* @param token
 */
func SetLoginByCookie(c *gin.Context, token string) {
	// 将 token 再写入 cookie 里，默认过期时间（14天）
	var cookieMaxAge = int(3600 * config.Config.Common.LoginExpireTime)
	if config.Config.Common.LoginExpireTime == 0 {
		cookieMaxAge = 3600 * 168
	}
	c.SetCookie(config.Config.Common.HeaderTokenName, token, cookieMaxAge, "/", "", false, true)
}

/*ClearLoginByCookie
* @Description: 	清除用户登录状态
* @param c
 */
func ClearLoginByCookie(c *gin.Context) {
	// 报错，则设置cookie为空
	c.SetCookie(config.Config.Common.HeaderTokenName, "", -1, "/", "", false, true)
}
