package BBSUserManage

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db/model"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*RegisterBBSUser
* @Description:  注册用户
* @param c
 */
func RegisterBBSUser(c *gin.Context) {
	var data userService.UserRegisterRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 执行注册前检查
	if isCanContinue, err := userService.BeforeRegisterBBSUser(c); isCanContinue == false {
		utils.ErrorJson(c, err.Error())
		return
	}

	// 2. 基于当前数据，生成注册用的数据
	var user model.BBSUser
	data.ConvertToBBSUser(&user)

	// 3. 注册前特殊处理。例如刚注册的账号，需要管理员审批后才允许登录，则在这里处理 user 数据
	if config.Config.Runtime.RegistrationRequiresAdminApproval {
		// 设置状态为待审核
		user.Status = model.UserStatusPendingReview
	}

	// 4. 执行注册
	if err := userService.CreateBBSUser(&user); err != nil {
		utils.ErrorJson(c, err.Error())
	} else {
		utils.SuccessJson(c, "注册成功", user)
	}
}
