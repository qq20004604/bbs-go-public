package userService

import (
	"errors"
	"main/db"
	"main/db/model"
)

/*UpdateUserInfo
* @Description: 更新用户数据
* @param userInfo
* @return error
 */
func UpdateUserInfo(userInfo UpdateUserInfoRequest) error {
	user := model.BBSUser{}
	if err := db.DbItem.Model(&model.BBSUser{}).Where("id = ?", userInfo.ID).First(&user).Error; err != nil {
		return errors.New("用户不存在")
	}

	// 更新提供的字段
	updateFields := make(map[string]interface{})
	if userInfo.Name != "" {
		updateFields["name"] = userInfo.Name
		// 在更新之前，需要判断数据库里有没有重复的名字，如果有则报错并返回
		var tempUser model.BBSUser
		db.DbItem.Model(&model.BBSUser{}).Where("name = ?", userInfo.Name).First(&tempUser)
		// 如果能查到这个同名用户，并且这个同名用户不是自己，则报错
		// 后面这个处理逻辑，主要是避免某些傻逼前端传一个同名 name 给后端
		if tempUser.ID != 0 && tempUser.ID != userInfo.ID {
			return errors.New("用户名已存在，请修改")
		}
	}
	if userInfo.Email != "" {
		updateFields["email"] = userInfo.Email
	}
	if userInfo.Mobile != "" {
		updateFields["mobile"] = userInfo.Mobile
	}
	if userInfo.Gender != 0 {
		updateFields["gender"] = userInfo.Gender
	}
	if !userInfo.Birthday.IsZero() {
		updateFields["birthday"] = userInfo.Birthday
	}
	if userInfo.Signature != "" {
		updateFields["signature"] = userInfo.Signature
	}
	if userInfo.Company != "" {
		updateFields["company"] = userInfo.Company
	}
	if userInfo.Website != "" {
		updateFields["website"] = userInfo.Website
	}

	if err := db.DbItem.Model(&user).Updates(updateFields).Error; err != nil {
		return errors.New("更新失败")
	}

	return nil
}
