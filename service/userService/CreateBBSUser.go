package userService

import (
	"errors"
	"main/db"
	"main/db/model"
)

/*CreateBBSUser
* @Description:	向数据库中插入一条新的用户记录
* @param user	用户数据
* @return error	报错信息
 */
func CreateBBSUser(user *model.BBSUser) error {
	// 再基于用户数据并插入到数据库表之前，先执行一些前置操作，例如密码加盐等
	if err := beforeCreateBBSUser(user); err != nil {
		return err
	}

	// 执行插入操作
	result := db.DbItem.Model(model.BBSUser{}).Create(&user)

	// 插入失败
	if result.Error != nil {
		return result.Error
	}

	// 插入成功
	return nil
}

/*beforeCreateBBSUser
* @Description: 在创建用户之前，进行一些特殊处理
* @param user	用户信息
* @return error
 */
func beforeCreateBBSUser(user *model.BBSUser) error {
	// 1. 对密码进行加密处理，只存储加密后的密码
	pwWithSalt := SaltAndHashPassword(user.Account, user.Password)
	user.Password = pwWithSalt

	// 2. 对用户进行查重，如果该用户重复的话，则不允许创建
	var userData = BBSUserExist{
		AuthID:  user.AuthID,
		Account: user.Account,
		Name:    user.Name,
		Email:   user.Email,
		Mobile:  user.Mobile,
	}

	if isExist, err := IsUserExistByAccount(&userData); err != nil {
		return err
	} else if isExist {
		return errors.New("该用户已存在")
	}

	// final. 一切正常
	return nil
}
