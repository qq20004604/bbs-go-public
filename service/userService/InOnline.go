package userService

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*IsOnline
* @Description: 判断用户当前是否在线。需要登录后才能调用的接口，都必须先行调用这个方法
*				如果需要判断权限的话，则需要调用其他接口
* @param c
* @return bool	在线返回 true，不在线返回 false
 */
func IsOnline(c *gin.Context) bool {
	// 1. 先拿到cookie
	token, errCookie := c.Cookie("bbs-token")
	// 如果报错了，一般是没找到这个 cookie，但无论如何，把cookie清除掉
	if errCookie != nil {
		ClearLoginByCookie(c)
		return false
	}

	// 2. 拿cookie去redis里查询
	isOnline, err := isTokenExistInRedis(c, token)
	// 报错的话，一般是系统服务出错
	if errCookie != nil {
		ClearLoginByCookie(c)
		log.Error(err)
		return false
	}

	// 3. 如果能查到（token未过期，则认为在登录，返回true）
	return isOnline
}
