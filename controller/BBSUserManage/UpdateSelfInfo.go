package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*UpdateSelfInfo
* @Description: 更新本人用户信息
* @param c
 */
func UpdateSelfInfo(c *gin.Context) {
	var data userService.UpdateSelfInfoRequest
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

	// 2. 生成一个对象，拷贝要被处理的数据
	var UpdateUserInfoData = userService.UpdateUserInfoRequest{
		ID:        userID,
		Name:      data.Name,
		Email:     data.Email,
		Mobile:    data.Mobile,
		Gender:    data.Gender,
		Birthday:  data.Birthday,
		Signature: data.Signature,
		Company:   data.Company,
		Website:   data.Website,
	}

	if err := userService.UpdateUserInfo(UpdateUserInfoData); err != nil {
		log.Error("用户信息更新失败，用户ID：", userID, "，数据：", data)
		utils.ErrorJson(c, err.Error())
		return
	} else {
		utils.SuccessJson(c, "更新成功", gin.H{})
	}
}
