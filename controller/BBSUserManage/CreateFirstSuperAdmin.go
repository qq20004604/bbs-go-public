package BBSUserManage

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

	var firstAdminName string
	if len(config.Config.Runtime.FirstAdminName) > 0 {
		firstAdminName = config.Config.Runtime.FirstAdminName
	} else {
		firstAdminName = "超级管理员"
	}
	// 2. 查看超级管理员账号是否已创建（避免重复创建）
	log.Info("初始化超级管理员账号：开始")
	var userData = userService.BBSUserExist{
		Account: config.Config.Runtime.FirstAdminAccount,
		Name:    firstAdminName,
	}

	if isExist, err := userService.IsUserExistByAccount(&userData); err != nil {
		log.Error("超级管理员：创建失败，", err)
		return
	} else if isExist {
		log.Info("超级管理员：已存在，无需创建")
		return
	}
	log.Info("超级管理员：开始创建")
	var firstAdmin = model.BBSUser{
		ID:      0,
		AuthID:  0,
		Account: config.Config.Runtime.FirstAdminAccount,
		Name:    config.Config.Runtime.FirstAdminName,
		IsAdmin: 20,
	}

	if err := userService.CreateUser(&firstAdmin); err != nil {
		log.Error("超级管理员创建失败")
	} else {
		log.Info("超级管理员创建成功")
	}
}
