package utils

import "github.com/gin-gonic/gin"

func ErrorJson(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": nil,
	})
}

func SuccessJson(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}
