package Topic

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/topicService"
	"main/service/userService"
	"main/utils"
)

/*GetTopicListByPage
* @Description: 批量获取所有帖子（只包含未删除的，后续可能会添加条件：必须审批通过）
* @param c
 */
func GetTopicListByPage(c *gin.Context) {
	var data topicService.GetTopicListRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 先判断是否已登录
	isOnline, _ := userService.IsOnline(c)
	if !isOnline {
		utils.ErrorJson(c, "请登录后再进行操作")
		return
	}

	// 2. 获取所有未被删除的帖子
	res := topicService.GetNormalTopicListByPage(data.Page)

	utils.SuccessJson(c, "查询成功", res)
}
