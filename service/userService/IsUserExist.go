package userService

import (
	"main/db"
	"main/db/model"
)

/*IsUserExistByAccount
* @Description: 用户是否存在。识别标准是用户名/密码/邮箱地址
* @param c
 */
func IsUserExistByAccount(userData *BBSUserExist) (bool, error) {
	var count int64
	query := db.DbItem.Model(model.BBSUser{})
	// 这3个字段是必填且不能为空，所以添加进来
	query.Where("account = ? OR name = ?", userData.Account, userData.Name)

	// Email和Mobile 可能为空，所以不能作为判重条件，只有在有值时才判重
	if userData.AuthID != 0 {
		query.Or("auth_id = ?", userData.AuthID)
	}
	if userData.Email != "" {
		query.Or("email = ?", userData.Email)
	}
	if userData.Mobile != "" {
		query.Or("email = ?", userData.Mobile)
	}

	err := query.Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
