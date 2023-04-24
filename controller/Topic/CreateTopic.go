package Topic

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/topicService"
	"main/service/userService"
	"main/utils"
)

/*CreateTopic
* @Description: 新增一个主题帖
* @param c
 */
func CreateTopic(c *gin.Context) {
	var data topicService.CreateTopicRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 获取到当前用户的ID
	isOnline, userID := userService.IsOnline(c)
	if !isOnline {
		utils.ErrorJson(c, "请登录后再进行操作")
		return
	}

	// 3. 执行发帖操作
	if err := topicService.CreateTopic(data, userID, c); err != nil {
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "发帖成功", gin.H{})
	}
}
