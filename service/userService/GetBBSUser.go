package userService

import (
	"errors"
	"main/db"
	"main/db/model"
)

/*GetBBSUserByID
* @Description: 根据用户ID获取用户信息
* @param c
* @param ID
* @return error
 */
func GetBBSUserByID(ID uint) (error, model.BBSUser) {
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("id = ?", ID).First(&user)
	if user.ID == 0 {
		return errors.New("该用户不存在"), model.BBSUser{}
	}
	if err := beforeGetBBSUser(&user); err != nil {
		return err, model.BBSUser{}
	}
	return nil, user
}

/*GetBBSUserByAccount
* @Description: 根据用户账号获取用户信息
* @param c
* @param account	用户账号
* @return error
* @return model.BBSUser
 */
func GetBBSUserByAccount(account string) (error, model.BBSUser) {
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return errors.New("该用户不存在"), model.BBSUser{}
	}
	if err := beforeGetBBSUser(&user); err != nil {
		return err, model.BBSUser{}
	}
	return nil, user
}

/*beforeGetBBSUser
* @Description: 	在成功获取到用户信息之后，返回查询到用户信息之前，判断能否正常获取查询结果
*					这个函数用于处理一些获取用户时的特殊需求
* @param user
* @return error
 */
func beforeGetBBSUser(user *model.BBSUser) error {
	return nil
}
