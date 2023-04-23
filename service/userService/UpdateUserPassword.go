package userService

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"main/db"
	"main/db/model"
)

/*UpdateUserPassword
* @Description: 更新用户密码
* @param userID		被更新用户的ID
* @param password	更新的密码
* @return error
 */
func UpdateUserPassword(userID uint, password string) error {
	// 1. 首先根据 userID 查出来这个人的相关信息
	var user model.BBSUser
	db.DbItem.Model(&model.BBSUser{}).Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return errors.New("该用户不存在，请检查用户ID 或 重试")
	}
	newPassword := SaltAndHashPassword(user.Account, password)
	log.Info(newPassword)
	user.Password = newPassword
	if err := db.DbItem.Save(&user).Error; err != nil {
		return errors.New("更新用户密码失败，请重试")
	}
	return nil
}
