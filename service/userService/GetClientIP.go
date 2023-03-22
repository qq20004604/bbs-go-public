package userService

import "github.com/gin-gonic/gin"

/*GetClientIP
* @Description: 获取当前用户的IP
* @param c
* @return string
 */
func GetClientIP(c *gin.Context) string {
	clientIP := c.Request.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = c.Request.Header.Get("X-Real-Ip")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	return clientIP
}
