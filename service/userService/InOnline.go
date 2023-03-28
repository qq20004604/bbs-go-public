package userService

import (
	"github.com/gin-gonic/gin"
	"main/config"
)

/*IsOnline
* @Description: 判断用户当前是否在线。需要登录后才能调用的接口，都必须先行调用这个方法
*				如果需要判断权限的话，则需要调用其他接口
* @param c
* @return bool	在线返回 true，不在线返回 false
* @return uint	在线返回用户ID，不在线则返回 0
 */
func IsOnline(c *gin.Context) (bool, uint) {
	// 1. 先从 cookie 拿到 token
	token, errCookie := c.Cookie(config.Config.Common.HeaderTokenName)
	// 如果报错了，一般是没找到这个 cookie，但无论如何，把cookie清除掉
	if errCookie != nil {
		ClearLoginByCookie(c)
		return false, 0
	}

	// 2. 拿 token 去 redis 里查询该 token 是否有效
	err, userId := CheckTokenAvailable(c, token)
	if err != nil {
		ClearLoginByCookie(c)
		return false, 0
	}

	// 3. 如果能查到（token未过期，则认为在登录，返回true）
	return true, userId
}
