package scripts

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db/model"
	"main/service/userService"
)

var Ctx = context.Background()

/*CreateFirstSuperAdmin
* @Description: 创建第一个超级管理员（这个会判断能否创建
 */
func CreateFirstSuperAdmin() {
	// 1. 判断是否需要创建第一个超级管理员账号
	if config.Config.Runtime.CreateFirstAdmin == false {
		// 不需要则返回
		log.Info("初始化超级管理员账号：无需创建")
		return
	}

	// 添加默认值：账号、用户名、密码
	var firstAdminName string
	if len(config.Config.Runtime.FirstAdminName) > 0 {
		firstAdminName = config.Config.Runtime.FirstAdminName
	} else {
		firstAdminName = "超级管理员"
	}
	var firstAdminAccount string
	if len(config.Config.Runtime.FirstAdminAccount) > 0 {
		firstAdminAccount = config.Config.Runtime.FirstAdminAccount
	} else {
		firstAdminAccount = "admin"
	}
	var firstAdminPassword string
	if len(config.Config.Runtime.FirstAdminPassword) > 0 {
		firstAdminPassword = config.Config.Runtime.FirstAdminPassword
	} else {
		firstAdminPassword = "12345678"
	}
	// 2. 查看超级管理员账号是否已创建（避免重复创建）
	log.Info("超级管理员账号：开始创建")
	var firstAdmin = model.BBSUser{
		ID:       0,
		AuthID:   0,
		Account:  firstAdminAccount,
		Name:     firstAdminName,
		Password: firstAdminPassword,
		IsAdmin:  20,
		Status:   0,
		//LastLoginAt: utils.DateTime{},
		LastLoginIP: "0.0.0.0",
		Email:       "test@test.test",
		Mobile:      "12345678901",
		Gender:      1,
		//Birthday:    utils.DateTime{},
		Signature: "测试 Signature",
		Company:   "测试 Company",
		Website:   "测试 Website",
	}

	if err := userService.CreateBBSUser(&firstAdmin); err != nil {
		log.Info("超级管理员账号：创建失败，", err)
	} else {
		log.Info("超级管理员账号：创建成功")
	}
}
